package web

import (
	"fmt"
	"net/http"

	"github.com/matherique/share/internal/usecase"
)

type createHandler struct {
	createUseCase usecase.CreateUseCase
}

func RegisterCreateHandler(createUseCase usecase.CreateUseCase) *createHandler {
	i := &createHandler{
		createUseCase: createUseCase,
	}

	return i
}

func (h createHandler) Do(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	link, err := h.createUseCase.Execute(r.Context(), r.Body, r.ContentLength)
	if err != nil {
		http.Error(w, "error on create snipet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, r.Host+"/"+link)
}
