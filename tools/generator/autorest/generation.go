// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package autorest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/tools/apidiff/exports"
	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest/model"
	"github.com/Azure/azure-sdk-for-go/tools/generator/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type GenerateContext interface {
	SDKRoot() string
	SpecRoot() string
	RepoContent() map[string]exports.Content
}

// GenerateInput describes the input information for a package generation
type GenerateInput struct {
	// Readme is the relative path of the readme file to the root directory of azure-sdk-for-go
	Readme string
	// Tag is the readme tag to be generated
	Tag string
	// CommitHash is the head commit hash of azure-rest-api-specs
	CommitHash string
	// Options specifies the options that this generation task will be using
	Options model.Options
}

// GenerateOptions describes the options for a package generation
type GenerateOptions struct {
	// MetadataOutputRoot specifies the root directory of all the metadata goes.
	// Metadata will be generated to a temp directory if not specified.
	// The metadataOutput directory will not be removed after the generation succeeded
	MetadataOutputRoot string
	// Stderr ...
	Stderr io.Writer
	// Stdout ...
	Stdout io.Writer
	// AutoRestLogPrefix ...
	AutoRestLogPrefix string
	// ChangelogTitle
	ChangelogTitle string
}

type GenerateResult struct {
	MetadataOutputRoot string
	Metadata           GenerationMetadata
	Package            ChangelogResult
}

func GeneratePackage(ctx GenerateContext, input GenerateInput, options GenerateOptions) (*GenerateResult, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}
	if err := options.validate(); err != nil {
		return nil, err
	}

	absReadme := filepath.Join(ctx.SpecRoot(), input.Readme)
	metadataOutput := filepath.Join(options.MetadataOutputRoot, input.Tag)
	g := NewGeneratorFromOptions(input.Options).WithTag(input.Tag).WithMetadataOutput(metadataOutput).WithReadme(absReadme)

	// generate
	if err := generate(g, options.Stdout, options.Stderr, options.AutoRestLogPrefix); err != nil {
		return nil, fmt.Errorf("failed to execute autorest: %+v", err)
	}
	// write the changelog and metadata file
	result, metadata, err := changelogAndMetadata(ctx, input, metadataOutput, options.ChangelogTitle, g.Arguments())
	if err != nil {
		return nil, err
	}

	return &GenerateResult{
		MetadataOutputRoot: options.MetadataOutputRoot,
		Metadata:           *metadata,
		Package:            *result,
	}, nil
}

func generate(generator *Generator, stdout, stderr io.Writer, prefix string) error {
	stdoutPipe, _ := generator.StdoutPipe()
	stderrPipe, _ := generator.StderrPipe()
	defer stdoutPipe.Close()
	defer stderrPipe.Close()
	var arguments []string
	for _, o := range generator.Arguments() {
		arguments = append(arguments, o.Format())
	}
	log.Printf("Generation parameters: %s", strings.Join(arguments, ", "))
	_ = generator.Start()
	// we put all the output from autorest to stderr since those are logs in order not to interrupt the proper output of the release command
	go utils.ScannerPrint(bufio.NewScanner(stdoutPipe), stdout, prefix)
	go utils.ScannerPrint(bufio.NewScanner(stderrPipe), stderr, prefix)
	return generator.Wait()
}

func changelogAndMetadata(ctx GenerateContext, input GenerateInput, metadataOutput, changelogTitle string, argument []model.Option) (*ChangelogResult, *GenerationMetadata, error) {
	result, err := changelog(ctx, metadataOutput, changelogTitle)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to write changelog file: %+v", err)
	}

	// write the metadata file
	metadata, err := metadata(input, *result, argument)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to write metadata file: %+v", err)
	}

	return result, metadata, nil
}

func changelog(ctx GenerateContext, metadataOutput, changelogTitle string) (*ChangelogResult, error) {
	// parse the metadata from autorest
	metadataMap, err := NewMetadataProcessorFromLocation(metadataOutput).Process()
	if err != nil {
		return nil, fmt.Errorf("failed to parse metadata in '%s': %+v", metadataOutput, err)
	}
	// process the changelog
	changelogResults, err := NewChangelogProcessorFromContext(ctx).Process(metadataMap)
	if err != nil {
		return nil, fmt.Errorf("failed to process the changelog: %+v", err)
	}
	// we should only have one changelog
	if len(changelogResults) != 1 {
		return nil, fmt.Errorf("expecting 1 changelog result, but got %d", len(changelogResults))
	}

	changelogPath, err := WriteChangelogFile(changelogTitle, changelogResults[0])
	if err != nil {
		return nil, fmt.Errorf("failed to write changelog file: %+v", err)
	}
	log.Printf("changelog file writes to '%s'", changelogPath)
	return &changelogResults[0], nil
}

func metadata(input GenerateInput, result ChangelogResult, arguments []model.Option) (*GenerationMetadata, error) {
	metadata := getMetadata(input, result, arguments)
	metadataPath, err := WriteMetadataFile(result.PackageFullPath, metadata)
	if err != nil {
		return nil, err
	}
	log.Printf("metadata file writes to '%s'", metadataPath)
	return &metadata, nil
}

func getMetadata(input GenerateInput, result ChangelogResult, arguments []model.Option) GenerationMetadata {
	options := AdditionalOptionsToString(arguments)
	codeGenVersion := input.Options.CodeGeneratorVersion()
	return GenerationMetadata{
		CommitHash:     input.CommitHash,
		Readme:         NormalizedSpecRoot + utils.NormalizePath(input.Readme),
		Tag:            input.Tag,
		CodeGenVersion: codeGenVersion,
		RepositoryURL:  "https://github.com/Azure/azure-rest-api-specs.git",
		AutorestCommand: fmt.Sprintf("autorest --use=%s --tag=%s --go-sdk-folder=/_/azure-sdk-for-go %s /_/azure-rest-api-specs/%s",
			codeGenVersion, result.Tag, strings.Join(options, " "), utils.NormalizePath(input.Readme)),
		AdditionalProperties: GenerationMetadataAdditionalProperties{
			AdditionalOptions: strings.Join(options, " "),
		},
	}
}

func (input GenerateInput) validate() error {
	if input.Readme == "" {
		return fmt.Errorf("`Readme` cannot be empty in input")
	}
	if filepath.IsAbs(input.Readme) {
		return fmt.Errorf("`Readme` must be a relative path")
	}
	if input.Tag == "" {
		return fmt.Errorf("`Tag` cannot be empty in input")
	}
	if input.Options == nil {
		return fmt.Errorf("`Options` cannot be nil")
	}
	return nil
}

func (options *GenerateOptions) validate() error {
	if options.MetadataOutputRoot == "" {
		options.MetadataOutputRoot = filepath.Join(os.TempDir(), fmt.Sprintf("generation-metadata-%v", time.Now().Unix()))
	}
	if options.ChangelogTitle == "" {
		options.ChangelogTitle = "Change History"
	}
	return nil
}

// WriteChangelogFile writes the changelog to the disk
func WriteChangelogFile(title string, result ChangelogResult) (string, error) {
	fileContent := fmt.Sprintf(`# %s

%s`, title, result.Changelog.ToMarkdown())
	path := filepath.Join(result.PackageFullPath, ChangelogFilename)
	changelogFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer changelogFile.Close()
	if _, err := changelogFile.WriteString(fileContent); err != nil {
		return "", err
	}
	return path, nil
}

// WriteMetadataFile writes the metadata to the disk
func WriteMetadataFile(packagePath string, metadata GenerationMetadata) (string, error) {
	metadataFilepath := filepath.Join(packagePath, MetadataFilename)
	metadataFile, err := os.Create(metadataFilepath)
	if err != nil {
		return "", err
	}
	defer metadataFile.Close()

	// marshal metadata
	b, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return "", fmt.Errorf("cannot marshal metadata: %+v", err)
	}

	if _, err := metadataFile.Write(b); err != nil {
		return "", err
	}
	return metadataFilepath, nil
}
