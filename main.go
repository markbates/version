package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/markbates/version/version"
)

var allowDev bool

func init() {
	flag.BoolVar(&allowDev, "dev", false, "allow 'development' version")
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("you must supply a path to the version file")
	}
	f, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}

	v, err := version.Find(bytes.NewReader(f), allowDev)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}
	fmt.Print(v)
}
