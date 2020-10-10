package model

import (
	"fmt"

	"github.com/kuritka/k8gb-tools/pkg/common/types"
)

//Stringr store validation attributes for string values
type Stringr struct {
	Values         []string
	Error          error
	Recommendation error
}

//Intr store validation attributes for int values
type Intr struct {
	Values []int64
	Error  error
}

func InitStringr() *Stringr {
	s := new(Stringr)
	s.Values = make([]string, 0)
	return s
}

func (s *Stringr) Append(value string) {
	s.Values = append(s.Values, value)
}

//ValuesAreUnique
func (s *Stringr) ValuesAreUnique() *Stringr {
	set := types.NewSet()
	for _, value := range s.Values {
		set[value] = true
	}
	if len(s.Values) != len(set) {
		s.Error = fmt.Errorf("expecting unique values %s", set)
	}
	return s
}

//ValuesAreNotEmpty
func (s *Stringr) ValuesAreNotEmpty() *Stringr {
	for _, value := range s.Values {
		if value == "" {
			s.Error = fmt.Errorf("field is empty \"\" ")
			break
		}
	}
	return s
}

//ValuesAreUnique
func (s *Stringr) ValuesAreEqual() *Stringr {
	set := types.NewSet()
	for _, value := range s.Values {
		set[value] = true
	}
	if len(s.Values) == 1 {
		s.Error = fmt.Errorf("expecting equal values %s", set)
	}
	return s
}

func (s *Stringr) ValuesAreIn(setOfValues ...string) *Stringr {
	for _, value := range s.Values {
		b := false
		for _, s := range setOfValues {
			if s == value {
				 b = true
				 break
			}
		}
		if b == false {
			s.Error = fmt.Errorf("%s is out of allowed definitions %s", value, setOfValues)
			return s
		}
	}
	return s
}
