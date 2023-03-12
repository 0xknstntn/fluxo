package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgValidatorRegistration = "validator_registration"

var _ sdk.Msg = &MsgValidatorRegistration{}

func NewMsgValidatorRegistration(creator string, amount string) *MsgValidatorRegistration {
	return &MsgValidatorRegistration{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgValidatorRegistration) Route() string {
	return RouterKey
}

func (msg *MsgValidatorRegistration) Type() string {
	return TypeMsgValidatorRegistration
}

func (msg *MsgValidatorRegistration) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgValidatorRegistration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgValidatorRegistration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
