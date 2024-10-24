package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/ignite/modules/x/fundraising/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreateFixedPriceAuction(),
		CmdCreateBatchAuction(),
		CmdPlaceBid(),
		CmdModifyBid(),
	)

	return cmd
}

// parseFixedPriceAuctionRequest reads the file and parses FixedPriceAuctionRequest.
func parseFixedPriceAuctionRequest(fileName string) (req FixedPriceAuctionRequest, err error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return req, err
	}

	if err = json.Unmarshal(contents, &req); err != nil {
		return req, err
	}

	return req, nil
}

// FixedPriceAuctionRequest defines CLI request for a fixed price auction.
type FixedPriceAuctionRequest struct {
	StartPrice       sdkmath.LegacyDec       `json:"start_price"`
	SellingCoin      sdk.Coin                `json:"selling_coin"`
	PayingCoinDenom  string                  `json:"paying_coin_denom"`
	VestingSchedules []types.VestingSchedule `json:"vesting_schedules"`
	StartTime        time.Time               `json:"start_time"`
	EndTime          time.Time               `json:"end_time"`
}

// String returns a human-readable string representation of the request.
func (req FixedPriceAuctionRequest) String() string {
	result, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}
	return string(result)
}

// BatchAuctionRequest defines CLI request for an batch auction.
type BatchAuctionRequest struct {
	StartPrice        sdkmath.LegacyDec       `json:"start_price"`
	MinBidPrice       sdkmath.LegacyDec       `json:"min_bid_price"`
	SellingCoin       sdk.Coin                `json:"selling_coin"`
	PayingCoinDenom   string                  `json:"paying_coin_denom"`
	MaxExtendedRound  uint32                  `json:"max_extended_round"`
	ExtendedRoundRate sdkmath.LegacyDec       `json:"extended_round_rate"`
	VestingSchedules  []types.VestingSchedule `json:"vesting_schedules"`
	StartTime         time.Time               `json:"start_time"`
	EndTime           time.Time               `json:"end_time"`
}

// parseBatchAuctionRequest reads the file and parses BatchAuctionRequest.
func parseBatchAuctionRequest(fileName string) (req BatchAuctionRequest, err error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return req, err
	}

	if err = json.Unmarshal(contents, &req); err != nil {
		return req, err
	}

	return req, nil
}

// String returns a human readable string representation of the request.
func (req BatchAuctionRequest) String() string {
	result, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}
	return string(result)
}

// parseBidType parses bid type string and returns types.BidType.
func parseBidType(s string) (types.BidType, error) {
	switch strings.ToLower(s) {
	case "fixed-price", "fp", "f":
		return types.BidTypeFixedPrice, nil
	case "batch-worth", "bw", "w":
		return types.BidTypeBatchWorth, nil
	case "batch-many", "bm", "m":
		return types.BidTypeBatchMany, nil
	}
	return 0, fmt.Errorf("invalid bid type: %s", s)
}
