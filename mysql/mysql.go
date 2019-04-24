package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/liyang31tg/ut/logger"

	_ "github.com/go-sql-driver/mysql"
)

var Client *sql.DB
var defaultopt *Options

type Options struct {
	User string
	Pwd  string
	Uri  string
	DB   string
}

func Register(opt *Options) *sql.DB {
	client, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", opt.User, opt.Pwd, opt.Uri, opt.DB))
	if err != nil {
		logger.Err.Println("请先把mysql给启动起来：", err)
		log.Fatal(err)
	} else {
		Client = client
		defaultopt = opt
	}
	go keepAlive()
	return client
}
func keepAlive() {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s", defaultopt.User, defaultopt.Pwd, defaultopt.Uri, defaultopt.DB)
	for {
		time.Sleep(2e9)
		err := Client.Ping()
		if err != nil {
			client, err := sql.Open("mysql", url)
			if err != nil {
				logger.Err.Println(err)
			} else {
				Client = client
			}
		}
	}
}

func Result(query *sql.Rows) []map[string]string {
	column, _ := query.Columns()              //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make([]map[string]string, 0) //最后得到的map

	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			logger.Err.Println(err)
			return results
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	return results
}
