package student

import (
	"context"
	"errors"
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

func (student StudentRepository) Create(ctx context.Context, user entities.CreateStudent) (int, error) {
	var id int
	transaction, err := student.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowContext(ctx, `insert into student (name, login, hashed_password, level, tg) values ($1, $2, $3, $4, $5) returning id;`,
		user.Name, user.Login, user.Password, user.Level, user.TG)
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
	row := student.db.QueryRowContext(ctx, `select name, login, level, tg from student where id = $1`, studentID)
	err := row.Scan(&user.Name, &user.Login, &user.Level, &user.TG)
	if err != nil {
		return nil, err
	}
	var mentors []int
	rows, err := student.db.QueryContext(ctx, `select id_mentor from mentor_student where id_student = $1`, studentID)
	for rows.Next() {
		var idMentor int
		err = rows.Scan(&idMentor)
		if err != nil {
			return nil, err
		}
		mentors = append(mentors, idMentor)
	}
	user.MentorIds = mentors

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (student StudentRepository) GetPassword(ctx context.Context, login string) (int, string, error) {
	var id int
	var password string
	row := student.db.QueryRowContext(ctx, `select id, hashed_password from student where login = $1`, login)
	err := row.Scan(&id, &password)
	if err != nil {
		return 0, "", err
	}
	return id, password, nil
}

func (student StudentRepository) UpdatePassword(ctx context.Context, studentID int, newPassword string) error {
	transaction, err := student.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	result, err := transaction.ExecContext(ctx, `UPDATE student SET hashed_password = $2 WHERE id = $1;`, studentID, newPassword)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return errors.New("failed rollback")
		}
		return errors.New("failed to update password")
	}
	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (student StudentRepository) Delete(ctx context.Context, studentID int) error {
	transaction, err := student.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `DELETE FROM student WHERE id=$1;`, studentID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return errors.New("failed rollback")
		}
		return errors.New("failed to update password")
	}
	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (student StudentRepository) AddMentor(ctx context.Context, studentID int, mentorID int) error {
	transaction, err := student.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	transaction.QueryRowContext(ctx, `insert into mentor_student (id_mentor, id_student) values ($1, $2);`,
		mentorID, studentID)

	if err := transaction.Commit(); err != nil {
		return err
	}
	return nil
}

func (student StudentRepository) AddTag(ctx context.Context, studentID int, tagID int) error {
	transaction, err := student.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	transaction.QueryRowContext(ctx, `insert into user_tag (id_mentor, id_student, id_tag) values ($1, $2, $3);`,
		0, studentID, tagID)

	if err := transaction.Commit(); err != nil {
		return err
	}
	return nil
}
