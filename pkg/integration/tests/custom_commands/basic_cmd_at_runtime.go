package custom_commands

import (
	"github.com/BSteffaniak/lazyaws/pkg/config"
	. "github.com/BSteffaniak/lazyaws/pkg/integration/components"
)

var BasicCmdAtRuntime = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Using a custom command provided at runtime to create a new file",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("blah")
	},
	SetupConfig: func(cfg *config.AppConfig) {},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Files().
			IsEmpty().
			IsFocused().
			Press(keys.Universal.ExecuteCustomCommand)

		t.ExpectPopup().Prompt().
			Title(Equals("Custom Command:")).
			Type("touch file.txt").
			Confirm()

		t.GlobalPress(keys.Files.RefreshFiles)

		t.Views().Files().
			IsFocused().
			Lines(
				Contains("file.txt"),
			)
	},
})
