package types

import (
	"fmt"
	"time"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
)

const (
	SellingReserveAddressPrefix string = "SellingReserveAddress"
	PayingReserveAddressPrefix  string = "PayingReserveAddress"
	VestingReserveAddressPrefix string = "VestingReserveAddress"
	ModuleAddressNameSplitter   string = "|"

	// ReserveAddressType is an address type of reserve for selling, paying, and vesting.
	// The module uses the address type of 32 bytes length, but it can be changed depending on Cosmos SDK's direction.
	ReserveAddressType = AddressType32Bytes
)

var (
	_ AuctionI = (*FixedPriceAuction)(nil)
	_ AuctionI = (*BatchAuction)(nil)
)

// NewBaseAuction creates a new BaseAuction object
//
//nolint:interfacer
func NewBaseAuction(
	auctionID uint64, typ AuctionType, auctioneerAddr string,
	sellingPoolAddr string, payingPoolAddr string,
	startPrice math.LegacyDec, sellingCoin sdk.Coin, payingCoinDenom string,
	vestingPoolAddr string, vestingSchedules []VestingSchedule,
	startTime time.Time, endTime []time.Time, status AuctionStatus,
) *BaseAuction {
	return &BaseAuction{
		AuctionId:             auctionID,
		Type:                  typ,
		Auctioneer:            auctioneerAddr,
		SellingReserveAddress: sellingPoolAddr,
		PayingReserveAddress:  payingPoolAddr,
		StartPrice:            startPrice,
		SellingCoin:           sellingCoin,
		PayingCoinDenom:       payingCoinDenom,
		VestingReserveAddress: vestingPoolAddr,
		VestingSchedules:      vestingSchedules,
		StartTime:             startTime,
		EndTime:               endTime,
		Status:                status,
	}
}

func (ba BaseAuction) GetId() uint64 { //nolint:golint
	return ba.AuctionId
}

func (ba *BaseAuction) SetId(auctionID uint64) error { //nolint:golint
	ba.AuctionId = auctionID
	return nil
}

func (ba BaseAuction) GetType() AuctionType {
	return ba.Type
}

func (ba *BaseAuction) SetType(typ AuctionType) error {
	ba.Type = typ
	return nil
}

func (ba BaseAuction) GetAuctioneer() string {
	return ba.Auctioneer
}

func (ba *BaseAuction) SetAuctioneer(addr string) error {
	ba.Auctioneer = addr
	return nil
}

func (ba BaseAuction) GetSellingReserveAddress() string {
	return ba.SellingReserveAddress
}

func (ba *BaseAuction) SetSellingReserveAddress(addr string) error {
	ba.SellingReserveAddress = addr
	return nil
}

func (ba BaseAuction) GetPayingReserveAddress() string {
	return ba.PayingReserveAddress
}

func (ba *BaseAuction) SetPayingReserveAddress(addr string) error {
	ba.PayingReserveAddress = addr
	return nil
}

func (ba BaseAuction) GetStartPrice() math.LegacyDec {
	return ba.StartPrice
}

func (ba *BaseAuction) SetStartPrice(price math.LegacyDec) error {
	ba.StartPrice = price
	return nil
}

func (ba BaseAuction) GetSellingCoin() sdk.Coin {
	return ba.SellingCoin
}

func (ba *BaseAuction) SetSellingCoin(coin sdk.Coin) error {
	ba.SellingCoin = coin
	return nil
}

func (ba BaseAuction) GetPayingCoinDenom() string {
	return ba.PayingCoinDenom
}

func (ba *BaseAuction) SetPayingCoinDenom(denom string) error {
	ba.PayingCoinDenom = denom
	return nil
}

func (ba BaseAuction) GetVestingReserveAddress() string {
	return ba.VestingReserveAddress
}

func (ba *BaseAuction) SetVestingReserveAddress(addr string) error {
	ba.VestingReserveAddress = addr
	return nil
}

func (ba BaseAuction) GetVestingSchedules() []VestingSchedule {
	return ba.VestingSchedules
}

func (ba *BaseAuction) SetVestingSchedules(schedules []VestingSchedule) error {
	ba.VestingSchedules = schedules
	return nil
}

func (ba BaseAuction) GetStartTime() time.Time {
	return ba.StartTime
}

func (ba *BaseAuction) SetStartTime(t time.Time) error {
	ba.StartTime = t
	return nil
}

func (ba BaseAuction) GetEndTime() []time.Time {
	return ba.EndTime
}

func (ba *BaseAuction) SetEndTime(t []time.Time) error {
	ba.EndTime = t
	return nil
}

func (ba BaseAuction) GetStatus() AuctionStatus {
	return ba.Status
}

func (ba *BaseAuction) SetStatus(status AuctionStatus) error {
	ba.Status = status
	return nil
}

// Validate checks for errors on the Auction fields
func (ba BaseAuction) Validate() error {
	if ba.Type != AuctionTypeFixedPrice && ba.Type != AuctionTypeBatch {
		return sdkerrors.Wrapf(ErrInvalidAuctionType, "unknown plan type: %s", ba.Type)
	}
	if _, err := sdk.AccAddressFromBech32(ba.Auctioneer); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid auctioneer address %q: %v", ba.Auctioneer, err)
	}
	if _, err := sdk.AccAddressFromBech32(ba.SellingReserveAddress); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid selling pool address %q: %v", ba.SellingReserveAddress, err)
	}
	if _, err := sdk.AccAddressFromBech32(ba.PayingReserveAddress); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid paying pool address %q: %v", ba.PayingReserveAddress, err)
	}
	if _, err := sdk.AccAddressFromBech32(ba.VestingReserveAddress); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid vesting pool address %q: %v", ba.VestingReserveAddress, err)
	}
	if !ba.StartPrice.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid start price: %f", ba.StartPrice)
	}
	if err := ba.SellingCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidCoins, "invalid selling coin: %v", ba.SellingCoin)
	}
	if ba.SellingCoin.Denom == ba.PayingCoinDenom {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "selling coin denom must not be the same as paying coin denom")
	}
	if err := sdk.ValidateDenom(ba.PayingCoinDenom); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid paying coin denom: %v", err)
	}
	if err := ValidateVestingSchedules(ba.VestingSchedules, ba.EndTime[len(ba.EndTime)-1]); err != nil {
		return err
	}
	return nil
}

// ShouldAuctionStarted returns true if the start time is equal or before the given time t.
func (ba BaseAuction) ShouldAuctionStarted(t time.Time) bool {
	return !ba.GetStartTime().After(t) // StartTime <= Time
}

// ShouldAuctionClosed returns true if the end time is equal or before the given time t.
func (ba BaseAuction) ShouldAuctionClosed(t time.Time) bool {
	ts := ba.GetEndTime()
	return !ts[len(ts)-1].After(t) // LastEndTime <= Time
}

// NewFixedPriceAuction returns a new fixed price auction.
func NewFixedPriceAuction(baseAuction *BaseAuction, remainingSellingCoin sdk.Coin) *FixedPriceAuction {
	return &FixedPriceAuction{
		BaseAuction:          baseAuction,
		RemainingSellingCoin: remainingSellingCoin,
	}
}

// NewBatchAuction returns a new batch auction.
func NewBatchAuction(baseAuction *BaseAuction, minBidPrice math.LegacyDec, matchedPrice math.LegacyDec, maxExtendedRound uint32, extendedRoundRate math.LegacyDec) *BatchAuction {
	return &BatchAuction{
		BaseAuction:       baseAuction,
		MinBidPrice:       minBidPrice,
		MatchedPrice:      matchedPrice,
		MaxExtendedRound:  maxExtendedRound,
		ExtendedRoundRate: extendedRoundRate,
	}
}

// AuctionI is an interface that inherits the BaseAuction and exposes common functions
// to get and set standard auction data.
type AuctionI interface {
	proto.Message

	GetId() uint64
	SetId(uint64) error

	GetType() AuctionType
	SetType(AuctionType) error

	GetAuctioneer() string
	SetAuctioneer(string) error

	GetSellingReserveAddress() string
	SetSellingReserveAddress(string) error

	GetPayingReserveAddress() string
	SetPayingReserveAddress(string) error

	GetStartPrice() math.LegacyDec
	SetStartPrice(math.LegacyDec) error

	GetSellingCoin() sdk.Coin
	SetSellingCoin(sdk.Coin) error

	GetPayingCoinDenom() string
	SetPayingCoinDenom(string) error

	GetVestingReserveAddress() string
	SetVestingReserveAddress(string) error

	GetVestingSchedules() []VestingSchedule
	SetVestingSchedules([]VestingSchedule) error

	GetStartTime() time.Time
	SetStartTime(time.Time) error

	GetEndTime() []time.Time
	SetEndTime([]time.Time) error

	GetStatus() AuctionStatus
	SetStatus(AuctionStatus) error

	ShouldAuctionStarted(t time.Time) bool
	ShouldAuctionClosed(t time.Time) bool

	Validate() error
}

// PackAuction converts AuctionI to Any.
func PackAuction(auction AuctionI) (*codectypes.Any, error) {
	any, err := codectypes.NewAnyWithValue(auction)
	if err != nil {
		return nil, err
	}
	return any, nil
}

// UnpackAuction converts Any to AuctionI.
func UnpackAuction(any *codectypes.Any) (AuctionI, error) {
	if any == nil {
		return nil, sdkerrors.Wrapf(errors.ErrInvalidType, "cannot unpack nil")
	}

	if any.TypeUrl == "" {
		return nil, sdkerrors.Wrap(errors.ErrInvalidType, "empty type url")
	}

	var auction AuctionI
	v := any.GetCachedValue()
	if v == nil {
		registry := codectypes.NewInterfaceRegistry()
		RegisterInterfaces(registry)
		if err := registry.UnpackAny(any, &auction); err != nil {
			return nil, err
		}
		return auction, nil
	}

	auction, ok := v.(AuctionI)
	if !ok {
		return nil, sdkerrors.Wrapf(errors.ErrInvalidType, "cannot unpack auction from %T", v)
	}

	return auction, nil
}

// UnpackAuctions converts Any slice to AuctionIs.
func UnpackAuctions(auctionsAny []*codectypes.Any) ([]AuctionI, error) {
	auctions := make([]AuctionI, len(auctionsAny))
	for i, any := range auctionsAny {
		p, err := UnpackAuction(any)
		if err != nil {
			return nil, err
		}
		auctions[i] = p
	}
	return auctions, nil
}

// MustMarshalAuction returns the marshalled auction bytes.
// It throws panic if it fails.
func MustMarshalAuction(cdc codec.BinaryCodec, auction AuctionI) []byte {
	bz, err := MarshalAuction(cdc, auction)
	if err != nil {
		panic(err)
	}
	return bz
}

// MustUnmarshalAuction return the unmarshalled auction from bytes.
// It throws panic if it fails.
func MustUnmarshalAuction(cdc codec.BinaryCodec, value []byte) AuctionI {
	pair, err := UnmarshalAuction(cdc, value)
	if err != nil {
		panic(err)
	}
	return pair
}

// MarshalAuction returns bytes from the auction interface.
func MarshalAuction(cdc codec.BinaryCodec, auction AuctionI) (value []byte, err error) {
	return cdc.MarshalInterface(auction)
}

// UnmarshalAuction returns the auction from the bytes.
func UnmarshalAuction(cdc codec.BinaryCodec, value []byte) (auction AuctionI, err error) {
	err = cdc.UnmarshalInterface(value, &auction)
	return auction, err
}

// SellingReserveAddress returns the selling reserve address with the given auction id.
func SellingReserveAddress(auctionID uint64) sdk.AccAddress {
	return DeriveAddress(ReserveAddressType, ModuleName, SellingReserveAddressPrefix+ModuleAddressNameSplitter+fmt.Sprint(auctionID))
}

// PayingReserveAddress returns the paying reserve address with the given auction id.
func PayingReserveAddress(auctionID uint64) sdk.AccAddress {
	return DeriveAddress(ReserveAddressType, ModuleName, PayingReserveAddressPrefix+ModuleAddressNameSplitter+fmt.Sprint(auctionID))
}

// VestingReserveAddress returns the vesting reserve address with the given auction id.
func VestingReserveAddress(auctionID uint64) sdk.AccAddress {
	return DeriveAddress(ReserveAddressType, ModuleName, VestingReserveAddressPrefix+ModuleAddressNameSplitter+fmt.Sprint(auctionID))
}
