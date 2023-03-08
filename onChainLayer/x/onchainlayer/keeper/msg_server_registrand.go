package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"onChainLayer/x/onchainlayer/types"
)

func (k msgServer) CreateRegistrand(goCtx context.Context, msg *types.MsgCreateRegistrand) (*types.MsgCreateRegistrandResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var registrand = types.Registrand{
		Creator:  msg.Creator,
		Ip:       msg.Ip,
		Nickname: msg.Nickname,
		Sig:      msg.Sig,
		Pk:       msg.Pk,
		Msg:      msg.Msg,
	}

	id := k.AppendRegistrand(
		ctx,
		registrand,
	)

	return &types.MsgCreateRegistrandResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateRegistrand(goCtx context.Context, msg *types.MsgUpdateRegistrand) (*types.MsgUpdateRegistrandResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var registrand = types.Registrand{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Ip:       msg.Ip,
		Nickname: msg.Nickname,
		Sig:      msg.Sig,
		Pk:       msg.Pk,
		Msg:      msg.Msg,
	}

	// Checks that the element exists
	val, found := k.GetRegistrand(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetRegistrand(ctx, registrand)

	return &types.MsgUpdateRegistrandResponse{}, nil
}

func (k msgServer) DeleteRegistrand(goCtx context.Context, msg *types.MsgDeleteRegistrand) (*types.MsgDeleteRegistrandResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetRegistrand(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRegistrand(ctx, msg.Id)

	return &types.MsgDeleteRegistrandResponse{}, nil
}
