package middleware

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

var logger = func(method, uri, name string, start time.Time) {
	log.Printf("\"method\":%q  \"uri\":%q    \"name\":%q   \"time\":%q", method, uri, name, time.Since(start))
}

func Logging(h httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		h(w, r, ps)
		logger(r.Method, r.URL.Path, name, start)
	}
}

//type access log
type LineOfLog struct {
	RemoteAddr  string
	ContentType string
	Path        string
	Query       string
	Method      string
	Body        string
}

//TemplateOfLog
//var TemplateOfLog = `
//Remote address:   {{.RemoteAddr}}
//Content-Type:     {{.ContentType}}
//HTTP method:      {{.Method}}
//path:
//{{.Path}}
//query string:
//{{.Query}}
//`
//
////Log using LineOfLog
//func Log(handler http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		bufbody := new(bytes.Buffer)
//		//bufbody.ReadFrom(r.Body)
//		body := bufbody.String()
//
//		line := LineOfLog{
//			r.RemoteAddr,
//			r.Header.Get("Content-Type"),
//			r.URL.Path,
//			r.URL.RawQuery,
//			r.Method, body,
//		}
//		tmpl, err := template.New("line").Parse(TemplateOfLog)
//		if err != nil {
//			panic(err)
//		}
//
//		bufline := new(bytes.Buffer)
//		err = tmpl.Execute(bufline, line)
//		if err != nil {
//			panic(err)
//		}
//
//		//log.Printf(bufline.String())
//		handler.ServeHTTP(w, r)
//	})
//}
