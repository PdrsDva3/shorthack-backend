package mentor

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"shorthack_backend/internal/entities"
	"shorthack_backend/internal/repository"
)

type MentorRepository struct {
	db *sqlx.DB
}

func NewMentorRepo(db *sqlx.DB) repository.MentorRepository {
	return MentorRepository{
		db: db,
	}
}

func (Mentor MentorRepository) Create(ctx context.Context, user entities.CreateMentor) (int, error) {
	var id int
	transaction, err := Mentor.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowContext(ctx, `insert into mentor (name, login, hashed_password) values ($1, $2, $3, $4) returning id;`,
		user.Name, user.Login, user.Password)
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	if err := transaction.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

func (Mentor MentorRepository) Get(ctx context.Context, MentorID int) (*entities.Mentor, error) {
	var user entities.Mentor
	row := Mentor.db.QueryRowContext(ctx, `select name, login from mentor where id = $1`, MentorID)
	err := row.Scan(&user.Name, &user.Login)
	if err != nil {
		return nil, err
	}
	var mentors []int
	rows, err := Mentor.db.QueryContext(ctx, `select id_student from mentor_student where id_Mentor = $1`, MentorID)
	for rows.Next() {
		var idMentor int
		err = rows.Scan(&idMentor)
		if err != nil {
			return nil, err
		}
		mentors = append(mentors, idMentor)
	}
	user.StudentIDs = mentors

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (Mentor MentorRepository) GetPassword(ctx context.Context, login string) (int, string, error) {
	var id int
	var password string
	row := Mentor.db.QueryRowContext(ctx, `select id, hashed_password from mentor where login = $1`, login)
	err := row.Scan(&id, &password)
	if err != nil {
		return 0, "", err
	}
	return id, password, nil
}

func (Mentor MentorRepository) UpdatePassword(ctx context.Context, MentorID int, newPassword string) error {
	transaction, err := Mentor.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `UPDATE mentor SET hashed_password = $2 WHERE id = $1;`, MentorID, newPassword)
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

func (Mentor MentorRepository) Delete(ctx context.Context, MentorID int) error {
	transaction, err := Mentor.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `DELETE FROM mentor WHERE id=$1;`, MentorID)
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

func (Mentor MentorRepository) AddTag(ctx context.Context, MentorID int, tagID int) error {
	transaction, err := Mentor.db.BeginTxx(ctx, nil) //todo может стать проблемой
	if err != nil {
		return err
	}
	transaction.QueryRowContext(ctx, `insert into user_tag (id_mentor, id_student, id_tag) values ($1, $2, $3);`,
		MentorID, 0, tagID)

	if err := transaction.Commit(); err != nil {
		return err
	}
	return nil
}

func (mentor MentorRepository) AddNewTag(ctx context.Context, MentorID int, tag string) (int, error) {
	var tagID int
	transaction, err := mentor.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	row := transaction.QueryRowxContext(ctx, `INSERT INTO tags (text) values ($1) RETURNING id;`, tag)

	if err := row.Scan(&tagID); err != nil {
		return 0, err
	}

	transaction.QueryRowContext(ctx, `insert into user_tag (id_mentor, id_student, id_tag) values ($1, $2, $3);`,
		MentorID, 0, tagID)

	if err := transaction.Commit(); err != nil {
		return 0, err
	}
	return tagID, nil

}
