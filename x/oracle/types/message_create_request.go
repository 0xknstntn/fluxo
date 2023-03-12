package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateRequest = "create_request"

var _ sdk.Msg = &MsgCreateRequest{}

func NewMsgCreateRequest(creator string, price string) *MsgCreateRequest {
	return &MsgCreateRequest{
		Creator: creator,
		Price:   price,
	}
}

func (msg *MsgCreateRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateRequest) Type() string {
	return TypeMsgCreateRequest
}

func (msg *MsgCreateRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	price, err := strconv.ParseInt(msg.Price, 10, 64)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "price is not an integer")
	}
	if price <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "price should be a positive integer")
	}
	return nil
}
