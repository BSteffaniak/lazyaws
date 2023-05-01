package misc

import (
	"github.com/BSteffaniak/lazyaws/pkg/config"
	. "github.com/BSteffaniak/lazyaws/pkg/integration/components"
)

var InitialOpen = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Confirms a popup appears on first opening Lazyaws",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
		config.UserConfig.DisableStartupPopups = false
	},
	SetupRepo: func(shell *Shell) {},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.ExpectPopup().Confirmation().
			Title(Equals("")).
			Content(Contains("Thanks for using lazyaws!")).
			Confirm()

		t.Views().Files().IsFocused()
	},
})
