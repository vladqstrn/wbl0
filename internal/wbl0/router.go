package wbl0

import "net/http"

func InitRouter(cm *CacheManager) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/delivery", &Handler{cm})
	return mux
}
