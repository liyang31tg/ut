package mgo

import (
	"SunPlayer/src/comm/logger"
	"errors"
	"sync"
	"time"

	"log"

	"gopkg.in/mgo.v2"
)

type client struct {
	url     string
	session *mgo.Session
	rwLock  sync.RWMutex
	db      string
}

func NewClient(dbName, uri string) *client {
	c := new(client)
	c.db = dbName
	c.url = uri
	c.dial()
	go c.keepAlive()
	return c
}

func (this *client) Coll(name string) (*mgo.Collection, error) {
	s := this.getSession()
	if s == nil {
		return nil, errors.New("net is bad")
	}
	return s.DB(this.db).C(name), nil
}

func (this *client) setSession(s *mgo.Session) {
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	this.session = s
}

func (this *client) getSession() *mgo.Session {
	this.rwLock.RLock()
	defer this.rwLock.RUnlock()
	return this.session
}

func (this *client) dial() {
	s, err := mgo.Dial(this.url)
	if err != nil {
		log.Fatal(err)
	} else {
		this.setSession(s)
	}
}

func (this *client) keepAlive() {
	for {
		s := this.getSession()
		if s == nil {
			this.dial()
		} else {
			err := s.Ping()
			if err != nil {
				logger.Err.Println("err:", err)
				this.setSession(nil)
				s.Close()
			}
		}
		time.Sleep(2 * time.Second)
	}

}
