package filter_by_path

import (
	"github.com/BSteffaniak/lazyaws/pkg/config"
	. "github.com/BSteffaniak/lazyaws/pkg/integration/components"
)

var CliArg = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Filter commits by file path, using CLI arg",
	ExtraCmdArgs: "-f filterFile",
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
	},
	SetupRepo: func(shell *Shell) {
		commonSetup(shell)
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		postFilterTest(t)
	},
})
