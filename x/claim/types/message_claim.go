package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewMsgClaim(claimer string, missionID uint64) *MsgClaim {
	return &MsgClaim{
		Claimer:   claimer,
		MissionID: missionID,
	}
}

func (msg *MsgClaim) Type() string {
	return sdk.MsgTypeURL(&MsgClaim{})
}
