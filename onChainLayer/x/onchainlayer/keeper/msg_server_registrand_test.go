package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"onChainLayer/x/onchainlayer/types"
)

func TestRegistrandMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateRegistrand(ctx, &types.MsgCreateRegistrand{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestRegistrandMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateRegistrand
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateRegistrand{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateRegistrand{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateRegistrand{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateRegistrand(ctx, &types.MsgCreateRegistrand{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateRegistrand(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRegistrandMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteRegistrand
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteRegistrand{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteRegistrand{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteRegistrand{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateRegistrand(ctx, &types.MsgCreateRegistrand{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteRegistrand(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
