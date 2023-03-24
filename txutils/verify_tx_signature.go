package txutils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bnb-chain/zkbnb-crypto/wasm/txtypes"
	"github.com/bnb-chain/zkbnb-go-sdk/types"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
)

type PublicKey = eddsa.PublicKey

func parsePk(pkStr string) (pk *PublicKey, err error) {
	pkBytes, err := hex.DecodeString(pkStr)
	if err != nil {
		return nil, err
	}
	pk = new(PublicKey)
	size, err := pk.SetBytes(pkBytes)
	if err != nil {
		return nil, err
	}
	if size != 32 {
		return nil, errors.New("invalid public key")
	}
	return pk, nil
}

func ConvertTransferNftTxInfo(tx *types.TransferNftTxReq, ops *types.TransactOpts) *txtypes.TransferNftTxInfo {
	return &txtypes.TransferNftTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		ToL1Address:       ops.ToAccountAddress,
		NftIndex:          tx.NftIndex,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		CallData:          ops.CallData,
		CallDataHash:      ops.CallDataHash,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertWithdrawNftTxInfo(tx *types.WithdrawNftTxReq, ops *types.TransactOpts) *txtypes.WithdrawNftTxInfo {
	return &txtypes.WithdrawNftTxInfo{
		AccountIndex:      tx.AccountIndex,
		NftIndex:          tx.NftIndex,
		ToAddress:         tx.ToAddress,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertOfferTxInfo(tx *types.OfferTxInfo) *txtypes.OfferTxInfo {
	return &txtypes.OfferTxInfo{
		Type:               tx.Type,
		OfferId:            tx.OfferId,
		AccountIndex:       tx.AccountIndex,
		NftIndex:           tx.NftIndex,
		AssetId:            tx.AssetId,
		AssetAmount:        tx.AssetAmount,
		ListedAt:           tx.ListedAt,
		ExpiredAt:          tx.ExpiredAt,
		RoyaltyRate:        tx.RoyaltyRate,
		ChanelAccountIndex: tx.ChanelAccountIndex,
		ChanelRate:         tx.ChanelRate,
		ProtocolRate:       tx.ProtocolRate,
		ProtocolAmount:     tx.ProtocolAmount,
		Sig:                tx.Sig,
		L1Sig:              tx.L1Sig,
	}
}

func ConvertChangePubKeyTxInfo(tx *types.ChangePubKeyReq, ops *types.TransactOpts) *txtypes.ChangePubKeyInfo {
	var (
		pubKeyX = make([]byte, 32)
		pubKeyY = make([]byte, 32)
	)
	copy(pubKeyX[:], tx.PubKeyX[:])
	copy(pubKeyY[:], tx.PubKeyY[:])
	return &txtypes.ChangePubKeyInfo{
		AccountIndex:      ops.FromAccountIndex,
		L1Address:         tx.L1Address,
		PubKeyX:           pubKeyX,
		PubKeyY:           pubKeyY,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertMintNftTxInfo(tx *types.MintNftTxReq, ops *types.TransactOpts) *txtypes.MintNftTxInfo {
	return &txtypes.MintNftTxInfo{
		CreatorAccountIndex: ops.FromAccountIndex,
		ToAccountIndex:      ops.ToAccountIndex,
		ToL1Address:         ops.ToAccountAddress,
		NftCollectionId:     tx.NftCollectionId,
		NftContentType:      tx.NftContentType,
		RoyaltyRate:         tx.RoyaltyRate,
		GasAccountIndex:     ops.GasAccountIndex,
		GasFeeAssetId:       ops.GasFeeAssetId,
		GasFeeAssetAmount:   ops.GasFeeAssetAmount,
		ExpiredAt:           ops.ExpiredAt,
		Nonce:               ops.Nonce,
		MetaData:            tx.MetaData,
		MutableAttributes:   tx.MutableAttributes,
	}
}

func ConvertTransferTx(tx *types.TransferTxReq, ops *types.TransactOpts) *txtypes.TransferTxInfo {
	return &txtypes.TransferTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		ToL1Address:       ops.ToAccountAddress,
		AssetId:           tx.AssetId,
		AssetAmount:       tx.AssetAmount,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		Memo:              ops.Memo,
		CallData:          ops.CallData,
		CallDataHash:      ops.CallDataHash,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertWithdrawTx(tx *types.WithdrawTxReq, ops *types.TransactOpts) *txtypes.WithdrawTxInfo {
	return &txtypes.WithdrawTxInfo{
		FromAccountIndex:  ops.FromAccountIndex,
		AssetId:           tx.AssetId,
		AssetAmount:       tx.AssetAmount,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ToAddress:         tx.ToAddress,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertCreateCollectionTxInfo(tx *types.CreateCollectionTxReq, ops *types.TransactOpts) *txtypes.CreateCollectionTxInfo {
	return &txtypes.CreateCollectionTxInfo{
		AccountIndex:      ops.FromAccountIndex,
		Name:              tx.Name,
		Introduction:      tx.Introduction,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func ConvertAtomicMatchTxInfo(tx *types.AtomicMatchTxReq, ops *types.TransactOpts) *txtypes.AtomicMatchTxInfo {
	return &txtypes.AtomicMatchTxInfo{
		AccountIndex: ops.FromAccountIndex,
		BuyOffer: &txtypes.OfferTxInfo{
			Type:               tx.BuyOffer.Type,
			OfferId:            tx.BuyOffer.OfferId,
			AccountIndex:       tx.BuyOffer.AccountIndex,
			NftIndex:           tx.BuyOffer.NftIndex,
			AssetId:            tx.BuyOffer.AssetId,
			AssetAmount:        tx.BuyOffer.AssetAmount,
			ListedAt:           tx.BuyOffer.ListedAt,
			ExpiredAt:          tx.BuyOffer.ExpiredAt,
			RoyaltyRate:        tx.BuyOffer.RoyaltyRate,
			ChanelAccountIndex: tx.BuyOffer.ChanelAccountIndex,
			ChanelRate:         tx.BuyOffer.ChanelRate,
			ProtocolRate:       tx.BuyOffer.ProtocolRate,
			ProtocolAmount:     tx.BuyOffer.ProtocolAmount,
			Sig:                tx.BuyOffer.Sig,
			L1Sig:              tx.BuyOffer.L1Sig,
		},
		SellOffer: &txtypes.OfferTxInfo{
			Type:               tx.SellOffer.Type,
			OfferId:            tx.SellOffer.OfferId,
			AccountIndex:       tx.SellOffer.AccountIndex,
			NftIndex:           tx.SellOffer.NftIndex,
			AssetId:            tx.SellOffer.AssetId,
			AssetAmount:        tx.SellOffer.AssetAmount,
			ListedAt:           tx.SellOffer.ListedAt,
			ExpiredAt:          tx.SellOffer.ExpiredAt,
			ChanelAccountIndex: tx.SellOffer.ChanelAccountIndex,
			ChanelRate:         tx.SellOffer.ChanelRate,
			Sig:                tx.SellOffer.Sig,
			L1Sig:              tx.SellOffer.L1Sig,
		},
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		Nonce:             ops.Nonce,
		ExpiredAt:         ops.ExpiredAt,
	}
}

func ConvertCancelOfferTxInfo(tx *types.CancelOfferTxReq, ops *types.TransactOpts) *txtypes.CancelOfferTxInfo {
	return &txtypes.CancelOfferTxInfo{
		AccountIndex:      ops.FromAccountIndex,
		OfferId:           tx.OfferId,
		GasAccountIndex:   ops.GasAccountIndex,
		GasFeeAssetId:     ops.GasFeeAssetId,
		GasFeeAssetAmount: ops.GasFeeAssetAmount,
		ExpiredAt:         ops.ExpiredAt,
		Nonce:             ops.Nonce,
	}
}

func VerifyCancelOfferTxSig(pubKey string, tx *types.CancelOfferTxInfo) error {
	message, err := tx.Hash(mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyWithdrawNftTxSig(pubKey string, tx *types.WithdrawNftTxInfo) error {
	message, err := tx.Hash(mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyTransferNftTxSig(pubKey string, tx *types.TransferNftTxInfo) error {
	message, err := tx.Hash(mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyOfferTxSig(pubKey string, tx *types.OfferTxInfo) error {
	convertedTx := ConvertOfferTxInfo(tx)
	message, err := convertedTx.Hash(mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyMintNftTxSig(pubKey string, tx *types.MintNftTxInfo) error {
	message, err := tx.Hash(mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyCreateCollectionTxSig(pubKey string, tx *types.CreateCollectionTxInfo) error {
	message, err := tx.Hash(mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func VerifyAtomicMatchTxSig(pubKey string, tx *types.AtomicMatchTxInfo) error {
	message, err := tx.Hash(mimc.NewMiMC())
	if err != nil {
		return err
	}

	pk, err := parsePk(pubKey)
	if err != nil {
		return err
	}
	hFunc := mimc.NewMiMC()
	valid, err := pk.Verify(tx.Sig, message, hFunc)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}
	return nil
}
