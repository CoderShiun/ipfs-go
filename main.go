package main

import (
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"os"
	"strings"
)

var ubOzoneClient = connUBOzoneClient()

func main() {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader("tttttttttttt"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", cid)

	data, err := sh.AddDir("/home/shiun/Music")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", data)

}