package util

import (
	"fmt"
)

func String(value any) string {
	if value == nil {
		return ""
	}
	switch x := value.(type) {
	case []byte:
		return string(value.([]byte))
	case string:
		return x
	default:
		return fmt.Sprintf("%+v", value)
	}
}
