package engram

import (
	"github.com/iugo-programming/iugo-ai/internal/installcmd"
	"github.com/iugo-programming/iugo-ai/internal/model"
	"github.com/iugo-programming/iugo-ai/internal/system"
)

func InstallCommand(profile system.PlatformProfile) ([][]string, error) {
	return installcmd.NewResolver().ResolveComponentInstall(profile, model.ComponentEngram)
}
