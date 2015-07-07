package main

import (
	"flag"
	"fmt"
	"github.com/sheldonh/primes/int64file"
	"io"
	"os"
)

var binPath string
var gzFlag bool

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	flag.StringVar(&binPath, "bin", "primes.bin", "binary file of primes to dump")
	flag.BoolVar(&gzFlag, "gzip", false, "enable gzip compression")
}

func main() {
	flag.Parse()

	fd, err := os.Open(binPath)
	check(err)
	f, err := int64file.NewReader(fd, gzFlag)
	check(err)
	defer f.Close()

	for {
		i, err := f.ReadInt64()
		if err == io.EOF {
			break
		}
		check(err)
		fmt.Println(i)
	}
}
