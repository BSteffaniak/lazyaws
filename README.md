![CI](https://github.com/BSteffaniak/lazyaws/workflows/Continuous%20Integration/badge.svg)
[![GitHub Releases](https://img.shields.io/github/downloads/BSteffaniak/lazyaws/total)](https://github.com/BSteffaniak/lazyaws/releases)
[![GitHub tag](https://img.shields.io/github/tag/BSteffaniak/lazyaws.svg)](https://github.com/BSteffaniak/lazyaws/releases/latest)

A simple terminal UI for aws cli commands, written in Go with the [gocui](https://github.com/jroimartin/gocui "gocui") library based off of [lazygit](https://github.com/jesseduffield/lazygit "lazygit") by Jesse Duffield.

## Table of contents

- [Installation](#installation)
  - [Binary releases](#binary-releases)
  - [Go](#go)
  - [Manual](#manual)
- [Usage](#usage)
  - [Keybindings](#keybindings)
- [Configuration](#configuration)
  - [Custom pagers](#configuration)
  - [Custom commands](#configuration)
- [Contributing](#contributing)
- [Donate](#donate)

## Installation

### Binary Releases

For Windows, Mac OS(10.12+) or Linux, you can download a binary release [here](../../releases).

### Go

```sh
go install github.com/BSteffaniak/lazyaws@latest
```

Please note:
If you get an error claiming that lazyaws cannot be found or is not defined, you
may need to add `~/go/bin` to your $PATH (MacOS/Linux), or `%HOME%\go\bin`
(Windows). Not to be mistaken for `C:\Go\bin` (which is for Go's own binaries,
not apps like lazyaws).

### Manual

You'll need to [install Go](https://golang.org/doc/install)

```
git clone https://github.com/BSteffaniak/lazyaws.git
cd lazyaws
go install
```

You can also use `go run main.go` to compile and run in one go (pun definitely intended)

## Usage

Call `lazyaws` in your terminal.

```sh
$ lazyaws
```

If you want, you can
also add an alias for this with `echo "alias la='lazyaws'" >> ~/.zshrc` (or
whichever rc file you're using).

### Keybindings

You can check out the list of keybindings [here](/docs/keybindings).

## Configuration

Check out the [configuration docs](docs/Config.md).

### Custom Pagers

See the [docs](docs/Custom_Pagers.md)

### Custom Commands

If lazyaws is missing a feature, there's a good chance you can implement it yourself with a custom command!

See the [docs](docs/Custom_Command_Keybindings.md)

## Contributing

We love your input! Please check out the [contributing guide](CONTRIBUTING.md).

### Debugging Locally

Run `lazyaws --debug` in one terminal tab and `lazyaws --logs` in another to view the program and its log output side by side

## Donate

If you would like to support the development of lazyaws, consider [sponsoring Jesse Duffield](https://github.com/sponsors/jesseduffield) (github is matching all donations dollar-for-dollar for 12 months)

