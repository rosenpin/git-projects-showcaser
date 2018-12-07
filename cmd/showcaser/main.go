package main

import (
	"flag"
	"fmt"
)

func main() {
	var gitPlatform string
	flag.StringVar(&gitPlatform, "git", "github", "git platform to use")
	flag.Parse()
	fmt.Println(gitPlatform)
}
