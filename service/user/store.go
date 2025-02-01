package user

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/MichaelGamel/ecom/types"
	"github.com/MichaelGamel/ecom/utils"
)

type Store struct {
	db    *sql.DB
	mysql sq.StatementBuilderType
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db,
		mysql: sq.StatementBuilder.PlaceholderFormat(sq.Question),
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	query, args, err := s.mysql.Select("*").From(utils.TablesConfig.Users).Where(sq.Eq{"email": email}).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil

}

func (s *Store) CreateUser(user types.User) error {
	stmt, args, err := s.mysql.Insert(utils.TablesConfig.Users).Columns("firstName", "lastName", "email", "password").Values(user.FirstName, user.LastName, user.Email, user.Password).ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(stmt, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	query, args, err := s.mysql.Select("*").From(utils.TablesConfig.Users).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
