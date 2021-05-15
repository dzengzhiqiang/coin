package types

import "fmt"

type BizCode int

const (
	BizCode_OK            BizCode = 0
	BizCode_ServerError   BizCode = 1
	BizCode_UnmashalError BizCode = 2
)

var bizCodes = map[BizCode]string{
	BizCode_OK:            "BizCode_OK",
	BizCode_ServerError:   "BizCode_ServerError",
	BizCode_UnmashalError: "BizCode_UnmashalError",
}

func (c BizCode) GoString() string {
	return c.String()
}

func (c BizCode) String() string {
	if strCode, ok := bizCodes[c]; ok {
		return strCode
	}
	return fmt.Sprintf("BizCode_Unknown<%d>", c)
}
