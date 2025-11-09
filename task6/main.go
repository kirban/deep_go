package main

import (
	"fmt"
	"unsafe"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		if len(name) > 42 {
			panic("name too long")
		}
		person.name = []byte(name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.coords[0] = int32(x)
		person.coords[1] = int32(y)
		person.coords[2] = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		if mana < 0 || mana > 1000 {
			panic("wrong mana")
		}
		person.hm[1] = uint16(mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		if health < 0 || health > 1000 {
			panic("wrong health")
		}
		person.hm[0] = uint16(health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		if respect < 0 || respect > 10 {
			panic("wrong respect")
		}
		person.characteristics[0] = uint8(respect)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		if strength < 0 || strength > 10 {
			panic("wrong strength")
		}
		person.characteristics[1] = uint8(strength)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		if experience < 0 || experience > 10 {
			panic("wrong experience")
		}
		person.characteristics[2] = uint8(experience)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		if level < 0 || level > 10 {
			panic("wrong level")
		}
		person.characteristics[3] = uint8(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags[0] = true
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags[1] = true
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags[2] = true
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		switch personType {
		case BuilderGamePersonType:
			person.personType = BuilderGamePersonType
		case BlacksmithGamePersonType:
			person.personType = BlacksmithGamePersonType
		case WarriorGamePersonType:
			person.personType = WarriorGamePersonType
		default:
			panic("unknown person type")
		}
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	gold            uint32
	coords          [3]int32
	hm              [2]uint16
	characteristics [4]uint8
	personType      uint8 // 0 1 2
	flags           [3]bool
	name            []byte
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}

	for _, option := range options {
		option(&person)
	}

	return person
}

func (p *GamePerson) Name() string {
	return string(p.name)
}

func (p *GamePerson) X() int {
	return int(p.coords[0])
}

func (p *GamePerson) Y() int {
	return int(p.coords[1])
}

func (p *GamePerson) Z() int {
	return int(p.coords[2])
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.hm[1])
}

func (p *GamePerson) Health() int {
	return int(p.hm[0])
}

func (p *GamePerson) Respect() int {
	return int(p.characteristics[0])
}

func (p *GamePerson) Strength() int {
	return int(p.characteristics[1])
}

func (p *GamePerson) Experience() int {
	return int(p.characteristics[2])
}

func (p *GamePerson) Level() int {
	return int(p.characteristics[3])
}

func (p *GamePerson) HasHouse() bool {
	return p.flags[0]
}

func (p *GamePerson) HasGun() bool {
	return p.flags[1]
}

func (p *GamePerson) HasFamily() bool {
	return p.flags[2]
}

func (p *GamePerson) Type() int {
	return int(p.personType)
}

func main() {
	warrior := NewGamePerson(
		WithName("Vasya"),
		WithCoordinates(19, 2, 42),
		WithGold(1000),
		WithMana(100),
		WithHealth(100),
		WithRespect(0),
		WithStrength(1),
		WithExperience(10),
		WithLevel(1),
		WithHouse(),
		WithGun(),
		WithFamily(),
		WithType(WarriorGamePersonType),
	)

	fmt.Printf("%+v\n", warrior)
	fmt.Println(unsafe.Sizeof(warrior))
	fmt.Println(unsafe.Alignof(warrior))
	fmt.Printf("%#v\n", (*[56]byte)(unsafe.Pointer(&warrior)))
}
