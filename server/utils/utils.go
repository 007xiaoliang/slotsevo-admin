package utils

import (
	"encoding/json"
	"errors"
	"strconv"
)

func ParseMap(aMap map[string]interface{}) (map[string]string, error) {
	var retMap = map[string]string{}
	var err error
	for key, val := range aMap {
		switch val.(type) {
		case map[string]interface{}:
			retMap[key] = MapToJson(val.(map[string]interface{}))
		case string:
			retMap[key] = val.(string)
		case float64:
			retMap[key] = strconv.Itoa(int(val.(float64)))
		case int:
			retMap[key] = strconv.Itoa(val.(int))
		default:
			err = errors.New("unsupported type")
		}
	}
	return retMap, err
}

func MapToJson(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
