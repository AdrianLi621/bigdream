package pkg

import (
	"context"
	"sync"

	"github.com/olivere/elastic/v7"
)
var client *elastic.Client

var err error

var once sync.Once

type Student struct {
	Name string
	Age int
}
/**
构造函数
 */
func init()  {
	NewES()
}

/**
实例化
 */
func NewES()*elastic.Client {
	if client == nil {
		once.Do(func() {
			client, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
			if err != nil {
				panic(err)
			}
		})
	}
	return client
}
/**
检查数据库是否存在
 */
func IndexExist(index string)(bool,error)  {
	return client.IndexExists(index).Do(context.Background())
}
/**
创建数据库
 */
func CreateIndex(index string)(bool,error)  {
	_, err := client.CreateIndex(index).Do(context.Background())
	if err != nil {
		return false,err
	}
	return true,nil
}
/**
插入数据
 */
func InsertDoc(index string,data string)(bool,error){
	_, err = client.Index().
		Index(index).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		return false,err
	}
	return true,nil
}
/**
查询文档
 */
func SelectDoc(index string)([]*elastic.SearchHit,error)  {
	termQuery := elastic.NewTermQuery("goods_name", "小")
	searchResult, err := client.Search().
		Index(index).   // search in index "twitter"
		Query(termQuery).   // specify the query
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(context.Background())             // execute
	if err != nil {
		return nil,err
	}
	if searchResult.TotalHits()>0 {
		return searchResult.Hits.Hits,nil
	}
	return searchResult.Hits.Hits,nil
}






















