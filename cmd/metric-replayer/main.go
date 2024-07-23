package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// metric-replayer will replay metrics slowly, base on the timestamp in the file.
// it will read fast until it gets to the current timestamp,
// then it will read slowly, matching the speed of the timestamps.

var (
	filename = flag.String("filename", "", "filename with metrics to read.")
)

func main() {
	flag.Parse()

	var err error
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	flag.Parse()
	if filename == nil || len(*filename) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Unable to read file %q", *filename)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		fmt.Println(s.Text())

		// read the timestamp off the end. if it's in the future, sleep.
		parts := strings.Split(s.Text(), " ")
		if len(parts) == 3 {
			ns, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				continue
			}
			sleepTime := time.Unix(ns/1e9, ns%1e9).Sub(time.Now())
			time.Sleep(sleepTime)
		}
	}
	if s.Err() != nil {
		log.Fatalf("Error reading input file: %s", err)
	}
	log.Println("Replay done!. Going to sleep until telegraf exits")
	s = bufio.NewScanner(os.Stdin)
	for s.Scan() {
	}
}
