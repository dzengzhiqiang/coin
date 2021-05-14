(window.webpackJsonp=window.webpackJsonp||[]).push([[64],{410:function(t,s,a){"use strict";a.r(s);var e=a(42),r=Object(e.a)({},(function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"websocket介绍"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#websocket介绍"}},[t._v("#")]),t._v(" Websocket介绍")]),t._v(" "),a("ul",[a("li",[t._v("本篇所列出的所有wss接口的baseurl为: "),a("strong",[t._v("wss://stream.binance.com:9443")])]),t._v(" "),a("li",[t._v("Streams有单一原始 stream 或组合 stream")]),t._v(" "),a("li",[t._v("单一原始 streams 格式为 "),a("strong",[t._v("/ws/<streamName>")])]),t._v(" "),a("li",[t._v("组合streams的URL格式为 "),a("strong",[t._v("/stream?streams=<streamName1>/<streamName2>/<streamName3>")])]),t._v(" "),a("li",[t._v("订阅组合streams时，事件payload会以这样的格式封装: "),a("strong",[t._v('{"stream":"<streamName>","data":<rawPayload>}')])]),t._v(" "),a("li",[t._v("stream名称中所有交易对均为 "),a("strong",[t._v("小写")])]),t._v(" "),a("li",[t._v("每个到 "),a("strong",[t._v("stream.binance.com")]),t._v(" 的链接有效期不超过24小时，请妥善处理断线重连。")]),t._v(" "),a("li",[t._v("每3分钟，服务端会发送ping帧，客户端应当在10分钟内回复pong帧，否则服务端会主动断开链接。允许客户端发送不成对的pong帧(即客户端可以以高于10分钟每次的频率发送pong帧保持链接)。")])]),t._v(" "),a("h2",{attrs:{id:"实时订阅-取消数据流"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#实时订阅-取消数据流"}},[t._v("#")]),t._v(" 实时订阅/取消数据流")]),t._v(" "),a("ul",[a("li",[t._v("以下数据可以通过websocket发送以实现订阅或取消订阅数据流。示例如下。")]),t._v(" "),a("li",[t._v("响应内容中的"),a("code",[t._v("id")]),t._v("是无符号整数，作为往来信息的唯一标识。")]),t._v(" "),a("li",[t._v("如果相应内容中的 "),a("code",[t._v("result")]),t._v(" 为 "),a("code",[t._v("null")]),t._v("，表示请求发送成功。")])]),t._v(" "),a("h3",{attrs:{id:"订阅一个信息流"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#订阅一个信息流"}},[t._v("#")]),t._v(" 订阅一个信息流")]),t._v(" "),a("blockquote",[a("p",[a("strong",[t._v("响应")])])]),t._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"result"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token null keyword"}},[t._v("null")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("1")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("ul",[a("li",[a("p",[a("strong",[t._v("请求")])]),t._v(" "),a("p",[t._v("{"),a("br"),t._v('\n"method": "SUBSCRIBE",'),a("br"),t._v('\n"params":'),a("br"),t._v("\n["),a("br"),t._v('\n"btcusdt@aggTrade",'),a("br"),t._v('\n"btcusdt@depth"'),a("br"),t._v("\n],"),a("br"),t._v('\n"id": 1'),a("br"),t._v("\n}")])])]),t._v(" "),a("h3",{attrs:{id:"取消订阅一个信息流"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#取消订阅一个信息流"}},[t._v("#")]),t._v(" 取消订阅一个信息流")]),t._v(" "),a("blockquote",[a("p",[a("strong",[t._v("响应")])])]),t._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"result"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token null keyword"}},[t._v("null")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("312")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("ul",[a("li",[a("p",[a("strong",[t._v("请求")])]),t._v(" "),a("p",[t._v("{"),a("br"),t._v('\n"method": "UNSUBSCRIBE",'),a("br"),t._v('\n"params":'),a("br"),t._v("\n["),a("br"),t._v('\n"btcusdt@depth"'),a("br"),t._v("\n],"),a("br"),t._v('\n"id": 312'),a("br"),t._v("\n}")])])]),t._v(" "),a("h3",{attrs:{id:"已订阅信息流"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#已订阅信息流"}},[t._v("#")]),t._v(" 已订阅信息流")]),t._v(" "),a("blockquote",[a("p",[a("strong",[t._v("响应")])])]),t._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"result"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),t._v("\n    "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"btcusdt@aggTrade"')]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("3")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("ul",[a("li",[a("p",[a("strong",[t._v("请求")])]),t._v(" "),a("p",[t._v("{"),a("br"),t._v('\n"method": "LIST_SUBSCRIPTIONS",'),a("br"),t._v('\n"id": 3'),a("br"),t._v("\n}")])])]),t._v(" "),a("h3",{attrs:{id:"设定属性"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#设定属性"}},[t._v("#")]),t._v(" 设定属性")]),t._v(" "),a("p",[t._v("当前，唯一可以设置的属性是设置是否启用"),a("code",[t._v("combined")]),t._v('("组合")信息流。'),a("br"),t._v("\n当使用"),a("code",[t._v("/ws/")]),t._v('("原始信息流")进行连接时，combined属性设置为'),a("code",[t._v("false")]),t._v("，而使用 "),a("code",[t._v("/stream/")]),t._v("进行连接时则将属性设置为"),a("code",[t._v("true")]),t._v("。")]),t._v(" "),a("blockquote",[a("p",[a("strong",[t._v("响应")])])]),t._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"result"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token null keyword"}},[t._v("null")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("5")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("ul",[a("li",[a("p",[a("strong",[t._v("请求")])]),t._v(" "),a("p",[t._v("{"),a("br"),t._v('\n"method": "SET_PROPERTY",'),a("br"),t._v('\n"params":'),a("br"),t._v("\n["),a("br"),t._v('\n"combined",'),a("br"),t._v("\ntrue"),a("br"),t._v("\n],"),a("br"),t._v('\n"id": 5'),a("br"),t._v("\n}")])])]),t._v(" "),a("h3",{attrs:{id:"检索属性"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#检索属性"}},[t._v("#")]),t._v(" 检索属性")]),t._v(" "),a("blockquote",[a("p",[a("strong",[t._v("响应")])])]),t._v(" "),a("div",{staticClass:"language-json extra-class"},[a("pre",{pre:!0,attrs:{class:"language-json"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"result"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("true")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// Indicates that combined is set to true.")]),t._v("\n  "),a("span",{pre:!0,attrs:{class:"token property"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("2")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("ul",[a("li",[a("p",[a("strong",[t._v("请求")])]),t._v(" "),a("p",[t._v("{"),a("br"),t._v('\n"method": "GET_PROPERTY",'),a("br"),t._v('\n"params":'),a("br"),t._v("\n["),a("br"),t._v('\n"combined"'),a("br"),t._v("\n],"),a("br"),t._v('\n"id": 2'),a("br"),t._v("\n}")])])]),t._v(" "),a("p",[t._v("###错误信息")]),t._v(" "),a("table",[a("thead",[a("tr",[a("th",[t._v("错误信息")]),t._v(" "),a("th",[t._v("描述")])])]),t._v(" "),a("tbody",[a("tr",[a("td",[t._v('{"code": 0, "msg": "Unknown property"}')]),t._v(" "),a("td",[a("code",[t._v("SET_PROPERTY")]),t._v(" 或 "),a("code",[t._v("GET_PROPERTY")]),t._v("中应用的参数无效")])]),t._v(" "),a("tr",[a("td",[t._v('{"code": 1, "msg": "Invalid value type: expected Boolean", "id": \'%s\'}')]),t._v(" "),a("td",[t._v("仅接受"),a("code",[t._v("true")]),t._v("或"),a("code",[t._v("false")])])]),t._v(" "),a("tr",[a("td",[t._v('{"code": 2, "msg": "Invalid request: property name must be a string"}')]),t._v(" "),a("td",[t._v("提供的属性名无效")])]),t._v(" "),a("tr",[a("td",[t._v('{"code": 2, "msg": "Invalid request: request ID must be an unsigned integer"}')]),t._v(" "),a("td",[t._v("参数"),a("code",[t._v("id")]),t._v("未提供或"),a("code",[t._v("id")]),t._v("值是无效类型")])]),t._v(" "),a("tr",[a("td",[t._v('{"code": 2, "msg": "Invalid request: unknown variant %s, expected one of '),a("code",[t._v("SUBSCRIBE")]),t._v(", "),a("code",[t._v("UNSUBSCRIBE")]),t._v(", "),a("code",[t._v("LIST_SUBSCRIPTIONS")]),t._v(", "),a("code",[t._v("SET_PROPERTY")]),t._v(", "),a("code",[t._v("GET_PROPERTY")]),t._v(' at line 1 column 28"}')]),t._v(" "),a("td",[t._v("错字提醒，或提供的值不是预期类型")])]),t._v(" "),a("tr",[a("td",[t._v('{"code": 2, "msg": "Invalid request: too many parameters"}')]),t._v(" "),a("td",[t._v("数据中提供了不必要参数")])]),t._v(" "),a("tr",[a("td",[t._v('{"code": 2, "msg": "Invalid request: property name must be a string"}')]),t._v(" "),a("td",[t._v("未提供属性名")])]),t._v(" "),a("tr",[a("td",[t._v('{"code": 2, "msg": "Invalid request: missing field '),a("code",[t._v("method")]),t._v(' at line 1 column 73"}')]),t._v(" "),a("td",[t._v("数据未提供"),a("code",[t._v("method")])])]),t._v(" "),a("tr",[a("td",[t._v('{"code":3,"msg":"Invalid JSON: expected value at line %s column %s"}')]),t._v(" "),a("td",[t._v("JSON 语法有误.")])])])])])}),[],!1,null,null,null);s.default=r.exports}}]);