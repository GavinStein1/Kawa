// Design and implement a code base that can upload and retrieve files to/from IPFS.

package main

import (
	"fmt"
	"net/http"
	"context"
	"os"
	"encoding/json"
	"io/ioutil"
	"crypto/sha256"

	ipfsClient "github.com/ipfs/go-ipfs-http-client"
	ipfslog "berty.tech/go-ipfs-log"
	orbitdb "berty.tech/go-orbit-db"
	accesscontroller "berty.tech/go-orbit-db/accesscontroller"
	"berty.tech/go-orbit-db/iface"
)

type Config struct {
	StoreString string
}

type SongDocument struct {
	Title string
	Artist string
	Album string
	Cid string
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

func CreateDBInstance(ctx context.Context, client *ipfsClient.HttpApi) (*orbitdb.OrbitDB, error) {
	// Create an instance of orbitdb
	db, err :=	orbitdb.NewOrbitDB(ctx, client, nil)
	if err != nil {
		fmt.Printf("Failed to create orbitdb instance: %s\n", err)
		return nil, err
	}
	return &db, nil
}

func ConnectToDocStore(ctx context.Context, db orbitdb.OrbitDB, address string) (*orbitdb.DocumentStore, error) {
	// Connect to Kawa orbit store
	options := orbitdb.CreateDBOptions{}
	ac := &accesscontroller.CreateAccessControllerOptions{Access: map[string][]string{"write": {"*"}}}
	options.AccessController = ac
	store, err := db.Docs(ctx, address, &options)
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func LoadStore(ctx context.Context, store iface.Store) error {
	err := store.Load(ctx, 0)
	if err != nil {
		return err
	}
	return nil
}

func AddDocument(doc SongDocument, ctx context.Context, store iface.DocumentStore) (ipfslog.Entry, error) {
	var docMap map[string]interface{}
	docMap = make(map[string]interface{})
	hashID := fmt.Sprintf("%x", sha256.Sum256([]byte(doc.Artist + doc.Title)))  // hashID is type string

	docMap["_id"] = hashID
	docMap["title"] = doc.Title
	docMap["artist"] = doc.Artist
	docMap["album"] = doc.Album
	docMap["cid"] = doc.Cid

	op, err := store.Put(ctx, docMap)
	if err != nil {
		return nil, err
	}
	
	return op.GetEntry(), nil
}

func PrintWithKey(ctx context.Context, store iface.DocumentStore, key string) {
	hashKey := fmt.Sprintf("%x", sha256.Sum256([]byte("Flume" + "Go")))
	op, err := store.Get(ctx, hashKey, &iface.DocumentStoreGetOptions{})
	if err != nil {
		fmt.Printf("failed to get doc: %s\n", err)
		return
	}
	fmt.Println(op)
}

// func GetAllDocuments() {

// }

func main() {
	fmt.Println("Starting Kawa v0.1")

	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("Failed opening configFile: %s\n", err)
		return
	}
	byteValue, _ := ioutil.ReadAll(configFile)
	var config Config
	json.Unmarshal(byteValue, &config)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Step 1: connect to IPFS
	Client, err := CreateIPFSNode()
	if err != nil {
		fmt.Println(err)
		return
	}

	DB, err := CreateDBInstance(ctx, Client)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	Store, err := ConnectToDocStore(ctx, *DB, config.StoreString)
	if err != nil {
		fmt.Println(err)
		return
	}

	LoadStore(ctx, *Store)

	// song := SongDocument{"Go", "Flume", "Palaces", "0x"}
	// entry, err := AddDocument(song, ctx, *Store)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// printEntry(entry)

	PrintWithKey(ctx, *Store, "sum string")


}

func printEntry(entry ipfslog.Entry) {
	fmt.Printf("Log ID: %s\n", entry.GetLogID())
	fmt.Printf("Hash: %s\n", entry.GetHash())
	// fmt.Printf("Identity: %s\n", entry.GetIdentity())
	fmt.Println()
}