package mock

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yahfiilham/student-api/pkg/domain"
)

type StudentStore struct {
	GetStudentResp   *domain.Student
	GetStudentErr    error
	ListStudentResp  []domain.Student
	ListStudentErr   error
	SaveStudentErr   error
	UpdateStudentErr error
	DeleteStudentErr error
}

func (ss *StudentStore) List() ([]domain.Student, error) {
	m := []domain.Student{{
		ID:   uuid.UUID{},
		Name: "yahfi ilham",
		Age:  20,
	}}

	return m, ss.ListStudentErr
}

func (ss *StudentStore) Get(id uuid.UUID) (*domain.Student, error) {
	if id != uuid.Nil {
		return nil, errors.New("no found student")
	}

	m := &domain.Student{
		ID:   uuid.UUID{},
		Name: "yahfi ilham",
		Age:  20,
	}

	return m, ss.GetStudentErr
}

func (ss *StudentStore) Save(s *domain.Student) error {
	return ss.SaveStudentErr
}

func (ss *StudentStore) Update(id uuid.UUID, s *domain.Student) error {
	return ss.UpdateStudentErr
}

func (ss *StudentStore) Delete(id uuid.UUID) error {
	return ss.DeleteStudentErr
}
