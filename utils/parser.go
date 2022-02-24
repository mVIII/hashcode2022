package utils

import (
	"bufio"
	"log"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseWithSize(path string, bufferSize int, cb func(string)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, bufferSize)
	scanner.Buffer(buf, bufferSize)
	for scanner.Scan() {
		cb(scanner.Text())
	}
	errCheck(scanner.Err())

}

func Parse(path string, cb func(string)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cb(scanner.Text())
	}
	errCheck(scanner.Err())

}
