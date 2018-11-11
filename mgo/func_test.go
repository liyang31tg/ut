package mgo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestFunc(t *testing.T) {
	var DB = NewClient("SunPlayerDB", "mongodb://localhost:27017")
	c, _ := DB.Coll("cc")
	c.Insert(bson.M{"name": "LiYang"})
}
