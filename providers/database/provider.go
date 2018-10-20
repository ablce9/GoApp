package database

import (
	"github.com/ablce9/go-assignment/engine"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var (
	pdb *pg.DB
)

type Provider struct {
	Db *pg.DB
}

func (provider *Provider) GetKnightRepository() engine.KnightRepository {
	return &knightRepository{}
}

// Close closes database connection.
func (provider *Provider) Close() {
	provider.Db.Close()
}

// NewProvider creates Provider.
func NewProvider(addr, user, password, database string) *Provider {
	pdb = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: database,
	})
	if _, err := createKnightSchema(pdb); err != nil {
		panic(err)
	}
	return &Provider{
		Db: pdb,
	}
}

func createKnightSchema(db *pg.DB) (orm.Result, error) {
	return pdb.Exec(`create table if not exists knights(id serial PRIMARY KEY, name varchar, strength integer, weapon_power float)`)
}
