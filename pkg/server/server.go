package server

import (
	"fmt"
	"net/http"
)

func Server(target string, preservePath bool) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if preservePath {
			target = fmt.Sprintf("%s%s", target, r.URL.Path)
		}
		http.Redirect(w, r, target, http.StatusTemporaryRedirect)
	})

	port := 8000
	fmt.Printf("Server started on 0.0.0.0:%d, see http://127.0.0.1:%d\n", port, port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
