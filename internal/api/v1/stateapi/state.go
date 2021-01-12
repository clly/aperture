package stateapi

import (
	"clly/apterture/pkg/storage"
	"io/ioutil"
	"log"
	"net/http"
)

type StateAPI struct {
	db storage.StateStorage
}

func NewStateAPI(db storage.StateStorage) *StateAPI {
	return &StateAPI{db: db}
}

func (s *StateAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		log.Printf("%#v", r)
		log.Println(string(b))
		// TODO: Figure out how this is named
		s.db.Put("workspace", b)
	case http.MethodGet:
		b := s.db.Get("workspace")
		log.Printf("%#v", r)
		w.Write(b)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
