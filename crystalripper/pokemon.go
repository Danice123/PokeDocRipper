package crystalripper

type Pokemon struct {
	Pokedex    int         `json:"pokedex"`
	Name       string      `json:"name"`
	Stats      Statblock   `json:"stats"`
	Types      []Type      `json:"types"`
	BaseExp    int         `json:"base_experience"`
	Evolutions []Evolution `json:"evolutions,omitempty"`
	Learnset   []LearnMove `json:"learnset"`
	TMHM       []string    `json:"tmhm"` // TODO
}

type Statblock struct {
	Health         int `json:"hp"`
	Attack         int `json:"atk"`
	Defense        int `json:"def"`
	SpecialAttack  int `json:"spa"`
	SpecialDefense int `json:"spd"`
	Speed          int `json:"spe"`
}

type Evolution struct {
	Into   string `json:"into"`
	Level  int    `json:"level,omitempty"`
	Item   string `json:"item,omitempty"`
	Method string `json:"method"`
}

type LearnMove struct {
	Level int    `json:"level"`
	Move  string `json:"move"`
}

var TMHM = []string{
	"dynamicpunch",
	"headbutt",
	"curse",
	"rollout",
	"roar",
	"toxic",
	"zap-cannon",
	"rock-smash",
	"psych-up",
	"hidden-power",
	"sunny-day",
	"sweet-scent",
	"snore",
	"blizzard",
	"hyper-beam",
	"icy-wind",
	"protect",
	"rain-dance",
	"giga-drain",
	"endure",
	"frustration",
	"solarbeam",
	"iron-tail",
	"dragonbreath",
	"thunder",
	"earthquake",
	"return",
	"dig",
	"psychic",
	"shadow-ball",
	"mud-slap",
	"double-team",
	"ice-punch",
	"swagger",
	"sleep-talk",
	"sludge-bomb",
	"sandstorm",
	"fire-blast",
	"swift",
	"defense-curl",
	"thunderpunch",
	"dream-eater",
	"detect",
	"rest",
	"attract",
	"thief",
	"steel-wing",
	"fire-punch",
	"fury-cutter",
	"nightmare",
	"cut",
	"fly",
	"surf",
	"strength",
	"flash",
	"whirlpool",
	"waterfall",
	"flamethrower",
	"thunderbolt",
	"ice-beam",
}
