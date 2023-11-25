package app

import (
	"context"
	"strings"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sdkversion "github.com/cosmos/cosmos-sdk/version"
)

func (app App) RegisterUpgradeHandlers() {
	var version = strings.TrimSpace(sdkversion.Version)

	app.UpgradeKeeper.SetUpgradeHandler(version, func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
	})

}
