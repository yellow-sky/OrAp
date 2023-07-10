package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/maltegrosse/go-modemmanager"
	"net/http"
	"path"
)

func (s ApiService) createModemsMiddleware(mmgr modemmanager.ModemManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				modemId := mux.Vars(r)["modem_id"]

				var modem modemmanager.Modem
				sysModems, err := mmgr.GetModems()
				if err != nil {
					resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get modem list: " + err.Error()}
					s.writeCommonJsonResponse(w, resp)
					return
				}
				for _, sysModem := range sysModems {
					sysModemId := path.Base(string(sysModem.GetObjectPath()))
					if modemId == sysModemId {
						modem = sysModem
						break
					}
				}
				if modem == nil {
					resp := CommonResponse{Status: http.StatusNotFound, Error: "Modem with such id not found"}
					s.writeCommonJsonResponse(w, resp)
					return
				}
				newCont := context.WithValue(r.Context(), "modem_id", modemId)
				newCont = context.WithValue(newCont, "modem", modem)
				newReq := r.WithContext(newCont)
				next.ServeHTTP(w, newReq)
			})
	}
}
