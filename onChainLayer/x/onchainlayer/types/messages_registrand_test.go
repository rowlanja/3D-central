package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"onChainLayer/testutil/sample"
)

func TestMsgCreateRegistrand_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateRegistrand
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateRegistrand{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateRegistrand{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateRegistrand_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateRegistrand
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateRegistrand{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateRegistrand{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteRegistrand_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteRegistrand
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteRegistrand{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteRegistrand{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
