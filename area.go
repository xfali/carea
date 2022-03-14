// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package carea

type AreaData struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Name       string  `json:"name"`
	Code       int     `json:"code"`
	ParentCode int     `json:"parentCode"`
	Level      int     `json:"level"`
}

type Area struct {
	AreaData
	Subareas []Area `json:"subareas"`
}

