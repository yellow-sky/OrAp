package api

import (
	"context"
	"github.com/gorilla/mux"
	nm_dev_man "github.com/yellow-sky/orap/nm_device_manager"
	"net/http"
)

func (s ApiService) createDevicesMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				deviceId := mux.Vars(r)["device_id"]

				devManager, err := nm_dev_man.NewNmDeviceManager()
				if err != nil {
					resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on init device manager: " + err.Error()}
					s.writeCommonJsonResponse(w, resp)
					return
				}

				device, err := devManager.GetDeviceById(deviceId)
				if err != nil {
					resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get device by id: " + err.Error()}
					s.writeCommonJsonResponse(w, resp)
					return
				}

				if device == nil {
					resp := CommonResponse{Status: http.StatusNotFound, Error: "Device with such id not found"}
					s.writeCommonJsonResponse(w, resp)
					return
				}
				newCont := context.WithValue(r.Context(), "device_id", deviceId)
				newCont = context.WithValue(newCont, "device", device)
				newReq := r.WithContext(newCont)
				next.ServeHTTP(w, newReq)
			})
	}
}
