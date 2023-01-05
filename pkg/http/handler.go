package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/yahfiilham/student-api/pkg/domain"
)

type Handler struct {
	StudentSvc domain.StudentSvc
}

func NewHandler(studentSvc domain.StudentSvc) *Handler {
	return &Handler{
		StudentSvc: studentSvc,
	}
}

/*
 * health check
 */
func (h *Handler) GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, "server up", nil)
}

/*
 * student
 */
func (h *Handler) ListStudent(w http.ResponseWriter, r *http.Request) {
	student, err := h.StudentSvc.List()
	if err != nil {
		log.Err(err).Msgf("error retrieve list student")
		respond(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	respond(w, http.StatusOK, "success", student)
}

func (h *Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	log.Info().Str("student id", id)

	studentID, err := uuid.Parse(id)
	if err != nil {
		log.Err(err).Msgf("error parse student id")
		respond(w, http.StatusBadRequest, "invalid student id provided in url param", nil)
		return
	}

	if studentID == uuid.Nil {
		respond(w, http.StatusBadRequest, "please provide the student id to retrieve", nil)
		return
	}

	student, err := h.StudentSvc.Get(studentID)
	if err != nil {
		log.Err(err).Msgf("error retrieve student")
		respond(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	respond(w, http.StatusOK, "success", student)
}

func (h *Handler) AddStudent(w http.ResponseWriter, r *http.Request) {
	var data domain.Student
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Err(err).Msgf("error decode request body")
		respond(w, http.StatusInternalServerError, "error while decode request body", nil)
		return
	}

	if err := h.StudentSvc.Save(&data); err != nil {
		log.Err(err).Msgf("error to add student")
		respond(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	respond(w, http.StatusCreated, "success", data)
}

func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	log.Info().Str("student id", id)

	var data domain.Student
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Err(err).Msgf("error decode request body")
		respond(w, http.StatusInternalServerError, "error while decode request body", nil)
		return
	}

	studentID, err := uuid.Parse(id)
	if err != nil {
		log.Err(err).Msgf("error parse student id")
		respond(w, http.StatusBadRequest, "invalid student id provided in url param", nil)
		return
	}

	if studentID == uuid.Nil {
		respond(w, http.StatusBadRequest, "please provide the student id to retrieve", nil)
		return
	}

	if err := h.StudentSvc.Update(studentID, &data); err != nil {
		log.Err(err).Msgf("error to update student")
		respond(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	respond(w, http.StatusCreated, "success", data)
}

func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	log.Info().Str("student id", id)

	studentID, err := uuid.Parse(id)
	if err != nil {
		log.Err(err).Msgf("error parse student id")
		respond(w, http.StatusBadRequest, "invalid student id provided in url param", nil)
		return
	}

	if studentID == uuid.Nil {
		respond(w, http.StatusBadRequest, "please provide the student id to retrieve", nil)
		return
	}

	if err := h.StudentSvc.Delete(studentID); err != nil {
		log.Err(err).Msgf("error to delete student")
		respond(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	respond(w, http.StatusOK, fmt.Sprintf("success delete student with id: %v", studentID), nil)
}
