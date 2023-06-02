package http_transport

import (
	"encoding/json"
	"fmt"
	"github.com/smallretardedfish/gql-federation/medication/transport/graph"
	"github.com/smallretardedfish/gql-federation/medication/transport/graph/model"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

func GetMedications(store graph.MedicationStore, patientServiceHost string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query()["limit"]
		var limit int64
		var err error
		if len(param) == 0 || param[0] == "" {
			limit = 20
		} else {
			limit, err = strconv.ParseInt(param[0], 10, 64)
		}
		if err != nil {
			return
		}

		includePatientsStr := r.URL.Query()["includePatients"]

		medications, err := store.GetMedications(r.Context(), &graph.MedicationFilter{Limit: &limit})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var includePatients bool
		if len(includePatientsStr) > 0 && includePatientsStr[0] != "" {
			includePatients, err = strconv.ParseBool(includePatientsStr[0])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)

				return
			}
		}
		enc := json.NewEncoder(w)
		if includePatients {
			res := make([]Medication, 0, len(medications))
			for _, medication := range medications {
				patient, err := getPatients(medication.ID, patientServiceHost)
				if err != nil {
					slog.Error("get patients from service:", err)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				res = append(res, MedicationMapper(medication, patient))
			}
			if err := enc.Encode(res); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				slog.Error("encoding medications with patients:", err)
				return
			}
			return
		}

		if err := enc.Encode(medications); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			slog.Error("marshaling medications with patients:", err)
			return
		}
	}
}

func MedicationMapper(medication *model.Medication, patient *Patient) Medication {
	mappedMedication := Medication{
		ID: medication.ID,
	}

	if medication.Code != nil {
		mappedMedication.Code = &Coding{
			System:  medication.Code.System,
			Code:    medication.Code.Code,
			Display: medication.Code.Display,
		}
	}

	if medication.Form != nil {
		mappedMedication.Form = &Coding{
			System:  medication.Form.System,
			Code:    medication.Form.Code,
			Display: medication.Form.Display,
		}
	}

	if medication.Manufacturer != nil {
		mappedMedication.Manufacturer = &Organization{
			ID: medication.Manufacturer.ID,
		}
	}

	if patient != nil {
		mappedMedication.Patient = &Patient{
			ID:        patient.ID,
			Name:      patient.Name,
			Gender:    patient.Gender,
			BirthDate: patient.BirthDate,
			Address:   patient.Address,
			Telecom:   patient.Telecom,
		}
	}

	return mappedMedication
}

func getPatients(id int64, patientServiceHost string) (*Patient, error) {
	url := fmt.Sprintf("http://%s/patients?id=%d", patientServiceHost, id)
	response, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var patient *Patient
	if err := json.NewDecoder(response.Body).Decode(&patient); err != nil {
		return nil, err
	}

	return patient, nil
}
