package storage

import (
	"fmt"
	"github.com/guregu/null"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rinatkh/test_2022/config"
)

type Comment struct {
	Id              int64    `db:"id"`
	UserId          int      `db:"user_id"`
	CommentId       null.Int `db:"comment_id"`
	Content         string   `db:"content"`
	Level           int
	VoteCount       int     `db:"voteCount"`
	UpdatedAt       float64 `db:"updated_at"`
	UpdatedAtNormal string
}

func InitPsqlDB(c *config.Config) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.DBName,
		c.Postgres.SSLMode)
	fmt.Println(connectionUrl)
	database, err := sqlx.Connect(c.Postgres.PgDriver, connectionUrl)
	if err != nil {
		return nil, err
	}
	database.SetMaxOpenConns(50)
	return database, nil
}
