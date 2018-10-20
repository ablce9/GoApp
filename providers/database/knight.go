package database

import (
	"strconv"

	"github.com/ablce9/go-assignment/domain"
	"github.com/go-pg/pg/orm"
)

type knightRepository struct{}

func (repository *knightRepository) Find(ID string) *domain.Knight {
	var knights []*domain.Knight
	id, err := strconv.Atoi(ID)
	if err != nil {
		return nil
	}
	if _, err := pdb.Query(&knights, `select * from knights where id = ?`, id); err != nil {
		panic(err)
	}
	if len(knights) == 0 {
		return nil
	}
	return knights[0]
}

func (repository *knightRepository) FindAll() []*domain.Knight {
	var knights []*domain.Knight
	if _, err := pdb.Query(&knights, "select * from knights"); err != nil {
		panic(err)
	}
	return knights
}

func (repository *knightRepository) Save(knight *domain.Knight) {
	if err := orm.Insert(pdb, knight); err != nil {
		panic(err)
	}
}
