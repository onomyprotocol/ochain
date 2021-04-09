package onomygen

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/onomyprotocol/ochain/onomygen/key"
	"github.com/onomyprotocol/ochain/onomygen/network"
	"github.com/onomyprotocol/ochain/onomygen/node"
	"github.com/onomyprotocol/ochain/onomygen/utils"
)

type OnomyGen struct {
	chainID *string
}

func NewOnomyGen(chainID *string) OnomyGen {
	return OnomyGen{
		chainID: chainID,
	}
}

func (s OnomyGen) NewNetwork() *network.Network {
	return &network.Network{
		ChainID: *s.chainID,
		CLI:     utils.NewCLI(*s.chainID),
	}
}

func (s OnomyGen) NetworkCreate(count int, outputDir, startingIPAddress string, outputFile string) {
	net := network.NewNetwork(*s.chainID)
	summary, err := net.Build(count, outputDir, startingIPAddress)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = ioutil.WriteFile(outputFile, []byte(*summary), 0600); err != nil {
		log.Fatal(err)
		return
	}
}

func (s OnomyGen) NetworkReset(networkDir string) {
	if err := network.Reset(*s.chainID, networkDir); err != nil {
		log.Fatal(err)
	}
}

func (s OnomyGen) NewNode() *node.Node {
	return &node.Node{
		ChainID: *s.chainID,
		CLI:     utils.NewCLI(*s.chainID),
	}
}

func (s OnomyGen) NodeReset(nodeHomeDir *string) {
	if err := node.Reset(*s.chainID, nodeHomeDir); err != nil {
		log.Fatal(err)
	}
}

func (s OnomyGen) KeyGenerateMnemonic(name, password *string) {
	k := key.NewKey(name, password)
	k.GenerateMnemonic()
	fmt.Println(k.Mnemonic)
}

func (s OnomyGen) KeyRecoverFromMnemonic(mnemonic string) {
	k := key.NewKey(nil, nil)
	if err := k.RecoverFromMnemonic(mnemonic); err != nil {
		log.Fatal(err)
	}

	fmt.Println(heredoc.Doc(`
		Address: ` + k.Address + `
		Validator Address: ` + k.ValidatorAddress + `
		Consensus Address: ` + k.ConsensusAddress + `
	`))
}
