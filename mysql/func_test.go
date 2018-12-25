package mysql

import (
	"fmt"
	"testing"
	"time"
	"github.com/liyang31tg/ut"
)

func TestFunc(t *testing.T) {
	Register(&Options{
		User: "root",
		Pwd:  "QuGuoChengAndQuLing",
		Uri:  "127.0.0.1:3307",
		DB:   "bfbdb",
	})
	for i := 0; i < 1000; i++ {
		result, err := Client.Exec("insert into user (name,age) values (?,?)", "liYang"+ut.ToString(i), i)
		fmt.Println(result, err) // 错误自行处理不保证没错误
		time.Sleep(1e9)
	}

}
