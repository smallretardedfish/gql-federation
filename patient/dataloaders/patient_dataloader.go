package dataloaders

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/smallretardedfish/gql-federation/patient/graph/model"
	"github.com/smallretardedfish/gql-federation/patient/storage"
	"net/http"
	"strconv"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type PatientReader struct {
	patientStore storage.PatientStore
}

type Loaders struct {
	PatientLoader *dataloader.Loader
}

func NewLoaders(store storage.PatientStore) *Loaders {
	// define the data loader
	patientReader := &PatientReader{patientStore: store}
	loaders := &Loaders{
		PatientLoader: dataloader.NewBatchedLoader(patientReader.GetPatients),
	}
	return loaders
}

func (r *PatientReader) GetPatients(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids := make([]string, len(keys))
	filterIds := make([]int64, len(keys))
	for ix, key := range keys {
		ids[ix] = key.String()
		filterIds[ix], _ = strconv.ParseInt(key.String(), 10, 64)
	}

	patients, err := r.patientStore.GetPatients(ctx, &storage.PatientFilter{IDs: filterIds})
	if err != nil {
		return nil
	}

	patientMap := make(map[string]*model.Patient, len(patients))
	for i := range patients {
		patientMap[strconv.FormatInt(patients[i].ID, 10)] = patients[i]
	}

	output := make([]*dataloader.Result, len(ids))
	for index, key := range keys {
		user, ok := patientMap[key.String()]
		if ok {
			output[index] = &dataloader.Result{Data: user, Error: nil}
		} else {
			err := fmt.Errorf("user not found %d", key.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}

	return output
}

// Middleware injects data loaders into the context
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func GetPatient(ctx context.Context, userID int64) (*model.Patient, error) {
	loaders := For(ctx)
	thunk := loaders.PatientLoader.Load(ctx, dataloader.StringKey(strconv.FormatInt(userID, 10)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.Patient), nil
}
