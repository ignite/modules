package types

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParams_Validate(t *testing.T) {
	tests := []struct {
		name    string
		params  Params
		wantErr bool
	}{
		{
			name:   "should validate valid params",
			params: DefaultParams(),
		},
		{
			name: "should prevent validate params with invalid decay information",
			params: NewParams(DecayInformation{
				Enabled:    true,
				DecayStart: time.UnixMilli(1001),
				DecayEnd:   time.UnixMilli(1000),
			}, time.Time{}),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateDecayInformation(t *testing.T) {
	tests := []struct {
		name             string
		decayInformation interface{}
		wantErr          bool
	}{
		{
			name: "should validate valid decay information",
			decayInformation: DecayInformation{
				Enabled: false,
			},
		},
		{
			name:             "should prevent validate decay information with invalid interface",
			decayInformation: "test",
			wantErr:          true,
		},
		{
			name: "should prevent validate invalid decay information",
			decayInformation: DecayInformation{
				Enabled:    true,
				DecayStart: time.UnixMilli(1001),
				DecayEnd:   time.UnixMilli(1000),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDecayInformation(tt.decayInformation)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateAirdropStart(t *testing.T) {
	tests := []struct {
		name              string
		maxMetadataLength interface{}
		err               error
	}{
		{
			name:              "invalid interface",
			maxMetadataLength: "test",
			err:               fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:              "invalid float type",
			maxMetadataLength: 0.5,
			err:               fmt.Errorf("invalid parameter type: float64"),
		},
		{
			name:              "invalid number type",
			maxMetadataLength: uint32(5),
			err:               fmt.Errorf("invalid parameter type: uint32"),
		},
		{
			name:              "valid param",
			maxMetadataLength: int64(1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAirdropStart(tt.maxMetadataLength)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
