package engine

import "github.com/ablce9/go-assignment/domain"

type Engine interface {
	GetKnight(ID string) (*domain.Knight, error)
	ListKnights() []*domain.Knight
	Fight(fighter1ID string, fighter2ID string) domain.Fighter
}

type KnightRepository interface {
	Find(ID string) *domain.Knight
	FindAll() []*domain.Knight
	Save(knight *domain.Knight)
}

type DatabaseProvider interface {
	GetKnightRepository() KnightRepository
}

type arenaEngine struct {
	arena            *domain.Arena
	knightRepository KnightRepository
}

func (a *arenaEngine) GetKnightRepository() KnightRepository {
	return a.knightRepository
}

func NewEngine(db DatabaseProvider) Engine {
	return &arenaEngine{
		arena:            &domain.Arena{},
		knightRepository: db.GetKnightRepository(),
	}
}
