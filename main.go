package main

import (
	"fmt"
	"tshockau/git"
)

func main() {
	s := git.GetVersion()
	fmt.Println(s)
}
