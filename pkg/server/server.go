package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Server(baseTarget string, preservePath bool) error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	hostname, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target := baseTarget
		if preservePath {
			target = fmt.Sprintf("%s%s", baseTarget, r.URL.Path)
		}
		log.Info().Str("hostname", hostname).Str("path", r.URL.Path).Msg(r.Host + r.URL.Path + " -> " + target)
		http.Redirect(w, r, target, http.StatusTemporaryRedirect)
	})

	port := 8000
	log.Info().Str("hostname", hostname).Msgf("Server started on 0.0.0.0:%d, see http://127.0.0.1:%d", port, port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
