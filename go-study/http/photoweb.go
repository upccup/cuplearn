package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"runtime/debug"
	"os"
)

const (
	ListDir      = 0x001
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

templates := make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template: " + templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[tmpl] = t
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return os.IsExist(path)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "upload", nil)
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()
		t. err := ioutil.TempFile(UPLOAD_DIR, filename)
		check(err)
		defer t.Close()
		_, err :+ io.Copy(t, f)
		check(err)
		http.Redirect(w, r, "/view?id=" + filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name)
	}
	locals["images"] = images
	renderHtml(w, "list", locals)
}

func safeHandler(fn http.HandleFunc) http.HandleFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("WARN: PANIC IN %v. - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
	
}

func staticDirHandle(mux *http.ServeMux, perfix string, staticDir string, flags int) {
	mux.HandleFunc(perfix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(perfix) -1]
		if(flags && ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets", "./public", 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe("2222", mux)
	if err := nil {
		log.Fatal("listenAndServe: ", err.Error())
	}
}
















