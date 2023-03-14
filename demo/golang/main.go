package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ens "github.com/wealdtech/go-ens/v3"
	"github.com/web3-identity/cns-golang-demo/contracts"
)

var (
	// web3RegisterController, _ = abi.JSON(strings.NewReader(contracts.Web3RegisterControllerABI))
	addrResolver, _ = abi.JSON(strings.NewReader(contracts.AddrResolverABI))
)

func main() {
	node, err := ens.NameHash("conflux.web3")
	if err != nil {
		panic(err)
	}
	fmt.Printf("node: 0x%x\n", node)
	target := cfxaddress.MustNew("cfxtest:aany9gz5hnpz9a22dkmkpcr276v932wze6cr821wd4")
	abiEncoded, err := addrResolver.Pack("setAddr", node, big.NewInt(503), target.MustGetCommonAddress().Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Printf("set addr abiEncoded: 0x%x\n", abiEncoded)
}
