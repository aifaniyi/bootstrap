package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	inputFile := flag.String("inputFile", "", "absolute path to schema definition file")
	lang := flag.String("lang", "", "programming language output to be generated")
	projectName := flag.String("projectName", "", "name for new project")
	outDirName := flag.String("outDirName", "", "output directory for models, repos, services and controllers")

	flag.Parse()

	if *lang == "" {
		log.Fatal("lang is required")
	}

	if *inputFile == "" {
		log.Fatal("inputFile schema file name is required")
	}

	if *projectName == "" {
		log.Fatal("projectName  is required")
	}

	switch *lang {
	case "golang":
		err := generateGolang(*inputFile, *outDirName, *projectName)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Printf("unknown language %s", *lang)
	}

}
