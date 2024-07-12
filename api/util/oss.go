package util

import (
	"api_go/config"
	"sync"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	client *oss.Client
	bucket *oss.Bucket
	once   sync.Once
)

func InitOss() error {
	var err error
	once.Do(func() {
		client, err = oss.New(config.OSS.Endpoint, config.OSS.AccessKeyID, config.OSS.AccessKeySecret)
		if err != nil {
			return
		}

		bucket, err = client.Bucket(config.OSS.BucketName)
	})
	return err
}

func GetBucket() (*oss.Bucket, error) {
	if bucket == nil {
		if err := InitOss(); err != nil {
			return nil, err
		}
	}
	return bucket, nil
}
