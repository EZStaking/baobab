package v13

import (
	"github.com/CosmosContracts/juno/v13/app/upgrades"

	// oracletypes "github.com/CosmosContracts/juno/v13/x/oracle/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	ibchookstypes "github.com/osmosis-labs/osmosis/x/ibc-hooks/types"
)

// UpgradeName defines the on-chain upgrade name for the upgrade.
const UpgradeName = "v13"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV13UpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			// oracletypes.ModuleName,
			// intertxtypes.ModuleName,
			ibchookstypes.StoreKey,
		},
	},
}
