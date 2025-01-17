package eth_test

import (
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/system/address/eth"
	"github.com/33cn/chain33/system/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/crypto"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatEthAddr(t *testing.T) {

	ethDriver, err := address.LoadDriver(eth.ID, -1)
	require.Nil(t, err)
	d := secp256k1.Driver{}
	for i := 0; i < 100; i++ {

		ethPriv, err := crypto.GenerateKey()
		require.Nil(t, err)
		ethAddr := crypto.PubkeyToAddress(ethPriv.PublicKey).Hex()
		chain33Priv, err := d.PrivKeyFromBytes(crypto.FromECDSA(ethPriv))
		require.Nil(t, err)
		err = ethDriver.ValidateAddr(ethAddr)
		require.Nil(t, err)
		addr := ethDriver.PubKeyToAddr(chain33Priv.PubKey().Bytes())
		require.Equal(t, ethAddr, addr)
	}
	require.Equal(t, eth.Name, ethDriver.GetName())
}
