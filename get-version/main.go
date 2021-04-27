package main

import (
	// My "gathering_info" package (library) in $GOROOT\src\mypackages\gathering_info
	// get "mypackages/gathering_info"
	// My "gathering_info" package (library) in $GOPATH/src/github.com/graypit/my-go-libs/gathering_info
	get "github.com/graypit/my-go-libs/gathering_info"
)

func main() {
	get.Version()
	get.NumOfCpu()
}
