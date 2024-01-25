package app

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type importApp struct {
	maxSizeAllowed int64
}

func newImport() *importApp {
	return &importApp{
		maxSizeAllowed: 1024 * 1024,
	}
}

func (a importApp) do(w http.ResponseWriter, r *http.Request) {

	if r.ContentLength > a.maxSizeAllowed {
		http.Error(w, "max size allowed is 1mb", http.StatusRequestEntityTooLarge)
		return
	}

	buff := new(bytes.Buffer)

	n, err := io.CopyN(buff, r.Body, r.ContentLength)
	if err != nil {
		slog.Error("error uploading file", "err", err)
		http.Error(w, "error uploading file", http.StatusInternalServerError)
		return
	}

	slog.Info("receive bytes", "size", n)

	w.Write([]byte("ok"))
}
