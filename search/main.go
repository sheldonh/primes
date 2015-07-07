package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/sheldonh/primes"
	"os"
	"strconv"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var path string
var count int
var quiet bool

func init() {
	flag.StringVar(&path, "file", "primes.txt", "file to read last prime from and append new primes to")
	flag.IntVar(&count, "primes", 0, "number of primes to find (0 runs until interrupted)")
	flag.BoolVar(&quiet, "quiet", false, "don't print progress")
}

func main() {
	flag.Parse()

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	check(err)
	scanner := bufio.NewScanner(file)
	i := int64(-1)
	for scanner.Scan() {
		s := scanner.Text()
		i, err = strconv.ParseInt(s, 10, 64)
		check(err)
	}
	check(scanner.Err())

	var t time.Time
	if !quiet {
		if count == 0 {
			fmt.Println("will find primes greater than", i, "until interrupted")
		} else {
			var plural string
			if count > 1 {
				plural = "s"
			}
			fmt.Printf("will find %d prime%s greater than %d\n", count, plural, i)
		}
		t = time.Now()
	}

	for n := 0; count == 0 || n < count; {
		i++
		if primes.IsPrime(i) {
			if !quiet {
				n := time.Now()
				fmt.Printf("new prime %d found in %v\n", i, n.Sub(t))
			}
			_, err = fmt.Fprintf(file, "%d\n", i)
			check(err)
			n++
			if !quiet {
				t = time.Now()
			}
		}
	}
}
