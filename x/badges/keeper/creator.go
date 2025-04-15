package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetCreator(ctx sdk.Context, creator string, creatorOverride string) (string, error) {
	// If creatorOverride is set, use it as the creator
	if creatorOverride != "" {
		//TODO: Validate creator is the contract owner

		return creatorOverride, nil
	}

	// Otherwise, use the original creator
	return creator, nil
}
