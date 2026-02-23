package server

import (
	"embed"
	"html/template"
	"net/http"
	"strings"

	"decoder/decoder"
	"decoder/encoder"
)

//go:embed templates
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

// PageData holds template rendering data.
type PageData struct {
	Input      string
	Mode       string
	Result     string
	StatusCode int
	HasResult  bool
	IsError    bool
	ErrorMsg   string
}

// NewServeMux creates and returns an http.ServeMux with all routes registered.
func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	tmpl := template.Must(template.ParseFS(templateFS, "templates/*.html"))

	mux.Handle("/static/", http.FileServer(http.FS(staticFS)))

	// GET / — main page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			_ = tmpl.ExecuteTemplate(w, "error.html", PageData{StatusCode: http.StatusNotFound, ErrorMsg: "Page not found"})
			return
		}
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = tmpl.ExecuteTemplate(w, "index.html", PageData{})
	})

	// POST /decoder — encode or decode input
	mux.HandleFunc("/decoder", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = tmpl.ExecuteTemplate(w, "index.html", PageData{
				StatusCode: http.StatusBadRequest,
				IsError:    true,
				ErrorMsg:   "Bad request: unable to parse form",
			})
			return
		}

		input := strings.TrimRight(r.FormValue("input"), "\r\n")
		mode := r.FormValue("mode")
		if mode == "" {
			mode = "decode"
		}

		if input == "" {
			w.WriteHeader(http.StatusBadRequest)
			_ = tmpl.ExecuteTemplate(w, "index.html", PageData{
				Mode:       mode,
				StatusCode: http.StatusBadRequest,
				IsError:    true,
				ErrorMsg:   "Bad request: input must not be empty",
			})
			return
		}

		var result string
		var processErr error

		if mode == "encode" {
			result = encoder.Encode(input)
		} else {
			result, processErr = decoder.Decode(input)
			if processErr != nil {
				w.WriteHeader(http.StatusBadRequest)
				_ = tmpl.ExecuteTemplate(w, "index.html", PageData{
					Input:      input,
					Mode:       mode,
					StatusCode: http.StatusBadRequest,
					IsError:    true,
					ErrorMsg:   "Bad request: invalid encoded string",
				})
				return
			}
		}

		w.WriteHeader(http.StatusAccepted)
		_ = tmpl.ExecuteTemplate(w, "index.html", PageData{
			Input:      input,
			Mode:       mode,
			Result:     result,
			StatusCode: http.StatusAccepted,
			HasResult:  true,
		})
	})

	return mux
}
