package redis

import (
	"sync"
	"time"

	"github.com/liyang31tg/ut/logger"

	"github.com/go-redis/redis"
)

type Client struct {
	client *redis.Client
	mtx    sync.RWMutex
	opt    *Options
}

type Options redis.Options

const tick = 2 * time.Second

func NewClient(opt *Options) *Client {
	c := &Client{opt: opt}
	go c.keepAlive()
	return c
}

/*
*保证99%的情况下都能拿到有效的client，这里是阻塞式拿
 */
func (this *Client) GetClient() *redis.Client {
	var c *redis.Client
	for {
		c = this.getClient()
		if c != nil {
			break
		}
		time.Sleep(tick)
	}
	return c
}

func (this *Client) HandleSuccess(f func() bool) {
	for {
		s := f()
		if s {
			return
		}
		time.Sleep(tick)
	}
}

func (this *Client) getClient() *redis.Client {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	return this.client
}

func (this *Client) setClient(c *redis.Client) {
	this.mtx.Lock()
	this.mtx.Unlock()
	this.client = c
}

func (this *Client) getOpt() *redis.Options {
	if this.opt == nil {
		return &redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}
	} else {
		return (*redis.Options)(this.opt)
	}

}

func (this *Client) keepAlive() {
	for {
		c := this.getClient()
		if c == nil {
			tmpC := redis.NewClient(this.getOpt())
			logger.Info.Println("链接redis中...:", this.getOpt())
			//心跳拨号为2秒，那么链接就是1秒内5次如果没连上就代表失败
			for i := 0; i < 5; i++ {
				_, err := tmpC.Ping().Result()
				if err == nil {
					this.setClient(tmpC)
				}
				time.Sleep(tick / 10)
			}

		} else {
			_, err := c.Ping().Result()
			if err != nil {
				logger.Err.Println("redis  :", err)
				c.Close()
				this.setClient(nil)
			}
		}
		time.Sleep(tick)
	}
}
