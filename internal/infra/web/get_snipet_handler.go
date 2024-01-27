package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matherique/share/internal/usecase"
	"github.com/matherique/share/pkg/utils"
)

type getSnipetError struct {
	Message string `json:"message"`
}

type getSnipetHandler struct {
	getSnipetUseCase usecase.GetSnipetUseCase
}

func NewGetSnipetHandler(router *chi.Mux, g usecase.GetSnipetUseCase) {
	getsnipets := &getSnipetHandler{
		getSnipetUseCase: g,
	}

	router.Get("/{hash}", getsnipets.get)
}

func (g getSnipetHandler) get(w http.ResponseWriter, r *http.Request) {
	h := chi.URLParam(r, "hash")
	if len(h) == 0 {
		utils.SendRespond(w, http.StatusBadRequest, getSnipetError{"missing link"})
		return
	}

	snipet, err := g.getSnipetUseCase.Execute(r.Context(), h)

	if err != nil {
		utils.SendRespond(w, http.StatusNotFound, getSnipetError{"not found"})
		return
	}

	utils.SendRespond(w, http.StatusOK, snipet)
}
