package v13

import (
	"fmt"
	"strings"

	"github.com/CosmosContracts/juno/v13/app/keepers"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	// ICA

	icacontrollertypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"

	// types

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// Returns "ujunox" if the chain is uni, else returns the standard ujuno token denom.
func GetChainsDenomToken(chainID string) string {
	if strings.HasPrefix(chainID, "uni-") {
		return "ujunox"
	}
	return "ujuno"
}

func CreateV13UpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		// transfer module consensus version has been bumped to 2
		// the above is https://github.com/cosmos/ibc-go/blob/v5.1.0/docs/migrations/v3-to-v4.md
		logger := ctx.Logger().With("upgrade", UpgradeName)

		nativeDenom := GetChainsDenomToken(ctx.ChainID())
		logger.Info(fmt.Sprintf("With native denom %s", nativeDenom))

		// ICA - https://github.com/CosmosContracts/juno/blob/integrate_ica_changes/app/app.go#L846-L885
		vm[icatypes.ModuleName] = mm.Modules[icatypes.ModuleName].ConsensusVersion()
		logger.Info("upgraded icatypes version")

		// Update ICS27 Host submodule params
		hostParams := icahosttypes.Params{
			HostEnabled: false,
			// https://github.com/cosmos/ibc-go/blob/v4.2.0/docs/apps/interchain-accounts/parameters.md#allowmessages
			AllowMessages: []string{"*"},
		}

		// Run migrations
		versionMap, err := mm.RunMigrations(ctx, cfg, vm)

		// New modules run AFTER the migrations, so to set the correct params after the default.

		// Set ICA Params
		keepers.ICAHostKeeper.SetParams(ctx, hostParams)
		keepers.ICAControllerKeeper.SetParams(ctx, icacontrollertypes.Params{ControllerEnabled: true})
		logger.Info("upgraded ICAHostKeeper params")

		// Oracle
		// newOracleParams := oracletypes.DefaultParams()

		// // add osmosis to the oracle params
		// osmosisDenom := oracletypes.Denom{
		// 	BaseDenom:   "ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518",
		// 	SymbolDenom: "OSMO",
		// 	Exponent:    uint32(6),
		// }

		// allDenoms := oracletypes.DefaultWhitelist
		// allDenoms = append(allDenoms, osmosisDenom)

		// newOracleParams.Whitelist = allDenoms
		// newOracleParams.TwapTrackingList = allDenoms
		// logger.Info(fmt.Sprintf("Oracle params set: %s", newOracleParams.String()))

		// keepers.OracleKeeper.SetParams(ctx, newOracleParams)

		return versionMap, err
	}
}
