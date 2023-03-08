package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"onChainLayer/x/onchainlayer/types"
)

func (k Keeper) RegistrandAll(goCtx context.Context, req *types.QueryAllRegistrandRequest) (*types.QueryAllRegistrandResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var registrands []types.Registrand
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	registrandStore := prefix.NewStore(store, types.KeyPrefix(types.RegistrandKey))

	pageRes, err := query.Paginate(registrandStore, req.Pagination, func(key []byte, value []byte) error {
		var registrand types.Registrand
		if err := k.cdc.Unmarshal(value, &registrand); err != nil {
			return err
		}

		registrands = append(registrands, registrand)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRegistrandResponse{Registrand: registrands, Pagination: pageRes}, nil
}

func (k Keeper) Registrand(goCtx context.Context, req *types.QueryGetRegistrandRequest) (*types.QueryGetRegistrandResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	registrand, found := k.GetRegistrand(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRegistrandResponse{Registrand: registrand}, nil
}
