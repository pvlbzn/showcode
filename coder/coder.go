package coder

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Img struct {
	W      int
	H      int
	Offset int
}

func Code(path string) {
	f, length := read(path)
	str := toStr(f)
	columns, rows, offset := calculate(length)

	fmt.Println(length, columns, rows, offset)
	m := &Img{columns, rows, offset}
	m.write(str)
}

func read(path string) ([]byte, int) {
	// Open a source file.
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var buf []uint8
	l := 0
	for scanner.Scan() {
		s := scanner.Bytes()
		for _, b := range s {
			l++
			buf = append(buf, uint8(b))
		}
		// Don't forget a new line.
		buf = append(buf, uint8('\n'))
	}
	return buf, l
}

// Convert byte buffer to a string. This hack is needed
// because output file of .ppm type needs a not string
// literals such as "j", but 87.
func toStr(buf []byte) string {
	s := []string{}
	for _, c := range buf {
		literal := fmt.Sprintf("%d ", c)
		s = append(s, literal)
	}
	str := strings.Join(s, "")
	return str
}

// Calculate an image size and return the image object.
// 'l' is a total number of bytes.
func calculate(l int) (int, int, int) {
	px := l
	// Pixel Offset, that means that this number of bytes
	// is needed to create a one pixel successfully.
	pxOffset := 0
	for (px % 3) != 0 {
		px += 1
		pxOffset += 1
	}
	px = px / 3
	// Calculate number of rows and columns.
	rows := int(math.Floor(math.Sqrt(float64(px))))
	for px%rows != 0 {
		rows -= 1
	}
	return (px / rows), rows, pxOffset
}

func (m *Img) write(s string) {
	// Add magic numbers
	magic := fmt.Sprintf("P3 %d %d 255 ", m.W, m.H)
	magic += s
	magic = magic[:len(magic)-1]

	for i := m.Offset; i > 0; i-- {
		magic += " 0 "
	}

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
