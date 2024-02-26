package config

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Start() {
	r := mux.NewRouter()

	h := func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte("<h1>Hello world</h1>"))
	}
	// page routes
	r.HandleFunc("/signup", h).Methods(http.MethodGet)
	r.HandleFunc("/signin", h).Methods(http.MethodGet)

	// protected page routes
	p := r.PathPrefix("/").Subrouter()
	p.Use(authMiddleware)
	p.HandleFunc("/", h).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(s.listenAddr, r))
}

// TODO: SHOULD BE DELETED
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside auth middleware")
		next.ServeHTTP(w, r)
	})
}
