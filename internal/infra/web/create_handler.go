package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matherique/share/internal/usecase"
)

type createHandler struct {
	createUseCase usecase.CreateUseCase
}

func RegisterCreateHandler(router *chi.Mux, createUseCase usecase.CreateUseCase) {
	i := &createHandler{
		createUseCase: createUseCase,
	}

	router.Post("/", i.do)
}

func (h createHandler) do(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	link, err := h.createUseCase.Execute(r.Context(), r.Body, r.ContentLength)
	if err != nil {
		http.Error(w, "error on create snipet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, r.Host+"/"+link)
}
