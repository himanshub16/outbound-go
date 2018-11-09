package main

import (
	"crypto/tls"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

type PostgreSQLRepository struct {
	db *pg.DB
}

func (p *PostgreSQLRepository) FindCounterById(id string) (*Counter, error) {
	counter := &Counter{ID: id}

	err := p.db.Select(counter)
	if err != nil {
		return nil, err
	}

	return counter, nil
}

func (p *PostgreSQLRepository) UpsertCounter(counter Counter) error {
	_, err := p.db.Model(&counter).
		OnConflict("(id) DO UPDATE").
		Set("count = ?count").
		Insert()
	return err
}

func (p *PostgreSQLRepository) InsertLink(link Link) error {
	return p.db.Insert(&link)
}

func (p *PostgreSQLRepository) FindLinkByShortIdInt(id uint) (*Link, error) {
	link := &Link{ShortIDInt: id}
	err := p.db.Model(link).
		Where("short_id_int = ?short_id_int").
		Select()
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (p *PostgreSQLRepository) UpdateLink(link Link) error {
	return p.db.Update(&link)
}

func (p *PostgreSQLRepository) close() {
	p.db.Close()
}

func NewPostgreSQLRepository() *PostgreSQLRepository {
	var tlsConfig *tls.Config = nil
	if config.UseSSL == true {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}
	}
	db := pg.Connect(&pg.Options{
		Addr:      config.DBURL,
		User:      config.DBUser,
		Password:  config.DBPass,
		Database:  config.DBName,
		TLSConfig: tlsConfig,
	})

	for _, model := range []interface{}{&Link{}, &Counter{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	return &PostgreSQLRepository{db}
}
