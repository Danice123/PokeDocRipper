package crystalripper

type Trainer struct {
	Name string           `json:"name"`
	Team []TrainerPokemon `json:"team"`
}

type TrainerPokemon struct {
	Name  string   `json:"name"`
	Level int      `json:"level"`
	Item  string   `json:"item,omitempty"`
	Moves []string `json:"moves,omitempty"`
}
