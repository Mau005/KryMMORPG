package models

type Items struct {
	Group            string
	RWInfo           uint
	ReadOnlyId       uint
	Stackable        bool
	Useable          bool
	Moveable         bool
	AlwaysOnTop      bool
	PickUpable       bool
	Rotable          bool
	RotateTo         uint8
	HasHeight        bool
	FloorChangeDown  bool
	FloorChangeNorth bool
	FloorChangeSouth bool
	FloorChangeEast  bool
	FloorChangeWest  bool
	IsDoor           bool
	NewChagers       uint
	NewTime          uint
	IsDeleter        bool
	BlockSolid       bool
	BlockProjectile  bool
	BlockPathFind    bool
	RuneMagLevel     int
	MagicFieldType   int
	Speed            uint
	Id               uint
	ClientId         uint
	MaxItem          uint
	Weight           uint
	WeaponType       bool
	SlotPosition     bool
	AmuType          bool
	ShootType        bool
	Attack           uint
	Defence          uint
	Armor            uint
	DecayTo          uint
	DecayTime        uint
	CanDecay         bool
	LightLevel       uint
	LightColor       uint
}

func (i *Items) IsBlocking() bool {
	return i.BlockSolid
}
