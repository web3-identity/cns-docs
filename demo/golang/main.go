package main

import (
	"fmt"
	"math/big"
	"strings"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ens "github.com/wealdtech/go-ens/v3"
	"github.com/web3-identity/cns-golang-demo/contracts"
)

var (
	// web3RegisterControllerABI, _ = abi.JSON(strings.NewReader(contracts.Web3RegisterControllerABI))
	addrResolverABI, _     = abi.JSON(strings.NewReader(contracts.AddrResolverABI))
	web3RegisterController *contracts.Web3RegisterController
)

func init() {
	client, err := sdk.NewClient("https://test.confluxrpc.com")
	if err != nil {
		panic(err)
	}

	web3RegisterController, _ = contracts.NewWeb3RegisterController(cfxaddress.MustNew("cfxtest:aca1858y5a9fnyx9rxd1c9knr517cd0e6afzzhgj01"), client)
}

func main() {
	fmt.Printf("commit hash: 0x%x", makeCommitment())
}

func makeCommitment() [32]byte {
	data := genData()
	name := "ywfynb11"
	owner := cfxaddress.MustNew("cfxtest:aasxcuy07j0hg3x48mxzjhkhk7krksejvapx43n351")
	duration := 31536000
	secret := "0x0000000000000000000000000000000000000000000000000000000000000000"
	resolver := cfxaddress.MustNew("cfxtest:acbfyf69zaxau5a23w10dgyrmb0hrz4p9pewn6sejp")
	reverse_record := true
	fuses := 0
	wrapper_expiry := 1691003455

	commitHash, err := web3RegisterController.MakeCommitment(nil, name,
		owner.MustGetCommonAddress(),
		big.NewInt(int64(duration)),
		common.HexToHash(secret),
		resolver.MustGetCommonAddress(),
		[][]byte{data},
		reverse_record, uint16(fuses),
		uint64(wrapper_expiry))
	if err != nil {
		panic(err)
	}
	return commitHash
}

func genData() []byte {
	node, err := ens.NameHash("conflux.web3")
	if err != nil {
		panic(err)
	}
	fmt.Printf("node: 0x%x\n", node)
	target := cfxaddress.MustNew("cfxtest:aany9gz5hnpz9a22dkmkpcr276v932wze6cr821wd4")
	abiEncoded, err := addrResolverABI.Pack("setAddr", node, big.NewInt(503), target.MustGetCommonAddress().Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Printf("set addr abiEncoded: 0x%x\n", abiEncoded)
	return abiEncoded
}
