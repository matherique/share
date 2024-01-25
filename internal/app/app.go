package app

import "net/http"

func Register() {
	a := newImport()

	http.HandleFunc("/", a.do)
}
