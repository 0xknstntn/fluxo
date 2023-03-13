package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintUSFX = "mint_usfx"

var _ sdk.Msg = &MsgMintUSFX{}

func NewMsgMintUSFX(creator string, amount string) *MsgMintUSFX {
	return &MsgMintUSFX{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgMintUSFX) Route() string {
	return RouterKey
}

func (msg *MsgMintUSFX) Type() string {
	return TypeMsgMintUSFX
}

func (msg *MsgMintUSFX) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintUSFX) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintUSFX) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
