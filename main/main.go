package main

import (
	"encoding/json"
	"os"

	"github.com/danice123/crystaldocripper/crystalripper"
)

func JSONRemarshal(bytes []byte) ([]byte, error) {
	var ifce interface{}
	err := json.Unmarshal(bytes, &ifce)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(ifce, "", "    ")
}

func main() {
	// rom, err := os.ReadFile("pokecrystal.gbc")
	rom, err := os.ReadFile("Pokemon - Crystal Kaizo+ 0.1.35.gbc")
	if err != nil {
		panic(err)
	}

	cr := crystalripper.New(rom)

	d, err := json.MarshalIndent(cr, "", "\t")
	if err != nil {
		panic(err)
	}

	d, err = JSONRemarshal(d)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("output.json", d, 0666)
	if err != nil {
		panic(err)
	}
}
