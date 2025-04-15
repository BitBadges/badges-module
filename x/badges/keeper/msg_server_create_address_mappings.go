package keeper

import (
	"context"

	"github.com/bitbadges/badges-module/x/badges/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAddressLists(goCtx context.Context, msg *types.MsgCreateAddressLists) (*types.MsgCreateAddressListsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := k.GetCreator(ctx, msg.Creator, msg.CreatorOverride)
	if err != nil {
		return nil, err
	}
	msg.Creator = creator

	for _, addressList := range msg.AddressLists {
		addressList.CreatedBy = msg.Creator
		if err := k.CreateAddressList(ctx, addressList); err != nil {
			return nil, err
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "badges"),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	)

	return &types.MsgCreateAddressListsResponse{}, nil
}
