package app

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sdkversion "github.com/cosmos/cosmos-sdk/version"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func (app App) RegisterUpgradeHandlers() {
	var version = strings.TrimSpace(sdkversion.Version)

	app.UpgradeKeeper.SetUpgradeHandler(version, func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
	})

}
