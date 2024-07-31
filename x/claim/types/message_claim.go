package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/pkg/errors"
)

const TypeMsgClaim = "claim"

var _ sdk.Msg = &MsgClaim{}

func NewMsgClaim(claimer string, missionID uint64) *MsgClaim {
	return &MsgClaim{
		Claimer:   claimer,
		MissionID: missionID,
	}
}

func (msg *MsgClaim) Route() string {
	return RouterKey
}

func (msg *MsgClaim) Type() string {
	return TypeMsgClaim
}

func (msg *MsgClaim) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Claimer); err != nil {
		return errors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
