package abi

// Code autogenerated. DO NOT EDIT.

import (
	"errors"
	"github.com/tonkeeper/tongo/boc"
	"github.com/tonkeeper/tongo/tlb"
)

func (j *NFTPayload) UnmarshalTLB(cell *boc.Cell, decoder *tlb.Decoder) error {
	if cell.BitsAvailableForRead() == 0 && cell.RefsAvailableForRead() == 0 {
		return nil
	}
	tempCell := cell.CopyRemaining()
	op64, err := tempCell.ReadUint(32)
	if errors.Is(err, boc.ErrNotEnoughBits) {
		j.SumType = UnknownNFTOp
		j.Value = cell.CopyRemaining()
		return nil
	}
	op := uint32(op64)
	j.OpCode = &op
	switch op {
	case TextCommentNFTOpCode: // 0x00000000
		var res TextCommentNFTPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = TextCommentNFTOp
			j.Value = res
			return nil
		}
	case EncryptedTextCommentNFTOpCode: // 0x2167da4b
		var res EncryptedTextCommentNFTPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = EncryptedTextCommentNFTOp
			j.Value = res
			return nil
		}
	case TeleitemBidInfoNFTOpCode: // 0x38127de1
		var res TeleitemBidInfoNFTPayload
		if err := tlb.Unmarshal(tempCell, &res); err == nil {
			j.SumType = TeleitemBidInfoNFTOp
			j.Value = res
			return nil
		}

	}
	j.SumType = UnknownNFTOp
	j.Value = cell.CopyRemaining()

	return nil
}

const (
	TextCommentNFTOp          NFTOpName = "TextComment"
	EncryptedTextCommentNFTOp NFTOpName = "EncryptedTextComment"
	TeleitemBidInfoNFTOp      NFTOpName = "TeleitemBidInfo"

	TextCommentNFTOpCode          NFTOpCode = 0x00000000
	EncryptedTextCommentNFTOpCode NFTOpCode = 0x2167da4b
	TeleitemBidInfoNFTOpCode      NFTOpCode = 0x38127de1
)

var KnownNFTTypes = map[string]any{
	TextCommentNFTOp:          TextCommentNFTPayload{},
	EncryptedTextCommentNFTOp: EncryptedTextCommentNFTPayload{},
	TeleitemBidInfoNFTOp:      TeleitemBidInfoNFTPayload{},
}
var nftOpCodes = map[NFTOpName]NFTOpCode{
	TextCommentNFTOp:          TextCommentNFTOpCode,
	EncryptedTextCommentNFTOp: EncryptedTextCommentNFTOpCode,
	TeleitemBidInfoNFTOp:      TeleitemBidInfoNFTOpCode,
}

type TextCommentNFTPayload struct {
	Text tlb.Text
}

type EncryptedTextCommentNFTPayload struct {
	CipherText tlb.Bytes
}

type TeleitemBidInfoNFTPayload struct {
	Bid   tlb.Grams
	BidTs uint32
}
