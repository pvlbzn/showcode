package coder

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func Decode(path string) {
	s := readImage(path)
	arr := arrify(s)
	arr = deleteMagic(arr)
	str := convert(arr)
	writeSource(str)
	fmt.Println(str)
}

func readImage(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	str := []byte{}
	for scanner.Scan() {
		s := scanner.Text()
		for _, b := range s {
			str = append(str, byte(b))
		}
	}
	return string(str)
}

func arrify(s string) []string {
	return strings.Split(s, " ")
}

func deleteMagic(arr []string) []string {
	return arr[4:]
}

func convert(arr []string) string {
	s := []string{}
	for _, c := range arr {
		if c != "" {
			i, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			s = append(s, string(i))
		}
	}
	return strings.Join(s, "")
}

func writeSource(s string) {
	data := []byte(s)
	err := ioutil.WriteFile("output.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
