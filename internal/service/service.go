package service

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"shorthack_backend/internal/entities"
)

type StudentServ interface {
	Create(ctx context.Context, studentCreate entities.CreateStudent) (int, string, error)
	Login(ctx context.Context, studentLogin entities.CreateStudent) (int, string, error)
	//Refresh(ctx context.Context, sessionID string, span trace.Span) (string, string, error)
	UpdatePassword(ctx context.Context, studentLogin entities.CreateStudent, newPassword string) (string, error)

	GetMe(ctx context.Context, studentID int, span trace.Span) (entities.Student, error)
	Delete(ctx context.Context, studentID int, sessionID string) error

	AddTag(ctx context.Context, studentID int, sessionID string, tag string) error
	AddMentor(ctx context.Context, mentorID int, studentID int, sessionID string) error
}
