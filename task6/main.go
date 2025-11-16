package main

import (
	"unsafe"
)

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)
const DefaultPersonType = WarriorGamePersonType

const (
	HasHouseFlag  uint8 = 1 << iota // 001
	HasGunFlag                      // 010
	HasFamilyFlag                   // 100
)

const (
	RespectOffset uint8 = 4 * iota
	StrengthOffset
	ExperienceOffset
	LevelOffset
)

const (
	RespectMask uint16 = 0xf << (4 * iota)
	StrengthMask
	ExperienceMask
	LevelMask
)

// константы для валидации
const (
	MinName           = 0
	MaxName           = 42
	MinGold           = 0
	MaxGold           = 2_000_000_000
	MinCoord          = -2_000_000_000
	MaxCoord          = 2_000_000_000
	MinHealth         = 0
	MaxHealth         = 1000
	MinMana           = 0
	MaxMana           = 1000
	MinCharacteristic = 0
	MaxCharacteristic = 10
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		if len(name) > MaxName {
			person.name = person.name[:MaxName]
		} else if len(name) < MinName {
			person.name = []byte{}
		}

		person.name = []byte(name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		if x < MinCoord {
			x = MinCoord
		} else if x > MaxCoord {
			x = MaxCoord
		}

		if y < MinCoord {
			y = MinCoord
		} else if y > MaxCoord {
			y = MaxCoord
		}

		if z < MinCoord {
			z = MinCoord
		} else if z > MaxCoord {
			z = MaxCoord
		}

		person.coords[0] = int32(x)
		person.coords[1] = int32(y)
		person.coords[2] = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		if gold < MinGold {
			gold = MinGold
		} else if gold > MaxGold {
			gold = MaxGold
		}

		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		if mana < MinMana {
			mana = MinMana
		} else if mana > MaxMana {
			mana = MaxMana
		}

		person.hm[1] = uint16(mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		if health < MinHealth {
			health = MinHealth
		} else if health > MaxHealth {
			health = MaxHealth
		}

		person.hm[0] = uint16(health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		if respect < MinCharacteristic {
			respect = MinCharacteristic
		} else if respect > MaxCharacteristic {
			respect = MaxCharacteristic
		}

		person.characteristics |= uint16(respect) << RespectOffset
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		if strength < MinCharacteristic {
			strength = MinCharacteristic
		} else if strength > MaxCharacteristic {
			strength = MaxCharacteristic
		}

		person.characteristics |= uint16(strength) << StrengthOffset
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		if experience < MinCharacteristic {
			experience = MinCharacteristic
		} else if experience > MaxCharacteristic {
			experience = MaxCharacteristic
		}

		person.characteristics |= uint16(experience) << ExperienceOffset
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		if level < MinCharacteristic {
			level = MinCharacteristic
		} else if level > MaxCharacteristic {
			level = MaxCharacteristic
		}

		person.characteristics |= uint16(level) << LevelOffset
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= HasHouseFlag
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= HasGunFlag
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.flags |= HasFamilyFlag
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
			person.personType = DefaultPersonType
		}
	}
}

type GamePerson struct {
	gold            uint32
	coords          [3]int32  // x | y | z
	hm              [2]uint16 // health | mana
	characteristics uint16    // respect | strength | experience | level
	flags           uint8     // hasHouse | hasGun | hasFamily
	personType      uint8
	name            []byte
}

var _ uintptr = 64 - unsafe.Sizeof(GamePerson{})

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
	return int(p.characteristics&RespectMask) >> RespectOffset
}

func (p *GamePerson) Strength() int {
	return int(p.characteristics&StrengthMask) >> StrengthOffset
}

func (p *GamePerson) Experience() int {
	return int(p.characteristics&ExperienceMask) >> ExperienceOffset
}

func (p *GamePerson) Level() int {
	return int(p.characteristics&LevelMask) >> LevelOffset
}

func (p *GamePerson) HasHouse() bool {
	return p.flags&HasHouseFlag != 0
}

func (p *GamePerson) HasGun() bool {
	return p.flags&HasGunFlag != 0
}

func (p *GamePerson) HasFamily() bool {
	return p.flags&HasFamilyFlag != 0
}

func (p *GamePerson) Type() int {
	return int(p.personType)
}
