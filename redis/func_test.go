package redis

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

/*
保障插入顺序不间断，c.GetClient()阻塞式的
*/
func Test_Hanle(t *testing.T) {
	c := NewClient(&Options{
		Addr:     "localhost:6380",
		Password: "QuGuoChengAndQuLing1",
		DB:       0,
	})
	time.Sleep(1e9)
	//第一种情况，有可能失败
	{
		go func() {
			var i = 0
			for {
				<-time.NewTicker(1e9).C
				i++
				fmt.Println("handle start :", i)
				re, err := c.GetClient().HSet("table", "k"+strconv.Itoa(i), i).Result()
				fmt.Println("handle end :", i, re, err)
			}
		}()
	}
	//100%保证插入顺序，不丢失
	{
		var i = 0
		for {
			<-time.NewTicker(1e9).C
			i++
			fmt.Println("handle start1 :", i)
			c.HandleSuccess(func(cc *redis.Client) bool {
				re, _ := cc.HSet("table1", "k"+strconv.Itoa(i), i).Result()
				return re
			})
			fmt.Println("handle end1 :", i)
		}
	}
}
