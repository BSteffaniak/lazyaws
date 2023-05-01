package filter_by_path

import (
	"github.com/BSteffaniak/lazyaws/pkg/config"
	. "github.com/BSteffaniak/lazyaws/pkg/integration/components"
)

var TypeFile = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Filter commits by file path, by finding file in UI and filtering on it",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
	},
	SetupRepo: func(shell *Shell) {
		commonSetup(shell)
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Files().
			IsFocused().
			Press(keys.Universal.FilteringMenu)

		t.ExpectPopup().Menu().
			Title(Equals("Filtering")).
			Select(Contains("enter path to filter by")).
			Confirm()

		t.ExpectPopup().Prompt().
			Title(Equals("Enter path:")).
			Type("filterF").
			SuggestionLines(Equals("filterFile")).
			ConfirmFirstSuggestion()

		postFilterTest(t)
	},
})
