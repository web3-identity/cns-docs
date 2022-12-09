---
description: 快速找到你想要的接入方式
---

# Quick Start

## App 中使用/展示 .web3 用户名

如果需要在您的 App/dApp 中使用/展示 .Web3 用户名，则需要实现正向解析和反向解析功能；

****

**正向解析：**通过 **.web3 用户名查询**到对应的区块链地址；通常用于提升转账和交易的便捷性；

* 比如：转账时用户通过输入 xxx.web3，即可向其解析到的区块链地址转账；

**反向解析：**通过**区块链地址**查询到对应的 **.web3 用户名**；通常用于作为地址昵称、社交昵称的展示；

* 比如：在用户账户中心或好友列表中，不再展示复杂的地址，而是以 .web3 用户名作为昵称；

实现正反向解析，请参阅：

{% embed url="https://docs.ens.domains/contract-api-reference/publicresolver" %}

{% content-ref url="Deployment.md" %}
[Deployment.md](Deployment.md)
{% endcontent-ref %}



## App 中支持 .web3 用户名注册与售卖

我们提供了便捷的 API 方便大家接入 .web3 用户名的注册与售卖；

* 您需要自行接入支付渠道，并完成收款流程；

详情请参阅：

{% content-ref url="tutorials/cns-backend.md" %}
[cns-backend.md](tutorials/cns-backend.md)
{% endcontent-ref %}





