// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
)

var _ = (*receiptMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r Receipt) MarshalJSON() ([]byte, error) {
	type Receipt struct {
		PostState               hexutil.Bytes  `json:"root"`
		Status                  hexutil.Uint64 `json:"status"`
		CumulativeResourceUsage hexutil.Uint64 `json:"cumulativeResourceUsage" gencodec:"required"`
		Bloom                   Bloom          `json:"logsBloom"               gencodec:"required"`
		Logs                    []*Log         `json:"logs"                    gencodec:"required"`
		TxHash                  common.Hash    `json:"transactionHash"         gencodec:"required"`
		ContractAddress         common.Address `json:"contractAddress"`
		ResourceUsage           hexutil.Uint64 `json:"resourceUsage"           gencodec:"required"`
	}
	var enc Receipt
	enc.PostState = r.PostState
	enc.Status = hexutil.Uint64(r.Status)
	enc.CumulativeResourceUsage = hexutil.Uint64(r.CumulativeResourceUsage)
	enc.Bloom = r.Bloom
	enc.Logs = r.Logs
	enc.TxHash = r.TxHash
	enc.ContractAddress = r.ContractAddress
	enc.ResourceUsage = hexutil.Uint64(r.ResourceUsage)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *Receipt) UnmarshalJSON(input []byte) error {
	type Receipt struct {
		PostState               *hexutil.Bytes  `json:"root"`
		Status                  *hexutil.Uint64 `json:"status"`
		CumulativeResourceUsage *hexutil.Uint64 `json:"cumulativeResourceUsage" gencodec:"required"`
		Bloom                   *Bloom          `json:"logsBloom"               gencodec:"required"`
		Logs                    []*Log          `json:"logs"                    gencodec:"required"`
		TxHash                  *common.Hash    `json:"transactionHash"         gencodec:"required"`
		ContractAddress         *common.Address `json:"contractAddress"`
		ResourceUsage           *hexutil.Uint64 `json:"resourceUsage"           gencodec:"required"`
	}
	var dec Receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.PostState != nil {
		r.PostState = *dec.PostState
	}
	if dec.Status != nil {
		r.Status = uint64(*dec.Status)
	}
	if dec.CumulativeResourceUsage == nil {
		return errors.New("missing required field 'cumulativeResourceUsage' for Receipt")
	}
	r.CumulativeResourceUsage = uint64(*dec.CumulativeResourceUsage)
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Receipt")
	}
	r.Bloom = *dec.Bloom
	if dec.Logs == nil {
		return errors.New("missing required field 'logs' for Receipt")
	}
	r.Logs = dec.Logs
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionHash' for Receipt")
	}
	r.TxHash = *dec.TxHash
	if dec.ContractAddress != nil {
		r.ContractAddress = *dec.ContractAddress
	}
	if dec.ResourceUsage == nil {
		return errors.New("missing required field 'resourceUsage' for Receipt")
	}
	r.ResourceUsage = uint64(*dec.ResourceUsage)
	return nil
}
