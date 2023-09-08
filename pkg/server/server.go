package server

import (
	"fmt"
	"net/http"
)

func Server(target string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target, http.StatusTemporaryRedirect)
	})

	port := 8000
	fmt.Printf("Server started on 0.0.0.0:%d, see http://127.0.0.1:%d\n", port, port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
