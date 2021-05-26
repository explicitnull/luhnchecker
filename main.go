package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/theplant/luhn"
)

const (
	salt = "C6qpl4nCgYhg08vTXaQs"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(bytes)))

	var valid, total int
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		total++

		num, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
		}

		res := luhn.Valid(num)
		if res {
			valid++
		}
	}

	log.Printf("tolal: %d, valid: %d", total, valid)
}
