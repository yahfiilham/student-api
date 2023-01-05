package app_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yahfiilham/student-api/pkg/app"
	"github.com/yahfiilham/student-api/pkg/domain"
	"github.com/yahfiilham/student-api/test/mock"
)

/*
 * list
 */

func Test_StudentSvc_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var store mock.StudentStore
		svc := app.NewStudentSvc(&store)

		students, err := svc.List()
		assert.NoError(t, err)
		assert.NotEmpty(t, students)
	})
}

/*
 * get
 */
func Test_StudentSvc_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var store mock.StudentStore
		svc := app.NewStudentSvc(&store)

		s, err := svc.Get(uuid.UUID{})
		assert.NoError(t, err)
		assert.NotEmpty(t, s)
	})
}

/*
 * create
 */

func Test_StudentSvc_Save(t *testing.T) {
	type data struct {
		s *domain.Student
	}

	var store mock.StudentStore
	svc := app.NewStudentSvc(&store)

	tests := []struct {
		name    string
		data    data
		wantErr bool
	}{
		{
			name: "success",
			data: data{
				s: &domain.Student{
					Name: "yahfi ilham",
					Age:  20,
				},
			},
			wantErr: false,
		},
		{
			name:    "error",
			data:    data{s: &domain.Student{}},
			wantErr: true,
		},
		{
			name: "error empty name",
			data: data{s: &domain.Student{
				Age: 20,
			}},
			wantErr: true,
		},
		{
			name: "error empty age",
			data: data{s: &domain.Student{
				Name: "yahfi ilham",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Save(tt.data.s)
			if tt.wantErr {
				assert.NotEmpty(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

/*
 * update
 */

func Test_StudentSvc_Update(t *testing.T) {
	type data struct {
		id uuid.UUID
		s  *domain.Student
	}

	var store mock.StudentStore
	svc := app.NewStudentSvc(&store)

	tests := []struct {
		name    string
		data    data
		wantErr bool
	}{
		{
			name: "success",
			data: data{
				id: uuid.UUID{},
				s: &domain.Student{
					Name: "yahfi ilham",
					Age:  20,
				},
			},
			wantErr: false,
		},
		{
			name: "success without data",
			data: data{
				id: uuid.UUID{},
				s:  &domain.Student{},
			},
			wantErr: false,
		},
		{
			name: "success without name",
			data: data{
				id: uuid.UUID{},
				s: &domain.Student{
					Age: 20,
				},
			},
			wantErr: false,
		},
		{
			name: "success without age",
			data: data{
				id: uuid.UUID{},
				s: &domain.Student{
					Name: "yahfi ilham",
				},
			},
			wantErr: false,
		},
		{
			name: "failed student not found",
			data: data{
				id: uuid.NameSpaceDNS,
				s: &domain.Student{
					Name: "yahfi ilham",
					Age:  20,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Update(tt.data.id, tt.data.s)
			if tt.wantErr {
				assert.NotEmpty(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

/*
 * delete
 */

func Test_StudentSvc_Delete(t *testing.T) {
	type data struct {
		id uuid.UUID
	}

	var store mock.StudentStore
	svc := app.NewStudentSvc(&store)

	tests := []struct {
		name    string
		data    data
		wantErr bool
	}{
		{
			name: "success",
			data: data{
				id: uuid.UUID{},
			},
			wantErr: false,
		},
		{
			name: "failed student not found",
			data: data{
				id: uuid.NameSpaceDNS,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Delete(tt.data.id)
			if tt.wantErr {
				assert.NotEmpty(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
