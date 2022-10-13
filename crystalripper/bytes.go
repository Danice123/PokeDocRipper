package crystalripper

import (
	"bytes"
	"fmt"
	"strings"
)

func searchBytes(data []byte, search []byte) int {
	searchSize := len(search)
	result := -1
	for i := 0; i+searchSize < len(data); i++ {
		if bytes.Equal(data[i:i+searchSize], search) {
			if result != -1 {
				panic("Multiple matches!")
			}
			result = i
		}
	}
	if result == -1 {
		panic(fmt.Sprintf("NO MATCH FOR: %x", search))
	}
	return result
}

func searchBytesWithWildcards(data []byte, search []*byte) int {
	searchSize := len(search)
	result := -1
	for i := 0; i+searchSize < len(data); i++ {
		found := true
		for j := 0; j < len(search); j++ {
			if search[j] != nil {
				if data[i+j] != *search[j] {
					found = false
					break
				}
			}
		}

		if found {
			if result != -1 {
				fmt.Printf("WARNING: ALSO MATCHES %x\n", i)
			}
			result = i

		}
	}
	if result == -1 {
		panic(fmt.Sprintf("NO MATCH FOR: %x", search))
	}
	return result
}

type RomCursor struct {
	Rom    []byte
	Cursor int
}

func (ths *RomCursor) ReadByte() byte {
	out := ths.Rom[ths.Cursor]
	ths.Cursor++
	return out
}

func (ths *RomCursor) ReadInt() int {
	out := int(ths.Rom[ths.Cursor])
	ths.Cursor++
	return out
}

func (ths *RomCursor) ReadBytes(length int) []byte {
	out := ths.Rom[ths.Cursor : ths.Cursor+length]
	ths.Cursor += length
	return out
}

func (ths *RomCursor) ReadTerminatedString() string {
	var s strings.Builder
	for {
		b := ths.ReadByte()
		if b != 0x50 {
			s.WriteString(Charmap(b))
		} else {
			break
		}
	}
	return s.String()
}

func (ths *RomCursor) PeekByte() byte {
	return ths.Rom[ths.Cursor]
}
