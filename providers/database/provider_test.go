package database

import (
	"os"
	"testing"

	"github.com/ablce9/go-assignment/domain"
	"github.com/go-pg/pg/orm"
	"strings"
)

// Test configuration
const (
	dbAddr     = "db:5432"
	dbUser     = "postgres"
	dbPassword = "bad-password"
	dbDatabase = "go_assignment_test"
)

var (
	myKnightRepository *knightRepository
	myknight           = domain.Knight{
		Name:        "foo",
		Strength:    10,
		WeaponPower: 20,
	}
)

func TestMain(m *testing.M) {
	provider := NewProvider(
		dbAddr,
		dbUser,
		dbPassword,
		dbDatabase,
	)

	provider.Db.Exec(`create table if not exists knights(id serial PRIMARY KEY, name varchar, strength integer, weapon_power float)`)

	code := m.Run()

	if err := orm.DropTable(provider.Db, interface{}((*domain.Knight)(nil)), nil); err != nil {
		panic(err)
	}
	provider.Close()
	os.Exit(code)
}

func TestSave(t *testing.T) {
	myKnightRepository.Save(&myknight)
	if len(myKnightRepository.FindAll()) != 1 {
		t.Fail()
	}
}

func TestFind(t *testing.T) {
	knight := myKnightRepository.Find("1")
	if strings.Compare(knight.Name, "foo") != 0 {
		t.Errorf("Find: Name should be %s\n", myknight.Name)
	}
	if knight.Strength != 10 {
		t.Errorf("Find: Strength should be %d\n", myknight.Strength)
	}
	if knight.WeaponPower != 20 {
		t.Errorf("Find: WeaponPower should be %2f\n", myknight.WeaponPower)
	}
}

func TestFindAll(t *testing.T) {
	knights := myKnightRepository.FindAll()
	if len(knights) != 1 {
		t.Fail()
	}
	knight := knights[0]
	if strings.Compare(knight.Name, "foo") != 0 {
		t.Errorf("FindAll: Name should be %s\n", myknight.Name)
	}
	if knight.Strength != 10 {
		t.Errorf("FindAll: Strength should be %d\n", myknight.Strength)
	}
	if knight.WeaponPower != 20 {
		t.Errorf("FindAll: WeaponPower should be %2f\n", myknight.WeaponPower)
	}
}
