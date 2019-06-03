package httprouter

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	// httprouter 使用压缩字典树
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oh no not found"))
	})
	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, c interface{}) {
		log.Printf("recovering from panic ,reason: %#v", c.(error))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(c.(error).Error()))
	}

}
