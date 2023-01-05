package http_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yahfiilham/student-api/pkg/app"
	ihttp "github.com/yahfiilham/student-api/pkg/http"
	"github.com/yahfiilham/student-api/pkg/store/memory"
)

func TestStudent(t *testing.T) {
	store := memory.NewStudentStore()
	svc := app.NewStudentSvc(store)
	h := ihttp.NewHandler(svc)
	r := ihttp.NewRoute(h)

	tests := []struct {
		name       string
		method     string
		uri        string
		payload    []byte
		headers    map[string]string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "save student success",
			method:     http.MethodPost,
			uri:        "/student",
			payload:    []byte(`{"name": "yahfi ilham","age": 20}`),
			headers:    nil,
			statusCode: http.StatusCreated,
		},
		{
			name:       "save student error empty payload",
			method:     http.MethodPost,
			uri:        "/student",
			payload:    nil,
			headers:    nil,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "save student error empty name",
			method:     http.MethodPost,
			uri:        "/student",
			payload:    []byte(`{"age": 20}`),
			headers:    nil,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "save student error empty age",
			method:     http.MethodPost,
			uri:        "/student",
			payload:    []byte(`{"name": "yahfi ilham"}`),
			headers:    nil,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "list student success",
			method:     http.MethodGet,
			uri:        "/student",
			payload:    nil,
			headers:    nil,
			statusCode: http.StatusOK,
		},
		{
			name:       "get student success",
			method:     http.MethodGet,
			uri:        "/student",
			payload:    nil,
			headers:    nil,
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p io.Reader
			if tt.payload != nil {
				p = bytes.NewBuffer(tt.payload)
			}

			req := httptest.NewRequest(tt.method, "http://localhost:8080"+tt.uri, p)

			record := httptest.NewRecorder()

			r.ServeHTTP(record, req)

			resp := record.Result()

			assert.NotEmpty(t, resp)
			assert.Equal(t, tt.statusCode, resp.StatusCode)
		})
	}
}
