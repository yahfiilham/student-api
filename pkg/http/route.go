package http

import "net/http"

func Routes(h *Handler) {
	http.HandleFunc("/health-check", h.GetHealthCheck)
}
