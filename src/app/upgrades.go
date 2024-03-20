package app

import (
        "context"
        "strings"

        storetypes "cosmossdk.io/store/types"
        upgradetypes "cosmossdk.io/x/upgrade/types"
        "github.com/cosmos/cosmos-sdk/types/module"
        sdkversion "github.com/cosmos/cosmos-sdk/version"
)

func (app App) RegisterUpgradeHandlers() {
        var UpgradeName = strings.TrimSpace(sdkversion.Version)

        app.UpgradeKeeper.SetUpgradeHandler(UpgradeName, func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

                return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
        })

        upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
        if err == nil {
                if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
                        storeUpgrades := storetypes.StoreUpgrades{
                                Deleted: []string{"capability"},
                        }

                        // configure store loader that checks if version == upgradeHeight and applies store upgrades
                        app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
                }
        }

}
