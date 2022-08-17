package errors_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/errors"
)

func TestCritical(t *testing.T) {
	require.ErrorIs(t, errors.ErrCritical, errors.Critical("foo"))
}

func TestCriticalf(t *testing.T) {
	require.ErrorIs(t, errors.ErrCritical, errors.Criticalf("foo %s", "bar"))
}
