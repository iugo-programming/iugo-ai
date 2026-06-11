package screens

import (
	"strings"
	"testing"

	"github.com/iugo-programming/iugo-ai/internal/model"
)

func TestPersonaOptionsIncludeIUGONeutralArtifacts(t *testing.T) {
	options := PersonaOptions()
	found := false
	for _, option := range options {
		if option == model.PersonaIUGONeutralArtifacts {
			found = true
		}
	}
	if !found {
		t.Fatalf("PersonaOptions() = %v, missing %q", options, model.PersonaIUGONeutralArtifacts)
	}
}

func TestRenderPersonaDescribesIUGONeutralArtifacts(t *testing.T) {
	out := RenderPersona(model.PersonaIUGONeutralArtifacts, 2)
	for _, want := range []string{
		"iugo-agent-neutral-artifacts",
		"IUGO conversation",
		"English technical artifacts",
		"context language",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("RenderPersona() missing %q; output:\n%s", want, out)
		}
	}
}
