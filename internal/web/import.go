package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matherique/share/internal/usecase"
)

type importHandler struct {
	importUseCase usecase.ImportUseCase
}

func RegisterImportHandler(router *chi.Mux, importUseCase usecase.ImportUseCase) {
	i := &importHandler{
		importUseCase: importUseCase,
	}

	router.Post("/", i.do)
}

func (h importHandler) do(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	link, err := h.importUseCase.Execute(r.Context(), r.Body, r.ContentLength)
	if err != nil {
		http.Error(w, "error on import file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(link))
}
