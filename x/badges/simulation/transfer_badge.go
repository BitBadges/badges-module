package simulation

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"

	"github.com/bitbadges/badges-module/x/badges/keeper"
	"github.com/bitbadges/badges-module/x/badges/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgTransferBadges(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		msg := &types.MsgTransferBadges{
			Creator:      simAccount.Address.String(),
			CollectionId: sdkmath.NewUint(uint64(r.Int63n(100))),
			Transfers:    GetRandomTransfers(r, 3, accs),
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}
