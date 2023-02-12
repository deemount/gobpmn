package pool

import "strings"

var (
	structPool = "pool"
)

func IsPool(field string) bool {
	return strings.ToLower(field) == structPool
}
