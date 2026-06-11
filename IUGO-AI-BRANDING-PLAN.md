# Plan de Branding: Gentle-AI → IUGO-AI

## Estrategia

**Brand config centralizado + commits aislados por fase.**

Todo valor visible se define una sola vez en `internal/brand/`. El código Go importa de ahí. Cada fase es un commit independiente para facilitar rebase con upstream.

## Mapping Rules

| Upstream | IUGO |
|----------|------|
| `gentle-ai` | `iugo-ai` |
| `gentle_ai` | `iugo_ai` |
| `gentleman-programming` | `iugo-programming` |
| `gentleman` (marca) | `iugo` |
| `gentle` (prefijo marca) | `iugo` |
| `Gentle-AI` | `IUGO-AI` |
| `Gentleman` (marca) | `IUGO` |
| `Gentle` (prefijo marca) | `IUGO` |
| `.gentle-ai` (dir config) | `.iugo-ai` |
| `gentle-pi` (npm) | `iugo-pi` |
| `gentle-engram` (npm) | `iugo-engram` |
| `gentle-orchestrator` | `iugo-orchestrator` |

### Exceptions (NO cambiar)

- `gentle` como adjetivo inglés en textos de ayuda/doc (no es marca)
- `gentleman` cuando se refiere al rol/persona del agente, no a la marca del producto (ej: "Senior Architect persona" vs "Gentleman AI product")
- Go import paths en `go.sum` (se regeneran solos)

---

## Fase 1: Brand Config Centralizado

**Commit**: `feat(brand): add centralized brand config`

Crear `internal/brand/brand.go`:

```go
package brand

const (
    ProductName       = "IUGO-AI"
    ProductNameLower  = "iugo-ai"
    BinaryName        = "iugo-ai"
    OrgName           = "IUGO"
    OrgNameLower      = "iugo"
    GitHubOwner       = "iugo-programming"
    GitHubRepo        = "iugo-ai"
    ModulePath        = "github.com/iugo-programming/iugo-ai"
    ConfigDir         = ".iugo-ai"
    Tagline           = "Ecosystem, Frameworks, Workflows"
    ThemeName         = "iugo-kanagawa"
    AgentName         = "iugo-agent"
    OrchestratorName  = "iugo-orchestrator"
    MCPServerName     = "iugo-ai-engram"
    NpmPiPackage      = "iugo-pi"
    NpmEngramPackage  = "iugo-engram"
    DocsURL           = "https://github.com/iugo-programming/iugo-ai"
)
```

**Por qué primero**: Todas las fases siguientes importan de aquí. Si upstream agrega un nuevo string visible, solo agregás una constante nueva.

---

## Fase 2: Module Path y Directorios

**Commit**: `refactor: rename module path and cmd directory`

Cambios:
1. `go.mod`: `module github.com/gentleman-programming/gentle-ai` → `github.com/iugo-programming/iugo-ai`
2. Renombrar directorio: `cmd/gentle-ai/` → `cmd/iugo-ai/`
3. **Replace all** en `.go` files: `github.com/gentleman-programming/gentle-ai` → `github.com/iugo-programming/iugo-ai`
4. Ejecutar `go mod tidy`

**Impacto**: ~615 occurrences en imports Go. Es find/replace mecánico.

---

## Fase 3: Config Paths del Sistema

**Commit**: `refactor: use brand config for system paths`

Archivos a modificar:
- `internal/state/state.go`: `const stateDir = ".gentle-ai"` → usar `brand.ConfigDir`
- `internal/app/app.go`: `filepath.Join(homeDir, ".gentle-ai", "backups")` → usar brand
- `internal/backup/manifest.go`: `filepath.Join(home, ".gentle-ai", "backups")` → usar brand
- `internal/cli/doctor.go`: `filepath.Join(homeDir, ".gentle-ai")` → usar brand
- `internal/cli/restore.go`: `homeDir + "/.gentle-ai/backups"` → usar brand
- `internal/cli/run.go`: backup paths → usar brand
- `internal/cli/sync.go`: backup paths → usar brand
- `internal/components/uninstall/service.go`: paths → usar brand
- `internal/opencode/models.go`: cache path → usar brand
- `internal/update/upgrade/executor.go`: backup paths → usar brand
- `internal/components/filemerge/writer.go`: tmp prefix `.gentle-ai-*.tmp` → usar brand

---

## Fase 4: TUI (Logo, Tagline, Estilos)

**Commit**: `feat(tui): rebrand TUI for IUGO`

Archivos:
- `internal/tui/styles/styles.go`:
  - Reemplazar `logoLines` con ASCII art de IUGO
  - `Tagline()`: usar `brand.ProductName` + brand tagline
  - Colores: mantener Rose Pine o cambiar a paleta IUGO (opcional)
- `internal/tui/screens/welcome.go`: sin cambios directos (usa styles)
- `internal/tui/screens/complete.go`: texto `"gentle-ai"` → usar brand
- `internal/tui/screens/uninstall.go`: texto → usar brand
- `internal/tui/screens/upgrade_sync.go`: texto → usar brand
- `internal/tui/screens/dependency_tree.go`: `"gentle-pi"` → usar brand

---

## Fase 5: CLI Output

**Commit**: `feat(cli): rebrand CLI help and messages`

Archivos:
- `internal/app/help.go`: todo el texto visible con `brand.ProductName`
- `internal/app/app.go`:
  - `fmt.Fprintf(stdout, "gentle-ai %s\n", Version)` → brand
  - Error messages → brand
  - `func gentleAIUpgradeVersionFromTUI` → rename (interno, no visible)
- `internal/app/selfupdate.go`: tool name `"gentle-ai"` → brand
- `internal/cli/uninstall.go`: texto visible → brand
- `internal/cli/doctor.go`: `knownTools` → brand

---

## Fase 6: Update Registry

**Commit**: `feat(update): rebrand update registry`

Archivos:
- `internal/update/registry.go`:
  - `Name: "gentle-ai"` → `brand.BinaryName`
  - `Owner: "Gentleman-Programming"` → `brand.GitHubOwner`
  - `Repo: "gentle-ai"` → `brand.GitHubRepo`
- `internal/update/instructions.go`: texto → brand
- `internal/update/types.go`: solo si hay strings visibles

---

## Fase 7: Componentes con Nombres de Marca

**Commit**: `feat(components): rebrand component names`

Archivos:
- `internal/catalog/components.go`:
  - `"Gentleman, neutral or custom behavior"` → brand
  - `"Gentleman Guardian Angel"` → brand
  - `"Gentleman Kanagawa theme overlay"` → brand
  - `"Claude Code Gentleman custom theme"` → brand
- `internal/components/persona/inject.go`:
  - Agent key `"gentleman"` → `brand.AgentName`
  - Description → brand
- `internal/components/theme/inject.go`:
  - `Name: "Gentleman"` → brand
  - theme `"gentleman-kanagawa"` → `brand.ThemeName`
  - filename `"gentleman.json"` → brand
- `internal/components/engram/inject.go`:
  - MCP name `"gentle-ai-engram"` → brand
- `internal/agentbuilder/sdd.go`:
  - HTML markers `<!-- gentle-ai:custom-agent:... -->` → brand
- `internal/model/types.go`:
  - `ComponentOpenCodeGentleLogo` → rename

---

## Fase 8: NPM Package References

**Commit**: `feat(deps): update npm package references`

Archivos:
- `internal/agents/pi/adapter.go`:
  - `"npm:gentle-pi"` → `brand.NpmPiPackage`
  - `"npm:gentle-engram"` → `brand.NpmEngramPackage`
- `internal/versions/versions.go`:
  - `// renovate: datasource=npm depName=gentle-engram` → nuevo nombre
  - `const GentleEngram` → rename const
- `internal/tui/screens/dependency_tree.go`: package names → brand
- `internal/skillregistry/registry.go`: comment → brand

**Nota**: Los packages npm en sí mismos (`gentle-pi`, `gentle-engram`) son repos externos. Esta fase solo cambia las referencias. Los packages necesitan fork aparte.

---

## Fase 9: Assets Embebidos (Markdown)

**Commit**: `docs(assets): rebrand embedded markdown assets`

Archivos en `internal/assets/`:
- Todos los `sdd-orchestrator.md` (~10 archivos): reemplazar "gentle-ai", "gentleman", "gentle-orchestrator"
- Todos los `persona-gentleman.md` (~6 archivos): rename archivo + contenido
- `output-style-gentleman.md` (~2 archivos): rename + contenido
- `claude/sdd-*.md` (agents + commands): contenido
- `kimi/agents/gentleman.yaml`: rename + contenido
- `skills/sdd-archive/SKILL.md`: contenido
- `codex/engram-*.md`: contenido
- Tests: `internal/assets/assets_test.go` → actualizar referencias

---

## Fase 10: Agent Adapters (File Names Generated)

**Commit**: `refactor(agents): rebrand generated file names`

Archivos:
- `internal/agents/cursor/adapter.go`: `gentle-ai.mdc` → brand
- `internal/agents/kiro/adapter.go`: `gentle-ai.md` → brand
- `internal/agents/vscode/adapter.go`: `gentle-ai.instructions.md` → brand
- `internal/agents/kimi/adapter.go`: `gentleman.yaml` → brand

---

## Fase 11: Imágenes

**Commit**: `docs(brand): replace logo and banner images`

- Copiar `internal/custom/docs/assets/brand/iugo-ai-banner.png` → `docs/assets/brand/`
- Copiar `internal/custom/docs/assets/brand/iugo-ai-logo.png` → `docs/assets/brand/`
- Renombrar archivos a `iugo-ai-*`
- Actualizar referencias en `README.md`

---

## Fase 12: Documentación

**Commit**: `docs: rebrand all documentation`

Archivos:
- `README.md`: banner, badges, título, links, texto
- `CONTRIBUTORS.md`: nombre del proyecto
- `AGENTS.md`: instrucciones
- `docs/*.md` (~15 archivos): todas las menciones
- `package.json`: name, description, repository, bugs, homepage

---

## Fase 13: GitHub / CI

**Commit**: `ci: rebrand GitHub workflows and templates`

Archivos:
- `.goreleaser.yaml`: `project_name`, `binary`, `main`
- `.github/workflows/release.yml`: si hay referencias
- `.github/workflows/ci.yml`: si hay referencias
- `.github/ISSUE_TEMPLATE/*.yml`: si hay referencias
- `.engram/config.json`: `project_name`

---

## Fase 14: Tests

**Commit**: `test: update tests for IUGO branding`

Archivos con más occurrences:
- `internal/components/persona/inject_test.go` (242 refs)
- `internal/components/sdd/inject_test.go` (159 refs)
- `internal/tui/model_test.go` (82 refs)
- `internal/update/upgrade/strategy_test.go` (75 refs)
- `internal/update/check_test.go` (63 refs)
- Otros ~20 archivos de test

**Estrategia**: find/replace mecánico después de las fases 1-13.

---

## Orden de Ejecución Recomendado

```
1.  Fase 1  (brand config)        ← base, sin dependencias
2.  Fase 2  (module path)         ← rompe todo, hacerlo temprano
3.  Fase 3  (config paths)        ← usa brand config
4.  Fase 4  (TUI)                 ← usa brand config
5.  Fase 5  (CLI output)          ← usa brand config
6.  Fase 6  (update registry)     ← usa brand config
7.  Fase 7  (components)          ← usa brand config
8.  Fase 8  (npm refs)            ← usa brand config
9.  Fase 9  (assets embebidos)    ← find/replace
10. Fase 10 (agent adapters)      ← usa brand config
11. Fase 11 (imágenes)            ← copy files
12. Fase 12 (documentación)       ← find/replace
13. Fase 13 (CI/GitHub)           ← find/replace
14. Fase 14 (tests)               ← find/replace masivo
```

---

## Git Strategy para Upstream Sync

```bash
# Setup inicial
git remote add upstream https://github.com/Gentleman-Programming/gentle-ai.git

# Para actualizar con upstream
git fetch upstream
git rebase upstream/main

# Si hay conflictos en branding:
# - Conflictos en import paths: resolver a iugo-programming
# - Conflictos en strings visibles: resolver usando brand config
# - Conflictos en archivos nuevos de upstream: adaptar al brand config
```

**Regla**: Cada commit de branding debe ser lo más pequeño y aislado posible. Así, si upstream mueve un archivo, el conflicto es contenido, no estructural.

---

## Validación Post-Implementación

```bash
# 1. Compila sin errores
go build ./cmd/iugo-ai/

# 2. Tests pasan
go test ./...

# 3. No quedan referencias a gentle-ai (excepto en go.sum y comentarios legítimos)
grep -r "gentle-ai\|gentleman-programming" --include="*.go" . | grep -v "_test.go" | grep -v "go.sum"

# 4. El binario se llama iugo-ai
./iugo-ai version

# 5. El TUI muestra IUGO-AI
./iugo-ai
```

---

## Notas sobre Packages NPM Externos

`gentle-pi` y `gentle-engram` son packages npm publicados en repos separados. Para branding completo:

1. Fork `gentle-pi` → `iugo-pi` (repo nuevo)
2. Fork `gentle-engram` → `iugo-engram` (repo nuevo)
3. Publicar en npm con nuevos nombres
4. Actualizar `internal/versions/versions.go` con nuevos dep names

Esto es **independiente** del fork de gentle-ai y puede hacerse después.
