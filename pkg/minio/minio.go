// Copyright The TBox Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package minio

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/s3utils"
	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/config/sections"
)

var errorNotConnected = fmt.Errorf("minio is not connected")

// MinIO specifies MinIO/S3 access structure
type MinIO struct {
	Endpoint         string
	Secure           bool
	AccessKeyID      string
	SecretAccessKey  string
	BucketAutoCreate bool
	Region           string
	Perm             os.FileMode

	client *minio.Client
}

// NewMinIO creates new MinIO
func NewMinIO(
	endpoint string,
	secure bool,
	insecureSkipVerify bool,
	accessKeyID,
	secretAccessKey string,
	bucketAutoCreate bool,
	region string,
) (*MinIO, error) {
	var err error

	if region == "" {
		region = "us-east-1"
	}

	min := &MinIO{
		Endpoint:         endpoint,
		Secure:           secure,
		AccessKeyID:      accessKeyID,
		SecretAccessKey:  secretAccessKey,
		BucketAutoCreate: bucketAutoCreate,
		Region:           region,
		Perm:             0666,
	}
	opts := &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: secure,
		Region: region,
	}
	if secure && insecureSkipVerify {
		// All this dance is for TLSClientConfig - set InsecureSkipVerify: true
		opts.Transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
	min.client, err = minio.New(endpoint, opts)
	if err != nil {
		log.Errorf("ERROR call minio.New() %v", err.Error())
	}

	return min, err
}

// NewMinIOFromConfig creates new MinIO from config
func NewMinIOFromConfig(cfg sections.MinIOConfigurator) (*MinIO, error) {
	return NewMinIO(
		cfg.GetMinIOEndpoint(),
		cfg.GetMinIOSecure(),
		cfg.GetMinIOInsecureSkipVerify(),
		cfg.GetMinIOAccessKeyID(),
		cfg.GetMinIOSecretAccessKey(),
		cfg.GetMinIOBucketAutoCreate(),
		cfg.GetMinIORegion(),
	)
}

// Put creates specified object from the reader
func (m *MinIO) Put(bucketName, objectName string, reader io.Reader) (int64, error) {
	if m.client == nil {
		return 0, errorNotConnected
	}

	if m.BucketAutoCreate {
		if err := m.CreateBucket(bucketName); err != nil {
			return 0, err
		}
	}

	ctx := context.Background()
	// Specify -1 in case object size in unknown in advance
	size := int64(-1)
	options := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	}

	info, err := m.client.PutObject(ctx, bucketName, objectName, reader, size, options)
	return info.Size, err
}

// PutA creates specified object from the reader
func (m *MinIO) PutA(addr *common.S3Address, reader io.Reader) (int64, error) {
	return m.Put(addr.Bucket, addr.Object, reader)
}

// PutUUID creates object in specified bucket named by generated UUID from the reader
func (m *MinIO) PutUUID(bucketName string, path string, reader io.Reader) (*common.S3Address, int64, error) {
	return m.PutUUIDA(common.NewS3Address(bucketName, path), reader)
}

// PutUUIDA creates object in specified bucket named by generated UUID from the reader
func (m *MinIO) PutUUIDA(addr *common.S3Address, reader io.Reader) (*common.S3Address, int64, error) {
	target := &common.S3Address{
		Bucket: addr.Bucket,
		Object: PathJoin(addr.Object, common.NewUuidRandom().String()),
	}
	n, err := m.PutA(target, reader)
	return target, n, err
}

// FPut creates specified object from the file
func (m *MinIO) FPut(bucketName, objectName, fileName string) (int64, error) {
	if m.client == nil {
		return 0, errorNotConnected
	}

	if m.BucketAutoCreate {
		if err := m.CreateBucket(bucketName); err != nil {
			return 0, err
		}
	}

	ctx := context.Background()
	options := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	}

	info, err := m.client.FPutObject(ctx, bucketName, objectName, fileName, options)
	return info.Size, err
}

// FPutA creates specified object from the file
func (m *MinIO) FPutA(addr *common.S3Address, fileName string) (int64, error) {
	return m.FPut(addr.Bucket, addr.Object, fileName)
}

// FPutUUID creates object in specified bucket named by generated UUID from the file
func (m *MinIO) FPutUUID(bucketName, path, fileName string) (*common.S3Address, int64, error) {
	return m.FPutUUIDA(common.NewS3Address(bucketName, path), fileName)
}

// FPutUUIDA creates object in specified bucket named by generated UUID from the file
func (m *MinIO) FPutUUIDA(addr *common.S3Address, fileName string) (*common.S3Address, int64, error) {
	target := &common.S3Address{
		Bucket: addr.Bucket,
		Object: PathJoin(addr.Object, common.NewUuidRandom().String()),
	}
	n, err := m.FPutA(target, fileName)
	return target, n, err
}

// Get returns reader for specified object
func (m *MinIO) Get(bucketName, objectName string) (io.Reader, error) {
	if m.client == nil {
		return nil, errorNotConnected
	}

	ctx := context.Background()
	opts := minio.GetObjectOptions{}
	//opts.SetModified(time.Now().Round(10 * time.Minute)) // get object if was modified within the last 10 minutes
	return m.client.GetObject(ctx, bucketName, objectName, opts)
}

// GetA returns reader for specified object
func (m *MinIO) GetA(addr *common.S3Address) (io.Reader, error) {
	return m.Get(addr.Bucket, addr.Object)
}

// BufferGet downloads specified object into memory buffer
func (m *MinIO) BufferGet(bucketName, objectName string) (*bytes.Buffer, error) {
	// Obtain reader
	r, err := m.Get(bucketName, objectName)
	if err != nil {
		return nil, err
	}
	// Download data into mem
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, r)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// BufferGetA downloads specified object into memory buffer
func (m *MinIO) BufferGetA(addr *common.S3Address) (*bytes.Buffer, error) {
	return m.BufferGet(addr.Bucket, addr.Object)
}

// FGet downloads specified object into specified file
func (m *MinIO) FGet(bucketName, objectName, fileName string) error {
	if m.client == nil {
		return errorNotConnected
	}

	ctx := context.Background()
	opts := minio.GetObjectOptions{}
	//opts.SetModified(time.Now().Round(10 * time.Minute)) // get object if was modified within the last 10 minutes
	err := m.client.FGetObject(ctx, bucketName, objectName, fileName, opts)
	if err != nil {
		return err
	}

	stats, err := os.Stat(fileName)
	if err != nil {
		return err
	}

	if stats.Mode() == m.Perm {
		// Already has required perms
		return nil
	}

	return os.Chmod(fileName, m.Perm)
}

// FGetA downloads specified object into specified file
func (m *MinIO) FGetA(addr *common.S3Address, fileName string) error {
	return m.FGet(addr.Bucket, addr.Object, fileName)
}

// FGetTempFile downloads specified object into temp file created in specified dir with name pattern
// See ioutil.TempFile() for more into about dir and name pattern
func (m *MinIO) FGetTempFile(bucketName, objectName, dir, pattern string) (string, error) {
	r, err := m.Get(bucketName, objectName)
	if err != nil {
		log.Errorf("unable to get MinIO object %s/%s err: %v", bucketName, objectName, err)
		return "", err
	}

	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		log.Errorf("unable to create tmp file dir: %s pattern: %s err: %v", dir, pattern, err)
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		log.Errorf("unable to get copy MinIO object %s/%s err: %v", bucketName, objectName, err)
		os.Remove(f.Name())
		return "", err
	}

	return f.Name(), nil
}

// FGetTempFileA downloads specified object into temp file created in specified dir with name pattern
// See ioutil.TempFile() for more into about dir and name pattern
func (m *MinIO) FGetTempFileA(addr *common.S3Address, dir, pattern string) (string, error) {
	return m.FGetTempFile(addr.Bucket, addr.Object, dir, pattern)
}

// Remove removes/deletes specified object
func (m *MinIO) Remove(bucketName, objectName string) error {
	if m.client == nil {
		return errorNotConnected
	}

	ctx := context.Background()
	return m.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})
}

// RemoveA removes/deletes specified object
func (m *MinIO) RemoveA(addr *common.S3Address) error {
	return m.Remove(addr.Bucket, addr.Object)
}

// Copy copies specified src object into specified dst object
func (m *MinIO) Copy(dstBucketName, dstObjectName, srcBucketName, srcObjectName string) error {
	if m.client == nil {
		return errorNotConnected
	}

	if m.BucketAutoCreate {
		if err := m.CreateBucket(dstBucketName); err != nil {
			return err
		}
	}

	src := minio.CopySrcOptions{
		Bucket: srcBucketName,
		Object: srcObjectName,
	}
	dst := minio.CopyDestOptions{
		Bucket: dstBucketName,
		Object: dstObjectName,
	}

	_, err := m.client.CopyObject(context.Background(), dst, src)
	return err
}

// CopyA copies specified src object into specified dst object
func (m *MinIO) CopyA(dst, src *common.S3Address) error {
	return m.Copy(dst.Bucket, dst.Object, src.Bucket, src.Object)
}

// Move moves specified src object into specified dst object
func (m *MinIO) Move(dstBucketName, dstObjectName, srcBucketName, srcObjectName string) error {
	if err := m.Copy(dstBucketName, dstObjectName, srcBucketName, srcObjectName); err != nil {
		return err
	}
	return m.Remove(srcBucketName, srcObjectName)
}

// MoveA moves specified src object into specified dst object
func (m *MinIO) MoveA(dst, src *common.S3Address) error {
	return m.Move(dst.Bucket, dst.Object, src.Bucket, src.Object)
}

// Digest calculates digest of the specified object
func (m *MinIO) Digest(bucketName, objectName string, _type common.DigestType) (*common.Digest, error) {
	reader, err := m.Get(bucketName, objectName)
	if err != nil {
		return nil, err
	}

	switch _type {
	case common.DigestType_DIGEST_SHA256:
		break
	default:
		return nil, fmt.Errorf("unable to calc digest - unknown digest type %v", _type)
	}

	h := sha256.New()
	_, err = io.Copy(h, reader)
	if err != nil {
		return nil, err
	}
	digest := h.Sum(nil)

	res := &common.Digest{}
	res.Type = _type
	res.Data = digest

	return res, nil
}

// DigestA calculates digest of the specified object
func (m *MinIO) DigestA(addr *common.S3Address, _type common.DigestType) (*common.Digest, error) {
	return m.Digest(addr.Bucket, addr.Object, _type)
}

// ListBuckets lists all buckets
func (m *MinIO) ListBuckets() ([]string, error) {
	if m.client == nil {
		return nil, errorNotConnected
	}

	buckets, err := m.client.ListBuckets(context.Background())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var res []string
	for _, bucket := range buckets {
		res = append(res, bucket.Name)
	}

	return res, nil
}

// CreateBucket creates specified bucket
func (m *MinIO) CreateBucket(bucketName string) error {
	if m.client == nil {
		return errorNotConnected
	}

	opts := minio.MakeBucketOptions{
		Region: m.Region,
	}

	if exists, err := m.client.BucketExists(context.Background(), bucketName); (exists == true) && (err == nil) {
		// Bucket exists and you have permission to access it
		return nil
	}

	return m.client.MakeBucket(context.Background(), bucketName, opts)
}

// RemoveBucket removes/deletes specified bucket
func (m *MinIO) RemoveBucket(bucketName string) error {
	if m.client == nil {
		return errorNotConnected
	}
	return m.client.RemoveBucket(context.Background(), bucketName)
}

// List at max n objects from a bucket having specified name prefix.
func (m *MinIO) List(bucket, prefix string, n int) ([]minio.ObjectInfo, error) {
	var res []minio.ObjectInfo

	i := 0

	opts := minio.ListObjectsOptions{
		Recursive: true,
		Prefix:    prefix,
	}
	for object := range m.client.ListObjects(context.Background(), bucket, opts) {
		if object.Err != nil {
			continue
		}
		res = append(res, object)
		i++
	}

	// Limit result size
	if (len(res) > 0) && (n >= 0) {
		res = res[0:n]
	}

	return res, nil
}

// PathJoin joins object path components
func PathJoin(elem ...string) string {
	return filepath.ToSlash(filepath.Join(elem...))
}

// GetBucketNameError returns an error with description
// in case bucket name is not valid according to AWS s3 bucket naming rules.
// https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html
func GetBucketNameError(bucket string) error {
	return s3utils.CheckValidBucketNameStrict(bucket)
}

// IsBucketNameValid checks whether bucket name is valid according to AWS s3 bucket naming rules.
// https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html
func IsBucketNameValid(bucket string) bool {
	return GetBucketNameError(bucket) == nil
}

// GetObjectNameError returns an error with description
// in case object name is not valid according to AWS s2 object naming rules.
func GetObjectNameError(object string) error {
	return s3utils.CheckValidObjectName(object)
}

// IsObjectNameValid checks whether object name is valid according to AWS s2 object naming rules.
func IsObjectNameValid(object string) bool {
	return GetObjectNameError(object) == nil
}
