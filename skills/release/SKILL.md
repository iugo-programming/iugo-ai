---
name: iugo-ai-release
description: "Create IUGO-AI release tags with upstream version tracking. Trigger: release, tag, version bump, publish."
license: Apache-2.0
metadata:
  author: iugo-programming
  version: "1.0"
---

# IUGO-AI — Release Skill

## When to Use

Load this skill whenever you need to:
- Create a new release tag for IUGO-AI
- Publish a new version to GitHub Releases
- Track which upstream version the release is based on

## Release Flow

```
┌─────────────────────────────────────────────────────────────┐
│  1. Merge customized into prod                              │
│     git checkout prod && git merge customized               │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│  2. Determine upstream version                              │
│     git describe --tags --abbrev=0 upstream/main            │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│  3. Create annotated tag on prod with upstream reference    │
│     git tag -a v1.0.0 -m "Based on upstream v1.39.4"       │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│  4. Push tag (triggers GitHub Actions)                      │
│     git push origin v1.0.0                                  │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│  5. GoReleaser builds binaries + updates brew/scoop         │
└─────────────────────────────────────────────────────────────┘
```

**Key rule**: Releases are ALWAYS tagged from the `prod` branch.

## Step 1: Check Current State

Before creating a tag, verify:

```bash
# Check you're on prod branch
git branch --show-current

# Check working tree is clean
git status

# Ensure prod is up to date with customized
git checkout prod
git merge customized

# Check latest upstream version
git fetch upstream
git describe --tags --abbrev=0 upstream/main

# Check what's new since last tag
git log --oneline $(git describe --tags --abbrev=0)..HEAD
```

## Step 2: Determine Version Numbers

### IUGO-AI version (your release):
- **Major** (v1.0.0 → v2.0.0): Breaking changes, major rebrand updates
- **Minor** (v1.0.0 → v1.1.0): New features, new upstream sync
- **Patch** (v1.0.0 → v1.0.1): Bug fixes, small adjustments

### Upstream version (reference):
Check the latest tag on upstream:
```bash
git describe --tags --abbrev=0 upstream/main
```

## Step 3: Create Annotated Tag

**ALWAYS use annotated tags** (`-a`) with a message that includes the upstream version.

### Format:
```bash
git tag -a vX.Y.Z -m "Based on upstream vA.B.C"
```

### Examples:

```bash
# First release
git tag -a v1.0.0 -m "Based on upstream v1.39.4"
git push origin v1.0.0

# After syncing with new upstream
git tag -a v1.1.0 -m "Based on upstream v1.40.0"
git push origin v1.1.0

# Bug fix release (same upstream)
git tag -a v1.0.1 -m "Based on upstream v1.39.4"
git push origin v1.0.1
```

### Tag message format (extended):

For more detail, use multi-line message:

```bash
git tag -a v1.0.0 -m "IUGO-AI v1.0.0

Based on upstream Gentle-AI v1.39.4
First release with full IUGO branding."
```

## Step 4: Push Tag

```bash
git push origin vX.Y.Z
```

This triggers the GitHub Actions release workflow, which:
1. Builds binaries for linux/darwin/windows (amd64 + arm64)
2. Creates GitHub Release with the tag message
3. Generates checksums.txt
4. Updates Homebrew tap (`iugo-programming/homebrew-tap`)
5. Updates Scoop bucket (`iugo-programming/scoop-bucket`)

## Step 5: Verify Release

After push, verify:

```bash
# Check GitHub Actions status
gh run list --limit 1

# Check release was created
gh release view vX.Y.Z

# Test install script (should point to prod branch)
curl -fsSL https://raw.githubusercontent.com/iugo-programming/iugo-ai/prod/scripts/install.sh | bash
```

### Pre-release Checklist

Before pushing the tag, verify:

```bash
# README.md install commands point to prod (not main)
grep -n "/main" README.md
# Should return no matches — all URLs should use /prod/

# Build and tests pass
go build ./cmd/iugo-ai/
go test ./... 2>&1 | grep -E "^(ok|FAIL)"
```

## Quick Reference

| Action | Command |
|--------|---------|
| Check upstream version | `git describe --tags --abbrev=0 upstream/main` |
| Create tag | `git tag -a vX.Y.Z -m "Based on upstream vA.B.C"` |
| Push tag | `git push origin vX.Y.Z` |
| List tags | `git tag --sort=-v:refname` |
| Delete tag (local) | `git tag -d vX.Y.Z` |
| Delete tag (remote) | `git push origin --delete vX.Y.Z` |

## Version History Template

Maintain a record of releases and their upstream versions:

```markdown
| IUGO-AI | Upstream | Date       | Notes                |
|---------|----------|------------|----------------------|
| v1.0.0  | v1.39.4  | 2026-06-12 | First release        |
| v1.1.0  | v1.40.0  | 2026-06-20 | Synced with upstream |
| v1.1.1  | v1.40.0  | 2026-06-22 | Bug fix              |
```

## Troubleshooting

### Tag already exists
```bash
# Delete local tag
git tag -d vX.Y.Z

# Delete remote tag
git push origin --delete vX.Y.Z

# Recreate with correct message
git tag -a vX.Y.Z -m "Based on upstream vA.B.C"
git push origin vX.Y.Z
```

### Wrong upstream version in message
```bash
# Delete and recreate
git tag -d vX.Y.Z
git push origin --delete vX.Y.Z
git tag -a vX.Y.Z -m "Based on upstream vCORRECT_VERSION"
git push origin vX.Y.Z
```

### GitHub Actions failed
```bash
# Check workflow run
gh run list --limit 1
gh run view RUN_ID

# Common issues:
# - HOMEBREW_TAP_TOKEN not set or expired
# - Binaries too large (check .goreleaser.yaml)
```
