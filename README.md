<p align="center">
  <img width="536" src="https://user-images.githubusercontent.com/8456633/174470852-339b5011-5800-4bb9-a628-ff230aa8cd4e.png">
</p>


![CI](https://github.com/BSteffaniak/lazyaws/workflows/Continuous%20Integration/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/BSteffaniak/lazyaws)](https://goreportcard.com/report/github.com/BSteffaniak/lazyaws)
[![GolangCI](https://golangci.com/badges/github.com/BSteffaniak/lazyaws.svg)](https://golangci.com)
[![GoDoc](https://godoc.org/github.com/BSteffaniak/lazyaws?status.svg)](http://godoc.org/github.com/BSteffaniak/lazyaws)
[![GitHub Releases](https://img.shields.io/github/downloads/BSteffaniak/lazyaws/total)](https://github.com/BSteffaniak/lazyaws/releases)
[![GitHub tag](https://img.shields.io/github/tag/BSteffaniak/lazyaws.svg)](https://github.com/BSteffaniak/lazyaws/releases/latest)
[![homebrew](https://img.shields.io/homebrew/v/lazyaws)](https://github.com/Homebrew/homebrew-core/blob/master/Formula/lazyaws.rb)

A simple terminal UI for git commands, written in Go with the [gocui](https://github.com/jroimartin/gocui "gocui") library.

![Gif](../assets/staging.gif)

## Sponsors

<p align="center">
 Maintenance of this project is made possible by all the <a href="https://github.com/BSteffaniak/lazyaws/graphs/contributors">contributors</a> and <a href="https://github.com/sponsors/jesseduffield">sponsors</a>. If you'd like to sponsor this project and have your avatar or company logo appear below <a href="https://github.com/sponsors/jesseduffield">click here</a>. 💙
</p>

<p align="center">
<!-- sponsors --><a href="https://github.com/intabulas"><img src="https://github.com/intabulas.png" width="60px" alt="" /></a><a href="https://github.com/piot"><img src="https://github.com/piot.png" width="60px" alt="" /></a><a href="https://github.com/rgwood"><img src="https://github.com/rgwood.png" width="60px" alt="" /></a><a href="https://github.com/oliverguenther"><img src="https://github.com/oliverguenther.png" width="60px" alt="" /></a><a href="https://github.com/pawanjay176"><img src="https://github.com/pawanjay176.png" width="60px" alt="" /></a><a href="https://github.com/bdach"><img src="https://github.com/bdach.png" width="60px" alt="" /></a><a href="https://github.com/davidklsn"><img src="https://github.com/davidklsn.png" width="60px" alt="" /></a><a href="https://github.com/naoey"><img src="https://github.com/naoey.png" width="60px" alt="" /></a><a href="https://github.com/jryom"><img src="https://github.com/jryom.png" width="60px" alt="" /></a><a href="https://github.com/carstengehling"><img src="https://github.com/carstengehling.png" width="60px" alt="" /></a><a href="https://github.com/ceuk"><img src="https://github.com/ceuk.png" width="60px" alt="" /></a><a href="https://github.com/akospwc"><img src="https://github.com/akospwc.png" width="60px" alt="" /></a><a href="https://github.com/Xetera"><img src="https://github.com/Xetera.png" width="60px" alt="" /></a><a href="https://github.com/HoldenLucas"><img src="https://github.com/HoldenLucas.png" width="60px" alt="" /></a><a href="https://github.com/barbados-clemens"><img src="https://github.com/barbados-clemens.png" width="60px" alt="" /></a><a href="https://github.com/nartc"><img src="https://github.com/nartc.png" width="60px" alt="" /></a><a href="https://github.com/"><img src="https://github.com/.png" width="60px" alt="" /></a><a href="https://github.com/matejcik"><img src="https://github.com/matejcik.png" width="60px" alt="" /></a><a href="https://github.com/lucatume"><img src="https://github.com/lucatume.png" width="60px" alt="" /></a><a href="https://github.com/zach-fuller"><img src="https://github.com/zach-fuller.png" width="60px" alt="" /></a><a href="https://github.com/davdroman"><img src="https://github.com/davdroman.png" width="60px" alt="" /></a><a href="https://github.com/KowalskiPiotr98"><img src="https://github.com/KowalskiPiotr98.png" width="60px" alt="" /></a><a href="https://github.com/nicholascloud"><img src="https://github.com/nicholascloud.png" width="60px" alt="" /></a><a href="https://github.com/topher200"><img src="https://github.com/topher200.png" width="60px" alt="" /></a><a href="https://github.com/PhotonQuantum"><img src="https://github.com/PhotonQuantum.png" width="60px" alt="" /></a><a href="https://github.com/GitSquared"><img src="https://github.com/GitSquared.png" width="60px" alt="" /></a><a href="https://github.com/ava1ar"><img src="https://github.com/ava1ar.png" width="60px" alt="" /></a><a href="https://github.com/pedropombeiro"><img src="https://github.com/pedropombeiro.png" width="60px" alt="" /></a><a href="https://github.com/minidfx"><img src="https://github.com/minidfx.png" width="60px" alt="" /></a><a href="https://github.com/JoeKlemmer"><img src="https://github.com/JoeKlemmer.png" width="60px" alt="" /></a><a href="https://github.com/ColonelBucket8"><img src="https://github.com/ColonelBucket8.png" width="60px" alt="" /></a><a href="https://github.com/mutewinter"><img src="https://github.com/mutewinter.png" width="60px" alt="" /></a><a href="https://github.com/tobi"><img src="https://github.com/tobi.png" width="60px" alt="" /></a><a href="https://github.com/benbfortis"><img src="https://github.com/benbfortis.png" width="60px" alt="" /></a><a href="https://github.com/jakewarren"><img src="https://github.com/jakewarren.png" width="60px" alt="" /></a><a href="https://github.com/tgpholly"><img src="https://github.com/tgpholly.png" width="60px" alt="" /></a><a href="https://github.com/jisantuc"><img src="https://github.com/jisantuc.png" width="60px" alt="" /></a><a href="https://github.com/zabil"><img src="https://github.com/zabil.png" width="60px" alt="" /></a><a href="https://github.com/bitprophet"><img src="https://github.com/bitprophet.png" width="60px" alt="" /></a><a href="https://github.com/tayleighr"><img src="https://github.com/tayleighr.png" width="60px" alt="" /></a><a href="https://github.com/Novakov"><img src="https://github.com/Novakov.png" width="60px" alt="" /></a><a href="https://github.com/mthuggett"><img src="https://github.com/mthuggett.png" width="60px" alt="" /></a><a href="https://github.com/portothree"><img src="https://github.com/portothree.png" width="60px" alt="" /></a><a href="https://github.com/farzadmf"><img src="https://github.com/farzadmf.png" width="60px" alt="" /></a><a href="https://github.com/nekhaevskiy"><img src="https://github.com/nekhaevskiy.png" width="60px" alt="" /></a><a href="https://github.com/reivilibre"><img src="https://github.com/reivilibre.png" width="60px" alt="" /></a><a href="https://github.com/shtirlic"><img src="https://github.com/shtirlic.png" width="60px" alt="" /></a><a href="https://github.com/andreaskurth"><img src="https://github.com/andreaskurth.png" width="60px" alt="" /></a><a href="https://github.com/froody"><img src="https://github.com/froody.png" width="60px" alt="" /></a><!-- sponsors -->
</p>

## Elevator Pitch

Rant time: You've heard it before, git is _powerful_, but what good is that power when everything is so damn hard to do? Interactive rebasing requires you to edit a goddamn TODO file in your editor? _Are you kidding me?_ To stage part of a file you need to use a command line program to step through each hunk and if a hunk can't be split down any further but contains code you don't want to stage, you have to edit an arcane patch file _by hand_? _Are you KIDDING me?!_ Sometimes you get asked to stash your changes when switching branches only to realise that after you switch and unstash that there weren't even any conflicts and it would have been fine to just checkout the branch directly? _YOU HAVE GOT TO BE KIDDING ME!_

If you're a mere mortal like me and you're tired of hearing how powerful git is when in your daily life it's a powerful pain in your ass, lazyaws might be for you.

## Table of contents

- [Installation](#installation)
  - [Binary releases](#binary-releases)
  - [Homebrew](#homebrew)
  - [MacPorts](#macports)
  - [Void Linux](#void-linux)
  - [Scoop (Windows)](#scoop-windows)
  - [Arch Linux](#arch-linux)
  - [Fedora and RHEL](#fedora-and-rhel)
  - [Solus Linux](#solus-linux)
  - [Ubuntu](#ubuntu)
  - [Funtoo Linux](#funtoo-linux)
  - [FreeBSD](#freebsd)
  - [Conda](#conda)
  - [Go](#go)
  - [Chocolatey (Windows)](#chocolatey-windows)
  - [Manual](#manual)
- [Usage](#usage)
  - [Keybindings](#keybindings)
  - [Changing directory on exit](#changing-directory-on-exit)
  - [Undo/Redo](#undoredo)
- [Configuration](#configuration)
  - [Custom pagers](#configuration)
  - [Custom commands](#configuration)
- [Tutorials](#tutorials)
- [Cool Features](#cool-features)
- [Contributing](#contributing)
- [Donate](#donate)
- [Alternatives](#alternatives)

Github Sponsors is matching all donations dollar-for-dollar for 12 months so if you're feeling generous consider [sponsoring me](https://github.com/sponsors/jesseduffield)

[<img src="https://i.imgur.com/sVEktDn.png">](https://youtu.be/CPLdltN7wgE)

## Installation

### Binary Releases

For Windows, Mac OS(10.12+) or Linux, you can download a binary release [here](../../releases).

### Homebrew

Normally the lazyaws formula can be found in the Homebrew core but we suggest you tap our formula to get the frequently updated one. It works with Linux, too.

Tap:

```
brew install BSteffaniak/lazyaws/lazyaws
```

Core:

```
brew install lazyaws
```

### MacPorts

Latest version built from github releases.
Tap:

```
sudo port install lazyaws
```

### Void Linux

Packages for Void Linux are available in the distro repo

They follow upstream latest releases

```sh
sudo xbps-install -S lazyaws
```

### Scoop (Windows)

You can install `lazyaws` using [scoop](https://scoop.sh/). It's in the `extras` bucket:

```sh
# Add the extras bucket
scoop bucket add extras

# Install lazyaws
scoop install lazyaws
```

### Arch Linux

Packages for Arch Linux are available via pacman and AUR (Arch User Repository).

There are two packages. The stable one which is built with the latest release
and the git version which builds from the most recent commit.

- Stable: `sudo pacman -S lazyaws`
- Development: <https://aur.archlinux.org/packages/lazyaws-git/>

Instruction of how to install AUR content can be found here:
<https://wiki.archlinux.org/index.php/Arch_User_Repository>

### Fedora and RHEL

Packages for Fedora/RHEL and CentOS Stream are available via [Copr](https://copr.fedorainfracloud.org/coprs/atim/lazyaws/) (Cool Other Package Repo).

```sh
sudo dnf copr enable atim/lazyaws -y
sudo dnf install lazyaws
```

### Solus Linux

```sh
sudo eopkg install lazyaws
```

### Ubuntu

```sh
LAZYAWS_VERSION=$(curl -s "https://api.github.com/repos/BSteffaniak/lazyaws/releases/latest" | grep -Po '"tag_name": "v\K[^"]*')
curl -Lo lazyaws.tar.gz "https://github.com/BSteffaniak/lazyaws/releases/latest/download/lazyaws_${LAZYAWS_VERSION}_Linux_x86_64.tar.gz"
tar xf lazyaws.tar.gz lazyaws
sudo install lazyaws /usr/local/bin
```

Verify the correct installation of lazyaws:

```sh
lazyaws --version
```

### Funtoo Linux

Funtoo Linux has an autogenerated lazyaws package in [dev-kit](https://github.com/funtoo/dev-kit/tree/1.4-release/dev-vcs/lazyaws):

```sh
sudo emerge dev-vcs/lazyaws
```

### FreeBSD

```sh
pkg install lazyaws
```

### Conda

Released versions are available for different platforms, see <https://anaconda.org/conda-forge/lazyaws>

```sh
conda install -c conda-forge lazyaws
```

### Go

```sh
go install github.com/BSteffaniak/lazyaws@latest
```

Please note:
If you get an error claiming that lazyaws cannot be found or is not defined, you
may need to add `~/go/bin` to your $PATH (MacOS/Linux), or `%HOME%\go\bin`
(Windows). Not to be mistaken for `C:\Go\bin` (which is for Go's own binaries,
not apps like lazyaws).

### Chocolatey (Windows)

You can install `lazyaws` using [Chocolatey](https://chocolatey.org/):

```sh
choco install lazyaws
```

### Manual

You'll need to [install Go](https://golang.org/doc/install)

```
git clone https://github.com/BSteffaniak/lazyaws.git
cd lazyaws
go install
```

You can also use `go run main.go` to compile and run in one go (pun definitely intended)

## Usage

Call `lazyaws` in your terminal inside a git repository.

```sh
$ lazyaws
```

If you want, you can
also add an alias for this with `echo "alias lg='lazyaws'" >> ~/.zshrc` (or
whichever rc file you're using).

### Keybindings

You can check out the list of keybindings [here](/docs/keybindings).

### Changing Directory On Exit

If you change repos in lazyaws and want your shell to change directory into that repo on exiting lazyaws, add this to your `~/.zshrc` (or other rc file):

```
lg()
{
    export LAZYAWS_NEW_DIR_FILE=~/.lazyaws/newdir

    lazyaws "$@"

    if [ -f $LAZYAWS_NEW_DIR_FILE ]; then
            cd "$(cat $LAZYAWS_NEW_DIR_FILE)"
            rm -f $LAZYAWS_NEW_DIR_FILE > /dev/null
    fi
}
```

Then `source ~/.zshrc` and from now on when you call `lg` and exit you'll switch directories to whatever you were in inside lazyaws. To override this behaviour you can exit using `shift+Q` rather than just `q`.

### Undo/Redo

See the [docs](/docs/Undoing.md)

## Configuration

Check out the [configuration docs](docs/Config.md).

### Custom Pagers

See the [docs](docs/Custom_Pagers.md)

### Custom Commands

If lazyaws is missing a feature, there's a good chance you can implement it yourself with a custom command!

See the [docs](docs/Custom_Command_Keybindings.md)

## Tutorials

- [Video Tutorial](https://youtu.be/VDXvbHZYeKY)
- [Rebase Magic Video Tutorial](https://youtu.be/4XaToVut_hs)
- [Twitch Stream](https://www.twitch.tv/jesseduffield)

## Cool features

- Adding files easily
- Resolving merge conflicts
- Easily check out recent branches
- Scroll through logs/diffs of branches/commits/stash
- Quick pushing/pulling
- Squash down and reword commits

### Resolving merge conflicts

![Gif](../assets/resolving-merge-conflicts.gif)

### Interactive Rebasing

![Interactive Rebasing](../assets/rebase.gif)

## Contributing

We love your input! Please check out the [contributing guide](CONTRIBUTING.md).
For contributor discussion about things not better discussed here in the repo, join the discord channel

<a href="https://discord.gg/ehwFt2t4wt"><img src='../assets/discord.png' width='75'></a>

Check out this [video](https://www.youtube.com/watch?v=kNavnhzZHtk) walking through the creation of a small feature in lazyaws if you want an idea of where to get started.

### Debugging Locally

Run `lazyaws --debug` in one terminal tab and `lazyaws --logs` in another to view the program and its log output side by side

## Donate

If you would like to support the development of lazyaws, consider [sponsoring me](https://github.com/sponsors/jesseduffield) (github is matching all donations dollar-for-dollar for 12 months)

## FAQ

### What do the commit colors represent?

- Green: the commit is included in the master branch
- Yellow: the commit is not included in the master branch
- Red: the commit has not been pushed to the upstream branch

## Shameless Plug

If you want to see what I (Jesse) am up to in terms of development, follow me on
[twitter](https://twitter.com/DuffieldJesse) or check out my [blog](https://jesseduffield.com/)

## Alternatives

If you find that lazyaws doesn't quite satisfy your requirements, these may be a better fit:

- [GitUI](https://github.com/Extrawurst/gitui)
- [tig](https://github.com/jonas/tig)
