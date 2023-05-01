package constants

type Docs struct {
	CustomPagers      string
	CustomCommands    string
	CustomKeybindings string
	Keybindings       string
	Undoing           string
	Config            string
	Tutorial          string
}

var Links = struct {
	Docs        Docs
	Issues      string
	Donate      string
	Discussions string
	RepoUrl     string
	Releases    string
}{
	RepoUrl:     "https://github.com/BSteffaniak/lazyaws",
	Issues:      "https://github.com/BSteffaniak/lazyaws/issues",
	Donate:      "https://github.com/sponsors/jesseduffield",
	Discussions: "https://github.com/BSteffaniak/lazyaws/discussions",
	Releases:    "https://github.com/BSteffaniak/lazyaws/releases",
	Docs: Docs{
		CustomPagers:      "https://github.com/BSteffaniak/lazyaws/blob/master/docs/Custom_Pagers.md",
		CustomKeybindings: "https://github.com/BSteffaniak/lazyaws/blob/master/docs/keybindings/Custom_Keybindings.md",
		CustomCommands:    "https://github.com/BSteffaniak/lazyaws/wiki/Custom-Commands-Compendium",
		Keybindings:       "https://github.com/BSteffaniak/lazyaws/blob/master/docs/keybindings",
		Undoing:           "https://github.com/BSteffaniak/lazyaws/blob/master/docs/Undoing.md",
		Config:            "https://github.com/BSteffaniak/lazyaws/blob/master/docs/Config.md",
		Tutorial:          "https://youtu.be/VDXvbHZYeKY",
	},
}
