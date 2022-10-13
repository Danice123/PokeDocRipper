package crystalripper

type PokemonMove struct {
	Index    int    `json:"index"`
	Name     string `json:"name"`
	Type     Type   `json:"type"`
	Power    int    `json:"power"`
	Accuracy int    `json:"accuracy"`
	PP       int    `json:"pp"`

	Extra interface{} `json:"extra"`
}
