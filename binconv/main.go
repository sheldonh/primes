package main

import (
	"bufio"
	"flag"
	"github.com/sheldonh/primes/int64file"
	"os"
	"strconv"
)

var textPath string
var binPath string
var wantGzip bool

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	flag.StringVar(&textPath, "text", "primes.txt", "text file of primes (input)")
	flag.StringVar(&binPath, "bin", "primes.bin", "binary file of primes (output)")
	flag.BoolVar(&wantGzip, "gzip", false, "enable gzip compression")
}

func main() {
	flag.Parse()

	fd, err := os.OpenFile(binPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	check(err)
	f, err := int64file.NewWriter(fd, wantGzip)
	check(err)
	defer f.Close()

	text, err := os.Open(textPath)
	check(err)
	scanner := bufio.NewScanner(text)
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.ParseInt(s, 10, 64)
		check(err)
		err = f.WriteInt64(i)
		check(err)
	}
}
