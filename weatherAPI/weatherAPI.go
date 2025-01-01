package main

import (
	"os"
	"weatherAPI/httpPackage"
)

func main() {
	httpPackage.RunServer()
	defer os.Exit(0) 
}



