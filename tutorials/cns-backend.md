# Cns-Backend

## Postman
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/22322698-cd32951a-a24f-4fd5-a9fb-2e26f057532c?action=collection%2Ffork&collection-url=entityId%3D22322698-cd32951a-a24f-4fd5-a9fb-2e26f057532c%26entityType%3Dcollection%26workspaceId%3D0df0c5b3-6c0a-47ee-ab26-8ba0139261e4)

## Swagger
<a href="http://101.42.88.184/swagger/"><img src="https://static1.smartbear.co/swagger/media/assets/images/swagger_logo.svg" alt="drawing" style="width:200px;" /></a>

## 第三方接入交互图

![](https://note.youdao.com/yws/api/personal/file/WEB3b48793e4e39f246359a449cbdecc7eb?method=download&shareKey=7747d123352bb7fbd4f1273639661bac)

## API相关枚举类型：
| 枚举           | 意义           | 值                                                       |
|----------------|----------------|----------------------------------------------------------|
| trade_provider | 法币支付运营商 | （wechat, alipay, ...）                                  |
| trade_type     | 法币支付类型   | native/h5                                                |
| order_state    | 订单状态       | init, made, success; null 表示全部                       |
| trade_state    | 支付状态       | SUCCESS REFUND NOTPAY CLOSED REVOKED USERPAYING PAYERROR |
| refund_state   | 支付退款状态   | NIL SUCCESS CLOSED PROCESSING ABNORMAL                   |
| tx_state       | web3交易状态   | INIT SEND_FAILED EXECUTE_FAILED EXECUTED_SUCCESS         |

### tx_state
- INIT 未开始
- SEND_FAILED 发送失败，表示交易未发送，通常是由于Estimate失败，tx_error表示错误详情
- EXECUTE_FAILED 执行失败，表示交易发送成功但交易执行失败（概率非常低，这种的失败原因是由于Estimate时的合约状态与交易执行时的合约状态不一致）
- EXECUTED_SUCCESS 执行成功，表示交易发送成功而且交易执行成功
## 第三方接入
### 注册域名
1. 调用合约Web3RegisterController的makeCommit方法生成commit hash
2. 调用合约Web3RegisterController的commit方法提交commithash
3. 调用cns-backend的api [POST /v0/registers](http://101.42.88.184:8090/swagger_ui_dist/#/Registers/MakeRegisterOrder) 注册域名
4. 调用cns-backend的api [GET /v0/registers/:commit_hash](http://101.42.88.184:8090/swagger_ui_dist/#/Registers/Register) 查询注册状态

### 域名续费
1. 调用cns-backend的api [POST /v0/renews](http://101.42.88.184:8090/swagger_ui_dist/#/Renews/Renew) 续费
2. 调用cns-backend的api [GET /v0/renews/:id](http://101.42.88.184:8090/swagger_ui_dist/#/Renews/GetRenew) 查询续费状态

### 身份认证
"注册域名"及"域名续费" API 都需要身份认证，使用ApiKey方式，ApiKey 由官方人工发放

使用方法：
- 在请求头中增加 `X-Api-Key`

## 其它
### 交易相关信息
交易信息相关字段有
- [tx_state 交易状态](#tx_state)
- tx_hash 交易哈希
- tx_error 交易错误信息

> 当交易成功时 tx_hash 才有值，交易失败时 tx_error 才有值
### CNS合约
- [cns合约](https://github.com/web3-identity/cns-contracts)
- [测试网合约部署地址](https://github.com/web3-identity/cns-contracts/tree/master/docs#conflux-core-testnet)