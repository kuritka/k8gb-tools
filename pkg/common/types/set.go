package types

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func NewSet() Set {
	return make(map[string]bool)
}

func (s Set) String() string {
	arr := make([]string, 0)
	for key := range s {
		arr = append(arr, key)
	}
	return fmt.Sprintf("{ %s }", strings.Join(arr, ","))
}
