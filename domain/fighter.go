package domain

// Fighter is to describe what can fight.
type Fighter interface {
	GetID() string
	GetPower() float64
}

// Knight is a struct. It's an entity that can be stored in the database.
type Knight struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Strength    int     `json:"strength,omitempty"`
	WeaponPower float64 `json:"weapon_power,omitempty"`
}

// GetPower returns a Float which represents the Power level of
// the Fighter.
func (k *Knight) GetPower() float64 {
	return float64(k.Strength) + k.WeaponPower
}

// GetID returns a String which represents the ID of the combatant.
func (k *Knight) GetID() string {
	return string(k.ID)
}
