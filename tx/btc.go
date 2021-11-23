package tx

import (
	"fmt"

	"github.com/btcsuite/btcutil"
)

// NOT YET DONE
func NewSignedBitcoinTx() *btcutil.Tx {

	tx, err := btcutil.NewTxFromBytes([]byte{})
	if err != nil {
		fmt.Println(err)
	}

	return tx
}
