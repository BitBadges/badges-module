package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetCreator(ctx sdk.Context, creator string, creatorOverride string) (string, error) {
	// If creatorOverride is set, we need to verify actual creator is an approved contract address
	// IMPORTANT: Approved contract addresses should never be allowed to specify alternate creators other than the initial signer themselves
	// This is to prevent malicious contracts from overriding the creator
	if creatorOverride != "" {
		approvedContractAddresses := k.ApprovedContractAddresses
		for _, address := range approvedContractAddresses {
			if address == creator {
				return creatorOverride, nil
			}
		}
		return "", fmt.Errorf("creatorOverride is not an approved contract address. If you want to override the creator, you must do so through an approved contract address. Otherwise, use the original creator.")
	}

	// Otherwise, use the original creator
	return creator, nil
}
