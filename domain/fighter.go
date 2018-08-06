package domain

type Fighter interface {
	GetID() string
	GetPower() float64
}

type Knight struct {
}
