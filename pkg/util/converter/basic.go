package converter

import (
	"encoding/json"
	"strconv"
	"strings"
)

func ToInt64(v interface{}) int64 {
	switch v := v.(type) {
	case string:
		str := strings.TrimSpace(v)
		result, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return int64(result)
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case []byte:
		result, err := strconv.Atoi(string(v))
		if err != nil {
			return 0
		}
		return int64(result)
	default:
		return 0
	}
}

func ToUInt64(v interface{}) uint64 {
	switch v := v.(type) {
	case string:
		str := strings.TrimSpace(v)
		result, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return uint64(result)
	case int:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint64:
		return v
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case []byte:
		result, err := strconv.Atoi(string(v))
		if err != nil {
			return 0
		}
		return uint64(result)
	default:
		return 0
	}
}

func ToInt(v interface{}) int {
	switch v := v.(type) {
	case string:
		str := strings.TrimSpace(v)
		result, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return result
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case []byte:
		result, err := strconv.Atoi(string(v))
		if err != nil {
			return 0
		}
		return result
	default:
		return 0
	}
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch v := v.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatInt(int64(v), 10)
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case []uint8:
		return string(v)
	default:
		resultJSON, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		return string(resultJSON)
	}
}
