package sim

import (
	_ "github.com/Tereneckla/wowsim-wotlk/sim/common"
	dpsDeathKnight "github.com/Tereneckla/wowsim-wotlk/sim/deathknight/dps"
	tankDeathKnight "github.com/Tereneckla/wowsim-wotlk/sim/deathknight/tank"
	"github.com/Tereneckla/wowsim-wotlk/sim/druid/balance"
	"github.com/Tereneckla/wowsim-wotlk/sim/druid/feral"
	restoDruid "github.com/Tereneckla/wowsim-wotlk/sim/druid/restoration"
	feralTank "github.com/Tereneckla/wowsim-wotlk/sim/druid/tank"
	_ "github.com/Tereneckla/wowsim-wotlk/sim/encounters"
	"github.com/Tereneckla/wowsim-wotlk/sim/hunter"
	"github.com/Tereneckla/wowsim-wotlk/sim/mage"
	holyPaladin "github.com/Tereneckla/wowsim-wotlk/sim/paladin/holy"
	protectionPaladin "github.com/Tereneckla/wowsim-wotlk/sim/paladin/protection"
	"github.com/Tereneckla/wowsim-wotlk/sim/paladin/retribution"
	healingPriest "github.com/Tereneckla/wowsim-wotlk/sim/priest/healing"
	"github.com/Tereneckla/wowsim-wotlk/sim/priest/shadow"
	"github.com/Tereneckla/wowsim-wotlk/sim/priest/smite"
	"github.com/Tereneckla/wowsim-wotlk/sim/rogue"
	"github.com/Tereneckla/wowsim-wotlk/sim/shaman/elemental"
	"github.com/Tereneckla/wowsim-wotlk/sim/shaman/enhancement"
	restoShaman "github.com/Tereneckla/wowsim-wotlk/sim/shaman/restoration"
	"github.com/Tereneckla/wowsim-wotlk/sim/warlock"
	dpsWarrior "github.com/Tereneckla/wowsim-wotlk/sim/warrior/dps"
	protectionWarrior "github.com/Tereneckla/wowsim-wotlk/sim/warrior/protection"
)

var registered = false

func RegisterAll() {
	if registered {
		return
	}
	registered = true

	balance.RegisterBalanceDruid()
	feral.RegisterFeralDruid()
	feralTank.RegisterFeralTankDruid()
	restoDruid.RegisterRestorationDruid()
	elemental.RegisterElementalShaman()
	enhancement.RegisterEnhancementShaman()
	restoShaman.RegisterRestorationShaman()
	hunter.RegisterHunter()
	mage.RegisterMage()
	healingPriest.RegisterHealingPriest()
	shadow.RegisterShadowPriest()
	smite.RegisterSmitePriest()
	rogue.RegisterRogue()
	dpsWarrior.RegisterDpsWarrior()
	protectionWarrior.RegisterProtectionWarrior()
	holyPaladin.RegisterHolyPaladin()
	protectionPaladin.RegisterProtectionPaladin()
	retribution.RegisterRetributionPaladin()
	warlock.RegisterWarlock()
	dpsDeathKnight.RegisterDpsDeathknight()
	tankDeathKnight.RegisterTankDeathknight()
}
