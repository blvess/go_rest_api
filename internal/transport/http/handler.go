package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler stores pointers to our services
type Handler struct {
	Router *mux.Router
}

// NewHandler will create a new Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes creates our api routes
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")

	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})
}
