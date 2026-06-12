// Package brand centralizes all visible brand constants for the IUGO-AI fork.
// Every user-facing string, path, and identifier should reference this package
// instead of hardcoding values. When syncing with upstream Gentle-AI, any new
// visible string should get a constant here.
package brand

const (
	// ProductName is the human-readable product name shown in TUI, help, etc.
	ProductName = "IUGO-AI"

	// ProductNameLower is the lowercase slug used in CLI commands and paths.
	ProductNameLower = "iugo-ai"

	// BinaryName is the compiled binary filename.
	BinaryName = "iugo-ai"

	// OrgName is the organization/brand display name.
	OrgName = "IUGO"

	// OrgNameLower is the lowercase org slug.
	OrgNameLower = "iugo"

	// GitHubOwner is the GitHub org/user that owns the repo.
	GitHubOwner = "iugo-programming"

	// GitHubRepo is the repository name.
	GitHubRepo = "iugo-ai"

	// ModulePath is the Go module import path.
	ModulePath = "github.com/iugo-programming/iugo-ai"

	// ConfigDir is the user-scoped config/state directory name (~/.ConfigDir/).
	ConfigDir = ".iugo-ai"

	// Tagline is the one-line product description shown in TUI welcome.
	Tagline = "Ecosystem, Frameworks, Workflows"

	// DocsURL is the canonical documentation/repository URL.
	DocsURL = "https://github.com/iugo-programming/iugo-ai"

	// ThemeName is the custom TUI theme identifier.
	ThemeName = "iugo-kanagawa"

	// ClaudeThemeName is the Claude Code custom theme name.
	ClaudeThemeName = "IUGO"

	// ClaudeThemeFile is the Claude theme JSON filename.
	ClaudeThemeFile = "iugo.json"

	// AgentName is the primary agent/persona key in OpenCode/Kilocode.
	AgentName = "iugo-agent"

	// OrchestratorName is the SDD orchestrator agent name.
	OrchestratorName = "iugo-orchestrator"

	// MCPServerName is the Engram MCP server identifier.
	MCPServerName = "iugo-ai-engram"

	// NpmPiPackage is the npm package name for the Pi integration.
	NpmPiPackage = "iugo-pi"

	// NpmEngramPackage is the npm package name for the Engram integration.
	NpmEngramPackage = "gentle-engram"

	// PersonaDescription is the agent persona description shown in TUI.
	PersonaDescription = "Senior Architect mentor - helpful first, challenging when it matters"

	// GGAFullName is the full name of the GGA component.
	GGAFullName = "IUGO Guardian Angel"

	// ComponentThemeDescription is the theme component display name.
	ComponentThemeDescription = "IUGO Kanagawa theme overlay"

	// ComponentClaudeThemeDescription is the Claude theme component display name.
	ComponentClaudeThemeDescription = "Claude Code IUGO custom theme"

	// ComponentPersonaDescription is the persona component display name.
	ComponentPersonaDescription = "IUGO persona, neutral or custom behavior"

	// CursorRuleFile is the Cursor rules filename.
	CursorRuleFile = "iugo-ai.mdc"

	// KiroPromptFile is the Kiro system prompt filename.
	KiroPromptFile = "iugo-ai.md"

	// VSCodePromptFile is the VS Code Copilot instructions filename.
	VSCodePromptFile = "iugo-ai.instructions.md"

	// KimiAgentFile is the Kimi agent definition filename.
	KimiAgentFile = "iugo-agent.yaml"

	// MarkerPrefix is the HTML comment prefix for SDD custom agent markers.
	MarkerPrefix = "<!-- iugo-ai:custom-agent:"

	// MarkerSuffix is the HTML comment suffix for SDD custom agent markers.
	MarkerSuffix = " -->"

	// MarkerEndPrefix is the HTML comment prefix for closing SDD markers.
	MarkerEndPrefix = "<!-- /iugo-ai:custom-agent:"
)
