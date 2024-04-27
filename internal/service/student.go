package service

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/crypto/bcrypt"
	"shorthack_backend/internal/entities"
	"shorthack_backend/internal/repository"
)

type StudentService struct {
	StudentRepository repository.StudentRepository
}

func (usrs StudentService) Login(ctx context.Context, studentLogin entities.CreateStudent) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (usrs StudentService) GetMe(ctx context.Context, studentID int, span trace.Span) (entities.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (usrs StudentService) AddTag(ctx context.Context, studentID int, tag int) error {
	err := usrs.StudentRepository.AddTag(ctx, studentID, tag)
	if err != nil {
		return err
	}
	return nil
}

func (usrs StudentService) AddMentor(ctx context.Context, mentorID int, studentID int, sessionID string) error {
	//TODO implement me
	panic("implement me")
}

func InitStudentService(StudentRepository repository.StudentRepository) StudentServ {
	return &StudentService{StudentRepository: StudentRepository}
}

func (usrs StudentService) Create(ctx context.Context, student entities.CreateStudent) (int, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(student.Password), 10)
	if err != nil {
		return 0, nil
	}

	newStudent := entities.CreateStudent{
		StudentBase: student.StudentBase,
		Password:    string(hashed_password),
	}

	id, err := usrs.StudentRepository.Create(ctx, newStudent)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (usrs StudentService) Get(ctx context.Context, id int) (*entities.Student, error) {
	Student, err := usrs.StudentRepository.Get(ctx, id)
	if err != nil {
		return &entities.Student{}, err
	}
	if Student == nil {
		return nil, err
	}

	return Student, nil

}

func (usrs StudentService) GetPassword(ctx context.Context, login string) (int, string, error) {
	id, password, err := usrs.StudentRepository.GetPassword(ctx, login)
	if err != nil {
		return 0, "", err
	}

	return id, password, nil
}

func (usrs StudentService) UpdatePassword(ctx context.Context, StudentID int, newPassword string) error {
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 10)

	err := usrs.StudentRepository.UpdatePassword(ctx, StudentID, string(hashed_password))
	if err != nil {
		return err
	}

	return nil
}

func (usrs StudentService) Delete(ctx context.Context, StudentID int) error {
	err := usrs.StudentRepository.Delete(ctx, StudentID)
	if err != nil {
		return err
	}

	return nil
}
