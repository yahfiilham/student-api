package http

import (
	"net/http"
	"strconv"
	"strings"
)

type Route struct {
	Handler *Handler
}

func NewRoute(h *Handler) *Route {
	return &Route{Handler: h}
}

func (rt *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var id string
	p := r.URL.Path

	switch {
	case match(p, "/health") && r.Method == http.MethodGet:
		rt.Handler.GetHealthCheck(w, r)
		return
	case match(p, "/student") && r.Method == http.MethodGet:
		rt.Handler.ListStudent(w, r)
		return
	case match(p, "/student") && r.Method == http.MethodPost:
		rt.Handler.AddStudent(w, r)
		return
	case match(p, "/student/+", &id) && r.Method == http.MethodGet:
		rt.Handler.GetStudent(w, r)
		return
	case match(p, "/student/+", &id) && r.Method == http.MethodPut:
		rt.Handler.UpdateStudent(w, r)
		return
	case match(p, "/student/+", &id) && r.Method == http.MethodDelete:
		rt.Handler.DeleteStudent(w, r)
		return
	default:
		respond(w, http.StatusNotFound, "service not found", nil)
		return
	}
}

func match(path, pattern string, vars ...interface{}) bool {
	for ; pattern != "" && path != ""; pattern = pattern[1:] {
		switch pattern[0] {
		case '+':
			// '+' matches till next slash in path
			slash := strings.IndexByte(path, '/')
			if slash < 0 {
				slash = len(path)
			}

			segment := path[:slash]
			path = path[slash:]

			switch p := vars[0].(type) {
			case *string:
				*p = segment
			case *int:
				n, err := strconv.Atoi(segment)
				if err != nil || n < 0 {
					return false
				}
				*p = n
			default:
				panic("vars must be *string or *int")
			}

			vars = vars[1:]
		case path[0]:
			// non-'+' pattern byte must match path byte
			path = path[1:]
		default:
			return false
		}
	}

	return path == "" && pattern == ""
}
