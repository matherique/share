package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matherique/share/internal/usecase"
	"github.com/matherique/share/pkg/utils"
)

type getSecureSnipetHandler struct {
	getSnipetUseCase usecase.GetSecureSnipetUseCase
}

func NewGetSecureSnipetHandler(g usecase.GetSecureSnipetUseCase) *getSecureSnipetHandler {
	getsnipets := &getSecureSnipetHandler{
		getSnipetUseCase: g,
	}

	return getsnipets
}

func (g getSecureSnipetHandler) Do(w http.ResponseWriter, r *http.Request) {
	h := chi.URLParam(r, "hash")
	if len(h) == 0 {
		utils.SendRespond(w, http.StatusBadRequest, getSnipetError{"missing link"})
		return
	}

	key := r.Header.Get("Authorization")
	if len(key) == 0 {
		utils.SendRespond(w, http.StatusUnauthorized, getSnipetError{"unathorized"})
		return
	}

	snipet, err := g.getSnipetUseCase.Execute(r.Context(), h, []byte(key))

	if err != nil {
		utils.SendRespond(w, http.StatusNotFound, getSnipetError{"not found"})
		return
	}

	if !utils.IsBrowerRequest(r) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, snipet.Content)
		return
	}

	utils.SendRespond(w, http.StatusOK, snipet)
}
