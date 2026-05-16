package grpc

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/fernoe1/WATEC/teacher/internal/domain"
	tchersvc "github.com/fernoe1/protogen/watec/service/teacher"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fakeTeacherUsecase struct {
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

func (f *fakeTeacherUsecase) Create(ctx context.Context, teacher *domain.Teacher) error {
	f.createCalled = true
	f.lastTeacher = teacher
	return f.createErr
}

func (f *fakeTeacherUsecase) Read(ctx context.Context, name string) (error, *domain.Free) {
	f.readCalled = true
	f.lastName = name
	return f.readErr, f.readFree
}

func (f *fakeTeacherUsecase) Update(ctx context.Context, teacher *domain.Teacher) (error, *domain.Teacher) {
	f.updateCalled = true
	f.lastTeacher = teacher
	if f.updatedTeacher != nil {
		return f.updateErr, f.updatedTeacher
	}
	return f.updateErr, teacher
}

func (f *fakeTeacherUsecase) Delete(ctx context.Context, name string) error {
	f.deleteCalled = true
	f.lastName = name
	return f.deleteErr
}

func TestTeacherServiceCreateMapsRequestToDomain(t *testing.T) {
	uc := &fakeTeacherUsecase{}
	svc := NewTeacherService(uc)

	_, err := svc.Create(context.Background(), &tchersvc.CreateRequest{
		Name: "John",
		Free: []*tchersvc.Free{{RoomNumber: 101, From: 9, To: 11}},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !uc.createCalled {
		t.Fatal("usecase Create was not called")
	}
	if uc.lastTeacher == nil || uc.lastTeacher.Name != "John" || len(uc.lastTeacher.Free) != 1 {
		t.Fatalf("unexpected teacher mapping: %#v", uc.lastTeacher)
	}
	free := uc.lastTeacher.Free[0]
	if free.TeacherName != "John" || free.RoomNumber != 101 || free.From != 9 || free.To != 11 {
		t.Fatalf("unexpected free slot mapping: %#v", free)
	}
}

func TestTeacherServiceReadReturnsIsFreeTrueWithSlot(t *testing.T) {
	uc := &fakeTeacherUsecase{readFree: &domain.Free{RoomNumber: 202, From: 12, To: 14}}
	svc := NewTeacherService(uc)

	resp, err := svc.Read(context.Background(), &tchersvc.ReadRequest{Name: "Alice"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !resp.GetIsFree() {
		t.Fatal("expected IsFree to be true")
	}
	if resp.GetFree().GetRoomNumber() != 202 {
		t.Fatalf("unexpected response free slot: %#v", resp.GetFree())
	}
}

func TestTeacherServiceReadReturnsIsFreeFalseWithoutSlot(t *testing.T) {
	uc := &fakeTeacherUsecase{}
	svc := NewTeacherService(uc)

	resp, err := svc.Read(context.Background(), &tchersvc.ReadRequest{Name: "Alice"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if resp.GetIsFree() {
		t.Fatal("expected IsFree to be false")
	}
	if resp.GetFree() != nil {
		t.Fatalf("expected nil free slot, got %#v", resp.GetFree())
	}
}

func TestTeacherServiceConvertsUsecaseErrorToGrpcInternal(t *testing.T) {
	uc := &fakeTeacherUsecase{readErr: errors.New("database down")}
	svc := NewTeacherService(uc)

	_, err := svc.Read(context.Background(), &tchersvc.ReadRequest{Name: "Alice"})
	if err == nil {
		t.Fatal("expected error")
	}
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected grpc status error, got %T", err)
	}
	if st.Code() != codes.Internal {
		t.Fatalf("expected Internal code, got %v", st.Code())
	}
	if !strings.Contains(st.Message(), "database down") {
		t.Fatalf("unexpected status message: %s", st.Message())
	}
}

func TestTeacherServiceDeletePassesTeacherName(t *testing.T) {
	uc := &fakeTeacherUsecase{}
	svc := NewTeacherService(uc)

	_, err := svc.Delete(context.Background(), &tchersvc.DeleteRequest{Name: "Eve"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !uc.deleteCalled || uc.lastName != "Eve" {
		t.Fatalf("delete was not called correctly, name=%q", uc.lastName)
	}
}
