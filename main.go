package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"decoder/decoder"
	"decoder/encoder"
)

func main() {
	multiline := flag.Bool("multiline", false, "Process multiple lines from stdin")
	flag.BoolVar(multiline, "m", false, "Process multiple lines from stdin (shorthand)")
	encode := flag.Bool("encode", false, "Encode input into bracket notation")
	flag.BoolVar(encode, "e", false, "Encode input into bracket notation (shorthand)")
	flag.Parse()

	if *multiline {
		var input *os.File
		if flag.NArg() > 0 {
			f, err := os.Open(flag.Arg(0))
			if err != nil {
				fmt.Println("Error")
				os.Exit(1)
			}
			defer f.Close()
			input = f
		} else {
			input = os.Stdin
		}

		scanner := bufio.NewScanner(input)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if *encode {
			for _, line := range encoder.EncodeLines(lines) {
				fmt.Println(line)
			}
		} else {
			decoded, err := decoder.DecodeLines(lines)
			if err != nil {
				fmt.Println("Error")
				os.Exit(1)
			}
			for _, line := range decoded {
				fmt.Println(line)
			}
		}
	} else {
		if flag.NArg() == 0 {
			fmt.Println("Error")
			os.Exit(1)
		}
		arg := flag.Arg(0)
		if *encode {
			fmt.Println(encoder.Encode(arg))
		} else {
			result, err := decoder.Decode(arg)
			if err != nil {
				fmt.Println("Error")
				os.Exit(1)
			}
			fmt.Println(result)
		}
	}
}
