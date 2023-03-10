package tom

import (
	"net/http"

	"github.com/xltank/tom/logger"
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
		logger.Log("Request in, r.URL:", r.URL)

		for i := 0; i < len(s.Routes); i++ {
			route := s.Routes[i]
			logger.Log("finding route:", r.Method, r.URL.Path, ",", route.Method, route.Path)
			if r.Method == route.Method && r.URL.Path == route.Path {
				route.Handles[0](w, r)
			}
		}

		// w.Write([]byte("Success"))
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
