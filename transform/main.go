package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/dhanusaputra/go-exercises/transform/primitive"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `<html><body>
		<form action="/upload" method="post" enctype="multipart/form-data">
			<input type="file" name="image" />
			<button type="submit">Upload Image</button>
		</form>
		</body></html>`
		fmt.Fprint(w, html)
	})
	mux.HandleFunc("/modify/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("/img/" + filepath.Base(r.URL.Path))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer f.Close()
		ext := filepath.Ext(f.Name())
		modeStr := r.FormValue("mode")
		if modeStr == "" {
			renderModeChoices(w, r, f, ext)
			return
		}
		mode, err := strconv.Atoi(modeStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_ = mode
	})
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		onDisk, err := tempfile("", filepath.Ext(header.Filename)[1:])
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		defer onDisk.Close()
		_, err = io.Copy(onDisk, file)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/modify/"+filepath.Base(onDisk.Name()), http.StatusFound)
	})
	fs := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img/", fs))
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func renderModeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string) {
	a, err := genImage(rs, ext, 33, primitive.Cicle)
	if err != nil {
		panic(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rs.Seek(0, 0)
	b, err := genImage(rs, ext, 33, primitive.Ellipse)
	if err != nil {
		panic(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rs.Seek(0, 0)
	c, err := genImage(rs, ext, 33, primitive.Polygon)
	if err != nil {
		panic(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rs.Seek(0, 0)
	d, err := genImage(rs, ext, 33, primitive.Combo)
	if err != nil {
		panic(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	html := `<html><body>
		{{range .}}
			<a href="/modify/{{.}}?mode={{.Mode}}>
				<img style="width: 20%;" src="/img/{{.Name}}">
			</a>
		{{end}}
		</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	data := []struct {
		Name string
		Mode primitive.Mode
	}{
		{Name: filepath.Base(a), Mode: primitive.Cicle},
		{Name: filepath.Base(b), Mode: primitive.Ellipse},
		{Name: filepath.Base(c), Mode: primitive.Polygon},
		{Name: filepath.Base(d), Mode: primitive.Combo},
	}
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func genImage(file io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {
	out, err := primitive.Transform(file, ext, 33, primitive.WithMode(mode))
	if err != nil {
		return "", err
	}
	outFile, err := tempfile("", ext)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	io.Copy(outFile, out)
	return outFile.Name(), nil
}

func tempfile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("./img/", prefix)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
