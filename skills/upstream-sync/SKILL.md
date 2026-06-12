---
name: iugo-ai-upstream-sync
description: "Sync IUGO-AI fork with upstream Gentle-AI. Trigger: rebase, sync, merge upstream, update fork, pull upstream changes."
license: Apache-2.0
metadata:
  author: iugo-programming
  version: "1.0"
---

# IUGO-AI — Upstream Sync Skill

## When to Use

Load this skill whenever you need to:
- Sync the `customized` branch with upstream Gentle-AI changes
- Rebase IUGO-AI branding commits on top of new upstream code
- Resolve conflicts after upstream updates
- Check what upstream changed that might affect branding

## Repository Structure

```
main        → synced with upstream Gentleman-Programming/gentle-ai
customized  → IUGO-AI branded fork (our working branch)
```

## Sync Flow (Visual)

```
┌─────────────────────────────────────────────────────────────┐
│  Gentleman-Programming/gentle-ai (upstream)                 │
│  ─────────────────────────────────────────────              │
│  Original repo, new features, bug fixes                     │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           │ git fetch upstream
                           │ git merge upstream/main
                           ▼
┌─────────────────────────────────────────────────────────────┐
│  main (your fork)                                           │
│  ────────────────────                                        │
│  Mirror of upstream, no branding changes                    │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           │ git rebase main
                           │ (resolve conflicts, rebrand)
                           ▼
┌─────────────────────────────────────────────────────────────┐
│  customized (working branch)                                │
│  ─────────────────────────                                  │
│  IUGO-AI branded fork with all customizations               │
└─────────────────────────────────────────────────────────────┘
```

**Key rule**: main is ALWAYS a clean mirror of upstream. All branding lives in customized.

## Critical Files — Branding Hotspots

These files are the **highest risk** during upstream sync. If upstream modifies them, branding may break.

### Tier 1: Central Brand Config (NEVER let upstream overwrite)

| File | Why |
|------|-----|
| `internal/brand/brand.go` | **THE** source of truth for all IUGO branding constants |

### Tier 2: High-Risk Files (check every sync)

| File | What it controls | Upstream risk |
|------|-----------------|---------------|
| `internal/tui/styles/logo.go` | ASCII art logo | Upstream may change logo art |
| `internal/tui/styles/styles.go` | Tagline, color palette | Upstream may change tagline |
| `internal/tui/screens/welcome.go` | Welcome menu options | Upstream may add/remove menu items |
| `internal/tui/screens/complete.go` | Install completion messages | Upstream may change success/error text |
| `internal/tui/screens/uninstall.go` | Uninstall flow text | Upstream may change uninstall messages |
| `internal/app/help.go` | CLI help text | Upstream may add new commands |
| `internal/app/app.go` | CLI entry point, version output | Upstream may add new command routing |
| `internal/update/registry.go` | Tool registry (names, owners, repos) | Upstream may add new tools |
| `internal/components/persona/inject.go` | Persona agent JSON overlay | Upstream may change agent structure |
| `internal/components/theme/inject.go` | Theme name and Claude theme | Upstream may change theme format |
| `internal/components/engram/download.go` | engramOwner constant | **Must stay "Gentleman-Programming"** |
| `internal/versions/versions.go` | Package version constants | Upstream may bump versions |

### Tier 3: Medium-Risk Files (check if upstream adds features)

| File | What it controls |
|------|-----------------|
| `internal/catalog/components.go` | Component descriptions shown in TUI |
| `internal/components/sdd/inject.go` | SDD orchestrator agent name |
| `internal/components/engram/inject.go` | MCP server name |
| `internal/agentbuilder/sdd.go` | HTML markers for custom agents |
| `internal/model/types.go` | PersonaID, PresetID, ComponentID constants |
| `internal/cli/sync.go` | Persona output style filename |
| `internal/cli/doctor.go` | knownTools list |
| `internal/installcmd/resolver.go` | GGA git clone URLs |

### Tier 4: Generated Files (auto-branded via sed)

These files are easy to rebrand — just run the sed commands from the Rebrand Checklist.

| Pattern | Files |
|---------|-------|
| Go imports | All `.go` files with `github.com/gentleman-programming/gentle-ai` |
| Test strings | All `*_test.go` files |
| Markdown docs | All `docs/*.md`, `README.md`, `CONTRIBUTORS.md` |
| Golden files | `testdata/golden/*.golden`, `internal/tui/testdata/*.golden` |
| Embedded assets | `internal/assets/**/*.md`, `internal/assets/**/*.yaml` |
| CI/GitHub | `.goreleaser.yaml`, `.github/**/*` |

## Upstream Sync Procedure

### Step 1: Prepare

```bash
# Ensure upstream remote exists
git remote -v | grep upstream
# If missing:
git remote add upstream https://github.com/Gentleman-Programming/gentle-ai.git

# Fetch latest upstream
git fetch upstream
git fetch origin
```

### Step 2: Update main branch

```bash
git checkout main
git merge upstream/main
# Resolve any conflicts (should be clean if main tracks upstream)
git push origin main
```

### Step 3: Rebase customized onto updated main

```bash
git checkout customized
git rebase main
```

### Step 4: Resolve Conflicts

Conflicts will fall into these categories:

#### Category A: Branding string conflict
**Example**: Upstream changed `"gentle-ai"` to `"gentle-ai-v2"` in a file.
**Resolution**: Replace with `"iugo-ai"` (use brand constant if possible).

```bash
# Edit the file, resolve to iugo-ai value
# Then:
git add <file>
git rebase --continue
```

#### Category B: New upstream file with branding
**Example**: Upstream added a new file that contains `"gentleman"`.
**Resolution**: After rebase completes, run the Rebrand Checklist (Step 5).

#### Category C: Structural conflict (code logic)
**Example**: Upstream refactored a function you modified for branding.
**Resolution**: Accept upstream logic, then re-apply branding changes.

### Step 5: Rebrand Checklist (after rebase)

Run these checks to verify branding is intact:

```bash
# Check for unbranded references (excluding engram, go.sum, node_modules)
grep -rn "gentle-ai\|gentleman\|Gentleman-Programming" \
  --include="*.go" . \
  | grep -v "_test.go" \
  | grep -v "import" \
  | grep -v "engram" \
  | grep -v "go.sum" \
  | grep -v "node_modules" \
  | grep -v "Gentleman-Programming/engram"
```

If any matches appear, determine if they need rebranding:

| Pattern | Action |
|---------|--------|
| `gentle-ai` in user-visible string | Replace with `iugo-ai` or use `brand.ProductNameLower` |
| `gentleman` in user-visible string | Replace with `iugo` or use `brand.OrgNameLower` |
| `Gentleman-Programming` in URL | Replace with `iugo-programming` **EXCEPT for engram** |
| `gentle-ai` in comment (not user-visible) | Leave as-is or update for clarity |
| `gentleman` in Go variable name | Rename to `iugo`-prefixed (careful with syntax) |

### Step 6: Verify Build and Tests

```bash
# Build
go build ./cmd/iugo-ai/

# Run tests (expect ~46 pass, 4 pre-existing failures)
go test ./... 2>&1 | grep -E "^(ok|FAIL)"

# Check version output
./bin/iugo-ai.exe version
./bin/iugo-ai.exe help
```

### Step 7: Commit and Push

```bash
git add -A
git commit -m "sync: rebase onto upstream/main and rebrand new changes"
git push origin customized --force-with-lease
```

**IMPORTANT**: After a rebase, the local and remote branches diverge because
rebase rewrites commit hashes. You MUST use `--force-with-lease` to push.

Example state after rebase:
```
Your branch and 'origin/customized' have diverged,
and have 31 and 20 different commits each, respectively.
```

This is normal. The force push replaces the old commits on remote with the
rebased ones. `--force-with-lease` is safer than `--force` because it fails
if someone else pushed while you were rebasing.

## Quick Rebrand Commands

If upstream added new files that need rebranding:

### Go source files (non-test)
```bash
find . -name "*.go" -not -name "*_test.go" -not -path "./.git/*" -exec sed -i \
  -e 's|gentle-ai|iugo-ai|g' \
  -e 's|gentleman|iugo-agent|g' \
  -e 's|Gentleman|IUGO|g' \
  -e 's|gentle-orchestrator|iugo-orchestrator|g' \
  -e 's|gentleman-programming|iugo-programming|g' \
  -e 's|Gentleman-Programming|iugo-programming|g' \
  {} \;
```

### Test files
```bash
find . -name "*_test.go" -not -path "./.git/*" -exec sed -i \
  -e 's|gentle-ai|iugo-ai|g' \
  -e 's|gentleman|iugo-agent|g' \
  -e 's|Gentleman|IUGO|g' \
  -e 's|gentle-orchestrator|iugo-orchestrator|g' \
  -e 's|gentleman-programming|iugo-programming|g' \
  -e 's|Gentleman-Programming|iugo-programming|g' \
  {} \;
```

### Markdown files
```bash
find . -name "*.md" -not -path "./.git/*" -not -path "./node_modules/*" -exec sed -i \
  -e 's|gentle-ai|iugo-ai|g' \
  -e 's|gentleman|iugo-agent|g' \
  -e 's|Gentleman|IUGO|g' \
  -e 's|gentleman-programming|iugo-programming|g' \
  -e 's|Gentleman-Programming|iugo-programming|g' \
  {} \;
```

### Golden files
```bash
find testdata -name "*.golden" -exec sed -i \
  -e 's|gentle-ai|iugo-ai|g' \
  -e 's|gentleman|iugo-agent|g' \
  -e 's|Gentleman|IUGO|g' \
  {} \;
```

### YAML/JSON config files
```bash
find . -maxdepth 1 -name "*.yaml" -o -name "*.yml" -o -name "*.json" | xargs sed -i \
  -e 's|gentle-ai|iugo-ai|g' \
  -e 's|Gentleman-Programming|iugo-programming|g'
```

## Post-Sed Manual Fixes

After running sed, check for these common issues:

### 1. Go variable names with spaces
```bash
# BAD: sed turns gentlemanYaml into "iugo-agentYaml" (syntax error)
# GOOD: should be iugoAgentYaml
grep -rn "iugo-agent[A-Z]" --include="*.go" . | grep -v "_test.go"
```

### 2. JSON keys in Go strings
```bash
# BAD: sed turns \"gentleman\" into \"iugo-agent\" but misses escaped quotes
# Check persona inject overlay:
grep -n "openCodeAgentOverlayJSON" internal/components/persona/inject.go
```

### 3. Engram owner must stay Gentleman-Programming
```bash
# NEVER change this:
grep "engramOwner" internal/components/engram/download.go
# Should show: engramOwner = "Gentleman-Programming"
```

### 4. Engram NPM package must stay gentle-engram
```bash
grep "NpmEngramPackage" internal/brand/brand.go
# Should show: NpmEngramPackage = "gentle-engram"
```

### 5. Golden file renames
```bash
# If upstream adds new golden files with "gentleman" in name:
find testdata -name "*gentleman*" -exec sh -c 'mv "$1" "$(echo $1 | sed s/gentleman/iugo-agent/g)"' _ {} \;
```

## Engram: DO NOT FORK

Engram stays with the original owner:
- **GitHub**: `Gentleman-Programming/engram`
- **NPM**: `gentle-engram`
- **Constant**: `GentleEngram` in `internal/versions/versions.go`
- **Owner**: `engramOwner = "Gentleman-Programming"` in `internal/components/engram/download.go`
- **Registry**: `Owner: "Gentleman-Programming"` for engram in `internal/update/registry.go`

## NPM Packages: Separate Fork (if needed)

These packages need separate forks if you want to publish IUGO-branded versions:
- `gentle-pi` → `iugo-pi` (referenced in `internal/agents/pi/adapter.go`)
- `gentle-engram` → stays as `gentle-engram` (no fork needed)

Currently `iugo-pi` is referenced in code but the npm package doesn't exist yet. Either:
1. Fork and publish `iugo-pi` to npm
2. Or keep using `gentle-pi` by reverting references

## Troubleshooting

### Build fails after rebase
```bash
# Check for syntax errors from sed:
go build ./cmd/iugo-ai/ 2>&1 | head -20

# Common cause: variable names with spaces (iugo-agentYaml)
grep -rn "iugo-agent[A-Z]" --include="*.go" .
```

### Tests fail after rebase
```bash
# Run specific failing test:
go test <package> -run <TestName> -v

# Check if it's a golden file mismatch:
go test <package> -run <TestName> -v 2>&1 | grep "golden mismatch"

# Update golden files if needed:
go test <package> -run <TestName> -update
```

### TUI shows wrong text
```bash
# Check brand constants:
grep -n "ProductName\|Tagline\|OrgName" internal/brand/brand.go

# Check TUI styles:
grep -n "gentle\|Gentle\|iugo\|IUGO" internal/tui/styles/*.go
```

### New upstream command not branded
```bash
# Check help text:
grep -A 30 "func printHelp" internal/app/help.go

# Check command routing:
grep -n "case \"" internal/app/app.go
```
