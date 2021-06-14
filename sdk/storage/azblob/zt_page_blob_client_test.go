// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

//
//import (
//	"bytes"
//	"context"
//	"crypto/md5"
//	chk "gopkg.in/check.v1"
//	"io/ioutil"
//	"strings"
//	"time"
//)
//
//func (s *azblobTestSuite) TestPutGetPages() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	contentSize := 1024
//	offset, end, count := int64(0), int64(contentSize-1), int64(contentSize)
//	uploadPagesOptions := UploadPagesOptions{PageRange: &HttpRange{offset, count}}
//	putResp, err := pbClient.UploadPages(context.Background(), getReaderToRandomBytes(1024), &uploadPagesOptions)
//	_assert.Nil(err)
//	c.Assert(putResp.RawResponse.StatusCode, chk.Equals, 201)
//	c.Assert(putResp.LastModified, chk.NotNil)
//	c.Assert((*putResp.LastModified).IsZero(), chk.Equals, false)
//	c.Assert(putResp.ETag, chk.NotNil)
//	c.Assert(putResp.ContentMD5, chk.IsNil)
//	c.Assert(*putResp.BlobSequenceNumber, chk.Equals, int64(0))
//	c.Assert(*putResp.RequestID, chk.NotNil)
//	c.Assert(*putResp.Version, chk.NotNil)
//	c.Assert(putResp.Date, chk.NotNil)
//	c.Assert((*putResp.Date).IsZero(), chk.Equals, false)
//
//	pageList, err := pbClient.GetPageRanges(context.Background(), HttpRange{0, 1023}, nil)
//	_assert.Nil(err)
//	c.Assert(pageList.RawResponse.StatusCode, chk.Equals, 200)
//	c.Assert(pageList.LastModified, chk.NotNil)
//	c.Assert((*pageList.LastModified).IsZero(), chk.Equals, false)
//	c.Assert(pageList.ETag, chk.NotNil)
//	c.Assert(*pageList.BlobContentLength, chk.Equals, int64(512*10))
//	c.Assert(*pageList.RequestID, chk.NotNil)
//	c.Assert(*pageList.Version, chk.NotNil)
//	c.Assert(pageList.Date, chk.NotNil)
//	c.Assert((*pageList.Date).IsZero(), chk.Equals, false)
//	c.Assert(pageList.PageList, chk.NotNil)
//	pageRangeResp := pageList.PageList.PageRange
//	c.Assert(*pageRangeResp, chk.HasLen, 1)
//	rawStart, rawEnd := (*pageRangeResp)[0].Raw()
//	c.Assert(rawStart, chk.Equals, offset)
//	c.Assert(rawEnd, chk.Equals, end)
//}
//
//func (s *azblobTestSuite) TestUploadPagesFromURL() {
//	bsu := getServiceClient(nil)
//	credential, err := getGenericCredential("")
//	if err != nil {
//		c.Fatal("Invalid credential")
//	}
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	contentSize := 4 * 1024 * 1024 // 4MB
//	r, sourceData := getRandomDataAndReader(contentSize)
//	ctx := context.Background() // Use default Background context
//	srcBlob, _ := createNewPageBlobWithSize(c, containerClient, int64(contentSize))
//	destBlob, _ := createNewPageBlobWithSize(c, containerClient, int64(contentSize))
//
//	offset, _, count := int64(0), int64(contentSize-1), int64(contentSize)
//	uploadPagesOptions := UploadPagesOptions{PageRange: &HttpRange{offset, count}}
//	uploadSrcResp1, err := srcBlob.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//	c.Assert(uploadSrcResp1.RawResponse.StatusCode, chk.Equals, 201)
//	c.Assert(uploadSrcResp1.LastModified, chk.NotNil)
//	c.Assert((*uploadSrcResp1.LastModified).IsZero(), chk.Equals, false)
//	c.Assert(uploadSrcResp1.ETag, chk.NotNil)
//	c.Assert(uploadSrcResp1.ContentMD5, chk.IsNil)
//	c.Assert(*uploadSrcResp1.BlobSequenceNumber, chk.Equals, int64(0))
//	c.Assert(*uploadSrcResp1.RequestID, chk.NotNil)
//	c.Assert(*uploadSrcResp1.Version, chk.NotNil)
//	c.Assert(uploadSrcResp1.Date, chk.NotNil)
//	c.Assert((*uploadSrcResp1.Date).IsZero(), chk.Equals, false)
//
//	// Get source pbClient URL with SAS for UploadPagesFromURL.
//	srcBlobParts := NewBlobURLParts(srcBlob.URL())
//
//	srcBlobParts.SAS, err = BlobSASSignatureValues{
//		Protocol:      SASProtocolHTTPS,                     // Users MUST use HTTPS (not HTTP)
//		ExpiryTime:    time.Now().UTC().Add(48 * time.Hour), // 48-hours before expiration
//		ContainerName: srcBlobParts.ContainerName,
//		BlobName:      srcBlobParts.BlobName,
//		Permissions:   BlobSASPermissions{Read: true}.String(),
//	}.NewSASQueryParameters(credential)
//	if err != nil {
//		c.Fatal(err)
//	}
//
//	srcBlobURLWithSAS := srcBlobParts.URL()
//
//	// Upload page from URL.
//	pResp1, err := destBlob.UploadPagesFromURL(ctx, srcBlobURLWithSAS, 0, 0, int64(contentSize), nil)
//	_assert.Nil(err)
//	c.Assert(pResp1.RawResponse.StatusCode, chk.Equals, 201)
//	c.Assert(pResp1.ETag, chk.NotNil)
//	c.Assert(pResp1.LastModified, chk.NotNil)
//	c.Assert(pResp1.ContentMD5, chk.NotNil)
//	c.Assert(pResp1.RequestID, chk.NotNil)
//	c.Assert(pResp1.Version, chk.NotNil)
//	c.Assert(pResp1.Date, chk.NotNil)
//	c.Assert((*pResp1.Date).IsZero(), chk.Equals, false)
//
//	// Check data integrity through downloading.
//	downloadResp, err := destBlob.Download(ctx, nil)
//	_assert.Nil(err)
//	destData, err := ioutil.ReadAll(downloadResp.Body(RetryReaderOptions{}))
//	_assert.Nil(err)
//	c.Assert(destData, chk.DeepEquals, sourceData)
//}
//
//func (s *azblobTestSuite) TestUploadPagesFromURLWithMD5() {
//	bsu := getServiceClient(nil)
//	credential, err := getGenericCredential("")
//	if err != nil {
//		c.Fatal("Invalid credential")
//	}
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	contentSize := 4 * 1024 * 1024 // 4MB
//	r, sourceData := getRandomDataAndReader(contentSize)
//	md5Value := md5.Sum(sourceData)
//	contentMD5 := md5Value[:]
//	ctx := context.Background() // Use default Background context
//	srcBlob, _ := createNewPageBlobWithSize(c, containerClient, int64(contentSize))
//	destBlob, _ := createNewPageBlobWithSize(c, containerClient, int64(contentSize))
//
//	// Prepare source pbClient for copy.
//	offset, _, count := int64(0), int64(contentSize-1), int64(contentSize)
//	uploadPagesOptions := UploadPagesOptions{PageRange: &HttpRange{offset, count}}
//	uploadSrcResp1, err := srcBlob.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//	c.Assert(uploadSrcResp1.RawResponse.StatusCode, chk.Equals, 201)
//
//	// Get source pbClient URL with SAS for UploadPagesFromURL.
//	srcBlobParts := NewBlobURLParts(srcBlob.URL())
//
//	srcBlobParts.SAS, err = BlobSASSignatureValues{
//		Protocol:      SASProtocolHTTPS,                     // Users MUST use HTTPS (not HTTP)
//		ExpiryTime:    time.Now().UTC().Add(48 * time.Hour), // 48-hours before expiration
//		ContainerName: srcBlobParts.ContainerName,
//		BlobName:      srcBlobParts.BlobName,
//		Permissions:   BlobSASPermissions{Read: true}.String(),
//	}.NewSASQueryParameters(credential)
//	if err != nil {
//		c.Fatal(err)
//	}
//
//	srcBlobURLWithSAS := srcBlobParts.URL()
//
//	// Upload page from URL with MD5.
//	uploadPagesFromURLOptions := UploadPagesFromURLOptions{
//		SourceContentMD5: &contentMD5,
//	}
//	pResp1, err := destBlob.UploadPagesFromURL(ctx, srcBlobURLWithSAS, 0, 0, int64(contentSize), &uploadPagesFromURLOptions)
//	_assert.Nil(err)
//	c.Assert(pResp1.RawResponse.StatusCode, chk.Equals, 201)
//	c.Assert(pResp1.ETag, chk.NotNil)
//	c.Assert(pResp1.LastModified, chk.NotNil)
//	c.Assert(pResp1.ContentMD5, chk.NotNil)
//	c.Assert(*pResp1.ContentMD5, chk.DeepEquals, contentMD5)
//	c.Assert(pResp1.RequestID, chk.NotNil)
//	c.Assert(pResp1.Version, chk.NotNil)
//	c.Assert(pResp1.Date, chk.NotNil)
//	c.Assert((*pResp1.Date).IsZero(), chk.Equals, false)
//	c.Assert(*pResp1.BlobSequenceNumber, chk.Equals, int64(0))
//
//	// Check data integrity through downloading.
//	downloadResp, err := destBlob.Download(ctx, nil)
//	_assert.Nil(err)
//	destData, err := ioutil.ReadAll(downloadResp.Body(RetryReaderOptions{}))
//	_assert.Nil(err)
//	c.Assert(destData, chk.DeepEquals, sourceData)
//
//	// Upload page from URL with bad MD5
//	_, badMD5 := getRandomDataAndReader(16)
//	badContentMD5 := badMD5[:]
//	uploadPagesFromURLOptions = UploadPagesFromURLOptions{
//		SourceContentMD5: &badContentMD5,
//	}
//	_, err = destBlob.UploadPagesFromURL(ctx, srcBlobURLWithSAS, 0, 0, int64(contentSize), &uploadPagesFromURLOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeMD5Mismatch)
//}
//
//func (s *azblobTestSuite) TestClearDiffPages() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	contentSize := 2 * 1024
//	r := getReaderToRandomBytes(contentSize)
//	offset, _, count := int64(0), int64(contentSize-1), int64(contentSize)
//	uploadPagesOptions := UploadPagesOptions{PageRange: &HttpRange{offset, count}}
//	_, err := pbClient.UploadPages(context.Background(), r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	snapshotResp, err := pbClient.CreateSnapshot(context.Background(), nil)
//	_assert.Nil(err)
//
//	offset1, end1, count1 := int64(contentSize), int64(2*contentSize-1), int64(contentSize)
//	uploadPagesOptions1 := UploadPagesOptions{PageRange: &HttpRange{offset1, count1}}
//	_, err = pbClient.UploadPages(context.Background(), getReaderToRandomBytes(2048), &uploadPagesOptions1)
//	_assert.Nil(err)
//
//	pageListResp, err := pbClient.GetPageRangesDiff(context.Background(), HttpRange{0, 4096}, *snapshotResp.Snapshot, nil)
//	_assert.Nil(err)
//	pageRangeResp := pageListResp.PageList.PageRange
//	c.Assert(pageRangeResp, chk.NotNil)
//	c.Assert(*pageRangeResp, chk.HasLen, 1)
//	// c.Assert((*pageRangeResp)[0], chk.DeepEquals, PageRange{Start: &offset1, End: &end1})
//	rawStart, rawEnd := (*pageRangeResp)[0].Raw()
//	c.Assert(rawStart, chk.Equals, offset1)
//	c.Assert(rawEnd, chk.Equals, end1)
//
//	clearResp, err := pbClient.ClearPages(context.Background(), HttpRange{2048, 2048}, nil)
//	_assert.Nil(err)
//	c.Assert(clearResp.RawResponse.StatusCode, chk.Equals, 201)
//
//	pageListResp, err = pbClient.GetPageRangesDiff(context.Background(), HttpRange{0, 4095}, *snapshotResp.Snapshot, nil)
//	_assert.Nil(err)
//	c.Assert(pageListResp.PageList.PageRange, chk.IsNil)
//}
//
//func waitForIncrementalCopy(c *chk.C, copyBlobClient PageBlobClient, blobCopyResponse *PageBlobCopyIncrementalResponse) *string {
//	status := *blobCopyResponse.CopyStatus
//	var getPropertiesAndMetadataResult BlobGetPropertiesResponse
//	// Wait for the copy to finish
//	start := time.Now()
//	for status != CopyStatusSuccess {
//		getPropertiesAndMetadataResult, _ = copyBlobClient.GetProperties(ctx, nil)
//		status = *getPropertiesAndMetadataResult.CopyStatus
//		currentTime := time.Now()
//		if currentTime.Sub(start) >= time.Minute {
//			c.Fail()
//		}
//	}
//	return getPropertiesAndMetadataResult.DestinationSnapshot
//}
//
//func (s *azblobTestSuite) TestIncrementalCopy() {
//	bsu := getServiceClient(nil)
//
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	accessType := PublicAccessBlob
//	setAccessPolicyOptions := SetAccessPolicyOptions{
//		ContainerSetAccessPolicyOptions: ContainerSetAccessPolicyOptions{Access: &accessType},
//	}
//	_, err := containerClient.SetAccessPolicy(context.Background(), &setAccessPolicyOptions)
//	_assert.Nil(err)
//
//	srcBlob, _ := createNewPageBlob(c, containerClient)
//	contentSize := 1024
//	r := getReaderToRandomBytes(contentSize)
//	offset, _, count := int64(0), int64(0)+int64(contentSize-1), int64(contentSize)
//	uploadPagesOptions := UploadPagesOptions{PageRange: &HttpRange{offset, count}}
//	_, err = srcBlob.UploadPages(context.Background(), r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	snapshotResp, err := srcBlob.CreateSnapshot(context.Background(), nil)
//	_assert.Nil(err)
//
//	dstBlob := containerClient.NewPageBlobClient(generateBlobName())
//
//	resp, err := dstBlob.StartCopyIncremental(context.Background(), srcBlob.URL(), *snapshotResp.Snapshot, nil)
//	_assert.Nil(err)
//	c.Assert(resp.RawResponse.StatusCode, chk.Equals, 202)
//	c.Assert(resp.LastModified, chk.NotNil)
//	c.Assert((*resp.LastModified).IsZero(), chk.Equals, false)
//	c.Assert(resp.ETag, chk.NotNil)
//	c.Assert(*resp.RequestID, chk.Not(chk.Equals), "")
//	c.Assert(*resp.Version, chk.Not(chk.Equals), "")
//	c.Assert(resp.Date, chk.NotNil)
//	c.Assert((*resp.Date).IsZero(), chk.Equals, false)
//	c.Assert(*resp.CopyID, chk.Not(chk.Equals), "")
//	c.Assert(*resp.CopyStatus, chk.Equals, CopyStatusPending)
//
//	waitForIncrementalCopy(c, dstBlob, &resp)
//}
//
//func (s *azblobTestSuite) TestResizePageBlob() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	pbClient, _ := createNewPageBlob(c, containerClient)
//	resp, err := pbClient.Resize(context.Background(), 2048, nil)
//	_assert.Nil(err)
//	c.Assert(resp.RawResponse.StatusCode, chk.Equals, 200)
//
//	resp, err = pbClient.Resize(context.Background(), 8192, nil)
//	_assert.Nil(err)
//	c.Assert(resp.RawResponse.StatusCode, chk.Equals, 200)
//
//	resp2, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(*resp2.ContentLength, chk.Equals, int64(8192))
//}
//
//func (s *azblobTestSuite) TestPageSequenceNumbers() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	defer deleteContainer(containerClient)
//
//	sequenceNumber := int64(0)
//	actionType := SequenceNumberActionIncrement
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	resp, err := pbClient.UpdateSequenceNumber(context.Background(), &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//	c.Assert(resp.RawResponse.StatusCode, chk.Equals, 200)
//
//	sequenceNumber = int64(7)
//	actionType = SequenceNumberActionMax
//	updateSequenceNumberPageBlob = UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	resp, err = pbClient.UpdateSequenceNumber(context.Background(), &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//	c.Assert(resp.RawResponse.StatusCode, chk.Equals, 200)
//
//	sequenceNumber = int64(11)
//	actionType = SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob = UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	resp, err = pbClient.UpdateSequenceNumber(context.Background(), &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//	c.Assert(resp.RawResponse.StatusCode, chk.Equals, 200)
//}
//
//func (s *azblobTestSuite) TestPutPagesWithMD5() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	// put page with valid MD5
//	contentSize := 1024
//	readerToBody, body := getRandomDataAndReader(contentSize)
//	offset, _, count := int64(0), int64(0)+int64(contentSize-1), int64(contentSize)
//	md5Value := md5.Sum(body)
//	contentMD5 := md5Value[:]
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange:               &HttpRange{offset, count},
//		TransactionalContentMD5: &contentMD5,
//	}
//
//	putResp, err := pbClient.UploadPages(context.Background(), readerToBody, &uploadPagesOptions)
//	_assert.Nil(err)
//	c.Assert(putResp.RawResponse.StatusCode, chk.Equals, 201)
//	c.Assert(putResp.LastModified, chk.NotNil)
//	c.Assert((*putResp.LastModified).IsZero(), chk.Equals, false)
//	c.Assert(putResp.ETag, chk.NotNil)
//	c.Assert(putResp.ContentMD5, chk.NotNil)
//	c.Assert(*putResp.ContentMD5, chk.DeepEquals, contentMD5)
//	c.Assert(*putResp.BlobSequenceNumber, chk.Equals, int64(0))
//	c.Assert(*putResp.RequestID, chk.NotNil)
//	c.Assert(*putResp.Version, chk.NotNil)
//	c.Assert(putResp.Date, chk.NotNil)
//	c.Assert((*putResp.Date).IsZero(), chk.Equals, false)
//
//	// put page with bad MD5
//	readerToBody, body = getRandomDataAndReader(1024)
//	_, badMD5 := getRandomDataAndReader(16)
//	basContentMD5 := badMD5[:]
//	uploadPagesOptions = UploadPagesOptions{
//		PageRange:               &HttpRange{offset, count},
//		TransactionalContentMD5: &basContentMD5,
//	}
//	putResp, err = pbClient.UploadPages(context.Background(), readerToBody, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeMD5Mismatch)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageSizeInvalid() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := getPageBlobClient(c, containerClient)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//	}
//	_, err := pbClient.Create(ctx, 1, &createPageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeInvalidHeaderValue)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageSequenceInvalid() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := getPageBlobClient(c, containerClient)
//
//	sequenceNumber := int64(-1)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.NotNil(err)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageMetadataNonEmpty() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := getPageBlobClient(c, containerClient)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &basicMetadata,
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.Nil(err)
//
//	resp, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(resp.Metadata, chk.NotNil)
//	c.Assert(resp.Metadata, chk.DeepEquals, basicMetadata)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageMetadataEmpty() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := getPageBlobClient(c, containerClient)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &map[string]string{},
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//
//	resp, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(resp.Metadata, chk.IsNil)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageMetadataInvalid() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := getPageBlobClient(c, containerClient)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &map[string]string{"In valid1": "bar"},
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.NotNil(err)
//	c.Assert(strings.Contains(err.Error(), invalidHeaderErrorSubstring), chk.Equals, true)
//
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageHTTPHeaders() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := getPageBlobClient(c, containerClient)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		BlobHTTPHeaders:    &basicHeaders,
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.Nil(err)
//
//	resp, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	h := resp.GetHTTPHeaders()
//	c.Assert(h, chk.DeepEquals, basicHeaders)
//}
//
//func validatePageBlobPut(c *chk.C, pbClient PageBlobClient) {
//	resp, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(resp.Metadata, chk.NotNil)
//	c.Assert(resp.Metadata, chk.DeepEquals, basicMetadata)
//	c.Assert(resp.GetHTTPHeaders(), chk.DeepEquals, basicHeaders)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageIfModifiedSinceTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient) // Originally created without metadata
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &basicMetadata,
//		BlobHTTPHeaders:    &basicHeaders,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.Nil(err)
//
//	validatePageBlobPut(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageIfModifiedSinceFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient) // Originally created without metadata
//
//	currentTime := getRelativeTimeGMT(10)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &basicMetadata,
//		BlobHTTPHeaders:    &basicHeaders,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageIfUnmodifiedSinceTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient) // Originally created without metadata
//
//	currentTime := getRelativeTimeGMT(10)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &basicMetadata,
//		BlobHTTPHeaders:    &basicHeaders,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.Nil(err)
//
//	validatePageBlobPut(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageIfUnmodifiedSinceFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient) // Originally created without metadata
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &basicMetadata,
//		BlobHTTPHeaders:    &basicHeaders,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageIfMatchTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient) // Originally created without metadata
//
//	resp, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &basicMetadata,
//		BlobHTTPHeaders:    &basicHeaders,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: resp.ETag,
//			},
//		},
//	}
//	_, err = pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.Nil(err)
//
//	validatePageBlobPut(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageIfMatchFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient) // Originally created without metadata
//
//	sequenceNumber := int64(0)
//	eTag := "garbage"
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &basicMetadata,
//		BlobHTTPHeaders:    &basicHeaders,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageIfNoneMatchTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient) // Originally created without metadata
//
//	sequenceNumber := int64(0)
//	eTag := "garbage"
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &basicMetadata,
//		BlobHTTPHeaders:    &basicHeaders,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.Nil(err)
//
//	validatePageBlobPut(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobCreatePageIfNoneMatchFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient) // Originally created without metadata
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	sequenceNumber := int64(0)
//	createPageBlobOptions := CreatePageBlobOptions{
//		BlobSequenceNumber: &sequenceNumber,
//		Metadata:           &basicMetadata,
//		BlobHTTPHeaders:    &basicHeaders,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.Create(ctx, PageBlobPageBytes, &createPageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesInvalidRange() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	contentSize := 1024
//	r := getReaderToRandomBytes(contentSize)
//	offset, count := int64(0), int64(contentSize/2)
//	uploadPagesOptions := UploadPagesOptions{PageRange: &HttpRange{offset, count}}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	c.Assert(err, chk.Not(chk.IsNil))
//}
//
//// Body cannot be nil check already added in the request preparer
////func (s *azblobTestSuite) TestBlobPutPagesNilBody() {
////	bsu := getServiceClient()
////	containerClient, _ := createNewContainer(c, bsu)
////	defer deleteContainer(containerClient)
////	pbClient, _ := createNewPageBlob(c, containerClient)
////
////	_, err := pbClient.UploadPages(ctx, nil, nil)
////	_assert.NotNil(err)
////}
//
//func (s *azblobTestSuite) TestBlobPutPagesEmptyBody() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	r := bytes.NewReader([]byte{})
//	offset, count := int64(0), int64(0)
//	uploadPagesOptions := UploadPagesOptions{PageRange: &HttpRange{offset, count}}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesNonExistentBlob() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := getPageBlobClient(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{PageRange: &HttpRange{offset, count}}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeBlobNotFound)
//}
//
//func validateUploadPages(c *chk.C, pbClient PageBlobClient) {
//	// This will only validate a single put page at 0-PageBlobPageBytes-1
//	resp, err := pbClient.GetPageRanges(ctx, HttpRange{0, CountToEnd}, nil)
//	_assert.Nil(err)
//	pageListResp := resp.PageList.PageRange
//	start, end := int64(0), int64(PageBlobPageBytes-1)
//	rawStart, rawEnd := (*pageListResp)[0].Raw()
//	c.Assert(rawStart, chk.Equals, start)
//	c.Assert(rawEnd, chk.Equals, end)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfModifiedSinceTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	validateUploadPages(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfModifiedSinceFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfUnmodifiedSinceTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	validateUploadPages(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfUnmodifiedSinceFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfMatchTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	validateUploadPages(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfMatchFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	eTag := "garbage"
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfNoneMatchTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	eTag := "garbage"
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	validateUploadPages(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfNoneMatchFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfSequenceNumberLessThanTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	ifSequenceNumberLessThan := int64(10)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThan: &ifSequenceNumberLessThan,
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	validateUploadPages(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfSequenceNumberLessThanFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	sequenceNumber := int64(10)
//	actionType := SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	ifSequenceNumberLessThan := int64(1)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThan: &ifSequenceNumberLessThan,
//		},
//	}
//	_, err = pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeSequenceNumberConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfSequenceNumberLessThanNegOne() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	ifSequenceNumberLessThanOrEqualTo := int64(-1)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThanOrEqualTo: &ifSequenceNumberLessThanOrEqualTo,
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeInvalidInput)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfSequenceNumberLTETrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	sequenceNumber := int64(1)
//	actionType := SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	ifSequenceNumberLessThanOrEqualTo := int64(1)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThanOrEqualTo: &ifSequenceNumberLessThanOrEqualTo,
//		},
//	}
//	_, err = pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	validateUploadPages(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfSequenceNumberLTEqualFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	sequenceNumber := int64(10)
//	actionType := SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	ifSequenceNumberLessThanOrEqualTo := int64(1)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThanOrEqualTo: &ifSequenceNumberLessThanOrEqualTo,
//		},
//	}
//	_, err = pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeSequenceNumberConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfSequenceNumberLTENegOne() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	ifSequenceNumberLessThanOrEqualTo := int64(-1)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThanOrEqualTo: &ifSequenceNumberLessThanOrEqualTo,
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	//validateStorageError(c, err, )
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfSequenceNumberEqualTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	sequenceNumber := int64(1)
//	actionType := SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	ifSequenceNumberEqualTo := int64(1)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberEqualTo: &ifSequenceNumberEqualTo,
//		},
//	}
//	_, err = pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	validateUploadPages(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfSequenceNumberEqualFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	ifSequenceNumberEqualTo := int64(1)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberEqualTo: &ifSequenceNumberEqualTo,
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeSequenceNumberConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobPutPagesIfSequenceNumberEqualNegOne() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	ifSequenceNumberEqualTo := int64(-1)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberEqualTo: &ifSequenceNumberEqualTo,
//		},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions) // This will cause the library to set the value of the header to 0
//	_assert.NotNil(err)
//
//	//validateStorageError(c, err, )
//}
//
//func setupClearPagesTest() (ContainerClient, PageBlobClient) {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	return containerClient, pbClient
//}
//
//func validateClearPagesTest(c *chk.C, pbClient PageBlobClient) {
//	resp, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, nil)
//	_assert.Nil(err)
//	pageListResp := resp.PageList.PageRange
//	c.Assert(pageListResp, chk.IsNil)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesInvalidRange() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes + 1}, nil)
//	c.Assert(err, chk.Not(chk.IsNil))
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfModifiedSinceTrue() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	clearPageOptions := ClearPagesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.Nil(err)
//
//	validateClearPagesTest(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfModifiedSinceFalse() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	clearPageOptions := ClearPagesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfUnmodifiedSinceTrue() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	clearPageOptions := ClearPagesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.Nil(err)
//
//	validateClearPagesTest(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfUnmodifiedSinceFalse() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	clearPageOptions := ClearPagesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfMatchTrue() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	clearPageOptions := ClearPagesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.Nil(err)
//
//	validateClearPagesTest(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfMatchFalse() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	eTag := "garbage"
//	clearPageOptions := ClearPagesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfNoneMatchTrue() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	eTag := "garbage"
//	clearPageOptions := ClearPagesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.Nil(err)
//
//	validateClearPagesTest(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfNoneMatchFalse() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	clearPageOptions := ClearPagesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfSequenceNumberLessThanTrue() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	ifSequenceNumberLessThan := int64(10)
//	clearPageOptions := ClearPagesOptions{
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThan: &ifSequenceNumberLessThan,
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.Nil(err)
//
//	validateClearPagesTest(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfSequenceNumberLessThanFalse() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	sequenceNumber := int64(10)
//	actionType := SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	ifSequenceNumberLessThan := int64(1)
//	clearPageOptions := ClearPagesOptions{
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThan: &ifSequenceNumberLessThan,
//		},
//	}
//	_, err = pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeSequenceNumberConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfSequenceNumberLessThanNegOne() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	ifSequenceNumberLessThan := int64(-1)
//	clearPageOptions := ClearPagesOptions{
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThan: &ifSequenceNumberLessThan,
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeInvalidInput)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfSequenceNumberLTETrue() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	ifSequenceNumberLessThanOrEqualTo := int64(10)
//	clearPageOptions := ClearPagesOptions{
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThanOrEqualTo: &ifSequenceNumberLessThanOrEqualTo,
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.Nil(err)
//
//	validateClearPagesTest(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfSequenceNumberLTEFalse() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	sequenceNumber := int64(10)
//	actionType := SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	ifSequenceNumberLessThanOrEqualTo := int64(1)
//	clearPageOptions := ClearPagesOptions{
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThanOrEqualTo: &ifSequenceNumberLessThanOrEqualTo,
//		},
//	}
//	_, err = pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeSequenceNumberConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfSequenceNumberLTENegOne() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	ifSequenceNumberLessThanOrEqualTo := int64(-1)
//	clearPageOptions := ClearPagesOptions{
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberLessThanOrEqualTo: &ifSequenceNumberLessThanOrEqualTo,
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions) // This will cause the library to set the value of the header to 0
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeInvalidInput)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfSequenceNumberEqualTrue() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	sequenceNumber := int64(10)
//	actionType := SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	ifSequenceNumberEqualTo := int64(10)
//	clearPageOptions := ClearPagesOptions{
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberEqualTo: &ifSequenceNumberEqualTo,
//		},
//	}
//	_, err = pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.Nil(err)
//
//	validateClearPagesTest(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfSequenceNumberEqualFalse() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	sequenceNumber := int64(10)
//	actionType := SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	ifSequenceNumberEqualTo := int64(1)
//	clearPageOptions := ClearPagesOptions{
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberEqualTo: &ifSequenceNumberEqualTo,
//		},
//	}
//	_, err = pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeSequenceNumberConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobClearPagesIfSequenceNumberEqualNegOne() {
//	containerClient, pbClient := setupClearPagesTest(c)
//	defer deleteContainer(containerClient)
//
//	ifSequenceNumberEqualTo := int64(-1)
//	clearPageOptions := ClearPagesOptions{
//		SequenceNumberAccessConditions: &SequenceNumberAccessConditions{
//			IfSequenceNumberEqualTo: &ifSequenceNumberEqualTo,
//		},
//	}
//	_, err := pbClient.ClearPages(ctx, HttpRange{0, PageBlobPageBytes}, &clearPageOptions) // This will cause the library to set the value of the header to 0
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeInvalidInput)
//}
//
//func setupGetPageRangesTest() (containerClient ContainerClient, pbClient PageBlobClient) {
//	bsu := getServiceClient(nil)
//	containerClient, _ = createNewContainer(c, bsu)
//	pbClient, _ = createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	return
//}
//
//func validateBasicGetPageRanges(c *chk.C, resp *PageList, err error) {
//	_assert.Nil(err)
//	c.Assert(resp.PageRange, chk.NotNil)
//	c.Assert(*resp.PageRange, chk.HasLen, 1)
//	start, end := int64(0), int64(PageBlobPageBytes-1)
//	rawStart, rawEnd := (*resp.PageRange)[0].Raw()
//	c.Assert(rawStart, chk.Equals, start)
//	c.Assert(rawEnd, chk.Equals, end)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesEmptyBlob() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	resp, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, nil)
//	_assert.Nil(err)
//	c.Assert(resp.PageList.PageRange, chk.IsNil)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesEmptyRange() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	resp, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, nil)
//	_assert.Nil(err)
//	validateBasicGetPageRanges(c, resp.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesInvalidRange() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	_, err := pbClient.GetPageRanges(ctx, HttpRange{-2, 500}, nil)
//	_assert.Nil(err)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesNonContiguousRanges() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(2*PageBlobPageBytes), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	resp, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, nil)
//	_assert.Nil(err)
//	pageListResp := resp.PageList.PageRange
//	c.Assert(pageListResp, chk.NotNil)
//	c.Assert(*pageListResp, chk.HasLen, 2)
//
//	start, end := int64(0), int64(PageBlobPageBytes-1)
//	rawStart, rawEnd := (*pageListResp)[0].Raw()
//	c.Assert(rawStart, chk.Equals, start)
//	c.Assert(rawEnd, chk.Equals, end)
//
//	start, end = int64(PageBlobPageBytes*2), int64((PageBlobPageBytes*3)-1)
//	rawStart, rawEnd = (*pageListResp)[1].Raw()
//	c.Assert(rawStart, chk.Equals, start)
//	c.Assert(rawEnd, chk.Equals, end)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesNotPageAligned() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	resp, err := pbClient.GetPageRanges(ctx, HttpRange{0, 2000}, nil)
//	_assert.Nil(err)
//	validateBasicGetPageRanges(c, resp.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesSnapshot() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	resp, err := pbClient.CreateSnapshot(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(resp.Snapshot, chk.NotNil)
//
//	snapshotURL := pbClient.WithSnapshot(*resp.Snapshot)
//	resp2, err := snapshotURL.GetPageRanges(ctx, HttpRange{0, 0}, nil)
//	_assert.Nil(err)
//	validateBasicGetPageRanges(c, resp2.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesIfModifiedSinceTrue() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	resp, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, &getPageRangesOptions)
//	_assert.Nil(err)
//	validateBasicGetPageRanges(c, resp.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesIfModifiedSinceFalse() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, &getPageRangesOptions)
//	_assert.NotNil(err)
//
//	//serr := err.(StorageError)
//	//c.Assert(serr.RawResponse.StatusCode, chk.Equals, 304)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesIfUnmodifiedSinceTrue() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	resp, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, &getPageRangesOptions)
//	_assert.Nil(err)
//	validateBasicGetPageRanges(c, resp.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesIfUnmodifiedSinceFalse() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, &getPageRangesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesIfMatchTrue() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	resp, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: resp.ETag,
//			},
//		},
//	}
//	resp2, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, &getPageRangesOptions)
//	_assert.Nil(err)
//	validateBasicGetPageRanges(c, resp2.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesIfMatchFalse() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	eTag := "garbage"
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, &getPageRangesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesIfNoneMatchTrue() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	eTag := "garbage"
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: &eTag,
//			},
//		},
//	}
//	resp, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, &getPageRangesOptions)
//	_assert.Nil(err)
//	validateBasicGetPageRanges(c, resp.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobGetPageRangesIfNoneMatchFalse() {
//	containerClient, pbClient := setupGetPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.GetPageRanges(ctx, HttpRange{0, 0}, &getPageRangesOptions)
//	_assert.NotNil(err)
//	//serr := err.(StorageError)
//	//c.Assert(serr.RawResponse.StatusCode, chk.Equals, 304) // Service Code not returned in the body for a HEAD
//}
//
//func setupDiffPageRangesTest() (containerClient ContainerClient, pbClient PageBlobClient, snapshot string) {
//	bsu := getServiceClient(nil)
//	containerClient, _ = createNewContainer(c, bsu)
//	pbClient, _ = createNewPageBlob(c, containerClient)
//
//	r := getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count := int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//	}
//	_, err := pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//
//	resp, err := pbClient.CreateSnapshot(ctx, nil)
//	_assert.Nil(err)
//	snapshot = *resp.Snapshot
//
//	r = getReaderToRandomBytes(PageBlobPageBytes)
//	offset, count = int64(0), int64(PageBlobPageBytes)
//	uploadPagesOptions = UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//	}
//	_, err = pbClient.UploadPages(ctx, r, &uploadPagesOptions)
//	_assert.Nil(err)
//	return
//}
//
//func validateDiffPageRanges(c *chk.C, resp *PageList, err error) {
//	_assert.Nil(err)
//	pageListResp := resp.PageRange
//	c.Assert(pageListResp, chk.NotNil)
//	c.Assert(*resp.PageRange, chk.HasLen, 1)
//	start, end := int64(0), int64(PageBlobPageBytes-1)
//	rawStart, rawEnd := (*pageListResp)[0].Raw()
//	c.Assert(rawStart, chk.DeepEquals, start)
//	c.Assert(rawEnd, chk.DeepEquals, end)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangesNonExistentSnapshot() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	snapshotTime, _ := time.Parse(SnapshotTimeFormat, snapshot)
//	snapshotTime = snapshotTime.Add(time.Minute)
//	_, err := pbClient.GetPageRangesDiff(ctx, HttpRange{0, 0}, snapshotTime.Format(SnapshotTimeFormat), nil)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodePreviousSnapshotNotFound)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangeInvalidRange() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//	_, err := pbClient.GetPageRangesDiff(ctx, HttpRange{-22, 14}, snapshot, nil)
//	_assert.Nil(err)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangeIfModifiedSinceTrue() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	resp, err := pbClient.GetPageRangesDiff(ctx, HttpRange{0, 0}, snapshot, &getPageRangesOptions)
//	_assert.Nil(err)
//	validateDiffPageRanges(c, resp.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangeIfModifiedSinceFalse() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.GetPageRangesDiff(ctx, HttpRange{0, 0}, snapshot, &getPageRangesOptions)
//	_assert.NotNil(err)
//
//	//stgErr := err.(StorageError)
//	//c.Assert(stgErr.Response().StatusCode, chk.Equals, 304)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangeIfUnmodifiedSinceTrue() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	resp, err := pbClient.GetPageRangesDiff(ctx, HttpRange{0, 0}, snapshot, &getPageRangesOptions)
//	_assert.Nil(err)
//	validateDiffPageRanges(c, resp.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangeIfUnmodifiedSinceFalse() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.GetPageRangesDiff(ctx, HttpRange{0, 0}, snapshot, &getPageRangesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangeIfMatchTrue() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: resp.ETag,
//			},
//		},
//	}
//	resp2, err := pbClient.GetPageRangesDiff(ctx, HttpRange{0, 0}, snapshot, &getPageRangesOptions)
//	_assert.Nil(err)
//	validateDiffPageRanges(c, resp2.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangeIfMatchFalse() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	eTag := "garbage"
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.GetPageRangesDiff(ctx, HttpRange{0, 0}, snapshot, &getPageRangesOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangeIfNoneMatchTrue() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	eTag := "garbage"
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: &eTag,
//			},
//		},
//	}
//	resp, err := pbClient.GetPageRangesDiff(ctx, HttpRange{0, 0}, snapshot, &getPageRangesOptions)
//	_assert.Nil(err)
//	validateDiffPageRanges(c, resp.PageList, err)
//}
//
//func (s *azblobTestSuite) TestBlobDiffPageRangeIfNoneMatchFalse() {
//	containerClient, pbClient, snapshot := setupDiffPageRangesTest(c)
//	defer deleteContainer(containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	getPageRangesOptions := GetPageRangesOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.GetPageRangesDiff(ctx, HttpRange{0, 0}, snapshot, &getPageRangesOptions)
//	_assert.NotNil(err)
//
//	//serr := err.(StorageError)
//	//c.Assert(serr.Response().StatusCode, chk.Equals, 304)
//}
//
//func (s *azblobTestSuite) TestBlobResizeZero() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	// The default pbClient is created with size > 0, so this should actually update
//	_, err := pbClient.Resize(ctx, 0, nil)
//	_assert.Nil(err)
//
//	resp, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(*resp.ContentLength, chk.Equals, int64(0))
//}
//
//func (s *azblobTestSuite) TestBlobResizeInvalidSizeNegative() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	_, err := pbClient.Resize(ctx, -4, nil)
//	_assert.NotNil(err)
//}
//
//func (s *azblobTestSuite) TestBlobResizeInvalidSizeMisaligned() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	_, err := pbClient.Resize(ctx, 12, nil)
//	_assert.NotNil(err)
//}
//
//func validateResize(c *chk.C, pbClient PageBlobClient) {
//	resp, _ := pbClient.GetProperties(ctx, nil)
//	c.Assert(*resp.ContentLength, chk.Equals, int64(PageBlobPageBytes))
//}
//
//func (s *azblobTestSuite) TestBlobResizeIfModifiedSinceTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	resizePageBlobOptions := ResizePageBlobOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.Resize(ctx, PageBlobPageBytes, &resizePageBlobOptions)
//	_assert.Nil(err)
//
//	validateResize(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobResizeIfModifiedSinceFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	resizePageBlobOptions := ResizePageBlobOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.Resize(ctx, PageBlobPageBytes, &resizePageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobResizeIfUnmodifiedSinceTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	resizePageBlobOptions := ResizePageBlobOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.Resize(ctx, PageBlobPageBytes, &resizePageBlobOptions)
//	_assert.Nil(err)
//
//	validateResize(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobResizeIfUnmodifiedSinceFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	resizePageBlobOptions := ResizePageBlobOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.Resize(ctx, PageBlobPageBytes, &resizePageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobResizeIfMatchTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	resizePageBlobOptions := ResizePageBlobOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.Resize(ctx, PageBlobPageBytes, &resizePageBlobOptions)
//	_assert.Nil(err)
//
//	validateResize(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobResizeIfMatchFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	eTag := "garbage"
//	resizePageBlobOptions := ResizePageBlobOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.Resize(ctx, PageBlobPageBytes, &resizePageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobResizeIfNoneMatchTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	eTag := "garbage"
//	resizePageBlobOptions := ResizePageBlobOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.Resize(ctx, PageBlobPageBytes, &resizePageBlobOptions)
//	_assert.Nil(err)
//
//	validateResize(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobResizeIfNoneMatchFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	resizePageBlobOptions := ResizePageBlobOptions{
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.Resize(ctx, PageBlobPageBytes, &resizePageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberActionTypeInvalid() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	sequenceNumber := int64(1)
//	actionType := SequenceNumberAction("garbage")
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeInvalidHeaderValue)
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberSequenceNumberInvalid() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	defer func() { // Invalid sequence number should panic
//		recover()
//	}()
//
//	sequenceNumber := int64(-1)
//	actionType := SequenceNumberActionUpdate
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		BlobSequenceNumber: &sequenceNumber,
//		ActionType:         &actionType,
//	}
//
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeInvalidHeaderValue)
//}
//
//func validateSequenceNumberSet(c *chk.C, pbClient PageBlobClient) {
//	resp, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(*resp.BlobSequenceNumber, chk.Equals, int64(1))
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberIfModifiedSinceTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	actionType := SequenceNumberActionIncrement
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		ActionType: &actionType,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	validateSequenceNumberSet(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberIfModifiedSinceFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	actionType := SequenceNumberActionIncrement
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		ActionType: &actionType,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfModifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberIfUnmodifiedSinceTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(10)
//
//	actionType := SequenceNumberActionIncrement
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		ActionType: &actionType,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	validateSequenceNumberSet(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberIfUnmodifiedSinceFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	currentTime := getRelativeTimeGMT(-10)
//
//	actionType := SequenceNumberActionIncrement
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		ActionType: &actionType,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfUnmodifiedSince: &currentTime,
//			},
//		},
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberIfMatchTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	actionType := SequenceNumberActionIncrement
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		ActionType: &actionType,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	validateSequenceNumberSet(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberIfMatchFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	eTag := "garbage"
//	actionType := SequenceNumberActionIncrement
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		ActionType: &actionType,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberIfNoneMatchTrue() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	eTag := "garbage"
//	actionType := SequenceNumberActionIncrement
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		ActionType: &actionType,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: &eTag,
//			},
//		},
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.Nil(err)
//
//	validateSequenceNumberSet(c, pbClient)
//}
//
//func (s *azblobTestSuite) TestBlobSetSequenceNumberIfNoneMatchFalse() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	resp, _ := pbClient.GetProperties(ctx, nil)
//
//	actionType := SequenceNumberActionIncrement
//	updateSequenceNumberPageBlob := UpdateSequenceNumberPageBlob{
//		ActionType: &actionType,
//		BlobAccessConditions: BlobAccessConditions{
//			ModifiedAccessConditions: &ModifiedAccessConditions{
//				IfNoneMatch: resp.ETag,
//			},
//		},
//	}
//	_, err := pbClient.UpdateSequenceNumber(ctx, &updateSequenceNumberPageBlob)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func setupStartIncrementalCopyTest() (containerClient ContainerClient, pbClient PageBlobClient, copyPBClient PageBlobClient, snapshot string) {
//	bsu := getServiceClient(nil)
//	containerClient, _ = createNewContainer(c, bsu)
//
//	accessType := PublicAccessBlob
//	setAccessPolicyOptions := SetAccessPolicyOptions{
//		ContainerSetAccessPolicyOptions: ContainerSetAccessPolicyOptions{Access: &accessType},
//	}
//	_, err := containerClient.SetAccessPolicy(context.Background(), &setAccessPolicyOptions)
//	_assert.Nil(err)
//
//	pbClient, _ = createNewPageBlob(c, containerClient)
//	resp, _ := pbClient.CreateSnapshot(ctx, nil)
//	copyPBClient, _ = getPageBlobClient(c, containerClient)
//
//	// Must create the incremental copy pbClient so that the access conditions work on it
//	resp2, err := copyPBClient.StartCopyIncremental(ctx, pbClient.URL(), *resp.Snapshot, nil)
//	_assert.Nil(err)
//	waitForIncrementalCopy(c, copyPBClient, &resp2)
//
//	resp, _ = pbClient.CreateSnapshot(ctx, nil) // Take a new snapshot so the next copy will succeed
//	snapshot = *resp.Snapshot
//	return
//}
//
//func validateIncrementalCopy(c *chk.C, copyBlobURL PageBlobClient, resp *PageBlobCopyIncrementalResponse) {
//	t := waitForIncrementalCopy(c, copyBlobURL, resp)
//
//	// If we can access the snapshot without error, we are satisfied that it was created as a result of the copy
//	copySnapshotURL := copyBlobURL.WithSnapshot(*t)
//	_, err := copySnapshotURL.GetProperties(ctx, nil)
//	_assert.Nil(err)
//}
//
//func (s *azblobTestSuite) TestBlobStartIncrementalCopySnapshotNotExist() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//	copyBlobURL, _ := getPageBlobClient(c, containerClient)
//
//	snapshot := time.Now().UTC().Format(SnapshotTimeFormat)
//	_, err := copyBlobURL.StartCopyIncremental(ctx, pbClient.URL(), snapshot, nil)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeCannotVerifyCopySource)
//}
//
//func (s *azblobTestSuite) TestBlobStartIncrementalCopyIfModifiedSinceTrue() {
//	containerClient, pbClient, copyBlobURL, snapshot := setupStartIncrementalCopyTest(c)
//
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(-20)
//
//	copyIncrementalPageBlobOptions := CopyIncrementalPageBlobOptions{
//		ModifiedAccessConditions: &ModifiedAccessConditions{
//			IfModifiedSince: &currentTime,
//		},
//	}
//	resp, err := copyBlobURL.StartCopyIncremental(ctx, pbClient.URL(), snapshot, &copyIncrementalPageBlobOptions)
//	_assert.Nil(err)
//
//	validateIncrementalCopy(c, copyBlobURL, &resp)
//}
//
//func (s *azblobTestSuite) TestBlobStartIncrementalCopyIfModifiedSinceFalse() {
//	containerClient, pbClient, copyBlobURL, snapshot := setupStartIncrementalCopyTest(c)
//
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(20)
//
//	copyIncrementalPageBlobOptions := CopyIncrementalPageBlobOptions{
//		ModifiedAccessConditions: &ModifiedAccessConditions{
//			IfModifiedSince: &currentTime,
//		},
//	}
//	_, err := copyBlobURL.StartCopyIncremental(ctx, pbClient.URL(), snapshot, &copyIncrementalPageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobStartIncrementalCopyIfUnmodifiedSinceTrue() {
//	containerClient, pbClient, copyBlobURL, snapshot := setupStartIncrementalCopyTest(c)
//
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(20)
//
//	copyIncrementalPageBlobOptions := CopyIncrementalPageBlobOptions{
//		ModifiedAccessConditions: &ModifiedAccessConditions{
//			IfUnmodifiedSince: &currentTime,
//		},
//	}
//	resp, err := copyBlobURL.StartCopyIncremental(ctx, pbClient.URL(), snapshot, &copyIncrementalPageBlobOptions)
//	_assert.Nil(err)
//
//	validateIncrementalCopy(c, copyBlobURL, &resp)
//}
//
//func (s *azblobTestSuite) TestBlobStartIncrementalCopyIfUnmodifiedSinceFalse() {
//	containerClient, pbClient, copyBlobURL, snapshot := setupStartIncrementalCopyTest(c)
//
//	defer deleteContainer(containerClient)
//
//	currentTime := getRelativeTimeGMT(-20)
//
//	copyIncrementalPageBlobOptions := CopyIncrementalPageBlobOptions{
//		ModifiedAccessConditions: &ModifiedAccessConditions{
//			IfUnmodifiedSince: &currentTime,
//		},
//	}
//	_, err := copyBlobURL.StartCopyIncremental(ctx, pbClient.URL(), snapshot, &copyIncrementalPageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobStartIncrementalCopyIfMatchTrue() {
//	containerClient, pbClient, copyBlobURL, snapshot := setupStartIncrementalCopyTest(c)
//
//	defer deleteContainer(containerClient)
//
//	resp, _ := copyBlobURL.GetProperties(ctx, nil)
//
//	copyIncrementalPageBlobOptions := CopyIncrementalPageBlobOptions{
//		ModifiedAccessConditions: &ModifiedAccessConditions{
//			IfMatch: resp.ETag,
//		},
//	}
//	resp2, err := copyBlobURL.StartCopyIncremental(ctx, pbClient.URL(), snapshot, &copyIncrementalPageBlobOptions)
//	_assert.Nil(err)
//
//	validateIncrementalCopy(c, copyBlobURL, &resp2)
//}
//
//func (s *azblobTestSuite) TestBlobStartIncrementalCopyIfMatchFalse() {
//	containerClient, pbClient, copyBlobURL, snapshot := setupStartIncrementalCopyTest(c)
//
//	defer deleteContainer(containerClient)
//
//	eTag := "garbage"
//	copyIncrementalPageBlobOptions := CopyIncrementalPageBlobOptions{
//		ModifiedAccessConditions: &ModifiedAccessConditions{
//			IfMatch: &eTag,
//		},
//	}
//	_, err := copyBlobURL.StartCopyIncremental(ctx, pbClient.URL(), snapshot, &copyIncrementalPageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeTargetConditionNotMet)
//}
//
//func (s *azblobTestSuite) TestBlobStartIncrementalCopyIfNoneMatchTrue() {
//	containerClient, pbClient, copyBlobURL, snapshot := setupStartIncrementalCopyTest(c)
//
//	defer deleteContainer(containerClient)
//
//	eTag := "garbage"
//	copyIncrementalPageBlobOptions := CopyIncrementalPageBlobOptions{
//		ModifiedAccessConditions: &ModifiedAccessConditions{
//			IfNoneMatch: &eTag,
//		},
//	}
//	resp, err := copyBlobURL.StartCopyIncremental(ctx, pbClient.URL(), snapshot, &copyIncrementalPageBlobOptions)
//	_assert.Nil(err)
//
//	validateIncrementalCopy(c, copyBlobURL, &resp)
//}
//
//func (s *azblobTestSuite) TestBlobStartIncrementalCopyIfNoneMatchFalse() {
//	containerClient, pbClient, copyBlobURL, snapshot := setupStartIncrementalCopyTest(c)
//
//	defer deleteContainer(containerClient)
//
//	resp, _ := copyBlobURL.GetProperties(ctx, nil)
//
//	copyIncrementalPageBlobOptions := CopyIncrementalPageBlobOptions{
//		ModifiedAccessConditions: &ModifiedAccessConditions{
//			IfNoneMatch: resp.ETag,
//		},
//	}
//	_, err := copyBlobURL.StartCopyIncremental(ctx, pbClient.URL(), snapshot, &copyIncrementalPageBlobOptions)
//	_assert.NotNil(err)
//
//	validateStorageError(c, err, StorageErrorCodeConditionNotMet)
//}
