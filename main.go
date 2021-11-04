package main

import (
	"log"
)

func main() {
	//Check command line args
	err := getFlags()
	if err != nil {
		log.Fatal(err)
	}

	//Validate args
	//Validate S3 access
	//Main Loop
	//	Validate access to git repo
	//	check/create temp dir
	//	do git bundle
	//	compress
	//	upload to S3
}
