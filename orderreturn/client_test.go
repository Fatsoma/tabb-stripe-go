package orderreturn

import (
	"testing"

	stripe "github.com/fatsoma/tabb-stripe-go"
	_ "github.com/fatsoma/tabb-stripe-go/testing"
	assert "github.com/stretchr/testify/require"
)

func TestOrderReturnList(t *testing.T) {
	i := List(&stripe.OrderReturnListParams{})

	// Verify that we can get at least one order return
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.OrderReturn())
}
