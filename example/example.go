package main

import (
	"flag"
	"fmt"

	wadparser "github.com/KuroseNil/doom-wad-parser-go"
)

var (
	WADPath *string
)

func init() {
	WADPath = flag.String("wad", "./DOOM1.wad", "path to .wad file")
	flag.Parse()
}

func main() {
	wad, err := wadparser.LoadWAD(*WADPath)
	if err != nil {
		panic(err)
	}
	// wadtype should be converted to string because it is recognized as 4byte array
	wadtype := fmt.Sprintf("%s", wad.Header.WADType)
	fmt.Printf("WADTYPE : %v\n", wadtype)
	fmt.Printf("DirectoryCount : %d\n", wad.Header.DirectoryCount)
	fmt.Printf("DirectoryOffset : %d\n", wad.Header.DirectoryOffset)
}
