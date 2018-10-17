package domain

type Arena struct{}

func (arena *Arena) Fight(fighter1 Fighter, fighter2 Fighter) Fighter {
	f1Power := fighter1.GetPower()
	p2Power := fighter2.GetPower()

	if (f1Power == p2Power) {
		return nil
	}
	if (f1Power > p2Power) {
		return fighter1
	}
	return fighter2
}
