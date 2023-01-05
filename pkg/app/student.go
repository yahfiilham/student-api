package app

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yahfiilham/student-api/pkg/domain"
)

// studentSvc implements domain.StudentSvc
type studentSvc struct {
	store domain.StudentStore
}

func NewStudentSvc(store domain.StudentStore) domain.StudentSvc {
	return &studentSvc{
		store: store,
	}
}

func (ss *studentSvc) List() ([]domain.Student, error) {
	return ss.store.List()
}

func (ss *studentSvc) Get(id uuid.UUID) (*domain.Student, error) {
	return ss.store.Get(id)
}

func (ss *studentSvc) Save(s *domain.Student) error {
	if err := validateSaveStudent(s); err != nil {
		return err
	}

	s.ID = uuid.New()
	return ss.store.Save(s)
}

func (ss *studentSvc) Update(id uuid.UUID, s *domain.Student) error {
	student, err := ss.Get(id)
	if err != nil {
		return err
	}

	validateUpdateStudent(student, s)

	s.ID = id
	return ss.store.Update(id, s)
}

func (ss *studentSvc) Delete(id uuid.UUID) error {
	if _, err := ss.Get(id); err != nil {
		return err
	}

	return ss.store.Delete(id)
}

func validateSaveStudent(s *domain.Student) error {
	if s.Name == "" {
		return errors.New("name cannot be empty")
	}

	if s.Age == 0 {
		return errors.New("age cannot be empty")
	}

	return nil
}

func validateUpdateStudent(s, r *domain.Student) {
	if r.Name == "" {
		r.Name = s.Name
	}

	if r.Age == 0 {
		r.Age = s.Age
	}
}
