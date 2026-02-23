package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"

	"decoder/decoder"
	"decoder/encoder"
	"decoder/server"
)

func main() {
	multiline := flag.Bool("multiline", false, "Process multiple lines from stdin")
	flag.BoolVar(multiline, "m", false, "Process multiple lines from stdin (shorthand)")
	encode := flag.Bool("encode", false, "Encode input into bracket notation")
	flag.BoolVar(encode, "e", false, "Encode input into bracket notation (shorthand)")
	serveMode := flag.Bool("server", false, "Start the web interface server")
	flag.BoolVar(serveMode, "s", false, "Start the web interface server (shorthand)")
	addr := flag.String("addr", ":8080", "Address for the web server (e.g. :8080)")
	flag.Parse()

	if *serveMode {
		fmt.Printf("Art Decoder server listening on %s\n", *addr)
		if err := http.ListenAndServe(*addr, server.NewServeMux()); err != nil {
			fmt.Fprintf(os.Stderr, "server error: %v\n", err)
			os.Exit(1)
		}
		return
	}

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
