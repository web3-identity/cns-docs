---
description: 若需在产品中支持域名售卖和注册，请参考此文档；
---

# CNS-Open-Service

## Postman

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/22322698-cd32951a-a24f-4fd5-a9fb-2e26f057532c?action=collection%2Ffork\&collection-url=entityId%3D22322698-cd32951a-a24f-4fd5-a9fb-2e26f057532c%26entityType%3Dcollection%26workspaceId%3D0df0c5b3-6c0a-47ee-ab26-8ba0139261e4)

## Swagger

[![drawing](https://static1.smartbear.co/swagger/media/assets/images/swagger\_logo.svg)](http://101.42.88.184/swagger/cns-backend/)

## 第三方接入交互图

![](https://note.youdao.com/yws/api/personal/file/WEB3b48793e4e39f246359a449cbdecc7eb?method=download\&shareKey=7747d123352bb7fbd4f1273639661bac)

## API相关枚举类型：

| 枚举              | 意义       | 值                                                        |
| --------------- | -------- | -------------------------------------------------------- |
| trade\_provider | 法币支付运营商  | （wechat, alipay, ...）                                    |
| trade\_type     | 法币支付类型   | native/h5                                                |
| order\_state    | 订单状态     | init, made, success; null 表示全部                           |
| trade\_state    | 支付状态     | SUCCESS REFUND NOTPAY CLOSED REVOKED USERPAYING PAYERROR |
| refund\_state   | 支付退款状态   | NIL SUCCESS CLOSED PROCESSING ABNORMAL                   |
| tx\_state       | web3交易状态 | INIT SEND\_FAILED EXECUTE\_FAILED EXECUTED\_SUCCESS      |

### tx\_state

* INIT 未开始
* SEND\_FAILED 发送失败，表示交易未发送，通常是由于Estimate失败，tx\_error表示错误详情
* EXECUTE\_FAILED 执行失败，表示交易发送成功但交易执行失败（概率非常低，这种的失败原因是由于Estimate时的合约状态与交易执行时的合约状态不一致）
* EXECUTED\_SUCCESS 执行成功，表示交易发送成功而且交易执行成功

## 第三方接入

### 注册域名

{% hint style="danger" %}
务必在调用 registers 前，对用户所输入的内容进行**敏感词**过滤；

域名注册目前只支持英文字符和 Emoji；
{% endhint %}

1. 调用合约Web3RegisterController的`rentPriceInFiat`方法计算价格，例如要计算 `conflux.web3` 域名一年期的价格，`rentPriceInFiat("conflux", 3600*24*365)`; 得到的结果有两个字段 `base` 和 `premium`; 实际价格为 `(base + premium)/1000000`, 单位为“分”
2. 在注册域名时需要同时设置正向解析，方法为将commit.data设置合约方法`PublicResolver.setAddr(bytes32 node, uint coinType, bytes memory a)`的ABI编码。 参数`node`为域名的`node`，`coinType`为`conflux`的`CoinType`值`503`, `a`为正向解析到的地址
   1. node 计算使用库`@ensdomains/eth-ens-namehash`的`hash`方法。
   2. ABI计算使用[iface.encodeFunctionData](https://docs.ethers.org/v5/api/utils/abi/interface/#Interface--encoding)
3.  调用合约Web3RegisterController的`makeCommit`方法生成commit hash

    > commit.data设置为第2步生成的值
4. 调用合约Web3RegisterController的`commit`方法提交commit hash
5. 用户支付
6. 调用cns-backend的api [POST /v0/registers](http://101.42.88.184:8090/swagger\_ui\_dist/#/Registers/MakeRegisterOrder) 注册域名
7. 调用cns-backend的api [GET /v0/registers/:commit\_hash](http://101.42.88.184:8090/swagger\_ui\_dist/#/Registers/Register) 查询注册状态

### 域名续费

1. 域名续费价格计算方式同注册域名相同，如 `conflux.web3` 域名续费一年期的价格为`rentPriceInFiat("conflux", 3600*24*365)`
2. 用户支付
3. 调用cns-backend的api [POST /v0/renews](http://101.42.88.184:8090/swagger\_ui\_dist/#/Renews/Renew) 续费
4. 调用cns-backend的api [GET /v0/renews/:id](http://101.42.88.184:8090/swagger\_ui\_dist/#/Renews/GetRenew) 查询续费状态

### 身份认证

"注册域名"及"域名续费" API 都需要身份认证，使用ApiKey方式，ApiKey 由官方人工发放

使用方法：

* 在请求头中增加 `X-Api-Key`

## 其它

### 交易相关信息

交易信息相关字段有

* [tx\_state 交易状态](cns-backend.md#tx\_state)
* tx\_hash 交易哈希
* tx\_error 交易错误信息

> 当交易成功时 tx\_hash 才有值，交易失败时 tx\_error 才有值
