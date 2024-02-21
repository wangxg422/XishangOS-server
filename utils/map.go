package utils

import "encoding/json"

func ObjsToMapList(s any) ([]map[string]any, error) {
	var newMap []map[string]any
	data, _ := json.Marshal(s)
	err := json.Unmarshal(data, &newMap)

	return newMap, err
}
