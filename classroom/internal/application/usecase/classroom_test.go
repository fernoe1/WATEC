package usecase

import (
	"context"
	"io"
	"log/slog"
	"testing"

	MOCK "github.com/fernoe1/WATEC/classroom/internal/application/usecase/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClassroomUsecase_Read(T *testing.T) {
	table := []struct {
		name      string
		prepare   func(*MOCK.ClassroomRepository, *MOCK.InMemoryRepository)
		result    []int64
		wantCache bool
		wantErr   bool
	}{
		{
			name: "read from db",
			prepare: func(r *MOCK.ClassroomRepository, imr *MOCK.InMemoryRepository) {
				r.On("Read", mock.Anything).Return([]int64{12227}, nil)
				imr.On("Get", mock.Anything).Return(nil, nil)
				imr.On("Set", mock.Anything, mock.AnythingOfType("[]int64")).Return(nil)
			},
			result: []int64{12227},
		},

		{
			name: "read from cache",
			prepare: func(r *MOCK.ClassroomRepository, imr *MOCK.InMemoryRepository) {
				imr.On("Get", mock.Anything).Return([]int64{12227}, nil)
			},
			result:    []int64{12227},
			wantCache: true,
		},
	}

	for _, t := range table {
		T.Run(t.name, func(T *testing.T) {
			r := new(MOCK.ClassroomRepository)
			imr := new(MOCK.InMemoryRepository)
			t.prepare(r, imr)

			uc := ClassroomUsecase{
				log: slog.New(slog.NewTextHandler(io.Discard, nil)),
				r:   r,
				imr: imr,
			}

			free, err := uc.Read(context.Background())

			if t.wantErr {
				require.Error(T, err)
			} else {
				require.NoError(T, err)
			}

			require.Equal(T, t.result, free)

			r.AssertExpectations(T)
			imr.AssertExpectations(T)
		})
	}
}
