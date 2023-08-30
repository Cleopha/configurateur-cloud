package db

import (
	"backend/ent"
	"backend/utils"
	"fmt"
	_ "github.com/lib/pq"
)

type postgres struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func (p *postgres) loadConfig() {
	p.host = utils.MustGetEnv("POSTGRES_HOST")
	p.port = utils.MustGetEnv("POSTGRES_PORT")
	p.user = utils.MustGetEnv("POSTGRES_USER")
	p.password = utils.MustGetEnv("POSTGRES_PASSWORD")
	p.dbName = utils.MustGetEnv("POSTGRES_DB")
}

func (p *postgres) Client() (*Client, error) {
	p.loadConfig()

	client, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.host, p.port, p.user, p.password, p.dbName))
	fmt.Println(p)
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
