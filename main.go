package main

import (
	"fmt"
	"tshockau/git"
	"tshockau/ziper"
)

func main() {
	s := git.GetVersion()
	err := git.Download(s, "test.zip")
	fmt.Println(err)
	ziper.Unpack("test.zip", ".")
}
