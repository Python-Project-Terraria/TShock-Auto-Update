package main

import (
	"fmt"
	"os"
	"tshockau/git"
	"tshockau/ziper"
)

func main() {
	s := git.GetVersion()
	err := git.Download(s, "test.zip")
	fmt.Println(err)
	ziper.Unzip("test.zip", ".")
	os.Remove("test.zip")
}
