package domain

import "github.com/google/uuid"

type Student struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  uint32    `json:"age"`
}

type StudentSvc interface {
	List() ([]Student, error)
	Get(id uuid.UUID) (*Student, error)
	Save(s *Student) error
	Update(id uuid.UUID, s *Student) error
	Delete(id uuid.UUID) error
}

type StudentStore interface {
	List() ([]Student, error)
	Get(id uuid.UUID) (*Student, error)
	Save(s *Student) error
	Update(id uuid.UUID, s *Student) error
	Delete(id uuid.UUID) error
}
