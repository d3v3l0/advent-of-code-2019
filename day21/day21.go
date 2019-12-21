package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	//"github.com/dhconnelly/advent-of-code-2019/geom"
	"github.com/dhconnelly/advent-of-code-2019/intcode"
)

func readLine(ch <-chan int64) string {
	var data []byte
	for c := byte(<-ch); c != '\n'; c = byte(<-ch) {
		data = append(data, c)
	}
	return string(data)
}

func writeLine(ch chan<- int64, s string) {
	for _, c := range s {
		ch <- int64(c)
	}
	ch <- int64('\n')
}

type springdroid struct {
	prog []int64
}

func (d springdroid) execute(r io.Reader) error {
	in := make(chan int64)
	out := intcode.RunProgram(d.prog, in)

	fmt.Println(readLine(out))
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		line := scan.Text()
		writeLine(in, line)
		if line == "WALK" {
			break
		}
	}
	for c, ok := <-out; ok; c, ok = <-out {
		if c <= math.MaxInt8 {
			fmt.Printf("%c", c)
		} else {
			fmt.Printf("%d", c)
		}
	}
	return scan.Err()
}

func main() {
	data, err := intcode.ReadProgram(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var r io.Reader
	if len(os.Args) > 2 {
		f, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		r = f
	} else {
		r = os.Stdin
	}
	d := springdroid{data}
	d.execute(r)
}