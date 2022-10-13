package crystalripper

type Type string

const Normal = Type("normal")
const Fighting = Type("fighting")
const Flying = Type("flying")
const Poison = Type("poison")
const Ground = Type("ground")
const Rock = Type("rock")
const Bird = Type("bird")
const Bug = Type("bug")
const Ghost = Type("ghost")
const Steel = Type("steel")

const Curse = Type("curse")

const Fire = Type("fire")
const Water = Type("water")
const Grass = Type("grass")
const Electric = Type("electric")
const Psychic = Type("psychic")
const Ice = Type("ice")
const Dragon = Type("dragon")
const Dark = Type("dark")

func ToType(b byte) Type {
	switch b {
	case 0x0:
		return Normal
	case 0x1:
		return Fighting
	case 0x2:
		return Flying
	case 0x3:
		return Poison
	case 0x4:
		return Ground
	case 0x5:
		return Rock
	case 0x6:
		return Bird
	case 0x7:
		return Bug
	case 0x8:
		return Ghost
	case 0x9:
		return Steel
	case 0x13:
		return Curse
	case 0x14:
		return Fire
	case 0x15:
		return Water
	case 0x16:
		return Grass
	case 0x17:
		return Electric
	case 0x18:
		return Psychic
	case 0x19:
		return Ice
	case 0x1A:
		return Dragon
	case 0x1B:
		return Dark
	default:
		return ""
	}
}
