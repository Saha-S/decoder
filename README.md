# art-decoder

A Go command-line tool that converts encoded art data into text-based art using bracket notation.

## Setup and Installation

**Requirements:** Go 1.21+

```sh
git clone <repo-url>
cd decoder
go build ./...
```

## Usage

### Basic Decode

Pass a single encoded string as the argument:

```sh
go run . "[5 #][5 -_]-[5 #]"
# Output: #####-_-_-_-_-_-#####
```

**Bracket notation:** `[N str]` repeats `str` exactly `N` times.

```sh
go run . "[3 A]"       # → AAA
go run . "ABC[10 D]EFG" # → ABCDDDDDDDDDDEFG
```

### Multi-line Decode (`--multiline` / `-m`)

Read multiple lines from stdin (or a file), decode each line:

```sh
echo -e "[3 A]\n[2 BC]" | go run . --multiline
# Output:
# AAA
# BCBC

go run . --multiline testdata/plane.encoded
```

### Encode Mode (`--encode` / `-e`)

Encode a plain text string into bracket notation:

```sh
go run . --encode "#####-_-_-_-_-_-#####"
# Output: [5 #][5 -_]-[5 #]
```

### Combining Flags

Encode multiple lines from stdin:

```sh
echo -e "AAABBB\nhello" | go run . --encode --multiline
# Output:
# [3 A][3 B]
# hello
```

### Error Handling

The tool prints `Error` and exits for:
- Non-positive or non-numeric repeat count: `[0 #]`, `[abc #]`
- Missing space separator: `[5#]`
- Empty repeat string: `[5 ]`
- Unbalanced brackets: `[5 #` or `5 #]`

