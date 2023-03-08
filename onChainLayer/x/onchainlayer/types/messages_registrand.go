package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateRegistrand = "create_registrand"
	TypeMsgUpdateRegistrand = "update_registrand"
	TypeMsgDeleteRegistrand = "delete_registrand"
)

var _ sdk.Msg = &MsgCreateRegistrand{}

func NewMsgCreateRegistrand(creator string, ip string, nickname string, sig string, pk string, msg string) *MsgCreateRegistrand {
	return &MsgCreateRegistrand{
		Creator:  creator,
		Ip:       ip,
		Nickname: nickname,
		Sig:      sig,
		Pk:       pk,
		Msg:      msg,
	}
}

func (msg *MsgCreateRegistrand) Route() string {
	return RouterKey
}

func (msg *MsgCreateRegistrand) Type() string {
	return TypeMsgCreateRegistrand
}

func (msg *MsgCreateRegistrand) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRegistrand) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRegistrand) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRegistrand{}

func NewMsgUpdateRegistrand(creator string, id uint64, ip string, nickname string, sig string, pk string, msg string) *MsgUpdateRegistrand {
	return &MsgUpdateRegistrand{
		Id:       id,
		Creator:  creator,
		Ip:       ip,
		Nickname: nickname,
		Sig:      sig,
		Pk:       pk,
		Msg:      msg,
	}
}

func (msg *MsgUpdateRegistrand) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRegistrand) Type() string {
	return TypeMsgUpdateRegistrand
}

func (msg *MsgUpdateRegistrand) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRegistrand) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRegistrand) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRegistrand{}

func NewMsgDeleteRegistrand(creator string, id uint64) *MsgDeleteRegistrand {
	return &MsgDeleteRegistrand{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteRegistrand) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRegistrand) Type() string {
	return TypeMsgDeleteRegistrand
}

func (msg *MsgDeleteRegistrand) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRegistrand) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRegistrand) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
