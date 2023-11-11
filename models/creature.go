package models

import "gorm.io/gorm"

type Tactic struct {
	// Fight modes
	// 1 - Offensive fighting
	// 2 - Balanced fighting
	// 3 - Defensive fighting
	FightMode uint8

	// 0 - Stand while fighting
	// 1 - Chase opponent
	ChaseOpponent uint8

	// 0 - You cannot attack unmarked players
	// 1 - You can attack any player
	AttackPlayers uint8
}

type Skills struct {
	First          uint8
	FirstTries     uint64
	Club           uint8
	ClubTries      uint64
	Sword          uint8
	SwordTries     uint64
	Axe            uint8
	AxeTries       uint64
	Dist           uint8
	DistTries      uint64
	Shielding      uint8
	ShieldingTries uint64
	Fishing        uint8
	FishingTries   uint64
}

type Outfits struct {
	Type uint8
	Head uint8
	Body uint8
	Legs uint8
	Feet uint8
}

type Creature struct {
	gorm.Model
	Name       string
	Group      uint8
	AccountsID uint
	Level      uint16
	Vocation   uint8
	Health     uint
	HealthMax  uint
	Mana       uint
	ManaMax    uint
	Experience uint
	TimeOnline float64 //Is count hours for time
	Soul       uint8
	Sex        uint8
	Position   Position
	Skills     Skills
	Tactic     Tactic
}
