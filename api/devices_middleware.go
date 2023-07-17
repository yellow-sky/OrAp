package api

import (
	"context"
	"github.com/Wifx/gonetworkmanager/v2"
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

func (s ApiService) createDevicesMiddleware(nmgr gonetworkmanager.NetworkManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				deviceId := mux.Vars(r)["device_id"]

				var device gonetworkmanager.Device
				devices, err := nmgr.GetPropertyAllDevices()
				if err != nil {
					resp := CommonResponse{Status: http.StatusInternalServerError, Error: "Error on get devices: " + err.Error()}
					s.writeCommonJsonResponse(w, resp)
					return
				}
				for _, sysDevice := range devices {
					sysDeviceId := path.Base(string(sysDevice.GetPath()))
					if deviceId == sysDeviceId {
						device = sysDevice
						break
					}
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
