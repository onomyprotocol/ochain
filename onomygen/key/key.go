package key

import (
	"github.com/cosmos/go-bip39"
	"github.com/onomyprotocol/ochain/app"
	"github.com/onomyprotocol/onomy-sdk/crypto/hd"
	"github.com/onomyprotocol/onomy-sdk/crypto/keyring"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
)

var (
	hdpath = hd.NewFundraiserParams(0, sdk.CoinType, 0)
)

type Key struct {
	Name             *string
	Mnemonic         string
	Password         *string
	Address          string
	ValidatorAddress string
	ConsensusAddress string
	Keybase          keyring.Keyring
}

func NewKey(name, password *string) *Key {
	return &Key{
		Name:     name,
		Password: password,
		Keybase:  keyring.NewInMemory(),
	}
}

func (k *Key) GenerateMnemonic() {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	k.Mnemonic = mnemonic
}

func (k *Key) RecoverFromMnemonic(mnemonic string) error {
	k.setConfig()
	k.Mnemonic = mnemonic

	account, err := k.Keybase.NewAccount(*k.Name, k.Mnemonic, *k.Password, hdpath.String(), hd.Secp256k1)
	if err != nil {
		return err
	}

	consensusAddress := sdk.ConsAddress(account.GetPubKey().Address()).String()
	validatorAddress := sdk.ValAddress(account.GetPubKey().Address()).String()

	k.Address = account.GetAddress().String()
	k.ValidatorAddress = validatorAddress
	k.ConsensusAddress = consensusAddress

	return nil
}

func (k *Key) setConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, app.AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(app.ValidatorAddressPrefix, app.ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(app.ConsNodeAddressPrefix, app.ConsNodePubKeyPrefix)
	config.Seal()
}
