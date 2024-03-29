package main

import (
	"flag"
	"fmt"
)

var cpuprofile string
var enableProfile bool

func init() {
	flag.StringVar(&cpuprofile, "cpuprofile", "cpuprofile", "cpu profile name")
	flag.BoolVar(&enableProfile, "enableProfile", false, "enableProfile")
}

// go run flag.go --cpuprofile=cpupp --enableProfile=false hello foo
func main() {
	flag.Parse()

	if cpuprofile != "" {
		fmt.Println(cpuprofile)
	}

	if enableProfile {
		fmt.Print(enableProfile)
	}
	fmt.Println(flag.Args())
}
