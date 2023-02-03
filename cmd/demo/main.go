package main

import (
	"log"

	"github.com/slatp/bolt"
)

func read() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin(false)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte("key1"))
	v := b.Get([]byte("hello"))
	log.Println(string(v))
}

func write() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin(true)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	b, err := tx.CreateBucketIfNotExists([]byte("key1"))
	if err != nil {
		panic(err)
	}

	if err := b.Put([]byte("hello"), []byte("world")); err != nil {
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

func main() {
	write()
}
