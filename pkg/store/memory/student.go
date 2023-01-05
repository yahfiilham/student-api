package memory

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/yahfiilham/student-api/pkg/domain"
)

// studentStore implements domain.StudentStore with an memory storage
type studentStore struct {
	mu      *sync.Mutex
	student map[uuid.UUID]*domain.Student
}

func NewStudentStore() domain.StudentStore {
	return &studentStore{
		mu:      &sync.Mutex{},
		student: make(map[uuid.UUID]*domain.Student),
	}
}

func (ss *studentStore) List() ([]domain.Student, error) {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	students := make([]domain.Student, 0, len(ss.student))
	for _, v := range ss.student {
		students = append(students, *v)
	}

	return students, nil
}

func (ss *studentStore) Get(id uuid.UUID) (*domain.Student, error) {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	if student, exist := ss.student[id]; exist {
		return student, nil
	}

	return nil, errors.New("no student found")
}

func (ss *studentStore) Save(s *domain.Student) error {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	ss.student[s.ID] = s
	return nil
}

func (ss *studentStore) Update(id uuid.UUID, s *domain.Student) error {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	ss.student[id] = s
	return nil
}

func (ss *studentStore) Delete(id uuid.UUID) error {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	delete(ss.student, id)
	return nil
}
