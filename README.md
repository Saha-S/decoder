# art-decoder

A Go command-line tool **and web interface** that converts encoded art data into text-based art using bracket notation.

## Setup and Installation

**Requirements:** Go 1.21+

```sh
git clone <repo-url>
cd decoder
go build ./...
```

## Web Interface (art-interface)

Start the web server with the `--server` (or `-s`) flag:

```sh
go run . --server
# Server listening on http://localhost:8080
```

Use `--addr` to specify a custom address:

```sh
go run . --server --addr :9000
```

Open your browser at **http://localhost:8080** to use the interface.

### Interface features

| Feature | Description |
|---|---|
| **Text input** | Paste or type an encoded / plain string |
| **Decode mode** | Expands bracket notation → text art |
| **Encode mode** | Compresses plain text → bracket notation (server-side) |
| **HTTP status badge** | Displays the latest response code (`202 Accepted` / `400 Bad Request`) |

### HTTP Endpoints

| Method | Path | Response |
|---|---|---|
| `GET` | `/` | `200 OK` — main web page |
| `POST` | `/decoder` | `202 Accepted` — result page; `400 Bad Request` — malformed input |
| `GET` | `/decoder` | `303 See Other` — redirects to `/` |
| `GET` | `/static/*` | `200 OK` — CSS and static assets |
| anything else | `/*` | `404 Not Found` |

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

