package student

import (
	"context"
	"github.com/jmoiron/sqlx"
	"shorthack_backend/internal/entities"
	"shorthack_backend/internal/repository"
)

type StudentRepository struct {
	db *sqlx.DB
}

func NewStudentRepo(db *sqlx.DB) repository.StudentRepository {
	return StudentRepository{
		db: db,
	}
}

func (student StudentRepository) Creat(ctx context.Context, user entities.CreateStudent) (int, error) {
	var id int
	transaction, err := student.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowContext(ctx, `insert into student (name, login, hashed_password, level) values ($1, $2, $3, $4) returning id;`,
		user.Name, user.Login, user.Password, user.Level)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	if err := transaction.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

func (student StudentRepository) Get(ctx context.Context, studentID int) (*entities.Student, error) {
	var user entities.Student
	row := student.db.QueryRowContext(ctx, `select name, login, level from student where id = $1`, studentID)
	err := row.Scan(&user.Name, &user.Login, &user.Level)
	if err != nil {
		return nil, err
	}
	var mentors []int
	rows := student.db.QueryRowxContext(ctx, `select id_mentor from mentor_student where id_student = $1`, studentID)
	err = rows.Scan(&mentors)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//GetPassword(ctx context.Context, login string) (int, string, error)
//UpdatePassword(ctx context.Context, studentID int, newPassword string) error
//Delete(ctx context.Context, studentID int) error
//AddMentor(ctx context.Context, studentID int, mentor entities.Mentor) (int, error)
