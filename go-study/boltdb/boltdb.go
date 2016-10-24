package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var dbName = []byte("tiger")

func main() {
	db, err := bolt.Open("bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}

	key := []byte("test")
	value := []byte("upccup")

	// store data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(dbName)
		if err != nil {
			return err
		}

		return bucket.Put(key, value)
	})

	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(dbName)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", dbName)
		}

		val := bucket.Get(key)
		fmt.Println(string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
