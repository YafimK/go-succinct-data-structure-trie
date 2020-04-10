package main

import (
	"flag"
	"fmt"
	"github.com/YafimK/go-succinct-data-structure-trie/domain_tree"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	allowedChars := " _abcdefghijklmnopqrstuvwxyz0123456789.-"

	fileMode := flag.NewFlagSet("file", flag.ExitOnError)
	dirMode := flag.NewFlagSet("dir", flag.ExitOnError)

	if len(os.Args) < 4 {
		fmt.Println("expected 'file' or 'dir' commands followed by source file / folder path and output file / folder" +
			" path")
		os.Exit(1)
	}

	sourcePath := []string{}
	outputPath := []string{}
	switch os.Args[1] {
	case "file":
		if err := fileMode.Parse(os.Args[2:]); err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println("load world list file")
		fmt.Println("source file:", fileMode.Arg(0))
		fmt.Println("output file:", fileMode.Arg(1))
		sourcePath = append(sourcePath, fileMode.Arg(0))
		outputPath = append(outputPath, dirMode.Arg(1))
	case "dir":
		if err := dirMode.Parse(os.Args[2:]); err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println("recurse word list directory")
		fmt.Println("source folder:", dirMode.Arg(0))
		fmt.Println("output folder:", dirMode.Arg(1))
		sourceDir := dirMode.Arg(0)
		outputDir := dirMode.Arg(1)
		files, err := ioutil.ReadDir(sourceDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			sourcePath = append(sourcePath, filepath.Join(sourceDir, f.Name()))
			outputPath = append(outputPath, filepath.Join(outputDir, f.Name(), ".tree"))
		}

	default:
		fmt.Println("expected 'file' or 'dir' commands")
		os.Exit(1)
	}
	if len(sourcePath) != len(outputPath) {
		log.Fatalf("oh oh - sourcepaths are smaller then outputpath")
	}
	for i, sourceFile := range sourcePath {
		domain_tree.WriteNewDomainTree(allowedChars, path.Base(sourceFile), sourceFile, outputPath[i])
	}

}
