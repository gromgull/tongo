package abi

// Code autogenerated. DO NOT EDIT.

import (
	"context"
	"fmt"
	"github.com/startfellows/tongo"
	"github.com/startfellows/tongo/boc"
	"github.com/startfellows/tongo/tlb"
)

type ContentData struct {
	tlb.SumType
	Snake struct {
		Data tlb.SnakeData
	} `tlbSumType:"#00"`
	Chunks struct {
		Data tlb.ChunkedData
	} `tlbSumType:"#01"`
}

type FullContent struct {
	tlb.SumType
	Onchain struct {
		Data tlb.HashmapE[tlb.Bits256, tlb.Ref[ContentData]]
	} `tlbSumType:"#00"`
	Offchain struct {
		Uri tlb.Text
	} `tlbSumType:"#01"`
}

type ClosingConfig struct {
	QuarantinDuration        uint32
	MisbehaviorFine          tlb.Grams
	ConditionalCloseDuration uint32
}

type PaymentConfig struct {
	ExcessFee tlb.Grams
	DestA     tlb.MsgAddress
	DestB     tlb.MsgAddress
}

type Storage struct {
	BalanceA       tlb.Grams
	BalanceB       tlb.Grams
	KeyA           tlb.Uint256
	KeyB           tlb.Uint256
	ChannelId      tlb.Uint128
	Config         tlb.Ref[ClosingConfig]
	CommitedSeqnoA uint32
	CommitedSeqnoB uint32
	Quarantin      tlb.Maybe[tlb.Ref[QuarantinedState]]
	Payments       tlb.Ref[PaymentConfig]
}

type ConditionalPayment struct {
	Amount    tlb.Grams
	Condition tlb.Any
}

type SemiChannelBody struct {
	Seqno        uint64
	Sent         tlb.Grams
	Conditionals tlb.HashmapE[tlb.Uint32, ConditionalPayment]
}

type SemiChannel struct {
	Magic            tlb.Magic `tlb:"#43685374"`
	ChannelId        tlb.Uint128
	Data             SemiChannelBody
	CounterpartyData tlb.Maybe[tlb.Ref[SemiChannelBody]]
}

type SignedSemiChannel struct {
	Signature tlb.Bits512
	State     SemiChannel
}

type QuarantinedState struct {
	StateA           SemiChannelBody
	StateB           SemiChannelBody
	QuarantineStarts uint32
	StateCommitedByA bool
}

type TextCommentTextCommentInternalMsgBody struct {
	Text tlb.Text
}

type PaymentChannelSettleConditionalsInternalMsgBody struct {
	FromA                bool
	Signature            tlb.Bits512
	Tag                  uint32
	ChannelId            tlb.Uint128
	ConditionalsToSettle tlb.HashmapE[tlb.Uint32, tlb.Any]
}

type Tep85SbtProveOwnershipInternalMsgBody struct {
	QueryId        uint64
	Dest           tlb.MsgAddress
	ForwardPayload tlb.Ref[tlb.Any]
	WithContent    bool
}

type Tep62NftItemOwnershipAssignedInternalMsgBody struct {
	QueryId        uint64
	PrevOwner      tlb.MsgAddress
	ForwardPayload tlb.EitherRef[tlb.Any]
}

type Tep85SbtOwnershipProofInternalMsgBody struct {
	QueryId   uint64
	ItemId    tlb.Uint256
	Owner     tlb.MsgAddress
	Data      tlb.Ref[tlb.Any]
	RevokedAt uint64
	Content   tlb.Maybe[tlb.Ref[tlb.Any]]
}

type Tep85SbtOwnerInfoInternalMsgBody struct {
	QueryId   uint64
	ItemId    tlb.Uint256
	Initiator tlb.MsgAddress
	Owner     tlb.MsgAddress
	Data      tlb.Ref[tlb.Any]
	RevokedAt uint64
	Content   tlb.Maybe[tlb.Ref[tlb.Any]]
}

type PaymentChannelCooperativeCloseInternalMsgBody struct {
	SigA      tlb.Ref[tlb.Bits512]
	SigB      tlb.Ref[tlb.Bits512]
	Tag       uint32
	ChannelId tlb.Uint128
	BalanceA  tlb.Grams
	BalanceB  tlb.Grams
	SeqnoA    uint64
	SeqnoB    uint64
}

type Tep74JettonWalletTransferInternalMsgBody struct {
	QueryId             uint64
	Amount              tlb.VarUInteger16
	Destination         tlb.MsgAddress
	ResponseDestination tlb.MsgAddress
	CustomPayload       tlb.Maybe[tlb.Ref[tlb.Any]]
	ForwardTonAmount    tlb.VarUInteger16
	ForwardPayload      tlb.EitherRef[tlb.Any]
}

type Tep85SbtDestroyInternalMsgBody struct {
	QueryId uint64
}

type PaymentChannelFinishUncooperativeCloseInternalMsgBody struct{}

type Tep62NftItemGetStaticDataInternalMsgBody struct {
	QueryId uint64
}

type PaymentChannelInitChannelInternalMsgBody struct {
	IsA       bool
	Signature tlb.Bits512
	Tag       uint32
	ChannelId tlb.Uint128
	BalanceA  tlb.Grams
	BalanceB  tlb.Grams
}

type Tep74JettonWalletBurnInternalMsgBody struct {
	QueryId             uint64
	Amount              tlb.VarUInteger16
	ResponseDestination tlb.MsgAddress
	CustomPayload       tlb.Maybe[tlb.Ref[tlb.Any]]
}

type Tep62NftItemTransferInternalMsgBody struct {
	QueryId             uint64
	NewOwner            tlb.MsgAddress
	ResponseDestination tlb.MsgAddress
	CustomPayload       tlb.Maybe[tlb.Ref[tlb.Any]]
	ForwardAmount       tlb.VarUInteger16
	ForwardPayload      tlb.EitherRef[tlb.Any]
}

type PaymentChannelTopUpBalanceInternalMsgBody struct {
	AddA tlb.Grams
	AddB tlb.Grams
}

type Tep66NftRoyaltyGetRoyaltyParamsInternalMsgBody struct {
	QueryId uint64
}

type Tep85SbtRevokeInternalMsgBody struct {
	QueryId uint64
}

type Tep74JettonWalletTransferNotificationInternalMsgBody struct {
	QueryId        uint64
	Amount         tlb.VarUInteger16
	Sender         tlb.MsgAddress
	ForwardPayload tlb.EitherRef[tlb.Any]
}

type Tep62NftItemReportStaticDataInternalMsgBody struct {
	QueryId    uint64
	Index      tlb.Uint256
	Collection tlb.MsgAddress
}

type Tep66NftRoyaltyReportRoyaltyParamsInternalMsgBody struct {
	QueryId     uint64
	Numerator   uint16
	Denominator uint16
	Destination tlb.MsgAddress
}

type PaymentChannelChallengeQuarantinedStateInternalMsgBody struct {
	ChallengedByA bool
	Signature     tlb.Bits512
	Tag           uint32
	ChannelId     tlb.Uint128
	SchA          tlb.Ref[SignedSemiChannel]
	SchB          tlb.Ref[SignedSemiChannel]
}

type PaymentChannelStartUncooperativeCloseInternalMsgBody struct {
	SignedByA bool
	Signature tlb.Bits512
	Tag       uint32
	ChannelId tlb.Uint128
	SchA      tlb.Ref[SignedSemiChannel]
	SchB      tlb.Ref[SignedSemiChannel]
}

type PaymentChannelCooperativeCommitInternalMsgBody struct {
	SigA      tlb.Ref[tlb.Bits512]
	SigB      tlb.Ref[tlb.Bits512]
	Tag       uint32
	ChannelId tlb.Uint128
	SeqnoA    uint64
	SeqnoB    uint64
}

type Tep85SbtRequestOwnerInternalMsgBody struct {
	QueryId        uint64
	Dest           tlb.MsgAddress
	ForwardPayload tlb.Ref[tlb.Any]
	WithContent    bool
}

type ExcessesInternalMsgBody struct {
	QueryId uint64
}

type PaymentChannelChannelClosedInternalMsgBody struct {
	ChannelId tlb.Uint128
}

func MessageDecoder(cell *boc.Cell) (string, any, error) {
	tag, err := cell.ReadUint(32)
	if err != nil {
		return "", nil, err
	}
	switch tag {
	case 0x0:
		var res TextCommentTextCommentInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "TextCommentTextComment", res, err
	case 0x4291097:
		var res PaymentChannelSettleConditionalsInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "PaymentChannelSettleConditionals", res, err
	case 0x4ded148:
		var res Tep85SbtProveOwnershipInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep85SbtProveOwnership", res, err
	case 0x5138d91:
		var res Tep62NftItemOwnershipAssignedInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep62NftItemOwnershipAssigned", res, err
	case 0x524c7ae:
		var res Tep85SbtOwnershipProofInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep85SbtOwnershipProof", res, err
	case 0xdd607e3:
		var res Tep85SbtOwnerInfoInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep85SbtOwnerInfo", res, err
	case 0xdf395b9:
		var res PaymentChannelCooperativeCloseInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "PaymentChannelCooperativeClose", res, err
	case 0xf8a7ea5:
		var res Tep74JettonWalletTransferInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep74JettonWalletTransfer", res, err
	case 0x1f04537a:
		var res Tep85SbtDestroyInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep85SbtDestroy", res, err
	case 0x25432a91:
		var res PaymentChannelFinishUncooperativeCloseInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "PaymentChannelFinishUncooperativeClose", res, err
	case 0x2fcb26a2:
		var res Tep62NftItemGetStaticDataInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep62NftItemGetStaticData", res, err
	case 0x56103bba:
		var res PaymentChannelInitChannelInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "PaymentChannelInitChannel", res, err
	case 0x595f07bc:
		var res Tep74JettonWalletBurnInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep74JettonWalletBurn", res, err
	case 0x5fcc3d14:
		var res Tep62NftItemTransferInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep62NftItemTransfer", res, err
	case 0x67c7d281:
		var res PaymentChannelTopUpBalanceInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "PaymentChannelTopUpBalance", res, err
	case 0x693d3950:
		var res Tep66NftRoyaltyGetRoyaltyParamsInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep66NftRoyaltyGetRoyaltyParams", res, err
	case 0x6f89f5e3:
		var res Tep85SbtRevokeInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep85SbtRevoke", res, err
	case 0x7362d09c:
		var res Tep74JettonWalletTransferNotificationInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep74JettonWalletTransferNotification", res, err
	case 0x8b771735:
		var res Tep62NftItemReportStaticDataInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep62NftItemReportStaticData", res, err
	case 0xa8cb00ad:
		var res Tep66NftRoyaltyReportRoyaltyParamsInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep66NftRoyaltyReportRoyaltyParams", res, err
	case 0xab15bc6a:
		var res PaymentChannelChallengeQuarantinedStateInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "PaymentChannelChallengeQuarantinedState", res, err
	case 0xb942f428:
		var res PaymentChannelStartUncooperativeCloseInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "PaymentChannelStartUncooperativeClose", res, err
	case 0xcf19aff5:
		var res PaymentChannelCooperativeCommitInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "PaymentChannelCooperativeCommit", res, err
	case 0xd0c3bfea:
		var res Tep85SbtRequestOwnerInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Tep85SbtRequestOwner", res, err
	case 0xd53276db:
		var res ExcessesInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "Excesses", res, err
	case 0xdddc88ba:
		var res PaymentChannelChannelClosedInternalMsgBody
		err = tlb.Unmarshal(cell, &res)
		return "PaymentChannelChannelClosed", res, err
	}
	return "", nil, fmt.Errorf("invalid message tag")
}

type Executor interface {
	RunSmcMethod(ctx context.Context, accountID tongo.AccountID, method string, params tlb.VmStack) (uint32, tlb.VmStack, error)
}

type GetNftDataResult struct {
	Init              int8
	Index             tlb.Int257
	CollectionAddress tongo.AccountID
	OwnerAddress      tongo.AccountID
	IndividualContent tlb.Any
}

func GetNftData(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_nft_data", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetNftDataResult(stack)
}

func decodeGetNftDataResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 5 || (stack[0].SumType != "VmStkTinyInt") || (stack[1].SumType != "VmStkTinyInt" && stack[1].SumType != "VmStkInt") || (stack[2].SumType != "VmStkSlice") || (stack[3].SumType != "VmStkSlice") || (stack[4].SumType != "VmStkCell") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var init int8
	init = int8(stack[0].Int64())
	var index tlb.Int257
	index = stack[1].Int257()
	var collectionAddress tongo.AccountID
	err = stack[2].VmStkSlice.UnmarshalToTlbStruct(&collectionAddress)
	if err != nil {
		return "", nil, err
	}
	var ownerAddress tongo.AccountID
	err = stack[3].VmStkSlice.UnmarshalToTlbStruct(&ownerAddress)
	if err != nil {
		return "", nil, err
	}
	var individualContent tlb.Any
	individualContentCell := &stack[4].VmStkCell.Value
	err = tlb.Unmarshal(individualContentCell, &individualContent)
	if err != nil {
		return "", nil, err
	}
	return "GetNftDataResult", GetNftDataResult{
		Init:              init,
		Index:             index,
		CollectionAddress: collectionAddress,
		OwnerAddress:      ownerAddress,
		IndividualContent: individualContent,
	}, nil
}

type GetNftContentResult struct {
	Content FullContent
}

func GetNftContent(ctx context.Context, executor Executor, reqAccountID tongo.AccountID, index tlb.Int257, individualContent tlb.Any) (string, any, error) {
	stack := tlb.VmStack{}
	var (
		val tlb.VmStackValue
		err error
	)
	val = tlb.VmStackValue{SumType: "VmStkInt", VmStkInt: index}
	stack.Put(val)
	val, err = tlb.TlbStructToVmCell(individualContent)
	if err != nil {
		return "", nil, err
	}
	stack.Put(val)

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_nft_content", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetNftContentResult(stack)
}

func decodeGetNftContentResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 1 || (stack[0].SumType != "VmStkCell") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var content FullContent
	contentCell := &stack[0].VmStkCell.Value
	err = tlb.Unmarshal(contentCell, &content)
	if err != nil {
		return "", nil, err
	}
	return "GetNftContentResult", GetNftContentResult{
		Content: content,
	}, nil
}

type GetCollectionDataResult struct {
	NextItemIndex     tlb.Int257
	CollectionContent tlb.Any
	OwnerAddress      tongo.AccountID
}

func GetCollectionData(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_collection_data", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetCollectionDataResult(stack)
}

func decodeGetCollectionDataResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 3 || (stack[0].SumType != "VmStkTinyInt" && stack[0].SumType != "VmStkInt") || (stack[1].SumType != "VmStkCell") || (stack[2].SumType != "VmStkSlice") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var nextItemIndex tlb.Int257
	nextItemIndex = stack[0].Int257()
	var collectionContent tlb.Any
	collectionContentCell := &stack[1].VmStkCell.Value
	err = tlb.Unmarshal(collectionContentCell, &collectionContent)
	if err != nil {
		return "", nil, err
	}
	var ownerAddress tongo.AccountID
	err = stack[2].VmStkSlice.UnmarshalToTlbStruct(&ownerAddress)
	if err != nil {
		return "", nil, err
	}
	return "GetCollectionDataResult", GetCollectionDataResult{
		NextItemIndex:     nextItemIndex,
		CollectionContent: collectionContent,
		OwnerAddress:      ownerAddress,
	}, nil
}

type GetNftAddressByIndexResult struct {
	Address tongo.AccountID
}

func GetNftAddressByIndex(ctx context.Context, executor Executor, reqAccountID tongo.AccountID, index tlb.Int257) (string, any, error) {
	stack := tlb.VmStack{}
	var (
		val tlb.VmStackValue
		err error
	)
	val = tlb.VmStackValue{SumType: "VmStkInt", VmStkInt: index}
	stack.Put(val)

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_nft_address_by_index", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetNftAddressByIndexResult(stack)
}

func decodeGetNftAddressByIndexResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 1 || (stack[0].SumType != "VmStkSlice") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var address tongo.AccountID
	err = stack[0].VmStkSlice.UnmarshalToTlbStruct(&address)
	if err != nil {
		return "", nil, err
	}
	return "GetNftAddressByIndexResult", GetNftAddressByIndexResult{
		Address: address,
	}, nil
}

type RoyaltyParamsResult struct {
	Numerator   uint16
	Denominator uint16
	Destination tongo.AccountID
}

func RoyaltyParams(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "royalty_params", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeRoyaltyParamsResult(stack)
}

func decodeRoyaltyParamsResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 3 || (stack[0].SumType != "VmStkTinyInt") || (stack[1].SumType != "VmStkTinyInt") || (stack[2].SumType != "VmStkSlice") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var numerator uint16
	numerator = uint16(stack[0].Int64())
	var denominator uint16
	denominator = uint16(stack[1].Int64())
	var destination tongo.AccountID
	err = stack[2].VmStkSlice.UnmarshalToTlbStruct(&destination)
	if err != nil {
		return "", nil, err
	}
	return "RoyaltyParamsResult", RoyaltyParamsResult{
		Numerator:   numerator,
		Denominator: denominator,
		Destination: destination,
	}, nil
}

type GetJettonDataResult struct {
	TotalSupply      tlb.Int257
	Mintable         int8
	AdminAddress     tongo.AccountID
	JettonContent    tlb.Any
	JettonWalletCode tlb.Any
}

func GetJettonData(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_jetton_data", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetJettonDataResult(stack)
}

func decodeGetJettonDataResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 5 || (stack[0].SumType != "VmStkTinyInt" && stack[0].SumType != "VmStkInt") || (stack[1].SumType != "VmStkTinyInt") || (stack[2].SumType != "VmStkSlice") || (stack[3].SumType != "VmStkCell") || (stack[4].SumType != "VmStkCell") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var totalSupply tlb.Int257
	totalSupply = stack[0].Int257()
	var mintable int8
	mintable = int8(stack[1].Int64())
	var adminAddress tongo.AccountID
	err = stack[2].VmStkSlice.UnmarshalToTlbStruct(&adminAddress)
	if err != nil {
		return "", nil, err
	}
	var jettonContent tlb.Any
	jettonContentCell := &stack[3].VmStkCell.Value
	err = tlb.Unmarshal(jettonContentCell, &jettonContent)
	if err != nil {
		return "", nil, err
	}
	var jettonWalletCode tlb.Any
	jettonWalletCodeCell := &stack[4].VmStkCell.Value
	err = tlb.Unmarshal(jettonWalletCodeCell, &jettonWalletCode)
	if err != nil {
		return "", nil, err
	}
	return "GetJettonDataResult", GetJettonDataResult{
		TotalSupply:      totalSupply,
		Mintable:         mintable,
		AdminAddress:     adminAddress,
		JettonContent:    jettonContent,
		JettonWalletCode: jettonWalletCode,
	}, nil
}

type GetWalletAddressResult struct {
	JettonWalletAddress tongo.AccountID
}

func GetWalletAddress(ctx context.Context, executor Executor, reqAccountID tongo.AccountID, ownerAddress tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}
	var (
		val tlb.VmStackValue
		err error
	)
	val, err = tlb.TlbStructToVmCellSlice(ownerAddress)
	if err != nil {
		return "", nil, err
	}
	stack.Put(val)

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_wallet_address", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetWalletAddressResult(stack)
}

func decodeGetWalletAddressResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 1 || (stack[0].SumType != "VmStkSlice") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var jettonWalletAddress tongo.AccountID
	err = stack[0].VmStkSlice.UnmarshalToTlbStruct(&jettonWalletAddress)
	if err != nil {
		return "", nil, err
	}
	return "GetWalletAddressResult", GetWalletAddressResult{
		JettonWalletAddress: jettonWalletAddress,
	}, nil
}

type GetWalletDataResult struct {
	Balance          tlb.Int257
	Owner            tongo.AccountID
	Jetton           tongo.AccountID
	JettonWalletCode tlb.Any
}

func GetWalletData(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_wallet_data", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetWalletDataResult(stack)
}

func decodeGetWalletDataResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 4 || (stack[0].SumType != "VmStkTinyInt" && stack[0].SumType != "VmStkInt") || (stack[1].SumType != "VmStkSlice") || (stack[2].SumType != "VmStkSlice") || (stack[3].SumType != "VmStkCell") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var balance tlb.Int257
	balance = stack[0].Int257()
	var owner tongo.AccountID
	err = stack[1].VmStkSlice.UnmarshalToTlbStruct(&owner)
	if err != nil {
		return "", nil, err
	}
	var jetton tongo.AccountID
	err = stack[2].VmStkSlice.UnmarshalToTlbStruct(&jetton)
	if err != nil {
		return "", nil, err
	}
	var jettonWalletCode tlb.Any
	jettonWalletCodeCell := &stack[3].VmStkCell.Value
	err = tlb.Unmarshal(jettonWalletCodeCell, &jettonWalletCode)
	if err != nil {
		return "", nil, err
	}
	return "GetWalletDataResult", GetWalletDataResult{
		Balance:          balance,
		Owner:            owner,
		Jetton:           jetton,
		JettonWalletCode: jettonWalletCode,
	}, nil
}

type RecordDnsresolveResult struct {
	ResolvedBits int64
	Result       tlb.DNSRecord
}

type RecordsDnsresolveResult struct {
	ResolvedBits int64
	Result       tlb.DNSRecordSet
}

func Dnsresolve(ctx context.Context, executor Executor, reqAccountID tongo.AccountID, domain []byte, category tlb.Int257) (string, any, error) {
	stack := tlb.VmStack{}
	var (
		val tlb.VmStackValue
		err error
	)
	val, err = tlb.TlbStructToVmCellSlice(domain)
	if err != nil {
		return "", nil, err
	}
	stack.Put(val)
	val = tlb.VmStackValue{SumType: "VmStkInt", VmStkInt: category}
	stack.Put(val)

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "dnsresolve", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	for _, f := range []func(tlb.VmStack) (string, any, error){decodeRecordDnsresolveResult, decodeRecordsDnsresolveResult} {
		s, r, err := f(stack)
		if err == nil {
			return s, r, nil
		}
	}
	return "", nil, fmt.Errorf("can not decode outputs")
}

func decodeRecordDnsresolveResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 2 || (stack[0].SumType != "VmStkTinyInt") || (stack[1].SumType != "VmStkCell") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var resolvedBits int64
	resolvedBits = int64(stack[0].Int64())
	var result tlb.DNSRecord
	resultCell := &stack[1].VmStkCell.Value
	err = tlb.Unmarshal(resultCell, &result)
	if err != nil {
		return "", nil, err
	}
	return "RecordDnsresolveResult", RecordDnsresolveResult{
		ResolvedBits: resolvedBits,
		Result:       result,
	}, nil
}

func decodeRecordsDnsresolveResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 2 || (stack[0].SumType != "VmStkTinyInt") || (stack[1].SumType != "VmStkCell") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var resolvedBits int64
	resolvedBits = int64(stack[0].Int64())
	var result tlb.DNSRecordSet
	resultCell := &stack[1].VmStkCell.Value
	err = tlb.Unmarshal(resultCell, &result)
	if err != nil {
		return "", nil, err
	}
	return "RecordsDnsresolveResult", RecordsDnsresolveResult{
		ResolvedBits: resolvedBits,
		Result:       result,
	}, nil
}

type BasicGetSaleDataResult struct {
	Marketplace    tongo.AccountID
	Nft            tongo.AccountID
	Owner          tongo.AccountID
	FullPrice      tlb.Int257
	MarketFee      uint64
	RoyaltyAddress tongo.AccountID
	RoyaltyAmount  uint64
}

type GetgemsGetSaleDataResult struct {
	FixPrice         uint64
	IsComplete       bool
	CreatedAt        uint64
	Marketplace      tongo.AccountID
	Nft              tongo.AccountID
	Owner            tongo.AccountID
	FullPrice        tlb.Int257
	MarketFeeAddress tongo.AccountID
	MarketFee        uint64
	RoyaltyAddress   tongo.AccountID
	RoyaltyAmount    uint64
}

func GetSaleData(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_sale_data", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	for _, f := range []func(tlb.VmStack) (string, any, error){decodeBasicGetSaleDataResult, decodeGetgemsGetSaleDataResult} {
		s, r, err := f(stack)
		if err == nil {
			return s, r, nil
		}
	}
	return "", nil, fmt.Errorf("can not decode outputs")
}

func decodeBasicGetSaleDataResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 7 || (stack[0].SumType != "VmStkSlice") || (stack[1].SumType != "VmStkSlice") || (stack[2].SumType != "VmStkSlice") || (stack[3].SumType != "VmStkTinyInt" && stack[3].SumType != "VmStkInt") || (stack[4].SumType != "VmStkTinyInt") || (stack[5].SumType != "VmStkSlice") || (stack[6].SumType != "VmStkTinyInt") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var marketplace tongo.AccountID
	err = stack[0].VmStkSlice.UnmarshalToTlbStruct(&marketplace)
	if err != nil {
		return "", nil, err
	}
	var nft tongo.AccountID
	err = stack[1].VmStkSlice.UnmarshalToTlbStruct(&nft)
	if err != nil {
		return "", nil, err
	}
	var owner tongo.AccountID
	err = stack[2].VmStkSlice.UnmarshalToTlbStruct(&owner)
	if err != nil {
		return "", nil, err
	}
	var fullPrice tlb.Int257
	fullPrice = stack[3].Int257()
	var marketFee uint64
	marketFee = uint64(stack[4].Int64())
	var royaltyAddress tongo.AccountID
	err = stack[5].VmStkSlice.UnmarshalToTlbStruct(&royaltyAddress)
	if err != nil {
		return "", nil, err
	}
	var royaltyAmount uint64
	royaltyAmount = uint64(stack[6].Int64())
	return "BasicGetSaleDataResult", BasicGetSaleDataResult{
		Marketplace:    marketplace,
		Nft:            nft,
		Owner:          owner,
		FullPrice:      fullPrice,
		MarketFee:      marketFee,
		RoyaltyAddress: royaltyAddress,
		RoyaltyAmount:  royaltyAmount,
	}, nil
}

func decodeGetgemsGetSaleDataResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 11 || (stack[0].SumType != "VmStkTinyInt") || (stack[1].SumType != "VmStkTinyInt") || (stack[2].SumType != "VmStkTinyInt") || (stack[3].SumType != "VmStkSlice") || (stack[4].SumType != "VmStkSlice") || (stack[5].SumType != "VmStkSlice") || (stack[6].SumType != "VmStkTinyInt" && stack[6].SumType != "VmStkInt") || (stack[7].SumType != "VmStkSlice") || (stack[8].SumType != "VmStkTinyInt") || (stack[9].SumType != "VmStkSlice") || (stack[10].SumType != "VmStkTinyInt") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var fixPrice uint64
	fixPrice = uint64(stack[0].Int64())
	var isComplete bool
	isComplete = stack[1].Int64() != 0
	var createdAt uint64
	createdAt = uint64(stack[2].Int64())
	var marketplace tongo.AccountID
	err = stack[3].VmStkSlice.UnmarshalToTlbStruct(&marketplace)
	if err != nil {
		return "", nil, err
	}
	var nft tongo.AccountID
	err = stack[4].VmStkSlice.UnmarshalToTlbStruct(&nft)
	if err != nil {
		return "", nil, err
	}
	var owner tongo.AccountID
	err = stack[5].VmStkSlice.UnmarshalToTlbStruct(&owner)
	if err != nil {
		return "", nil, err
	}
	var fullPrice tlb.Int257
	fullPrice = stack[6].Int257()
	var marketFeeAddress tongo.AccountID
	err = stack[7].VmStkSlice.UnmarshalToTlbStruct(&marketFeeAddress)
	if err != nil {
		return "", nil, err
	}
	var marketFee uint64
	marketFee = uint64(stack[8].Int64())
	var royaltyAddress tongo.AccountID
	err = stack[9].VmStkSlice.UnmarshalToTlbStruct(&royaltyAddress)
	if err != nil {
		return "", nil, err
	}
	var royaltyAmount uint64
	royaltyAmount = uint64(stack[10].Int64())
	return "GetgemsGetSaleDataResult", GetgemsGetSaleDataResult{
		FixPrice:         fixPrice,
		IsComplete:       isComplete,
		CreatedAt:        createdAt,
		Marketplace:      marketplace,
		Nft:              nft,
		Owner:            owner,
		FullPrice:        fullPrice,
		MarketFeeAddress: marketFeeAddress,
		MarketFee:        marketFee,
		RoyaltyAddress:   royaltyAddress,
		RoyaltyAmount:    royaltyAmount,
	}, nil
}

type GetAuthorityAddressResult struct {
	Address tongo.AccountID
}

func GetAuthorityAddress(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_authority_address", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetAuthorityAddressResult(stack)
}

func decodeGetAuthorityAddressResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 1 || (stack[0].SumType != "VmStkSlice") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var address tongo.AccountID
	err = stack[0].VmStkSlice.UnmarshalToTlbStruct(&address)
	if err != nil {
		return "", nil, err
	}
	return "GetAuthorityAddressResult", GetAuthorityAddressResult{
		Address: address,
	}, nil
}

type GetRevokedTimeResult struct {
	Time uint64
}

func GetRevokedTime(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_revoked_time", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetRevokedTimeResult(stack)
}

func decodeGetRevokedTimeResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 1 || (stack[0].SumType != "VmStkTinyInt") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var time uint64
	time = uint64(stack[0].Int64())
	return "GetRevokedTimeResult", GetRevokedTimeResult{
		Time: time,
	}, nil
}

type GetChannelStateResult struct {
	State uint64
}

func GetChannelState(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error) {
	stack := tlb.VmStack{}

	errCode, stack, err := executor.RunSmcMethod(ctx, reqAccountID, "get_channel_state", stack)
	if err != nil {
		return "", nil, err
	}
	if errCode != 0 && errCode != 1 {
		return "", nil, fmt.Errorf("method execution failed with code: %v", errCode)
	}
	return decodeGetChannelStateResult(stack)
}

func decodeGetChannelStateResult(stack tlb.VmStack) (resultType string, resultAny any, err error) {
	if len(stack) != 1 || (stack[0].SumType != "VmStkTinyInt") {
		return "", nil, fmt.Errorf("invalid stack format")
	}
	var state uint64
	state = uint64(stack[0].Int64())
	return "GetChannelStateResult", GetChannelStateResult{
		State: state,
	}, nil
}
