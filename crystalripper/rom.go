package crystalripper

import (
	"encoding/binary"
	"math"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/bits-and-blooms/bitset"
)

type CrystalRipper struct {
	Pokemon  []Pokemon     `json:"pokemon"`
	Moves    []PokemonMove `json:"moves"`
	Trainers []Trainer     `json:"trainers"`
	Items    []string      `json:"items"`
}

func New(rom []byte) *CrystalRipper {
	cr := &CrystalRipper{}
	cr.readBaseData(rom)
	cr.readMoveList(rom)
	cr.readItemList(rom)
	cr.readLearnsets(rom)
	cr.readTrainers(rom)
	return cr
}

func (ths *CrystalRipper) readBaseData(rom []byte) {
	search := make([]*byte, 14)
	search[0] = pointer.ToByte(0x01)
	search[7] = pointer.ToByte(0x16)
	search[8] = pointer.ToByte(0x03)
	search[13] = pointer.ToByte(0x1F)
	// Search for 01 pokemon data (bulbasaur). Assumes Grass/Poison typing and canonical growth rate
	c := RomCursor{
		Rom:    rom,
		Cursor: searchBytesWithWildcards(rom, search),
	}

	for i := 1; ; i++ {
		poke := Pokemon{}
		poke.Pokedex = c.ReadInt()
		if poke.Pokedex != i {
			break
		}

		poke.Stats.Health = c.ReadInt()
		poke.Stats.Attack = c.ReadInt()
		poke.Stats.Defense = c.ReadInt()
		poke.Stats.Speed = c.ReadInt()
		poke.Stats.SpecialAttack = c.ReadInt()
		poke.Stats.SpecialDefense = c.ReadInt()
		poke.Types = []Type{ToType(c.ReadByte())}
		secondType := ToType(c.ReadByte())
		if poke.Types[0] != secondType {
			poke.Types = append(poke.Types, secondType)
		}
		_ = c.ReadInt() // catch rate
		poke.BaseExp = c.ReadInt()
		_ = c.ReadInt() // item one
		_ = c.ReadInt() // item two
		_ = c.ReadInt() // gender ratio
		_ = c.ReadInt() // unused
		_ = c.ReadInt() // step cycles
		_ = c.ReadInt() // unused
		_ = c.ReadInt() // front dimensions
		c.Cursor += 4   // unused data
		_ = c.ReadInt() // growth rate
		_ = c.ReadInt() // egg groups

		tmhm := bitset.From([]uint64{binary.LittleEndian.Uint64(c.ReadBytes(8))})
		poke.TMHM = []string{}
		for i, move := range TMHM {
			if tmhm.Test(uint(i)) {
				poke.TMHM = append(poke.TMHM, move)
			}
		}

		ths.Pokemon = append(ths.Pokemon, poke)
	}

	// Search for BULBASAUR@, assumes first mon in list
	c = RomCursor{Rom: rom, Cursor: searchBytes(rom, []byte{0x81, 0x94, 0x8B, 0x81, 0x80, 0x92, 0x80, 0x94, 0x91, 0x50})}
	for i := 0; i < len(ths.Pokemon); i++ {
		name := c.ReadBytes(10)
		ths.Pokemon[i].Name = strings.ToLower(ConvertCharmaps(name))
	}
}

func (ths *CrystalRipper) readMoveList(rom []byte) {
	search := make([]*byte, 16)       // 0x00041AFB
	search[0] = pointer.ToByte(0x01)  // First move
	search[1] = pointer.ToByte(0x00)  // Is normal hit
	search[7] = pointer.ToByte(0x02)  // Second move
	search[8] = pointer.ToByte(0x00)  // Is normal hit
	search[14] = pointer.ToByte(0x03) // Third move
	search[15] = pointer.ToByte(0x1D) // Is multi hit
	c := RomCursor{Rom: rom, Cursor: searchBytesWithWildcards(rom, search)}

	for i := 1; ; i++ {
		move := PokemonMove{}
		move.Index = c.ReadInt()
		if move.Index != i {
			break
		}
		_ = c.ReadInt() // Effect
		move.Power = c.ReadInt()
		move.Type = ToType(c.ReadByte())
		move.Accuracy = int(math.Round(float64(c.ReadByte()) / float64(0xFF) * 100.0))
		move.PP = c.ReadInt()
		_ = int(math.Round(float64(c.ReadByte()) / float64(0xFF) * 100.0)) // Effect chance

		ths.Moves = append(ths.Moves, move)
	}

	// Searching for word: POUND, assumes first move in list
	c = RomCursor{Rom: rom, Cursor: searchBytes(rom, []byte{0x8f, 0x8e, 0x94, 0x8d, 0x83})}

	for i := 1; i <= len(ths.Moves); i++ {
		ths.Moves[i-1].Name = strings.ReplaceAll(
			strings.ToLower(c.ReadTerminatedString()),
			" ",
			"-",
		)
	}
}

func (ths *CrystalRipper) readItemList(rom []byte) {
	// Searching for "MASTER BALL@" with a null before it
	c := RomCursor{
		Rom:    rom,
		Cursor: searchBytes(rom, []byte{0x00, 0x8c, 0x80, 0x92, 0x93, 0x84, 0x91, 0x7f, 0x81, 0x80, 0x8b, 0x8b, 0x50}), // 0x001C7FFF
	}

	for {
		item := strings.Title(strings.ToLower(c.ReadTerminatedString()))
		if item == "?" {
			break
		}
		ths.Items = append(ths.Items, item)
	}
}

func (ths *CrystalRipper) readLearnsets(rom []byte) {
	// Search for EVOLVE_LEVEL, level 16, IVYSAUR, NIL
	c := RomCursor{
		Rom:    rom,
		Cursor: searchBytes(rom, []byte{0x01, 0x10, 0x02, 0x00}), // 0x000427A7
	}

	for i := 0; i < len(ths.Pokemon); i++ {
		// Evolution table
	evoloop:
		for {
			evo := Evolution{}
			switch c.ReadByte() {
			case 0x00:
				break evoloop
			case 0x01:
				evo.Method = "level"
				evo.Level = c.ReadInt()
				evo.Into = ths.Pokemon[c.ReadInt()-1].Name
			case 0x02:
				evo.Method = "item"
				evo.Item = ths.Items[c.ReadInt()-1]
				evo.Into = ths.Pokemon[c.ReadInt()-1].Name
			case 0x03:
				evo.Method = "trade"
				evo.Item = ths.Items[c.ReadInt()-1]
				evo.Into = ths.Pokemon[c.ReadInt()-1].Name
			case 0x04:
				evo.Method = "happiness"
				c.ReadByte() // TIME OF DAY
				evo.Into = ths.Pokemon[c.ReadInt()-1].Name
			case 0x05:
				evo.Method = "stat"
				evo.Level = c.ReadInt()
				c.ReadByte() // ATK_*_DEF constant (LT, GT, EQ)
				evo.Into = ths.Pokemon[c.ReadInt()-1].Name
			}
			ths.Pokemon[i].Evolutions = append(ths.Pokemon[i].Evolutions, evo)
		}
		// Learnsets
		for {
			lm := LearnMove{}
			lm.Level = c.ReadInt()
			if lm.Level == 0 {
				break
			}
			lm.Move = ths.Moves[c.ReadInt()-1].Name

			ths.Pokemon[i].Learnset = append(ths.Pokemon[i].Learnset, lm)
		}
	}
}

func (ths *CrystalRipper) readTrainers(rom []byte) {
	// Search for FALKNER@
	// c := RomCursor{
	// 	Rom:    rom,
	// 	Cursor: searchBytes(rom, []byte{0x85, 0x80, 0x8b, 0x8a, 0x8d, 0x84, 0x91, 0x50}), // 0x00039A1F
	// }

	// Search for idk, ck+ is hard
	c := RomCursor{
		Rom:    rom,
		Cursor: int(0x001D8000), // 0x001D8000
	}

	for {
		if c.PeekByte() == 0x0 {
			break // idk if this works for hacks, works for stock rom
		}

		trainer := Trainer{}
		trainer.Name = c.ReadTerminatedString()
		tType := c.ReadByte()

		// fmt.Printf("Trainer %s - Type %d\n", trainer.Name, int(tType))

		for {
			lb := c.ReadByte()
			if lb == 0xFF {
				break
			}
			tp := TrainerPokemon{}
			tp.Level = int(lb)
			tp.Name = ths.Pokemon[c.ReadInt()-1].Name

			if tType > 0x04 {
				c.ReadBytes(2) // unknown ck+ bytes
			}

			if tType > 0x01 && tType != 0x05 {
				i := c.ReadInt()
				if i != 0 {
					tp.Item = ths.Items[i-1]
				}
			}

			if tType == 0x01 || tType > 0x02 {
				for i := 0; i < 4; i++ {
					move := c.ReadInt()
					if move == 0 {
						continue
					}
					tp.Moves = append(tp.Moves, ths.Moves[move-1].Name)
				}
			}
			trainer.Team = append(trainer.Team, tp)
			// fmt.Printf("%s - %d - %s - %v\n", tp.Name, tp.Level, tp.Item, tp.Moves)
		}
		ths.Trainers = append(ths.Trainers, trainer)
	}

	// c = RomCursor{
	// 	Rom:    rom,
	// 	Cursor: int(0x0002C1EF),
	// }

	// for i := 0; i < len(ths.Trainers); i++ {
	// 	title := c.ReadTerminatedString()
	// 	ths.Trainers[i].Name = title + " " + ths.Trainers[i].Name
	// }
}
