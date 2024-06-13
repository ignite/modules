package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/pkg/errors"
)

var _ sdk.Msg = &MsgClaim{}

func NewMsgClaim(creator string, missionID uint64) *MsgClaim {
	return &MsgClaim{
		Claimer:   creator,
		MissionID: missionID,
	}
}

func (msg *MsgClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaim) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Claimer); err != nil {
		return errors.Wrapf(errors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
