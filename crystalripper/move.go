package crystalripper

import (
	"fmt"
	"strings"
)

type PokemonMove struct {
	Index    int      `json:"index"`
	Name     string   `json:"name"`
	Type     Type     `json:"type"`
	Power    int      `json:"power"`
	Accuracy int      `json:"accuracy"`
	PP       int      `json:"pp"`
	Extra    []string `json:"extra"`
}

func GetEffect(index int, id int, chance int, power int) []string {
	if id == 0 {
		return []string{}
	}

	if index == 18 { // Whirlwind
		return []string{EffectMap[id] + "\nHits opponents in the semi-invulnerable turn of Fly"}
	}

	if index == 90 { // Fissure
		return []string{EffectMap[id] + "\nHits opponents in the semi-invulnerable turn of Dig"}
	}

	if index == 91 { // Dig
		return []string{"Becomes semi-invulnerable for a turn, then attacks on the next\nWhile semi-invulnerable, invulnerable to all moves except Earthquake, Magnitude, and Fissure\nEarthquake and Magnitude will deal double damage"}
	}

	if index == 107 { // Minimise
		return []string{EffectMap[id] + "\nCauses Stomp to deal double damage afterwards"}
	}

	if index == 156 { // Rest
		return []string{"Restores 100% HP and cures status, goes to sleep for 2 turns"}
	}

	if id == 41 { // level damage
		return []string{fmt.Sprintf(EffectMap[id], power)}
	}

	effect := EffectMap[id]
	if strings.Contains(effect, "%d") {
		effect = fmt.Sprintf(effect, chance)
	} else {
		effect = fmt.Sprintf(effect)
	}

	return []string{effect}
}

var EffectMap = []string{
	"",
	"Puts opponent to sleep",
	"%d%% chance to poison opponent",
	"User heals 50%% of damage dealt",
	"%d%% chance to burn opponent",
	"%d%% chance to freeze opponent",
	"%d%% chance to paralyze opponent",
	"User faints",
	"Can only hit sleeping targets, user heals 50%% damage dealt",
	"If the last damage taken was special, do twice the damage",
	"Raises attack",
	"%d%% chance to raise defense",
	"Raises speed",
	"Raises special attack",
	"Raises special defense",
	"Raises accuracy",
	"Raises evasion",
	"Always hits",
	"Lowers opponent's attack",
	"Lowers opponent's defense",
	"Lowers opponent's speed",
	"Lowers opponent's special attack",
	"Lowers opponent's special defense",
	"Lowers opponent's accuracy",
	"Lowers opponent's evasion",
	"Clears the opponent's stat changes",
	"Absorbs damage for 2-3 turns, then deals twice the damage taken to the opponent",
	"Lasts 2-3 turns, then confuses the user",
	"Forces the opponent to switch",
	"Hits 2-5 times",
	"Changes the user's type into one of its moves' types",
	"%d%% chance to flinch opponent",
	"Restores 50%% HP",
	"Toxic poisons opponent",
	"Scatters coins that are collected after battle",
	"Double's the team's special defense for 5 turns",
	"%d%% to either paralyze, freeze, or burn the opponent",
	"",
	"OHKO",
	"Charges for a turn before attacking",
	"Deals half of the opponent's current HP",
	"Deals %d damage",
	"For 2-5 turns, prevents the opponent from switching and deals 1/16 HP at the end of each turn",
	"",
	"Hits 2 times",
	"If it misses or fails, deals 1/8 the damage it would've dealt to the user instead",
	"Prevents the opponent from lowering the user's stats until the user switches out",
	"Increases critical hit rate",
	"User takes 25%% of damage dealt as recoil",
	"Confuses opponent",
	"Sharply raises attack",
	"Sharply raises defense",
	"Sharply raises speed",
	"Sharply raises special attack",
	"Sharply raises special defense",
	"Sharply raises accuracy",
	"Sharply raises evasion",
	"Transforms into the opponent, copying types, all stats except HP, and all moves with at most 5PP each",
	"Sharply lowers opponent's attack",
	"Sharply lowers opponent's defense",
	"Sharply lowers opponent's speed",
	"Sharply lowers opponent's special attack",
	"Sharply lowers opponent's special defense",
	"Sharply lowers opponent's accuracy",
	"Sharply lowers opponent's evasion",
	"Double's the team's defense for 5 turns",
	"Poisons opponent",
	"Paralyzes opponent",
	"%d%% chance to lower opponent's attack",
	"%d%% chance to lower opponent's defense",
	"%d%% chance to lower opponent's speed",
	"%d%% chance to lower opponent's special attack",
	"%d%% chance to lower opponent's special defense",
	"%d%% chance to lower opponent's accuracy",
	"%d%% chance to lower opponent's evasion",
	"User charges for one turn before attacking",
	"%d%% chance to confuse opponent",
	"Hits 2-5 times\n20%% chance to poison opponents (including steel types)",
	"",
	"If the user is above 25%% HP, removes 25%% HP and adds a Substitute that will take that much damage for the user",
	"Requires a turn to recharge after being used",
	"Increases in damage every time when used back to back and the user takes damage",
	"Copies the opponent's last used move for the duration of the battle",
	"Uses a random move",
	"Afflicts the opponent with Leech Seed, transfering 1/8 of its health to the opponent's opponent, even if the user switches out",
	"Does nothing",
	"Disables opponent's the last used move for 2-8 turns",
	"Deals user's level in damage",
	"Deals between 50%% and 150%% user's level in damage",
	"If the last damage taken was physical, do twice the damage",
	"Prevents the opponent from selecting any move except the last one it used for 2-6 turns",
	"Sets the HP of both the user and the opponent to the average of their HP",
	"Can only be used when asleep",
	"Changes the user's type into one that resists the last move it was hit by",
	"The next move used by the user will always hit, including semi-invulnerable turns of Fly and Dig",
	"Copies the opponent's last used move permanently",
	"",
	"Can only be used when asleep, uses another known move randomly",
	"Will KO the opponent if the user dies before its next turn",
	"Base power is higher when the user has less HP, from 20 to 200",
	"Decreases the PP of the opponent's last used move by 2-5",
	"Will leave opponent on 1 HP",
	"Cures status for entire team",
	"Hits before other attacks",
	"Hits 3 times, doubling in damage each hit",
	"If no item is held, steals the opponent's",
	"Prevents the opponent from switching out until the user does",
	"If the opponent is sleeping, inflicts a nightmare",
	"10%% chance to burn opponent\nWill instantly thaw the user",
	"If the user is a Ghost type, removes half of its max HP and inflicts a curse on the opponent that will deal 1/4 HP a turn\nOtherwise, raises attack and defense, lowers speed",
	"",
	"Goes first, prevents being attacked for the turn",
	"Places a layer of spikes on the opponent's side that will deal 1/8 HP when an opponent switches in\nFlying type pokemon are unaffected",
	"Ignores accuracy drops and evasion increases and removes Ghost's immunities until the opponent switches out",
	"Will cause both the user and opponent to faint in 3 turns unless they switch out",
	"Sets up a sandstorm for 5 turns",
	"Survive all damage on 1 HP until the end of the turn\nWill occur before attacks",
	"Lasts 5 turns, doubling in power each turn\nDoubles in damage if the user has used Defense Curl",
	"Sharply raise's the opponent's attack, confuses opponent",
	"Doubles in damage each turn it is used back to back",
	"Infatuates the opponent",
	"Power increases with happiness, up to a maximum of 102",
	"Can damage with 40, 80, or 120 BP, or heal the opponent for 25%% of their max HP",
	"Power increases as happiness decreases, up to a maximum of 102",
	"Prevents the user's team from being afflicted with status conditions or confusion for 5 turns",
	"50%% chance to burn opponent\nWill instantly thaw the user",
	"Ranges in power from 10 to 150, 70 is most common\nHits and deals double damage opponents in the semi-invulnerable turn of Dig",
	"Switches to another pokemon, keeping stat changes and effects like Mean Look",
	"If the target attempts to switch out, doubles in power and hits first",
	"Removes the effects of binding moves, Leech Seed, and Spikes on the user's side of the field",
	"82",
	"83",
	"Restores 50%% HP",
	"Restores 50%% HP",
	"Restores 50%% HP",
	"Varies in type and damage based on the user's DVs",
	"Sets up rain for 5 turns",
	"Sets up harsh sunlight for 5 turns",
	"%d%% chance to raise defense",
	"%d%% chance to raise attack",
	"%d%% chance to raise all stats",
	"FAKE OUT",
	"If user is above half health, removes half of its max HP and sets the user to +6 attack, otherwise removes no health and sharply raises attack",
	"Copies opponent's stat changes",
	"If the last damage taken was special, do twice the damage",
	"Raises the user's defense on the first turn, then deals damage on the second",
	"%d%% chance to flinch opponent\nHits an deals double damage against opponents in the semi-invulnerable turn of Fly",
	"Hits and deals double damage opponents in the semi-invulnerable turn of Dig",
	"Hits 2 turns after it is uesed, regardless if the user and opponent switch out",
	"Hits an deals double damage against opponents in the semi-invulnerable turn of Fly",
	"Does double damage to opponents who have used Minimize",
	"Charges on the first turn, then deals damage on the second\nDoes not need to charge in harsh sunlight\nDeals half damage in rain",
	"%d%% chance to paralyze opponent\n100%% accurate in rain, 50%% accurage in sun\nHits opponents in the semi-invulnerable turn of Fly",
	"Flees wild encounters",
	"Performs an attack with 10 BP for ever party member with HP and without a status",
	"Becomes semi-invulnerable for a turn, then attacks on the next\nWhile semi-invulnerable, invulnerable to all moves except Gust, Thunder, Twister, and Whirlwind\nGust and Twister will deal double damage",
	"Raises defense\nDoubles the power of Rollout",
}
