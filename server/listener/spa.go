package listener

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	html "html/template"

	"pacstall.dev/webserver/ssr"
)

type spaHandler struct {
	staticPath string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filePath := filepath.Join(h.staticPath, path)
	if strings.Contains(filePath, "..") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	_, err = os.Stat(filePath)
	fileNotFound := os.IsNotExist(err)
	isAPI := strings.Contains(filePath, "/api/")

	if (fileNotFound && !isAPI) || (!fileNotFound && !isAPI && filePath == h.staticPath) {
		serveIndexHtml(w, r, h.staticPath)
		return
	} else if fileNotFound && isAPI {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func serveIndexHtml(w http.ResponseWriter, r *http.Request, staticPath string) {
	templateData := ssr.GetTemplateForPath(r.URL.Path)
	template := html.Must(html.ParseFiles(filepath.Join(staticPath, "index.html")))

	w.Header().Add("Content-Type", "text/html")
	if template.Execute(w, templateData) != nil {
		http.Error(w, "Could not execute template", http.StatusInternalServerError)
	}
}
