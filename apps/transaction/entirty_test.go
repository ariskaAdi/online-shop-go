package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubTotal(t *testing.T) {
	var trx = TransactionEntity{
		ProductPrice: 10_000,
		Amount: 10,
	} 

	expected := uint(100_000)
	
	trx.SetSubTotal()
	require.Equal(t, expected, trx.SubTotal)
}