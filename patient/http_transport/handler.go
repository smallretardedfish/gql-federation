package http_transport

import (
	"encoding/json"
	"github.com/smallretardedfish/gql-federation/patient/storage"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

func GetPatient(store storage.PatientStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query()["id"]
		var id int64
		var err error
		if param[0] == "" {
			id = 0
		} else {
			id, err = strconv.ParseInt(param[0], 10, 64)
		}
		if err != nil {
			slog.Error("parsing patient id:", err)
			return
		}

		pat, err := store.GetPatient(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bytes, err := json.Marshal(pat)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(bytes); err != nil {
			return
		}
	}
}
