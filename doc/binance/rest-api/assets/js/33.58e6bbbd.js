(window.webpackJsonp=window.webpackJsonp||[]).push([[33],{379:function(s,t,e){"use strict";e.r(t);var a=e(42),n=Object(a.a)({},(function(){var s=this,t=s.$createElement,e=s._self._c||t;return e("ContentSlotsDistributor",{attrs:{"slot-key":s.$parent.slotKey}},[e("h1",{attrs:{id:"general-wss-information"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#general-wss-information"}},[s._v("#")]),s._v(" General WSS Information")]),s._v(" "),e("ul",[e("li",[s._v("The base endpoint is: "),e("strong",[s._v("wss://stream.binance.com:9443")])]),s._v(" "),e("li",[s._v("Streams can be accessed either in a single raw stream or in a combined stream")]),s._v(" "),e("li",[s._v("Raw streams are accessed at "),e("strong",[s._v("/ws/<streamName>")])]),s._v(" "),e("li",[s._v("Combined streams are accessed at "),e("strong",[s._v("/stream?streams=<streamName1>/<streamName2>/<streamName3>")])]),s._v(" "),e("li",[s._v("Combined stream events are wrapped as follows: "),e("strong",[s._v('{"stream":"<streamName>","data":<rawPayload>}')])]),s._v(" "),e("li",[s._v("All symbols for streams are "),e("strong",[s._v("lowercase")])]),s._v(" "),e("li",[s._v("A single connection to "),e("strong",[s._v("stream.binance.com")]),s._v(" is only valid for 24 hours; expect to be disconnected at the 24 hour mark")]),s._v(" "),e("li",[s._v("The websocket server will send a "),e("code",[s._v("ping frame")]),s._v(" every 3 minutes. If the websocket server does not receive a "),e("code",[s._v("pong frame")]),s._v(" back from the connection within a 10 minute period, the connection will be disconnected. Unsolicited "),e("code",[s._v("pong frames")]),s._v(" are allowed.")])]),s._v(" "),e("h2",{attrs:{id:"live-subscribing-unsubscribing-to-streams"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#live-subscribing-unsubscribing-to-streams"}},[s._v("#")]),s._v(" Live Subscribing/Unsubscribing to streams")]),s._v(" "),e("ul",[e("li",[s._v("The following data can be sent through the websocket instance in order to subscribe/unsubscribe from streams. Examples can be seen below.")]),s._v(" "),e("li",[s._v("The "),e("code",[s._v("id")]),s._v(" used in the JSON payloads is an unsigned INT used as an identifier to uniquely identify the messages going back and forth.")]),s._v(" "),e("li",[s._v("In the response, if the "),e("code",[s._v("result")]),s._v(" received is "),e("code",[s._v("null")]),s._v(" this means the request sent was a success for non-query requests (e.g. Subscribing/Unsubscribing).")])]),s._v(" "),e("h3",{attrs:{id:"subscribe-to-a-stream"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#subscribe-to-a-stream"}},[s._v("#")]),s._v(" Subscribe to a stream")]),s._v(" "),e("ul",[e("li",[e("p",[s._v("Request")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"method"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"SUBSCRIBE"')]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"params"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("[")]),s._v("\n    "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"btcusdt@aggTrade"')]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n    "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"btcusdt@depth"')]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("]")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("1")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])]),s._v(" "),e("li",[e("p",[s._v("Response")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"result"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token null keyword"}},[s._v("null")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("1")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])])]),s._v(" "),e("h3",{attrs:{id:"unsubscribe-to-a-stream"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#unsubscribe-to-a-stream"}},[s._v("#")]),s._v(" Unsubscribe to a stream")]),s._v(" "),e("ul",[e("li",[e("p",[s._v("Request")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"method"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"UNSUBSCRIBE"')]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"params"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("[")]),s._v("\n    "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"btcusdt@depth"')]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("]")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("312")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])]),s._v(" "),e("li",[e("p",[s._v("Response")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"result"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token null keyword"}},[s._v("null")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("312")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])])]),s._v(" "),e("h3",{attrs:{id:"listing-subscriptions"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#listing-subscriptions"}},[s._v("#")]),s._v(" Listing Subscriptions")]),s._v(" "),e("ul",[e("li",[e("p",[s._v("Request")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"method"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"LIST_SUBSCRIPTIONS"')]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("3")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])]),s._v(" "),e("li",[e("p",[s._v("Response")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"result"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("[")]),s._v("\n    "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"btcusdt@aggTrade"')]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("]")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("3")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])])]),s._v(" "),e("h3",{attrs:{id:"setting-properties"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#setting-properties"}},[s._v("#")]),s._v(" Setting Properties")]),s._v(" "),e("p",[s._v("Currently, the only property that can be set is whether "),e("code",[s._v("combined")]),s._v(" stream payloads are enabled or not.\nThe combined property is set to "),e("code",[s._v("false")]),s._v(" when connecting using "),e("code",[s._v("/ws/")]),s._v(' ("raw streams") and '),e("code",[s._v("true")]),s._v(" when connecting using "),e("code",[s._v("/stream/")]),s._v(".")]),s._v(" "),e("ul",[e("li",[e("p",[s._v("Request")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"method"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"SET_PROPERTY"')]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"params"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("[")]),s._v("\n    "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"combined"')]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n    "),e("span",{pre:!0,attrs:{class:"token boolean"}},[s._v("true")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("]")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("5")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])]),s._v(" "),e("li",[e("p",[s._v("Response")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"result"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token null keyword"}},[s._v("null")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("5")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])])]),s._v(" "),e("h3",{attrs:{id:"retrieving-properties"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#retrieving-properties"}},[s._v("#")]),s._v(" Retrieving Properties")]),s._v(" "),e("ul",[e("li",[e("p",[s._v("Request")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"method"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"GET_PROPERTY"')]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"params"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("[")]),s._v("\n    "),e("span",{pre:!0,attrs:{class:"token string"}},[s._v('"combined"')]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("]")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("2")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])]),s._v(" "),e("li",[e("p",[s._v("Response")]),s._v(" "),e("div",{staticClass:"language-json extra-class"},[e("pre",{pre:!0,attrs:{class:"language-json"}},[e("code",[e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("{")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"result"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token boolean"}},[s._v("true")]),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(",")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token comment"}},[s._v("// Indicates that combined is set to true.")]),s._v("\n  "),e("span",{pre:!0,attrs:{class:"token property"}},[s._v('"id"')]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v(":")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("2")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("}")]),s._v("\n")])])])])]),s._v(" "),e("h3",{attrs:{id:"error-messages"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#error-messages"}},[s._v("#")]),s._v(" Error Messages")]),s._v(" "),e("table",[e("thead",[e("tr",[e("th",[s._v("Error Message")]),s._v(" "),e("th",[s._v("Description")])])]),s._v(" "),e("tbody",[e("tr",[e("td",[s._v('{"code": 0, "msg": "Unknown property","id": %s}')]),s._v(" "),e("td",[s._v("Parameter used in the "),e("code",[s._v("SET_PROPERTY")]),s._v(" or "),e("code",[s._v("GET_PROPERTY")]),s._v(" was invalid")])]),s._v(" "),e("tr",[e("td",[s._v('{"code": 1, "msg": "Invalid value type: expected Boolean"}')]),s._v(" "),e("td",[s._v("Value should only be "),e("code",[s._v("true")]),s._v(" or "),e("code",[s._v("false")])])]),s._v(" "),e("tr",[e("td",[s._v('{"code": 2, "msg": "Invalid request: property name must be a string"}')]),s._v(" "),e("td",[s._v("Property name provided was invalid")])]),s._v(" "),e("tr",[e("td",[s._v('{"code": 2, "msg": "Invalid request: request ID must be an unsigned integer"}')]),s._v(" "),e("td",[s._v("Parameter "),e("code",[s._v("id")]),s._v(" had to be provided or the value provided in the "),e("code",[s._v("id")]),s._v(" parameter is an unsupported type")])]),s._v(" "),e("tr",[e("td",[s._v('{"code": 2, "msg": "Invalid request: unknown variant %s, expected one of '),e("code",[s._v("SUBSCRIBE")]),s._v(", "),e("code",[s._v("UNSUBSCRIBE")]),s._v(", "),e("code",[s._v("LIST_SUBSCRIPTIONS")]),s._v(", "),e("code",[s._v("SET_PROPERTY")]),s._v(", "),e("code",[s._v("GET_PROPERTY")]),s._v(' at line 1 column 28"}')]),s._v(" "),e("td",[s._v("Possible typo in the provided method or provided method was neither of the expected values")])]),s._v(" "),e("tr",[e("td",[s._v('{"code": 2, "msg": "Invalid request: too many parameters"}')]),s._v(" "),e("td",[s._v("Unnecessary parameters provided in the data")])]),s._v(" "),e("tr",[e("td",[s._v('{"code": 2, "msg": "Invalid request: property name must be a string"}')]),s._v(" "),e("td",[s._v("Property name was not provided")])]),s._v(" "),e("tr",[e("td",[s._v('{"code": 2, "msg": "Invalid request: missing field '),e("code",[s._v("method")]),s._v(' at line 1 column 73"}')]),s._v(" "),e("td",[e("code",[s._v("method")]),s._v(" was not provided in the data")])]),s._v(" "),e("tr",[e("td",[s._v('{"code":3,"msg":"Invalid JSON: expected value at line %s column %s"}')]),s._v(" "),e("td",[s._v("JSON data sent has incorrect syntax.")])])])])])}),[],!1,null,null,null);t.default=n.exports}}]);