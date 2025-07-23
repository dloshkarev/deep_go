package hw6

type Option func(*GamePerson)

type GamePerson struct {
	// |--Name--|---X---|---Y---|---Z---|--Gold-|-Mana--|-Health|Strength|Respect|-Level|Experience|PersonType|Family|-Gun--|-House|
	// |42 bytes|4 bytes|4 bytes|4 bytes|4 bytes|12 bits|12 bits| 4 bits | 4 bits|4 bits|--4 bits--|--4 bits--|-1 bit|-1 bit|-1 bit|
	data [DataLength]byte
}

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		if len(name) > NameLength {
			panic("name is too long")
		}
		copy(person.data[:], name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.writeInt(x, XIdx)
		person.writeInt(y, YIdx)
		person.writeInt(z, ZIdx)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.writeInt(gold, GoldIdx)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.data[ManaIdx] = byte(mana)
		person.data[ManaIdx+1] |= byte(mana>>8) << 4
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.data[HealthIdx+1] = byte(health)
		person.data[HealthIdx] |= byte(health >> 8)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.writeNibble(respect, RespectStrengthIdx, false)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.writeNibble(strength, RespectStrengthIdx, true)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.writeNibble(experience, ExpLevelIdx, false)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.writeNibble(level, ExpLevelIdx, true)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.data[FlagsIdx] |= HouseMask
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.data[FlagsIdx] |= GunMask
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.data[FlagsIdx] |= FamilyMask
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.data[FlagsIdx] |= byte(personType) << 4
	}
}

const (
	DataLength            = 64
	NameLength            = 42
	XIdx                  = 42
	YIdx                  = 46
	ZIdx                  = 50
	GoldIdx               = 54
	ManaIdx               = 58
	HealthIdx             = 59
	RespectStrengthIdx    = 61
	ExpLevelIdx           = 62
	FlagsIdx              = 63
	HouseMask             = 1 << 0
	GunMask               = 1 << 1
	FamilyMask            = 1 << 2
	HighNibbleMask        = 0b00001111
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}

	for _, option := range options {
		option(&person)
	}

	return person
}

func (p *GamePerson) Name() string {
	nameBytes := [NameLength]byte{}
	length := 0
	for i := 0; i < NameLength; i++ {
		nameByte := p.data[i]
		if nameByte != 0 {
			nameBytes[i] = nameByte
			length++
		} else {
			break
		}
	}
	return string(nameBytes[:length])
}

func (p *GamePerson) X() int {
	return p.readInt(XIdx)
}

func (p *GamePerson) Y() int {
	return p.readInt(YIdx)
}

func (p *GamePerson) Z() int {
	return p.readInt(ZIdx)
}

func (p *GamePerson) Gold() int {
	return p.readInt(GoldIdx)
}

func (p *GamePerson) Mana() int {
	return int(p.data[ManaIdx+1])>>4<<8 | int(p.data[ManaIdx])
}

func (p *GamePerson) Health() int {
	return int(p.data[HealthIdx]&HighNibbleMask)<<8 | int(p.data[HealthIdx+1])
}

func (p *GamePerson) Respect() int {
	return p.readNibble(RespectStrengthIdx, false)
}

func (p *GamePerson) Strength() int {
	return p.readNibble(RespectStrengthIdx, true)
}

func (p *GamePerson) Experience() int {
	return p.readNibble(ExpLevelIdx, false)
}

func (p *GamePerson) Level() int {
	return p.readNibble(ExpLevelIdx, true)
}

func (p *GamePerson) HasHouse() bool {
	return p.data[FlagsIdx]&HouseMask == HouseMask
}

func (p *GamePerson) HasGun() bool {
	return p.data[FlagsIdx]&GunMask == GunMask
}

func (p *GamePerson) HasFamily() bool {
	return p.data[FlagsIdx]&FamilyMask == FamilyMask
}

func (p *GamePerson) Type() int {
	return int(p.data[FlagsIdx] >> 4)
}

func (p *GamePerson) writeInt(x int, idx int) {
	p.data[idx+3] = byte(x)
	p.data[idx+2] = byte(x >> 8)
	p.data[idx+1] = byte(x >> 16)
	p.data[idx] = byte(x >> 24)
}

func (p *GamePerson) readInt(idx int) int {
	return int(int8(p.data[idx]))<<24 | int(p.data[idx+1])<<16 | int(p.data[idx+2])<<8 | int(p.data[idx+3])
}

func (p *GamePerson) writeNibble(x int, idx int, isLeft bool) {
	if isLeft {
		p.data[idx] |= byte(x) << 4
	} else {
		p.data[idx] |= byte(x)
	}
}

func (p *GamePerson) readNibble(idx int, isLeft bool) int {
	if isLeft {
		return int(p.data[idx] >> 4)
	} else {
		return int(p.data[idx] & HighNibbleMask)
	}
}
