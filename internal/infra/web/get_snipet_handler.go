package web

import (
	"fmt"
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

func NewGetSnipetHandler(g usecase.GetSnipetUseCase) *getSnipetHandler {
	getsnipets := &getSnipetHandler{
		getSnipetUseCase: g,
	}

	return getsnipets
}

func (g getSnipetHandler) Do(w http.ResponseWriter, r *http.Request) {
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

	if snipet.IsSecure {
		utils.SendRespond(w, http.StatusUnauthorized, getSnipetError{"unauthorized"})
		return
	}

	if !utils.IsBrowerRequest(r) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, snipet.Content)
		return
	}

	utils.SendRespond(w, http.StatusOK, snipet)
}
