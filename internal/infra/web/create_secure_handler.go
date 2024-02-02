package web

import (
	"fmt"
	"net/http"

	"github.com/matherique/share/internal/usecase"
)

type createSecureHandler struct {
	createSecureSnipet usecase.CreateSecureSnipetUseCase
}

func RegisterCreateSecureHandler(createUseCase usecase.CreateSecureSnipetUseCase) *createSecureHandler {
	i := &createSecureHandler{
		createSecureSnipet: createUseCase,
	}

	return i
}

func (h createSecureHandler) Do(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	key, link, err := h.createSecureSnipet.Execute(r.Context(), r.Body, r.ContentLength)
	if err != nil {
		http.Error(w, "error on create snipet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "link: "+r.Host+"/"+link+"\nkey: "+key)
}
