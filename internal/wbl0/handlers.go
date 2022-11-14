package wbl0

import (
	"fmt"
	"net/http"
)

type Handler struct {
	cm *CacheManager
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	value, ok := h.cm.Get(id)
	if !ok {
		fmt.Fprint(w, "this identifier does not exist")
		return
	}
	fmt.Fprint(w, value)
}
