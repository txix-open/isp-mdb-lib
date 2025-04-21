package pdp_diff

import (
	"github.com/pkg/errors"
)

func isArrayOfObjects(v any) bool {
	arr, ok := v.([]any)
	if !ok {
		return false
	}
	for _, el := range arr {
		_, ok = el.(map[string]any)
		if !ok {
			return false
		}
	}

	return true
}

func isObject(v any) bool {
	_, ok := v.(map[string]any)
	return ok
}

func isFlatObject(obj map[string]any) bool {
	for _, v := range obj {
		if isObject(v) || isArrayOfObjects(v) {
			return false
		}
	}
	return true
}

func isString(v any) bool {
	_, ok := v.(string)
	return ok
}

func castToString(v any) (string, error) {
	s, ok := v.(string)
	if !ok {
		return "", errors.New("unable to cast to string")
	}
	return s, nil
}

func toArrayMap(v any) ([]map[string]any, error) {
	arr, ok := v.([]any)
	if !ok {
		return nil, errors.New("object is not an array")
	}
	res := make([]map[string]any, 0)
	for _, el := range arr {
		m, ok := el.(map[string]any)
		if !ok {
			return nil, errors.New("unable to cast array object to map")
		}
		res = append(res, m)
	}

	return res, nil
}

func toItemIdMap(arr []map[string]any) (map[string]map[string]any, []string, error) {
	dataByItemId := make(map[string]map[string]any)
	itemIds := make([]string, 0, len(arr))
	for i, item := range arr {
		val, ok := item[itemIdPath]
		if !ok {
			return nil, nil, errors.Errorf("item_id not found index %d", i)
		}
		itemId, ok := val.(string)
		if !ok {
			return nil, nil, errors.Errorf("item_id is not a string index %d", i)
		}
		dataByItemId[itemId] = item
		itemIds = append(itemIds, itemId)
	}

	return dataByItemId, itemIds, nil
}
