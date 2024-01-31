package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matherique/share/internal/usecase"
)

type createSecureHandler struct {
	createSecureSnipet usecase.CreateSecureSnipetUseCase
}

func RegisterCreateSecureHandler(router *chi.Mux, createUseCase usecase.CreateSecureSnipetUseCase) {
	i := &createSecureHandler{
		createSecureSnipet: createUseCase,
	}

	router.Post("/secure", i.do)
}

func (h createSecureHandler) do(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	key, link, err := h.createSecureSnipet.Execute(r.Context(), r.Body, r.ContentLength)
	if err != nil {
		http.Error(w, "error on create snipet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "link: "+r.Host+"/"+link+"\nkey: "+key)
}
