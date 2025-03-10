package core

// **************************************
// AUTO GENERATED BY BASE_STATS_PARSER.PY
// **************************************

import (
	"github.com/Tereneckla/wowsim-wotlk/sim/core/proto"
	"github.com/Tereneckla/wowsim-wotlk/sim/core/stats"
)

const ExpertisePerQuarterPercentReduction = 8.197496
const HasteRatingPerHastePercent = 32.789989
const CritRatingPerCritChance = 45.905987
const MeleeHitRatingPerHitChance = 32.789989
const SpellHitRatingPerHitChance = 26.231993
const DefenseRatingPerDefense = 4.918498
const DodgeRatingPerDodgeChance = 45.250187
const ParryRatingPerParryChance = 45.250187
const BlockRatingPerBlockChance = 16.394995
const ResilienceRatingPerCritReductionChance = 94.271225

var CritPerAgiMaxLevel = map[proto.Class]float64{
	proto.Class_ClassUnknown:     0.0,
	proto.Class_ClassWarrior:     0.0160,
	proto.Class_ClassPaladin:     0.0192,
	proto.Class_ClassHunter:      0.0120,
	proto.Class_ClassRogue:       0.0120,
	proto.Class_ClassPriest:      0.0192,
	proto.Class_ClassDeathknight: 0.0160,
	proto.Class_ClassShaman:      0.0120,
	proto.Class_ClassMage:        0.0196,
	proto.Class_ClassWarlock:     0.0198,
	proto.Class_ClassDruid:       0.0120,
}
var ExtraClassBaseStats = map[proto.Class]stats.Stats{
	proto.Class_ClassUnknown: {},
	proto.Class_ClassWarrior: {
		stats.Mana:      0.0000,
		stats.SpellCrit: 0.0000 * CritRatingPerCritChance,
		stats.MeleeCrit: 3.1891 * CritRatingPerCritChance,
	},
	proto.Class_ClassPaladin: {
		stats.Mana:      4394.0000,
		stats.SpellCrit: 3.3355 * CritRatingPerCritChance,
		stats.MeleeCrit: 3.2685 * CritRatingPerCritChance,
	},
	proto.Class_ClassHunter: {
		stats.Mana:      5046.0000,
		stats.SpellCrit: 3.6020 * CritRatingPerCritChance,
		stats.MeleeCrit: -1.5320 * CritRatingPerCritChance,
	},
	proto.Class_ClassRogue: {
		stats.Mana:      0.0000,
		stats.SpellCrit: 0.0000 * CritRatingPerCritChance,
		stats.MeleeCrit: -0.2950 * CritRatingPerCritChance,
	},
	proto.Class_ClassPriest: {
		stats.Mana:      3863.0000,
		stats.SpellCrit: 1.2375 * CritRatingPerCritChance,
		stats.MeleeCrit: 3.1765 * CritRatingPerCritChance,
	},
	proto.Class_ClassDeathknight: {
		stats.Mana:      1000.0000,
		stats.SpellCrit: 0.0000 * CritRatingPerCritChance,
		stats.MeleeCrit: 3.1891 * CritRatingPerCritChance,
	},
	proto.Class_ClassShaman: {
		stats.Mana:      4396.0000,
		stats.SpellCrit: 2.2010 * CritRatingPerCritChance,
		stats.MeleeCrit: 2.9220 * CritRatingPerCritChance,
	},
	proto.Class_ClassMage: {
		stats.Mana:      3268.0000,
		stats.SpellCrit: 0.9075 * CritRatingPerCritChance,
		stats.MeleeCrit: 3.4540 * CritRatingPerCritChance,
	},
	proto.Class_ClassWarlock: {
		stats.Mana:      3856.0000,
		stats.SpellCrit: 1.7000 * CritRatingPerCritChance,
		stats.MeleeCrit: 2.6220 * CritRatingPerCritChance,
	},
	proto.Class_ClassDruid: {
		stats.Mana:      3496.0000,
		stats.SpellCrit: 1.8515 * CritRatingPerCritChance,
		stats.MeleeCrit: 7.4755 * CritRatingPerCritChance,
	},
}
