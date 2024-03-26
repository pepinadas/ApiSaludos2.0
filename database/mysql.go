package database

import (
	"RestGolang/models"
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(url string) (*MySQLRepository, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return &MySQLRepository{db}, nil
}

func (repo *MySQLRepository) InsertGreet(ctx context.Context, greet *models.Greet) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO greets (id, greet_content, lenguage, user_id) VALUES (?, ?, ?, ?);", greet.Id, greet.GreetingContent, greet.Lenguage, greet.UserId)
	return err
}

func (repo *MySQLRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES (?, ?, ?);", user.Id, user.Email, user.Password)
	return err
}

func (repo *MySQLRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = ? ", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *MySQLRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = ? ", email)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *MySQLRepository) Close() error {
	return repo.db.Close()
}

func (repo *MySQLRepository) GetGreetById(ctx context.Context, id string) (*models.Greet, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, greet_content,lenguage,  created_at, user_id  FROM greets WHERE id = ? ", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var greet = models.Greet{}
	for rows.Next() {
		if err = rows.Scan(&greet.Id, &greet.GreetingContent, &greet.Lenguage, &greet.CreatedAt, &greet.UserId); err == nil {
			return &greet, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &greet, nil
}

func (repo *MySQLRepository) UpdateGreet(ctx context.Context, greet *models.Greet) error {
	_, err := repo.db.ExecContext(ctx, " UPDATE greets SET greet_content = ? WHERE id = ? and user_id = ? ", greet.GreetingContent, greet.Id, greet.UserId)
	return err
}

func (repo *MySQLRepository) DeleteGreet(ctx context.Context, id string, userId string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM greets WHERE id = ? and user_id = ?", id, userId)
	return err
}

func (repo *MySQLRepository) ListGreet(ctx context.Context, page uint64) ([]*models.Greet, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, greet_content, user_id, created_at FROM greets LIMIT ? OFFSET ?", 5, page*2)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var greets []*models.Greet
	for rows.Next() {
		var greet = models.Greet{}
		if err = rows.Scan(&greet.Id, &greet.GreetingContent, &greet.UserId, &greet.CreatedAt); err == nil {
			greets = append(greets, &greet)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return greets, nil
}
