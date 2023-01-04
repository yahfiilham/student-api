package http

import (
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := Resp{
		Code: http.StatusOK,
		Msg:  "server up",
	}
	respond(w, r, &resp)
}
