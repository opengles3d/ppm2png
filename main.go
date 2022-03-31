package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

func parseSize(line string) (width int, height int, err error) {
	ints := strings.Split(line, " ")
	if len(ints) != 2 {
		err = fmt.Errorf("parseSize error: %s", line)
		return 0, 0, err
	}
	width, err = strconv.Atoi(ints[0])
	if err != nil {
		return 0, 0, err
	}

	height, err = strconv.Atoi(ints[1])
	if err != nil {
		return 0, 0, err
	}
	return width, height, nil
}

func parseColor(line string) (r, g, b uint8) {
	ints := strings.Split(line, " ")
	if len(ints) != 3 {
		return 0, 0, 0
	}
	ri, err := strconv.Atoi(ints[0])
	if err != nil {
		return 0, 0, 0
	}

	gi, err := strconv.Atoi(ints[1])
	if err != nil {
		return 0, 0, 0
	}

	bi, err := strconv.Atoi(ints[2])
	if err != nil {
		return 0, 0, 0
	}
	return uint8(ri), uint8(gi), uint8(bi)
}

func main() {
	r, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()
	if !strings.HasPrefix(line, "P3") {
		fmt.Println("Error format")
		return
	}

	scanner.Scan()
	line = scanner.Text()
	width, height, err := parseSize(line)
	if err != nil {
		fmt.Println("Error format")
		return
	}
	scanner.Scan()
	line = scanner.Text()
	if !strings.HasPrefix(line, "255") {
		fmt.Println("Error format")
		return
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer file.Close()

	m := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			scanner.Scan()
			line = scanner.Text()
			r, g, b := parseColor(line)
			m.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	png.Encode(file, m)
}
