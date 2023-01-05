package http

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func respond(w http.ResponseWriter, code int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	bts, err := json.Marshal(resp)
	if err != nil {
		log.Err(err).Msgf("marshalling response to json failed")
		return
	}

	_, err = w.Write(bts)
	if err != nil {
		log.Err(err).Msgf("writing response to response-writer failed")
	}
}
