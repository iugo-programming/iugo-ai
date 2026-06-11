package cli

import "github.com/iugo-programming/iugo-ai/internal/model"

func isGentlemanConversationPersona(persona model.PersonaID) bool {
	return persona == model.PersonaIugo || persona == model.PersonaIugoNeutralArtifacts
}
