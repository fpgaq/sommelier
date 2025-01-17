package types

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/crypto"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

func (c *RebalanceVote) InvalidationScope() tmbytes.HexBytes {
	return crypto.Keccak256Hash(c.ABIEncodedRebalanceBytes()).Bytes()
}

func (rv *RebalanceVote) Equals(other RebalanceVote) bool {
	if !rv.Cellar.Equals(*other.Cellar) {
		return false
	}

	if rv.CurrentPrice != other.CurrentPrice {
		return false
	}

	return true
}

func (c *RebalanceVote) Hash(salt string, val sdk.ValAddress) ([]byte, error) {
	databytes, err := c.Marshal()

	if err != nil {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrJSONMarshal, "failed to marshal rebalance vote",
		)
	}

	hexbytes := hex.EncodeToString(databytes)

	// calculate the vote hash on the server
	commitHash := DataHash(salt, hexbytes, val)

	return commitHash, nil
}
