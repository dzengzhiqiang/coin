# API地址
```text
正式网 https://api.binance.com/api	
测试网 https://testnet.binance.vision/api

正式网 wss://stream.binance.com:9443/ws	
测试网 wss://testnet.binance.vision/ws

正式网 wss://stream.binance.com:9443/stream
测试网 wss://testnet.binance.vision/stream
```

# 1. 现货接口

## 1.1 查询现货账户信息

```json
{
  "makerCommission": 15,
  "takerCommission": 15,
  "buyerCommission": 0,
  "sellerCommission": 0,
  "canTrade": true,
  "canWithdraw": true,
  "canDeposit": true,
  "updateTime": 123456789,
  "accountType": "SPOT",
  "balances": [
    {
      "asset": "BTC",
      "free": "4723846.89208129",
      "locked": "0.00000000"
    },
    {
      "asset": "LTC",
      "free": "4763368.68006011",
      "locked": "0.00000000"
    }
  ],
  "permissions": [
    "SPOT"
  ]
}
```