package simulation

import (
	"math/rand"

	"github.com/bitbadges/badges-module/x/badges/keeper"
	"github.com/bitbadges/badges-module/x/badges/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdateUserApprovals(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateUserApprovals{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateUserOutgoingApprovals simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UpdateUserOutgoingApprovals simulation not implemented"), nil, nil
	}
}
