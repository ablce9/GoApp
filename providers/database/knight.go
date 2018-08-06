package database

import "gitlab.com/zenport.io/go-assignment/domain"

type knightRepository struct{}

func (repository *knightRepository) Find(ID string) *domain.Knight {
	return nil
}
