# Hats

**A simple CLI tool for managing multiple Git identities and configuration**

Hats makes it easy to manage multiple Git configuration without manually editing complex `.gitconfig` files, It is especially useful for freelancers, contractors, or developers who need to seamlessly switch between personal, work, and client Git identities based on the directory they are working in.

## Feature & Roadmap

Hats is actively maintained and being develop in phases. Here is the current status of the project:

### Currently implement

- [x] Safe global integration `hats init`: inject a single, non destructive `include` hook into your global `~/.gitconfig` without altering your existing settings.
- [x] Profile management `hats profile`: create, list, and remove isolated Git identities (name, email) without manually editing INI files.
- [x] Directory mapping `hats rule`: bind profile to specific directories (e.g, `~/work/`) to automatically switch identities based on where your repository lives.
- [x] Smart compilation engine: automatically sorts overlapping directory rules (longest/most specific path wins) and handle trailing slashes to guarantee expected Git behavior.
- [x] Idempotent state sync: safely generate `.gitconfig` shards using a temp-and-swap filesystem strategy to prevent corruption if the CLI crashes mid-write.
- [x] Native Git evaluation: relies on Git's native `includeIf` mechanism rather than background daemons or shell aliases.

### Next features (coming soon)

- [x] Active identity detection `hats current`: instantly verify which profile is active being applied to your current working directory.
- [x] Commit signing support: pass `signingKey` and `commit.gpgsign` values through to profiles for developers using SSH and GPG commit signing.
- [ ] Version command: view current install CLI version (injected via GoReleaser).
- [x] CLI autocompletion: tab completion for Bash, Zsh, and Fish shells.

### Good to have (backlog)

- [ ] Remote URL matching: apply profiles based on the Git remote URL (e.g, automatically use a profile for any `github.com/my-company/*` clone) using Git's `hasconfig:remote.*.url:` syntax.
- [ ] Health check `hats doctor`: a diagnostic command to detect broken symlinks, syntax errors in global configuration, or overlapping manual Git configuration.
- [ ] Interactive mode: guide terminal prompts for creating profiles and rules instead of passing flags.
- [ ] windows support: full path normalization and support for windows filesystem (`C:\` vs `/`).
- [ ] Homebrew tap: 1-click installation and update for macOS/Linux users via `brew install sen9kuni/tap/hats`
- [ ] Terminal UI (TUI) Mode: a rich, Interactive, full-screen terminal interface (e.g, using Bubble Tea) for managing profiles and rules.

## ⌨️ Autocompletion

Hats supports autocompletion for Bash, Zsh, and Fish shells.

### Bash

**Temporary (current session):**

```bash
source <(hats completion bash)
```

**Permanent (Linux):**

```bash
hats completion bash | sudo tee /etc/bash_completion.d/hats > /dev/null
```

**Permanent (macOS):**

```bash
hats completion bash > $(brew --prefix)/etc/bash_completion.d/hats
```

### Zsh

**Temporary (current session):**

```bash
source <(hats completion zsh)
```

**Permanent:**
If shell completion is not already enabled in your environment, you will need to enable it by executing `echo "autoload -U compinit; compinit" >> ~/.zshrc`. Then run:

```bash
hats completion zsh > "${fpath[1]}/_hats"
```

### Fish

**Temporary (current session):**

```bash
hats completion fish | source
```

**Permanent:**

```bash
hats completion fish > ~/.config/fish/completions/hats.fish
```
