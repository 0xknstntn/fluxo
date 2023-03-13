package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnUSFX = "burn_usfx"

var _ sdk.Msg = &MsgBurnUSFX{}

func NewMsgBurnUSFX(creator string, amount string) *MsgBurnUSFX {
	return &MsgBurnUSFX{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBurnUSFX) Route() string {
	return RouterKey
}

func (msg *MsgBurnUSFX) Type() string {
	return TypeMsgBurnUSFX
}

func (msg *MsgBurnUSFX) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnUSFX) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnUSFX) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
