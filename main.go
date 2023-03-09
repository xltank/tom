package tom

import (
	"fmt"
	"net/http"

	"github.com/xltank/tom/model"
)

func New() *Tom {
	return new(Tom)
}

type Tom struct {
	Address string
	Routes  []model.Route
}

func (s *Tom) Listen(address string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request in, ", r.URL)
		w.Write([]byte("Success"))
	})
	http.ListenAndServe(address, nil)
}

func (s *Tom) Get(path string, handFunc ...http.HandlerFunc) {
	s.Routes = append(s.Routes, model.Route{
		Path:    path,
		Method:  http.MethodGet,
		Handles: handFunc,
	})
}
