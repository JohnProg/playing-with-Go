package main

import (
	"html/template"
	"io"
	"net/http"
	// "net/url"
	"os"
	"log"
	"regexp"
)

const (
	WEBSITE           = "http://blueimp.github.com/jQuery-File-Upload/"
	MIN_FILE_SIZE     = 1       // bytes
	MAX_FILE_SIZE     = 5000000 // bytes
	IMAGE_TYPES       = "image/(gif|p?jpeg|(x-)?png)"
	ACCEPT_FILE_TYPES = IMAGE_TYPES
	EXPIRATION_TIME   = 300 // seconds
	THUMBNAIL_PARAM   = "=s80"
)

var (
	imageTypes      = regexp.MustCompile(IMAGE_TYPES)
	acceptFileTypes = regexp.MustCompile(ACCEPT_FILE_TYPES)
)

type FileInfo struct {
	Type         string            `json:"type"`
	Size         int64             `json:"size"`
	Error        string            `json:"error,omitempty"`
}

func (fi *FileInfo) ValidateType() (valid bool) {
	if acceptFileTypes.MatchString(fi.Type) {
		return true
	}
	fi.Error = "Filetype not allowed"
	return false
}

func (fi *FileInfo) ValidateSize() (valid bool) {
	if fi.Size < MIN_FILE_SIZE {
		fi.Error = "File is too small"
	} else if fi.Size > MAX_FILE_SIZE {
		fi.Error = "File is too big"
	} else {
		return true
	}
	return false
}

//Compile templates on start
var templates = template.Must(template.ParseFiles("tmpl/upload.html"))

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl+".html", data)
}

//This is where the action happens.
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// fi = &FileInfo{}

	// params, _ := url.ParseQuery(r.URL.RawQuery)
	switch r.Method {
	//GET displays the upload form.
	case "GET":
		display(w, "upload", nil)

	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//get a ref to the parsed multipart form
		m := r.MultipartForm
		//get the *fileheaders
		files := m.File["myfiles"]
		for i, _ := range files {

			// if !fi.ValidateType() {
			// 	return
			// }

			// if !fi.ValidateSize() {
			// 	return
			// }
			// fh := &FileHeader{
			// 	Filename: filename,
			// 	Header:   p.Header,
			// }
			log.Println(files[i].Filename)
			log.Println(files[i].Header)
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.
			dst, err := os.Create("/Users/johnmachahuay/Documents/Projects/Go-MultipleFileUpload/assets/images/" + files[i].Filename)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}
		//display success message.
		display(w, "upload", files)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	//static file handler.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
