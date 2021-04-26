# 1. 钱包接口

## 1.1 系统状态 

GET /wapi/v3/systemStatus.html

status 0: 正常 1：系统维护

```json
{ 
    "status": 0,             
    "msg": "normal"           
}
```

## 1.2 获取账户币种信息

GET /sapi/v1/capital/config/getall (HMAC SHA256)

| 参数       | 类型 | 必填 | 描述 |
| ---------- | ---- | ---- | ---- |
| recvWindow | LONG | NO   |      |
| timestamp  | LONG | YES  |      |




```json
[
	{
		"coin": "BTC",
 		"depositAllEnable": true,
 		"free": "0.08074558",
 		"freeze": "0.00000000",
 		"ipoable": "0.00000000",
 		"ipoing": "0.00000000",
 		"isLegalMoney": false,
 		"locked": "0.00000000",
 		"name": "Bitcoin",
 		"networkList": [
	 		{
	 			"addressRegex": "^(bnb1)[0-9a-z]{38}$",
	   			"coin": "BTC",
	   			"depositDesc": "Wallet Maintenance, Deposit Suspended", // 仅在充值关闭时返回
	   			"depositEnable": false,
	   			"isDefault": false,        
	   			"memoRegex": "^[0-9A-Za-z\\-_]{1,120}$",
	   			"minConfirm": 1,  // 上账所需的最小确认数
	   			"name": "BEP2",
	   			"network": "BNB",            
	   			"resetAddressStatus": false,
	   			"specialTips": "Both a MEMO and an Address are required to successfully deposit your BEP2-BTCB tokens to Binance.",
	   			"unLockConfirm": 0,  // 解锁需要的确认数 
	   			"withdrawDesc": "Wallet Maintenance, Withdrawal Suspended", // 仅在提现关闭时返回
	   			"withdrawEnable": false,
	   			"withdrawFee": "0.00000220",
	   			"withdrawMin": "0.00000440"
	   		},
	  		{
	  			"addressRegex": "^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$|^(bc1)[0-9A-Za-z]{39,59}$",
	   			"coin": "BTC",
	   			"depositEnable": true,
	   			"insertTime": 1563532929000,
	   			"isDefault": true,
	   			"memoRegex": "",
	   			"minConfirm": 1,  // 上账所需的最小确认数
	   			"name": "BTC",
	   			"network": "BTC",
	   			"resetAddressStatus": false,
	   			"specialTips": "",
	   			"unLockConfirm": 0,  // 解锁需要的确认数 
	   			"updateTime": 1571014804000, 
	   			"withdrawEnable": true,
	   			"withdrawFee": "0.00050000",
	   			"withdrawIntegerMultiple": "0.00000001",
	   			"withdrawMin": "0.00100000"
	   		}
   		],
 		"storage": "0.00000000",
 		"trading": true,
 		"withdrawAllEnable": true,
 		"withdrawing": "0.00000000"
 	}
]
```