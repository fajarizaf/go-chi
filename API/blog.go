package API

import "net/http"

func GetBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ini get blog"))
}

func GetListBlog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ini LIST blog"))
}
