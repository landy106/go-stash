package filter

import (
	"maps"

	jsoniter "github.com/json-iterator/go"
)

func TransferFilter(field, target string) FilterFunc {
	return func(m map[string]any) map[string]any {
		val, ok := m[field]
		if !ok {
			return m
		}

		s, ok := val.(string)
		if !ok {
			return m
		}

		var nm map[string]any
		if err := jsoniter.Unmarshal([]byte(s), &nm); err != nil {
			return m
		}

		delete(m, field)
		if len(target) > 0 {
			m[target] = nm
		} else {
			maps.Copy(m, nm)
		}

		return m
	}
}
