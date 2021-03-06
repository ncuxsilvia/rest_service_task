package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rest_service_task/impl/structs"
)

type DBConnection interface {
	Open(dbname string, host string, port uint, user string, pass string) error
	CreateUser(firstName string, lastName string, login string, passHash string, age int, phone int64) error
	AuthUser(login string, passHash string) (*structs.User, error)
	GetFilms(pageSize int, pageNumber int, genre *string, releaseYear *int) (*[]structs.Film, error)
	GetFilmsByUser(id int64, pageSize int, pageNumber int) (*[]structs.Film, error)
	StartRent(login string, filmID int) error
	EndRent(filmID int, userID int64) error
	Close() error
}

type DB struct {
	connection *sql.DB
}

func NewDB() DBConnection {
	return &DB{nil}
}

func (db *DB) Open(dbname string, host string, port uint, user string, pass string) error {
	var err error
	connString := fmt.Sprintf("user=%s dbname=%s host=%s port=%d password=%s sslmode=disable", user, dbname, host, port, pass)

	db.connection, err = sql.Open("postgres", connString)
	if err != nil {
		return err
	}

	err = db.connection.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) CreateUser(firstName string, lastName string, login string, passHash string, age int, phone int64) error {
	if db.connection == nil {
		return fmt.Errorf("Database connection not opened.")
	}

	_, err := db.connection.Query("SELECT public.create_user($1::VARCHAR, $2::VARCHAR, $3::VARCHAR, $4::VARCHAR, $5::INT, $6::BIGINT);", firstName, lastName, login, passHash, age, phone)
	return err
}

func (db *DB) AuthUser(login string, passHash string) (*structs.User, error) {
	if db.connection == nil {
		return nil, fmt.Errorf("Database connection not opened.")
	}

	// result, err := db.connection.Query("SELECT public.auth_user($1::VARCHAR, $2::VARCHAR)", login, passHash)
	// if err != nil {
	// 	return nil, err
	// }

	var user structs.User
	err := db.connection.QueryRow("SELECT * FROM public.auth_user($1::VARCHAR, $2::VARCHAR)", login, passHash).
		Scan(&user.ID, &user.Login, &user.FirstName, &user.LastName, &user.Age, &user.Phone)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) GetFilms(pageSize int, pageNumber int, genre *string, releaseYear *int) (*[]structs.Film, error) {
	if db.connection == nil {
		return nil, fmt.Errorf("Database connection not opened.")
	}

	result, err := db.connection.Query("SELECT * FROM public.get_films($1::INT, $2::INT, $3::VARCHAR, $4::INT)", pageSize, pageNumber, genre, releaseYear)

	if err != nil {
		return nil, err
	}

	films := []structs.Film{}
	for result.Next() {
		var film structs.Film
		err = result.Scan(&film.ID, &film.Title, &film.Genres, &film.ReleaseYear)
		if err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	return &films, nil
}

func (db *DB) GetFilmsByUser(id int64, pageSize int, pageNumber int) (*[]structs.Film, error) {
	if db.connection == nil {
		return nil, fmt.Errorf("Database connection not opened.")
	}

	result, err := db.connection.Query("SELECT * FROM public.rented_by_user_id($1::BIGINT, $2::INT, $3::INT)", id, pageSize, pageNumber)

	if err != nil {
		return nil, err
	}

	films := []structs.Film{}
	for result.Next() {
		var film structs.Film
		err = result.Scan(&film.ID, &film.Title, &film.Genres, &film.ReleaseYear)
		if err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	return &films, nil
}

func (db *DB) StartRent(login string, filmID int) error {
	if db.connection == nil {
		return fmt.Errorf("Database connection not opened.")
	}

	_, err := db.connection.Query("SELECT * FROM public.start_rent($1::VARCHAR, $2::INT)", login, filmID)
	return err
}

func (db *DB) EndRent(filmID int, userID int64) error {
	if db.connection == nil {
		return fmt.Errorf("Database connection not opened.")
	}

	_, err := db.connection.Query("SELECT public.end_rent($1::INT, $2::BIGINT)", filmID, userID)
	return err
}

func (db *DB) Close() error {
	if db.connection == nil {
		return fmt.Errorf("Database connection not opened.")
	}
	return db.connection.Close()
}
