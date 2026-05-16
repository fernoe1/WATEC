package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/fernoe1/WATEC/teacher/internal/domain"
)

type fakeTeacherRepository struct {
	createCalled bool
	readCalled   bool
	updateCalled bool
	deleteCalled bool

	createErr error
	readErr   error
	updateErr error
	deleteErr error

	readFree       *domain.Free
	updatedTeacher *domain.Teacher
	lastTeacher    *domain.Teacher
	lastName       string
}

func (f *fakeTeacherRepository) Create(ctx context.Context, teacher *domain.Teacher) error {
	f.createCalled = true
	f.lastTeacher = teacher
	return f.createErr
}

func (f *fakeTeacherRepository) Read(ctx context.Context, name string) (error, *domain.Free) {
	f.readCalled = true
	f.lastName = name
	return f.readErr, f.readFree
}

func (f *fakeTeacherRepository) Update(ctx context.Context, teacher *domain.Teacher) (error, *domain.Teacher) {
	f.updateCalled = true
	f.lastTeacher = teacher
	if f.updatedTeacher != nil {
		return f.updateErr, f.updatedTeacher
	}
	return f.updateErr, teacher
}

func (f *fakeTeacherRepository) Delete(ctx context.Context, name string) error {
	f.deleteCalled = true
	f.lastName = name
	return f.deleteErr
}

type fakeInMemoryTeacherRepository struct {
	getCalled bool
	setCalled bool

	getErr error
	setErr error

	cachedFree  *domain.Free
	lastName    string
	lastTeacher *domain.Teacher
}

func (f *fakeInMemoryTeacherRepository) Get(ctx context.Context, name string) (error, *domain.Free) {
	f.getCalled = true
	f.lastName = name
	return f.getErr, f.cachedFree
}

func (f *fakeInMemoryTeacherRepository) Set(ctx context.Context, teacher *domain.Teacher) error {
	f.setCalled = true
	f.lastTeacher = teacher
	return f.setErr
}

func TestTeacherUsecaseReadReturnsCachedFreeSlot(t *testing.T) {
	ctx := context.Background()
	repo := &fakeTeacherRepository{}
	cache := &fakeInMemoryTeacherRepository{cachedFree: &domain.Free{RoomNumber: 101, From: 9, To: 11}}
	uc := NewTeacherUsecase(repo, cache)

	err, free := uc.Read(ctx, "John")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if free == nil || free.RoomNumber != 101 {
		t.Fatalf("expected cached free slot, got %#v", free)
	}
	if repo.readCalled {
		t.Fatal("database repository should not be called when cache has value")
	}
}

func TestTeacherUsecaseReadFallsBackToDatabaseAndCachesResult(t *testing.T) {
	ctx := context.Background()
	repo := &fakeTeacherRepository{readFree: &domain.Free{RoomNumber: 202, From: 12, To: 13}}
	cache := &fakeInMemoryTeacherRepository{}
	uc := NewTeacherUsecase(repo, cache)

	err, free := uc.Read(ctx, "Alice")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if free == nil || free.RoomNumber != 202 {
		t.Fatalf("expected database free slot, got %#v", free)
	}
	if !repo.readCalled {
		t.Fatal("database repository should be called on cache miss")
	}
	if !cache.setCalled {
		t.Fatal("cache should be refreshed after database hit")
	}
	if cache.lastTeacher == nil || cache.lastTeacher.Name != "Alice" || len(cache.lastTeacher.Free) != 1 {
		t.Fatalf("unexpected cached teacher: %#v", cache.lastTeacher)
	}
}

func TestTeacherUsecaseCreateStopsWhenDatabaseFails(t *testing.T) {
	ctx := context.Background()
	expectedErr := errors.New("db failed")
	repo := &fakeTeacherRepository{createErr: expectedErr}
	cache := &fakeInMemoryTeacherRepository{}
	uc := NewTeacherUsecase(repo, cache)

	err := uc.Create(ctx, &domain.Teacher{Name: "Bob"})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
	if cache.setCalled {
		t.Fatal("cache should not be updated when database create fails")
	}
}

func TestTeacherUsecaseUpdateRefreshesCache(t *testing.T) {
	ctx := context.Background()
	updated := &domain.Teacher{Name: "Dana", Free: []domain.Free{{RoomNumber: 303, From: 10, To: 12}}}
	repo := &fakeTeacherRepository{updatedTeacher: updated}
	cache := &fakeInMemoryTeacherRepository{}
	uc := NewTeacherUsecase(repo, cache)

	err, teacher := uc.Update(ctx, updated)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if teacher != updated {
		t.Fatal("expected updated teacher to be returned")
	}
	if !cache.setCalled || cache.lastTeacher != updated {
		t.Fatalf("cache was not refreshed with updated teacher")
	}
}

func TestTeacherUsecaseDeleteClearsCache(t *testing.T) {
	ctx := context.Background()
	repo := &fakeTeacherRepository{}
	cache := &fakeInMemoryTeacherRepository{}
	uc := NewTeacherUsecase(repo, cache)

	if err := uc.Delete(ctx, "Eve"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !repo.deleteCalled || repo.lastName != "Eve" {
		t.Fatal("database delete was not called correctly")
	}
	if !cache.setCalled || cache.lastTeacher == nil || cache.lastTeacher.Name != "Eve" {
		t.Fatalf("cache clear was not requested correctly: %#v", cache.lastTeacher)
	}
}
