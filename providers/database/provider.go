package database

import "gitlab.com/zenport.io/go-assignment/engine"

type Provider struct {
}

func (provider *Provider) GetKnightRepository() engine.KnightRepository {
	return &knightRepository{}
}

func (provider *Provider) Close() {

}

func NewProvider() *Provider {
	return &Provider{}
}
