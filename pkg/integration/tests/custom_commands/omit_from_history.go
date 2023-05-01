package custom_commands

import (
	"github.com/BSteffaniak/lazyaws/pkg/config"
	. "github.com/BSteffaniak/lazyaws/pkg/integration/components"
)

var OmitFromHistory = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Omitting a runtime custom command from history if it begins with space",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("blah")
	},
	SetupConfig: func(cfg *config.AppConfig) {},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.GlobalPress(keys.Universal.ExecuteCustomCommand)
		t.ExpectPopup().Prompt().
			Title(Equals("Custom Command:")).
			Type("echo aubergine").
			Confirm()

		t.GlobalPress(keys.Universal.ExecuteCustomCommand)
		t.ExpectPopup().Prompt().
			Title(Equals("Custom Command:")).
			SuggestionLines(Contains("aubergine")).
			SuggestionLines(DoesNotContain("tangerine")).
			Type(" echo tangerine").
			Confirm()

		t.GlobalPress(keys.Universal.ExecuteCustomCommand)
		t.ExpectPopup().Prompt().
			Title(Equals("Custom Command:")).
			SuggestionLines(Contains("aubergine")).
			SuggestionLines(DoesNotContain("tangerine")).
			Cancel()
	},
})
