package main

import (
	"fmt"
	"context"
	"os"
	"os/exec"
	"io"
	// "time"
	// "net/http"

	orbitdb "berty.tech/go-orbit-db"
	// shell "github.com/ipfs/go-ipfs-api"
	ipfsClient "github.com/ipfs/go-ipfs-http-client"
)

type SongDocument struct {
	title string
	artist string
	album string
	cid string
}

func CreateSongDocument(title string, artist string, album string, cid string) *SongDocument {
	song := SongDocument{title, artist, album, cid}
	return &song
}

func CreateIPFSNode() error {
	// Create/connect to an IPFS node
	// TODO: Start a ipfs node in a seperate process: (ipfs daemon --enable-pubsub-experiment)

	return nil

}

func CreateDBInstance() error {
	// Create an instance of orbitdb

}

func ConnectToStore() error {
	// Connect to Kawa orbit store
}



// func CreateDBInstance() error {
// 	// Create an instance of orbitdb
// }

// func ConnectToStore() error {
// 	// Connect to Kawa orbit store
// }

// func createDB() (orbitdb.KeyValueStore) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	c, err := client.NewURLApiWithClient("localhost:5001", &http.Client{})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}

// 	db, err := orbitdb.NewOrbitDB(ctx, c, &orbitdb.NewOrbitDBOptions{})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}
	
// 	dbStore, err := db.Create(ctx, "test", "keyvalue", &orbitdb.CreateDBOptions{})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}

// 	KvStore, err := db.KeyValue(ctx, dbStore.Address().String(), &orbitdb.CreateDBOptions{})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}

// 	return KvStore
// }

func connectDB(url string, client *ipfsClient.HttpApi) (orbitdb.KeyValueStore) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := orbitdb.NewOrbitDB(ctx, client, &orbitdb.NewOrbitDBOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	KvStore, err := db.KeyValue(ctx, url, &orbitdb.CreateDBOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	return KvStore
}

func main() {
	fmt.Println("hello world")
	CreateIPFSNode()
	// kvStore := createDB()
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// kvStore.Put(ctx, "message", []byte("test1"))

	// v, err := kvStore.Get(ctx, "message")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(v))
}

