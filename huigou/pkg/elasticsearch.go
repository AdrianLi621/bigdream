package pkg

import (
	"context"
	"sync"

	"github.com/olivere/elastic/v7"
)
type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}
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
func NewES() {
	if client == nil {
		once.Do(func() {
			client, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
			if err != nil {
				panic(err)
			}
		})
	}
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
	//n := 0
	//bulkRequest := client.Bulk()
	//for j := 0; j < 10000; j++ {
	//	n++
	//	tweet := Tweet{
	//		User: "olivere",
	//		Message: "Package strconv implements conversions to and from string representations of basic data types. " + strconv.Itoa(n),
	//	}
	//	req := elastic.NewBulkIndexRequest().Index("twitter").Type("tweet").Id(strconv.Itoa(n)).Doc(tweet)
	//	bulkRequest = bulkRequest.Add(req)
	//}
	//bulkResponse, err := bulkRequest.Do(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if bulkResponse != nil {
	//
	//}
	//return true,nil
}
/**
查询文档
 */
func SelectDoc(index string,key string,value string)([]*elastic.SearchHit,error)  {
	termQuery := elastic.NewTermQuery(key, value)
	searchResult, err := client.Search().
		Index(index).   // search in index "twitter"
		Query(termQuery).   // specify the query
		From(0).Size(100).   // take documents 0-9
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






















