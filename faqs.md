# FAQs

## CNS 与 ENS 有什么区别？

CNS 的核心顶级域名是 `.web3`, 采用最新的 NameWrapper 实现，所有的域名通过 NameWrapper 封装了一个 1155 NFT。

## 如何进行域名转移？

CNS 域名转移，可以通过 NameWrapper 的 `safeTransferFrom` 方法进行转移。

## 如何配置域名，包括 resolver， ttl，subrecord 等？

可调用 NameWrapper 的相关方法即可。