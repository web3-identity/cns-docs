# PublicResolver

PublicResolver 是一个套公共的、符合 .web3 用户名的标准解析器，允许所有人查询正反向解析记录，允许所有者更新解析记录。

.web3 域名的公共解析器与 ENS 的基本一致，您可以通过以下文档了解公共解析器。

{% embed url="https://docs.ens.domains/contract-api-reference/publicresolver" %}

### 需要注意的是

.web3 用户名对**查询反向解析记录**做了改进，为了解决一些容错性的问题；

在查询反向解析记录时，建议您使用 <mark style="color:red;">`REVERSE_RECORDS`</mark> 工具合约的 <mark style="color:red;">`getNames`</mark> <mark style="color:red;"></mark><mark style="color:red;"></mark> 来查询用户的反向解析记录；

[REVERSE\_RECORDS](https://github.com/web3-identity/cns-contracts/blob/master/contracts/utils/ReverseRecords.sol) 工具合约有以下特性：

1. 支持批量查询反向解析记录；
2. 支持正反向解析一致性的校验；当正反向解析设置不一致，则反向解析会返回空；

详见：[Github CNS ReverseRecords.sol](https://github.com/web3-identity/cns-contracts/blob/master/contracts/utils/ReverseRecords.sol)

{% hint style="danger" %}
为了防止正反向解析设置不一致，给普通用户带来困扰，建议您使用<mark style="color:red;">`REVERSE_RECORDS`</mark>工具合约查询用户反向解析记录
{% endhint %}

****

**具体原因如下（可以不看）：**

在 ENS 的合约中，是允许正向解析与反向解析设置的不一致的（比如：abc.eth 正向解析是地址 A，但是地址 A 的反向解析可以设置成 qwe.eth），这是合约层的一个特性；

但是在 ENS App 端，这种行为是不被允许的，即正反向解析不能设置的不一致，这是为了防止普通用户错误的使用的一种处理方式。

在 ENS App 端：

1. 必须先设置正向解析（abc.eth -> address A），反向解析的选项列表中才出现 abc.eth 的选项。
2. 将 address A 的反向解析设置为 abc.eth 后，取消 abc.eth -> addressA 的正向解析，在 App 端的反向解析会失效（但合约仍然是生效的），依赖于 SDK 中的方法，并且需要在每个产品端解决。

我们通过提供<mark style="color:red;">`REVERSE_RECORDS`</mark>工具合约，来减少需要通过 SDK 做额外判断的工作；

* 当用户的正反向解析设置不一致时，<mark style="color:red;">`REVERSE_RECORDS`</mark>的反向解析记录会返回空

故，为了防止正反向解析设置不一致，给普通用户带来困扰，建议您使用<mark style="color:red;">`REVERSE_RECORDS`</mark>工具合约查询用户反向解析记录。
