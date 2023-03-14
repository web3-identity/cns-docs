const ethers = require("ethers");
const { hash } = require("@ensdomains/eth-ens-namehash")
const { Conflux, Transaction, format } = require('js-conflux-sdk');


const conflux = new Conflux({
    url: 'https://test.confluxrpc.com',
    networkId: 1,
});
const { abi, bytecode } = require('./assets/web3RegisterController.json');
const web3RegController = conflux.Contract({ abi, bytecode, address: "cfxtest:aca1858y5a9fnyx9rxd1c9knr517cd0e6afzzhgj01" });

main()

async function main() {
    const commitHash = await web3RegController["makeCommitment(string,address,uint256,bytes32,address,bytes[],bool,uint16,uint64)"](
        "test3",
        "cfx:aapf8tkbyaw0sxbv30zh9s2nckn2c4ktajkhn5aha4",
        31536000,
        "0x00a28927897fddc247757c1d693760b8bc88017abf7293fd36d08cc3c5a57171",
        "cfx:acasaruvgf44ss67pxzfs1exvj7k2vyt863f72n6up",
        [genData()],
        true,
        0,
        1
    );
    console.log("maked commit hash: 0x%s" , commitHash.toString("hex"))
}

function genData() {
    const target = "cfxtest:aany9gz5hnpz9a22dkmkpcr276v932wze6cr821wd4"; // 要正向解析到的地址
    const node = hash("conflux.web3");
    const iface = new ethers.Interface(["function setAddr(bytes32 node, uint coinType, bytes memory a)"]);
    const data = iface.encodeFunctionData("setAddr", [node, 503, format.hexAddress(target)])
    return data
}