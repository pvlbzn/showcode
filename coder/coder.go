package coder

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"log"
	"os"
	"strings"
)

type Img struct {
	W int
	H int
}

func Code(path string) {
	f := read(path)
	str := toStr(f)
	write(str)
}

func read(path string) []byte {
	// Open source file.
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	/*
		1.  Find total number of chars / 3, that will be
			a number of pixels. If number % 3 != 0, then
			add 1 to it and try again. Record this additions
			they will be an offset.
		2. 	Find a total number of pixels RGB RGB RGB etc.
	*/

	scanner := bufio.NewScanner(f)
	var buf []uint8
	for scanner.Scan() {
		s := scanner.Bytes()
		for _, b := range s {
			buf = append(buf, uint8(b))
		}
		// Don't forget a new line.
		buf = append(buf, uint8('\n'))
	}
	return buf
}

func toStr(buf []byte) string {
	s := []string{}
	for _, c := range buf {
		// Take a byte and convert in into string, literally.
		literal := fmt.Sprintf("%d ", c)
		s = append(s, literal)
	}
	str := strings.Join(s, "")
	return str
}

func write(s string, m Img) {
	magic := fmt.Sprintf("P3 %d %d ", m.W, m.H)
	magic += s

	fmt.Println(magic)

	f, err := os.Create("output.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(magic)
	if err != nil {
		log.Fatal(err)
	}
}
