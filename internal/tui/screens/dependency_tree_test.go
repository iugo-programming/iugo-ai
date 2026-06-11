package screens

import (
	"fmt"
	"strings"
	"testing"

	"github.com/iugo-programming/iugo-ai/internal/model"
	"github.com/iugo-programming/iugo-ai/internal/planner"
	"github.com/iugo-programming/iugo-ai/internal/versions"
)

func TestRenderDependencyTreePiOnlyEngramPlanShowsComponentAndPiInstallCopy(t *testing.T) {
	selection := model.Selection{
		Agents:     []model.AgentID{model.AgentPi},
		Preset:     model.PresetFullIugo,
		Components: []model.ComponentID{model.ComponentEngram},
	}
	plan := planner.ResolvedPlan{
		Agents:            []model.AgentID{model.AgentPi},
		OrderedComponents: []model.ComponentID{model.ComponentEngram},
	}

	out := RenderDependencyTree(plan, selection, 0)

	if strings.Contains(out, "No components selected yet.") {
		t.Fatalf("RenderDependencyTree() showed generic empty copy for Pi-only Engram plan; output:\n%s", out)
	}
	for _, want := range []string{
		"Components to install",
		"engram",
		"Pi agent support will be installed.",
		"pi install npm:iugo-pi",
		"pi install npm:iugo-engram",
		"pi install npm:pi-mcp-adapter",
		fmt.Sprintf("npm exec --yes --package iugo-engram@%s -- pi-engram init", versions.IugoEngram),
		"pi install npm:pi-subagents",
		"pi install npm:pi-intercom",
		"pi install npm:@juicesharp/rpiv-ask-user-question",
		"pi install npm:pi-web-access",
		"pi install npm:@juicesharp/rpiv-todo",
		"pi install npm:pi-btw",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("RenderDependencyTree() missing %q for Pi-only plan; output:\n%s", want, out)
		}
	}
}

func TestRenderDependencyTreeGenericEmptyPlanKeepsExistingCopy(t *testing.T) {
	selection := model.Selection{Preset: model.PresetFullIugo}

	out := RenderDependencyTree(planner.ResolvedPlan{}, selection, 0)

	if !strings.Contains(out, "No components selected yet.") {
		t.Fatalf("RenderDependencyTree() missing generic empty copy; output:\n%s", out)
	}
	if strings.Contains(out, "Pi agent support will be installed.") {
		t.Fatalf("RenderDependencyTree() showed Pi copy for generic empty plan; output:\n%s", out)
	}
}

func TestRenderDependencyTreeMixedPiEmptyPlanShowsPiInstallCopy(t *testing.T) {
	selection := model.Selection{
		Agents: []model.AgentID{model.AgentPi, model.AgentOpenCode},
		Preset: model.PresetFullIugo,
	}
	plan := planner.ResolvedPlan{Agents: selection.Agents}

	out := RenderDependencyTree(plan, selection, 0)

	if strings.Contains(out, "No components selected yet.") {
		t.Fatalf("RenderDependencyTree() showed generic empty copy for mixed Pi plan; output:\n%s", out)
	}
	for _, want := range []string{
		"Pi agent support will be installed.",
		"pi install npm:iugo-pi",
		"pi install npm:iugo-engram",
		"pi install npm:pi-mcp-adapter",
		fmt.Sprintf("npm exec --yes --package iugo-engram@%s -- pi-engram init", versions.IugoEngram),
		"pi install npm:pi-subagents",
		"pi install npm:pi-intercom",
		"pi install npm:@juicesharp/rpiv-ask-user-question",
		"pi install npm:pi-web-access",
		"pi install npm:@juicesharp/rpiv-todo",
		"pi install npm:pi-btw",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("RenderDependencyTree() missing %q for mixed Pi plan; output:\n%s", want, out)
		}
	}
}
