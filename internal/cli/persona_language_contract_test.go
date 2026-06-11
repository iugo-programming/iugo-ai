package cli

import (
	"testing"

	"github.com/iugo-programming/iugo-ai/internal/model"
)

func TestNormalizePersonaAcceptsIUGONeutralArtifacts(t *testing.T) {
	got, err := normalizePersona("iugo-neutral-artifacts")
	if err != nil {
		t.Fatalf("normalizePersona() error = %v", err)
	}
	if got != model.PersonaIugoNeutralArtifacts {
		t.Fatalf("normalizePersona() = %q, want %q", got, model.PersonaIugoNeutralArtifacts)
	}
}
