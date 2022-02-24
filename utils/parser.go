package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseWithSize(path string, bufferSize int, cb func(line int, content string)) {
	file, err := os.Open(os.ExpandEnv("$HOME") + path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, bufferSize)
	scanner.Buffer(buf, bufferSize)
	count := 0
	for scanner.Scan() {
		cb(count, scanner.Text())
		count++
	}
	errCheck(scanner.Err())

}

func Parse(path string, cb func(line int, words []string)) {
	file, err := os.Open(os.ExpandEnv("$HOME") + path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		cb(count, strings.Split(scanner.Text(), " "))
		count++
	}
	errCheck(scanner.Err())

}
