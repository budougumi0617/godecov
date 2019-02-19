// Copyright (c) 2018.  Yoichiro Shimizu @budougumi0617

package godecov

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// TotalsArray is included many Codcov responses.
type TotalsArray struct {
	// C    int           `json:"C"`
	Files         int    `json:"f"` // [0] files count
	Lines         int    `json:"n"` // [1] lines count
	Hits          int    `json:"h"` // [2] hits count
	Partials      int    `json:"p"` // [3] partials count
	Missed        int    `json:"m"` // [4] missed count
	CoverageRatio string `json:"c"` // [5] coverage ratio
	Sessions      int    `json:"s"` // [9] sessions count
	Messages      int    `json:"M"` // messages count
	N             int    `json:"N"` // Unknown count
	Branches      int    `json:"b"` // branches count
	Methods       int    `json:"d"` // methods count
	// Diff []interface{} `json:"diff"`
}

// UnmarshalJSON creates TotalsArray from []interface{}.
func (ta *TotalsArray) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var arr []interface{}
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	// TODO Use NullInt
	f, ok := arr[0].(float64)
	if !ok {
		return errors.New("failed type cast f")
	}
	ta.Files = int(f)
	n, ok := arr[1].(float64)
	if !ok {
		return errors.New("failed type cast n")
	}
	ta.Lines = int(n)
	h, ok := arr[2].(float64)
	if !ok {
		return errors.New("failed type cast h")
	}
	ta.Hits = int(h)
	p, ok := arr[3].(float64)
	if !ok {
		return errors.New("failed type cast p")
	}
	ta.Partials = int(p)
	m, ok := arr[4].(float64)
	if !ok {
		return errors.New("failed type cast m")
	}
	ta.Missed = int(m)
	c, ok := arr[5].(string)
	if !ok {
		return errors.New("failed type cast c")
	}
	ta.CoverageRatio = c

	// 6-9, Unknown values

	s, ok := arr[9].(float64)
	if !ok {
		return errors.New("failed type cast s")
	}
	ta.Sessions = int(s)
	return nil
}

// MarshalJSON TODO Define MarshalJSON for TotalsArray
func (ta *TotalsArray) MarshalJSON() ([]byte, error) {

	return []byte(fmt.Sprintf("[%d, %d, %d, %d, %d, \"%s\", %d, %d, %d, %d, %d, %d]",
		ta.Files,
		ta.Lines,
		ta.Hits,
		ta.Partials,
		ta.Missed,
		ta.CoverageRatio,
		0,
		0,
		0,
		ta.Sessions,
		0,
		0,
	)), nil
}
