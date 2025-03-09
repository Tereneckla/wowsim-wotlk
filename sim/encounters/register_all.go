package encounters

import (
	"github.com/Tereneckla/wowsim-wotlk/sim/core"
	"github.com/Tereneckla/wowsim-wotlk/sim/encounters/icc"
	"github.com/Tereneckla/wowsim-wotlk/sim/encounters/naxxramas"
	"github.com/Tereneckla/wowsim-wotlk/sim/encounters/toc"
	"github.com/Tereneckla/wowsim-wotlk/sim/encounters/ulduar"
)

func init() {
	naxxramas.Register()
	ulduar.Register()
	toc.Register()
	icc.Register()
}

func AddSingleTargetBossEncounter(presetTarget *core.PresetTarget) {
	core.AddPresetTarget(presetTarget)
	core.AddPresetEncounter(presetTarget.Config.Name, []string{
		presetTarget.Path(),
	})
}
