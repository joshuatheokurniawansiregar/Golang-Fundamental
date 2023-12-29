package pkg

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type person struct {
	name string
	age  int64
}

var persons = []person{
	{name: "Joshua", age: 23},
	{name: "JJ", age: 32},
}

func Hello(write http.ResponseWriter, req *http.Request) {
	for _, data := range persons {
		var string_format string = strconv.FormatInt(data.age, 10)
		fmt.Fprintf(write, data.name)
		fmt.Fprintf(write, string(string_format))
	}
}

func Headers(write http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(write, "%v: %v\n", name, h)
		}
	}
}

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		Hello(w, r)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	if r.Method == "GET" {

		ttemp, _ := template.ParseFiles("login.gtpl")
		ttemp.Execute(w, nil)
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
		if err != nil {
			log.Fatal(err)
		}
	}
	// else {
	// 	r.ParseForm()
	// 	get_int, err := strconv.Atoi(r.Form.Get("age"))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if get_int > 100 {
	// 		fmt.Println("bigger than 100")
	// 	}
	// 	if len(r.Form["username"][0]) == 0 {
	// 		fmt.Println("test")
	// 	} else {
	// 		fmt.Println("username: ", r.Form.Get("username"))
	// 		fmt.Println("password", r.Form.Get("password"))
	// 	}
	// }
}

func login_another(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
	}
	fmt.Println(template.HTMLEscapeString(r.Form.Get("username")))
	fmt.Println(template.HTMLEscapeString(r.Form.Get("password")))
	template.HTMLEscape(w, []byte(r.Form.Get("username")))
	template.HTMLEscape(w, []byte(r.Form.Get("password")))

}

func upload(w http.ResponseWriter, r *http.Request) {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	t, _ := template.ParseFiles("upload.gtpl")
	fmt.Println(r.Body)
	t.Execute(w, token)
}

// func upload_another(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("upload.gtpl")
// 		t.Execute(w, nil)
// 	} else {
// 		r.Body = http.MaxBytesReader(w, r.Body, 32<<20+1024)
// 		reader, err := r.MultipartReader()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 		}
// 		text := make([]byte, 512)
// 		p, err := reader.NextPart()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		_, err = p.Read(text)
// 		if err != nil && err != io.EOF {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		p, err = reader.NextPart()
// 		if err != nil && err != io.EOF {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		if p.FormName() != "uploadfile" {
// 			http.Error(w, "uploadfile is expected", http.StatusBadRequest)
// 			return
// 		}
// 		buf := bufio.NewReader(p)
// 		sniff, _ := buf.Peek(512)
// 		contentType := http.DetectContentType(sniff)
// 		if contentType != "text/plain" {
// 			http.Error(w, "file type not allowed", http.StatusBadRequest)
// 			return
// 		}
// 		f, err := ioutil.TempFile("", "")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		defer f.Close()
// 		var maxSize int64 = 32 << 20
// 		lmt := io.MultiReader(buf, io.LimitReader(p, maxSize-511))
// 		written, err := io.Copy(f, lmt)
// 		if err != nil && err != io.EOF {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		if written > maxSize {
// 			os.Remove(f.Name())
// 			http.Error(w, "file size over limit", http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }
