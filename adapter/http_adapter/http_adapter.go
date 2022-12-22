package http_adapter

import "net/http"

type httpHandler struct {
	handler *http.ServeMux
}

func New() *httpHandler {
	router := http.NewServeMux()
	return &httpHandler{
		handler: router,
	}
}

func (h *httpHandler) GetRouter() *http.ServeMux {
	return h.handler

}

var Routes map[string]http.Handler

func InitRoutes() map[string]http.Handler {
	return map[string]http.Handler{}
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	if _, err := rw.Write([]byte("home")); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
