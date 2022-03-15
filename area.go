// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package carea

import (
	"encoding/json"
	"strconv"
)

type AreaCode string
type AreaLevel string

type AreaData struct {
	Latitude   string    `json:"latitude"`
	Longitude  string    `json:"longitude"`
	Name       string    `json:"name"`
	Code       AreaCode  `json:"code"`
	ParentCode AreaCode  `json:"parentCode"`
	Level      AreaLevel `json:"level"`
}

type Area struct {
	AreaData
	Subareas []Area `json:"subareas"`
}

func (a Area) String() string {
	d, _ := json.MarshalIndent(a, "", "\t")
	return string(d)
}

func (lv AreaLevel) Int() int {
	ret, _ := strconv.Atoi(string(lv))
	return int(ret)
}

func String2AreaCode(code string) AreaCode {
	return AreaCode(code)
}

func Int2AreaLevel(level int) AreaLevel {
	return AreaLevel(strconv.Itoa(level))
}
