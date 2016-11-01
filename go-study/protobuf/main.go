package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
)

var dbName = []byte("tiger")
var key = []byte("test")

func PutMessages(db *bolt.DB, msg *TestMessage) error {
	// store data
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(dbName)
		if err != nil {
			return err
		}

		p, err := proto.Marshal(msg)
		if err != nil {
			return err
		}
		return bucket.Put(key, p)
	})
}

func GetMessage(db *bolt.DB) TestMessage {
	var msg TestMessage
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(dbName)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", dbName)
		}

		val := bucket.Get(key)
		return proto.Unmarshal(val, &msg)
	})

	if err != nil {
		log.Println("read error: ", err)
		return msg
	}

	return msg
}

func main() {
	items := []TestMessage_MsgItem{
		TestMessage_MsgItem{
			Id:        1,
			ItemName:  "aaaaa",
			ItemValue: 333,
			ItemType:  111,
		},
	}
	msg := TestMessage{
		ClientName:   "upccup",
		Description:  "test proto",
		ClientId:     1,
		Messageitems: items,
	}

	db, err := bolt.Open("bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := PutMessages(db, &msg); err != nil {
		log.Fatal(err)
	}

	queryMsg := GetMessage(db)
	log.Printf("query msg: %+v", queryMsg)
	log.Printf("put msg: %+v", msg)
}
