package api

import (
	"encoding/json"
	"net/http"
)

type CommonResponse struct {
	Status int         `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func (s ApiService) writeCommonJsonResponse(w http.ResponseWriter, resp CommonResponse) {
	respJson, _ := json.Marshal(resp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	w.Write(respJson)
}
