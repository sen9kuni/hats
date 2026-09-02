# Hats

**A simple CLI tool for managing multiple Git identities and configuration**

Hats makes it easy to manage multiple Git configuration without manually editing complex `.gitconfig` files, It is especially useful for freelancers, contractors, or developers who need to seamlessly switch between personal, work, and client Git identities based on the directory they are working in.

## Why Hats?

Instead of running background daemons, relying on brittle shell aliases, or overwriting your global
identity, Hats uses Git's native `includeIf` mechanism. It safely inject exactly **one** line into your global config and manages the rest dynamically, guaranteeing 100% compatibility with standard Git behavior.

## Installation

Download the latest pre-compiled binary for macOS or Linux from [Releases page](https://github.com/sen9kuni/hats/releases).

_(Homebrew tap coming soon!)_

## Quick Start

**1. Initialize the environment**
Safely injects the Hats router into your global `~/.gitconfig`.

```bash
hats init
```

**2. Create a profile**
Define a Git identity (name, email, and optionally a GPG signing key).

```bash
hats profile add work -n "Your Name" -e "you@work.com"
```

**3. Map into a directory**
Tell Git to automatically use the `work` profile anytime you are inside `~/Projects/work`

```bash
hats rule add ~/Projects/work work
```

That's it. When you `cd` into that directory and type `git commit`, Git will automatically use `you@work.com`.

## Commands Overview

### Profiles

- `hats profile add <id> -n "Name" -e "Email"` : Create a new identity. Add `-k "Key"` to enforce GPG commit signing.

- `hats profile list` : View all managed profiles.
- `hats profile remove <id>` : Delete a profile and clean up its files.

### Rules

- `hats rule add <path> <profile-id>` : Map a profile to a directory path.
- `hats rule list` : View all active directory mappings.
- `hats rule remove <path>` : Remove a directory mapping.

### Utility

- `hats current` : Asks Git which profile is actively being applied to your current working directory.
- `hats version` : Print the installed CLI version.

## Autocompletion

Hats suppors Autocompletion for Bash, Zsh, and Fish shells.

### Bash

**Temporary (current session)**

```bash
source <(hats completion bash)
```

**Permanent (Linux)**

```bash
hats completion bash | sudo tee /etc/bash_completion.d/hats > /dev/null
```

**Permanent (macOS)**

```bash
hats completion bash > $(brew --prefix)/etc/bash_completion.d/hats
```

### Zsh

**Temporary (current session)**

```bash
source <(hats completion zsh)
```

**Permanent**
If shell completion is not already enabled in your environment, you will need to enable it by executing
`echo "autoload -U compinit; compinit" >> ~/.zshrc`. Then run:

```bash
hats completion zsh > "${fpath[1]}/_hats"
```

### Fish

**Temporary (current session)**

```bash
hats completion fish | source
```

**Permanent**

```bash
hats completion fish > ~/.config/fish/completions/hats.fish
```

## Features & Roadmap

Hats is actively maintained and being developed in phases.

### Currently Implemented

- [x] **Safe Global Integration**: Injects a single, non-destructive hook into `~/.gitconfig`.
- [x] **Profile & Rule Management**: Create, list, and remove isolated Git identities mapped to directories.
- [x] **Smart Compilation Engine**: Handles overlapping directory rules and trailing slashes perfectly.
- [x] **Idempotent State Sync**: Generates config shards safely using a temp-and-swap filesystem strategy.
- [x] **Active Identity Detection**: Verify the active profile using `hats current`.
- [x] **Commit Signing Support**: Pass `signingkey` and enforce GPG/SSH signatures per-profile.
- [x] **CLI Autocompletion**: Native tab completion for Bash, Zsh, and Fish.

### Backlog

- [ ] **Health Check (`hats doctor`)**: A diagnostic command to detect broken symlinks, syntax errors, or missing config shards.
- [ ] **Terminal UI (TUI) Mode**: A rich, interactive terminal interface for managing profiles visually.
- [ ] **Interactive Mode**: Guided terminal prompts for creating profiles (no flags required).
- [ ] **Remote URL Matching**: Apply profiles based on the Git remote URL instead of the local path.
- [ ] **Homebrew Tap**: 1-click installation via `brew install`.
- [ ] **Windows Support**: Full path normalization for Windows filesystems.

## 📄 License

MIT License
