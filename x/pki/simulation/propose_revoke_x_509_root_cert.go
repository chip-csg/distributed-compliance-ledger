package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func SimulateMsgProposeRevokeX509RootCert(
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgProposeRevokeX509RootCert{
			Signer: simAccount.Address.String(),
		}

		// TODO: Handling the ProposeRevokeX509RootCert simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ProposeRevokeX509RootCert simulation not implemented"), nil, nil
	}
}
