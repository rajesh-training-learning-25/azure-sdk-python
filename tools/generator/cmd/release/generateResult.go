package release

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest"
	"github.com/Azure/azure-sdk-for-go/tools/generator/config"
	"github.com/ahmetb/go-linq/v3"
)

type GenerateResult struct {
	Readme     string
	Tag        string
	CommitHash string
	Info       []config.ReleaseRequestInfo
	Package    autorest.ChangelogResult
}

func (r GenerateResult) PackageInfo() string {
	var messages []string
	linq.From(r.Info).Select(func(item interface{}) interface{} {
		return item.(config.ReleaseRequestInfo).String()
	}).ToSlice(&messages)

	return fmt.Sprintf("%s from [%s]", r.Package.PackageName, strings.Join(messages, ", "))
}
