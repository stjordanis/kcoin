// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
)

var _ = (*votedataMarshalling)(nil)

// MarshalJSON marshals as JSON.
func (v votedata) MarshalJSON() ([]byte, error) {
	type votedata struct {
		BlockHash   common.Hash    `json:"blockHash"    gencodec:"required"`
		BlockNumber *hexutil.Big   `json:"blockNumber"  gencodec:"required"`
		Round       hexutil.Uint64 `json:"round"        gencodec:"required"`
		Type        VoteType       `json:"type"         gencodec:"required"`
		V           *hexutil.Big   `json:"v"   gencodec:"required"`
		R           *hexutil.Big   `json:"r"   gencodec:"required"`
		S           *hexutil.Big   `json:"s"   gencodec:"required"`
	}
	var enc votedata
	enc.BlockHash = v.BlockHash
	enc.BlockNumber = (*hexutil.Big)(v.BlockNumber)
	enc.Round = hexutil.Uint64(v.Round)
	enc.Type = v.Type
	enc.V = (*hexutil.Big)(v.V)
	enc.R = (*hexutil.Big)(v.R)
	enc.S = (*hexutil.Big)(v.S)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (v *votedata) UnmarshalJSON(input []byte) error {
	type votedata struct {
		BlockHash   *common.Hash    `json:"blockHash"    gencodec:"required"`
		BlockNumber *hexutil.Big    `json:"blockNumber"  gencodec:"required"`
		Round       *hexutil.Uint64 `json:"round"        gencodec:"required"`
		Type        *VoteType       `json:"type"         gencodec:"required"`
		V           *hexutil.Big    `json:"v"   gencodec:"required"`
		R           *hexutil.Big    `json:"r"   gencodec:"required"`
		S           *hexutil.Big    `json:"s"   gencodec:"required"`
	}
	var dec votedata
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.BlockHash == nil {
		return errors.New("missing required field 'blockHash' for votedata")
	}
	v.BlockHash = *dec.BlockHash
	if dec.BlockNumber == nil {
		return errors.New("missing required field 'blockNumber' for votedata")
	}
	v.BlockNumber = (*big.Int)(dec.BlockNumber)
	if dec.Round == nil {
		return errors.New("missing required field 'round' for votedata")
	}
	v.Round = uint64(*dec.Round)
	if dec.Type == nil {
		return errors.New("missing required field 'type' for votedata")
	}
	v.Type = *dec.Type
	if dec.V == nil {
		return errors.New("missing required field 'v' for votedata")
	}
	v.V = (*big.Int)(dec.V)
	if dec.R == nil {
		return errors.New("missing required field 'r' for votedata")
	}
	v.R = (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for votedata")
	}
	v.S = (*big.Int)(dec.S)
	return nil
}
