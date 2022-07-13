// Design and implement a code base that can upload and retrieve files to/from IPFS.

package main

import (
	"fmt"
	"net/http"
	"context"

	ipfsClient "github.com/ipfs/go-ipfs-http-client"
	orbitdb "berty.tech/go-orbit-db"
)

type SongDocument struct {
	title string
	artist string
	album string
	cid string
}

func CreateSongDocument(title string, artist string, album string, cid string) *SongDocument {
	songDoc := SongDocument{title, artist, album, cid}
	return &songDoc
}

func CreateIPFSNode() (*ipfsClient.HttpApi, error) {
	// Create/connect to an IPFS node
	// TODO: Start a ipfs node in a seperate process: (ipfs daemon --enable-pubsub-experiment)
	client, err := ipfsClient.NewURLApiWithClient("localhost:5001", &http.Client{}) // uses client package
	if err != nil {
		return nil, err
	}
	return client, nil

}

func CreateDBInstance(ctx context.Context, client *ipfsClient.HttpApi) (orbitdb.OrbitDB, error) {
	// Create an instance of orbitdb
	db, err :=	orbitdb.NewOrbitDB(ctx, client, nil)
	if err != nil {
		fmt.Printf("Failed to create orbitdb instance: %s\n", err)
		return nil, err
	}
	fmt.Printf("%T\n", db)

	return db, nil
}

// func ConnectToStore() error {
// 	// Connect to Kawa orbit store
// }

func main() {
	fmt.Println("Starting Kawa v0.1")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Step 1: connect to IPFS
	client, err := CreateIPFSNode()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := CreateDBInstance(ctx, client)
	fmt.Println(db)

}