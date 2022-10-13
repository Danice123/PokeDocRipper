package crystalripper

import (
	"fmt"
	"strings"
)

func ConvertCharmaps(bytes []byte) string {
	s := strings.Builder{}
	for _, b := range bytes {
		if b == 0x50 {
			break
		} else {
			s.WriteString(Charmap(b))
		}
	}
	return s.String()
}

func Charmap(b byte) string {
	if b >= 0x80 && b <= 0x99 {
		return string(byte('A') + b - 0x80)
	} else if b >= 0xa0 && b <= 0xb9 {
		return string(byte('a') + b - 0xa0)
	} else if b >= 0xf6 {
		return string(byte('0') + b - 0xf6)
	} else if b == 0xe3 || b == 0xe8 {
		return "-"
	} else if b == 0xe6 {
		return "?"
	} else if b == 0xe0 {
		return "'"
	} else if b == 0x7f {
		return " "
	} else if b == 0xf5 {
		return "-f"
	} else if b == 0xef {
		return "-m"
	}
	return fmt.Sprintf("%x", b)
}
