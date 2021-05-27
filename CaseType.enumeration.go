////////////////////////////////////////////////////////////////////////
// Generated by bounded/enumeration
////////////////////////////////////////////////////////////////////////

package caser

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CaseType string

const (
	CaseType_Camel      CaseType = "camel"
	CaseType_Kebab      CaseType = "kebab"
	CaseType_KebabUpper CaseType = "kebabUpper"
	CaseType_Pascal     CaseType = "pascal"
	CaseType_Phrase     CaseType = "phrase"
	CaseType_Snake      CaseType = "snake"
	CaseType_SnakeUpper CaseType = "snakeUpper"
	CaseType_Unknown    CaseType = "unknown"
)

var (
	CaseTypes = []CaseType{
		CaseType_Camel,
		CaseType_Kebab,
		CaseType_KebabUpper,
		CaseType_Pascal,
		CaseType_Phrase,
		CaseType_Snake,
		CaseType_SnakeUpper,
		CaseType_Unknown,
	}
)

func IsCaseType(v string) bool {
	var f bool

	for _, e := range CaseTypes {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func CaseTypeParse(v string) (CaseType, error) {
	var o CaseType
	var f bool
	n := strings.ToLower(v)

	for _, e := range CaseTypes {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, ErrCaseTypeNotFound(v)
	}

	return o, nil
}

func ErrCaseTypeNotFound(v string) error {
	var ss []string

	for _, e := range CaseTypes {
		ss = append(ss, string(e))
	}

	return fmt.Errorf(
		"invalid enumeration type '%v', must be one of %v",
		v, strings.Join(ss, ","),
	)
}

func (t CaseType) String() string {
	return string(t)
}

func (t CaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *CaseType) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := CaseTypeParse(s)

	if err != nil {
		return err
	}

	t = &e

	return nil
}
