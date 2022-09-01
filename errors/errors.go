package errors

import sdkerrors "cosmossdk.io/errors"

// Codespace is the codespace for all errors defined in this package
const Codespace = "ignite-modules"

var (
	// ErrTxDecode is returned if we cannot parse a transaction
	ErrTxDecode = sdkerrors.Register(Codespace, 2, "tx parse error")

	// ErrInvalidSequence is used the sequence number (nonce) is incorrect
	// for the signature
	ErrInvalidSequence = sdkerrors.Register(Codespace, 3, "invalid sequence")

	// ErrUnauthorized is used whenever a request without sufficient
	// authorization is handled.
	ErrUnauthorized = sdkerrors.Register(Codespace, 4, "unauthorized")

	// ErrInsufficientFunds is used when the account cannot pay requested amount.
	ErrInsufficientFunds = sdkerrors.Register(Codespace, 5, "insufficient funds")

	// ErrUnknownRequest to doc
	ErrUnknownRequest = sdkerrors.Register(Codespace, 6, "unknown request")

	// ErrInvalidAddress to doc
	ErrInvalidAddress = sdkerrors.Register(Codespace, 7, "invalid address")

	// ErrInvalidPubKey to doc
	ErrInvalidPubKey = sdkerrors.Register(Codespace, 8, "invalid pubkey")

	// ErrUnknownAddress to doc
	ErrUnknownAddress = sdkerrors.Register(Codespace, 9, "unknown address")

	// ErrInvalidCoins to doc
	ErrInvalidCoins = sdkerrors.Register(Codespace, 10, "invalid coins")

	// ErrOutOfGas to doc
	ErrOutOfGas = sdkerrors.Register(Codespace, 11, "out of gas")

	// ErrMemoTooLarge to doc
	ErrMemoTooLarge = sdkerrors.Register(Codespace, 12, "memo too large")

	// ErrInsufficientFee to doc
	ErrInsufficientFee = sdkerrors.Register(Codespace, 13, "insufficient fee")

	// ErrTooManySignatures to doc
	ErrTooManySignatures = sdkerrors.Register(Codespace, 14, "maximum number of signatures exceeded")

	// ErrNoSignatures to doc
	ErrNoSignatures = sdkerrors.Register(Codespace, 15, "no signatures supplied")

	// ErrJSONMarshal defines an ABCI typed JSON marshalling error
	ErrJSONMarshal = sdkerrors.Register(Codespace, 16, "failed to marshal JSON bytes")

	// ErrJSONUnmarshal defines an ABCI typed JSON unmarshalling error
	ErrJSONUnmarshal = sdkerrors.Register(Codespace, 17, "failed to unmarshal JSON bytes")

	// ErrInvalidRequest defines an ABCI typed error where the request contains
	// invalid data.
	ErrInvalidRequest = sdkerrors.Register(Codespace, 18, "invalid request")

	// ErrTxInMempoolCache defines an ABCI typed error where a tx already exists
	// in the mempool.
	ErrTxInMempoolCache = sdkerrors.Register(Codespace, 19, "tx already in mempool")

	// ErrMempoolIsFull defines an ABCI typed error where the mempool is full.
	ErrMempoolIsFull = sdkerrors.Register(Codespace, 20, "mempool is full")

	// ErrTxTooLarge defines an ABCI typed error where tx is too large.
	ErrTxTooLarge = sdkerrors.Register(Codespace, 21, "tx too large")

	// ErrKeyNotFound defines an error when the key doesn't exist
	ErrKeyNotFound = sdkerrors.Register(Codespace, 22, "key not found")

	// ErrWrongPassword defines an error when the key password is invalid.
	ErrWrongPassword = sdkerrors.Register(Codespace, 23, "invalid account password")

	// ErrorInvalidSigner defines an error when the tx intended signer does not match the given signer.
	ErrorInvalidSigner = sdkerrors.Register(Codespace, 24, "tx intended signer does not match the given signer")

	// ErrorInvalidGasAdjustment defines an error for an invalid gas adjustment
	ErrorInvalidGasAdjustment = sdkerrors.Register(Codespace, 25, "invalid gas adjustment")

	// ErrInvalidHeight defines an error for an invalid height
	ErrInvalidHeight = sdkerrors.Register(Codespace, 26, "invalid height")

	// ErrInvalidVersion defines a general error for an invalid version
	ErrInvalidVersion = sdkerrors.Register(Codespace, 27, "invalid version")

	// ErrInvalidChainID defines an error when the chain-id is invalid.
	ErrInvalidChainID = sdkerrors.Register(Codespace, 28, "invalid chain-id")

	// ErrInvalidType defines an error an invalid type.
	ErrInvalidType = sdkerrors.Register(Codespace, 29, "invalid type")

	// ErrTxTimeoutHeight defines an error for when a tx is rejected out due to an
	// explicitly set timeout height.
	ErrTxTimeoutHeight = sdkerrors.Register(Codespace, 30, "tx timeout height")

	// ErrUnknownExtensionOptions defines an error for unknown extension options.
	ErrUnknownExtensionOptions = sdkerrors.Register(Codespace, 31, "unknown extension options")

	// ErrWrongSequence defines an error where the account sequence defined in
	// the signer info doesn't match the account's actual sequence number.
	ErrWrongSequence = sdkerrors.Register(Codespace, 32, "incorrect account sequence")

	// ErrPackAny defines an error when packing a protobuf message to Any fails.
	ErrPackAny = sdkerrors.Register(Codespace, 33, "failed packing protobuf message to Any")

	// ErrUnpackAny defines an error when unpacking a protobuf message from Any fails.
	ErrUnpackAny = sdkerrors.Register(Codespace, 34, "failed unpacking protobuf message from Any")

	// ErrLogic defines an internal logic error, e.g. an invariant or assertion
	// that is violated. It is a programmer error, not a user-facing error.
	ErrLogic = sdkerrors.Register(Codespace, 35, "internal logic error")

	// ErrConflict defines a conflict error, e.g. when two goroutines try to access
	// the same resource and one of them fails.
	ErrConflict = sdkerrors.Register(Codespace, 36, "conflict")

	// ErrNotSupported is returned when we call a branch of a code which is currently not
	// supported.
	ErrNotSupported = sdkerrors.Register(Codespace, 37, "feature not supported")

	// ErrNotFound defines an error when requested entity doesn't exist in the state.
	ErrNotFound = sdkerrors.Register(Codespace, 38, "not found")

	// ErrIO should be used to wrap internal errors caused by external operation.
	// Examples: not DB domain error, file writing etc...
	ErrIO = sdkerrors.Register(Codespace, 39, "Internal IO error")

	// ErrAppConfig defines an error occurred if min-gas-prices field in BaseConfig is empty.
	ErrAppConfig = sdkerrors.Register(Codespace, 40, "error in app.toml")

	// ErrInvalidGasLimit defines an error when an invalid GasWanted value is
	// supplied.
	ErrInvalidGasLimit = sdkerrors.Register(Codespace, 41, "invalid gas limit")

	// ErrPanic is only set when we recover from a panic, so we know to
	// redact potentially sensitive system info
	ErrPanic = sdkerrors.ErrPanic
)
