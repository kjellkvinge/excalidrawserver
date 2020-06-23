package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

//go:generate go run generator.go
func main() {
	//	fs := http.FileServer(http.Dir("./excalidraw/build/"))

	r := mux.NewRouter()
	fs := http.FileServer(assets)

	// API
	r.Handle("/api/v2/post/", new(countHandler))
	r.Handle("/api/v2/{id}", new(countHandler))

	// catchall - serve static files
	r.PathPrefix("/").Handler(fs)

	log.Println("Listening on :3001...")
	err := http.ListenAndServe(":3001", r)
	if err != nil {
		log.Fatal(err)
	}
}

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	filenamebase := "/tmp/drawing"

	if r.Method == "POST" {
		h.n++ // increate id counter
		filename := fmt.Sprintf("%s_%d", filenamebase, h.n)
		for fileExists(filename) {
			h.n++
			filename = fmt.Sprintf("%s_%d", filenamebase, h.n)
		}

		f, err := os.Create(filename)
		defer f.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		f.Write(b)
		fmt.Fprintf(w, `{"data":"http://%s/api/v2/%d","id":"%d"}`, r.Host, h.n, h.n)
		return
	}

	id := mux.Vars(r)["id"]
	dat, err := ioutil.ReadFile(filenamebase + "_" + id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(dat))

}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
