package pkg

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
)

/**
创建桶列表
*/
func CreateBucket(bucket string) (bool, error) {
	var appid = "1257886963"
	var region = "ap-guangzhou"
	var bucket_name = bucket + "-" + appid
	buckets_list := BucketsList()
	if len(buckets_list) > 0 {
		for _, v := range buckets_list {
			if v == bucket_name {
				return true, nil
			}
		}
	}
	urls := fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket_name, region)
	u, _ := url.Parse(urls)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDSkqjPs9Azb2GH5u7JI4e1B7sQltn5yFJ",
			SecretKey: "0nNmyxjIemhhKqAHILJI8lBNPDoERSGr",
		},
	})

	_, err := c.Bucket.Put(context.Background(), nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
获取桶列表
*/
func BucketsList() []string {
	c := cos.NewClient(nil, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDSkqjPs9Azb2GH5u7JI4e1B7sQltn5yFJ",
			SecretKey: "0nNmyxjIemhhKqAHILJI8lBNPDoERSGr",
		},
	})

	s, _, err := c.Service.Get(context.Background())
	if err != nil {
		panic(err)
	}
	var buckets []string
	for _, b := range s.Buckets {
		buckets = append(buckets, b.Name)
	}
	return buckets
}

/**
上传文件
*/
func UploadFile(filename string) {

	var appid = "1257886963"
	var region = "ap-guangzhou"
	var bucket_name = "huigoumall" + "-" + appid

	urls := fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket_name, region)
	u, _ := url.Parse(urls)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDSkqjPs9Azb2GH5u7JI4e1B7sQltn5yFJ",
			SecretKey: "0nNmyxjIemhhKqAHILJI8lBNPDoERSGr",
		},
	})

	key := "exampleobject66.jpg"
	f, err := os.Open(filename)
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "image/jpeg",
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			// 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
			XCosACL: "default",
		},
	}
	_, err = c.Object.Put(context.Background(), key, f, opt)
	if err != nil {
		panic(err)
	}

}
