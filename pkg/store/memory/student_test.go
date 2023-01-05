package memory_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yahfiilham/student-api/pkg/domain"
	"github.com/yahfiilham/student-api/pkg/store/memory"
)

/*
 * list
 */

func Test_StudentStore_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		store := memory.NewStudentStore()

		data := &domain.Student{
			ID:   uuid.UUID{},
			Name: "yahfi ilham",
			Age:  20,
		}
		err := store.Save(data)
		assert.NoError(t, err)

		students, err := store.List()
		assert.NoError(t, err)
		assert.NotEmpty(t, students)
	})
}

/*
 * get
 */

func Test_StudentStore_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		store := memory.NewStudentStore()

		data := &domain.Student{
			ID:   uuid.UUID{},
			Name: "yahfi ilham",
			Age:  20,
		}
		err := store.Save(data)
		assert.NoError(t, err)

		students, err := store.Get(data.ID)
		assert.NoError(t, err)
		assert.NotEmpty(t, students)
	})

	t.Run("failed", func(t *testing.T) {
		store := memory.NewStudentStore()

		students, err := store.Get(uuid.UUID{})
		assert.NotEmpty(t, err)
		assert.Empty(t, students)
	})
}

/*
 * save
 */

func Test_StudentStore_Save(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		store := memory.NewStudentStore()

		data := &domain.Student{
			ID:   uuid.UUID{},
			Name: "yahfi ilham",
			Age:  20,
		}
		err := store.Save(data)
		assert.NoError(t, err)
	})
}

/*
 * update
 */

func Test_StudentStore_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		store := memory.NewStudentStore()

		data := &domain.Student{
			ID:   uuid.UUID{},
			Name: "yahfi ilham",
			Age:  20,
		}
		err := store.Update(data.ID, data)
		assert.NoError(t, err)
	})
}

/*
 * delete
 */

func Test_StudentStore_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		store := memory.NewStudentStore()

		err := store.Delete(uuid.UUID{})
		assert.NoError(t, err)
	})
}
