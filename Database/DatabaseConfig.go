package Database

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var (
	PostgresDB *sqlx.DB

	DBContext   = context.TODO()
	sqlDBDriver = "postgres"
	sqlDBSource = "postgresql://postgres:asdasd@localhost:6432/pgbouncer_test?sslmode=disable"
)

type DatabaseConfig struct {
	MaxIdleConnections int `env:"DATABASE_MAX_IDLE_CONNECTIONS" envDefault:"10"`
	MaxOpenConnections int `env:"DATABASE_MAX_OPEN_CONNECTIONS" envDefault:"100"`
}

func PostgresConnection() {
	conn, err := sqlx.Open(sqlDBDriver, sqlDBSource)
	//conn, err := sqlx.Open(sqlDBDriver, sqlDBSourceJSON)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	fmt.Println("all ok in database")
	config := DatabaseConfig{}
	if parseError := cleanenv.ReadConfig(".env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}
	conn.SetMaxIdleConns(config.MaxIdleConnections)
	conn.SetMaxOpenConns(config.MaxOpenConnections)
	PostgresDB = conn
}
