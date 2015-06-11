package dbu

import (
	"encoding/json"
	"github.com/shaalx/merbership/logu"
)

func ConvStruct(i interface{}, ret interface{}) bool {
	b := I2JsonBytes(i)
	err := json.Unmarshal(b, &ret)
	return !logu.CheckErr(err)
}
