package repository

import (
	"context"
	"shorthack_backend/internal/entities"
)

type MentorRepository interface {
	Create(ctx context.Context, student entities.CreateMentor) (int, error)
	Get(ctx context.Context, studentID int) (*entities.Mentor, error)
	GetPassword(ctx context.Context, login string) (int, string, error)
	UpdatePassword(ctx context.Context, studentID int, newPassword string) error
	Delete(ctx context.Context, studentID int) error
	AddTag(ctx context.Context, mentorID int, tagID int) error
}

type StudentRepository interface {
	Create(ctx context.Context, student entities.CreateStudent) (int, error)
	Get(ctx context.Context, studentID int) (*entities.Student, error)
	GetPassword(ctx context.Context, login string) (int, string, error)
	UpdatePassword(ctx context.Context, studentID int, newPassword string) error
	Delete(ctx context.Context, studentID int) error
	AddMentor(ctx context.Context, studentID int, mentor entities.Mentor) (int, error)
	AddTag(ctx context.Context, studentID int, tagID int) error
}
