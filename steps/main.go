package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/sheldonh/primes"
	"math/big"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var threshPath string
var statePath string
var count int
var quiet bool

func init() {
	flag.StringVar(&threshPath, "thresh", "thresh.txt", "file to write step tresholds to")
	flag.StringVar(&statePath, "state", "thresh.txt.state", "file to write state to on exit")
	flag.BoolVar(&quiet, "quiet", false, "don't print progress")
}

func updateState(p int64, n int) error {
	state, err := os.OpenFile(statePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	_, err = fmt.Fprintf(state, "%d:%d\n", p, n)
	if err != nil {
		return err
	}
	return state.Close()
}

func main() {
	flag.Parse()

	var p = int64(0)
	var n = 1
	if _, err := os.Stat(statePath); err == nil {
		state, err := os.Open(statePath)
		check(err)
		scanner := bufio.NewScanner(state)
		scanner.Scan()
		check(scanner.Err())
		x := strings.Split(scanner.Text(), ":")
		p, err = strconv.ParseInt(x[0], 10, 64)
		check(err)
		n, err = strconv.Atoi(x[1])
		check(err)
		check(state.Close())
	}
	fmt.Println("starting at", p, "with", n, "tests")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Fprintf(os.Stderr, "received %v, saving state and exiting...\n", sig)
		updateState(p, n)
		os.Exit(1)
	}()

	thresh, err := os.OpenFile(threshPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

	bigInt := new(big.Int)
	for {
		if bigInt.SetInt64(p).ProbablyPrime(n) && !primes.IsPrime(p) {
			fmt.Printf("non-prime %d is ProbablyPrime(%d)\n", p, n)
			_, err = fmt.Fprintf(thresh, "%d:%d\n", p, n)
			check(err)
			n++
			// Sanity check: I haven't proven that incrementing tests by one always fixes the guess
			if bigInt.ProbablyPrime(n) {
				panic(fmt.Sprintf("%d is ProbablyPrime(%d) and ProbablyPrime(%d)", p, n-1, n))
			}
		}
		p++
	}
}
