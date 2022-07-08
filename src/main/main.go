// Design and implement a code base that can upload and retrieve files to/from IPFS.

package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	fmt.Println("Starting Kawa v0.1")
	sh := shell.NewShell("localhost:5001")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command (type help for more info): ")
		text, _ := reader.ReadString('\n')
		switch strings.TrimSpace(text) {
		case "help":
			fmt.Println("enter \"upload\" if you wish to upload a file to IPFS.")
			fmt.Println("enter \"retrieve\" if you wish to retrieve a file from IPFS. Have the CID ready.")
			fmt.Println("use quit or exit to finish")
		case "upload":
			fmt.Print("Enter filename: ")
			filename, _ := reader.ReadString('\n')
			filename = strings.TrimSpace(filename)
			upload(filename, sh)
		case "retrieve":
			fmt.Print("Enter CID: ")
			cid, _ := reader.ReadString('\n')
			cid = strings.TrimSpace(cid)
			retrieve(cid, sh)
		case "exit":
			os.Exit(1)
		case "quit":
			os.Exit(1)
		}
		
	}
	// upload("./go.mp3")

}

func upload(filename string, sh *shell.Shell) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	
	cid, err := sh.Add(strings.NewReader(string(file)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s \n", cid)
}

func retrieve(cid string, sh *shell.Shell) {
	wd, _ := os.Getwd()
	err := sh.Get(cid, wd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println("file saved to working directory")
} 


 