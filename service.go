// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package carea

import (
	"encoding/json"
	"fmt"
	"github.com/xfali/carea/static"
)

const TopLevel = 1

type AreaService interface {
	Data() ([]AreaData, error)

	AreaLevel() int
	Areas() ([]Area, error)
	AreaByLevel(level int) ([]Area, error)
	AreaByName(name string) ([]Area, error)
	AreaByCode(code int) (Area, error)
}

type defaultAreaService struct {
	areas [][]AreaData
}

func NewAreaService() *defaultAreaService {
	ret := &defaultAreaService{}
	err := ret.parse()
	if err != nil {
		return nil
	}
	return ret
}

func (s *defaultAreaService) Data() ([]AreaData, error) {
	var ret []AreaData
	err := json.Unmarshal([]byte(static.Areas), &ret)
	return ret, err
}

func (s *defaultAreaService) AreaLevel() int {
	return len(s.areas)
}

func (s *defaultAreaService) Areas() ([]Area, error) {
	return s.AreaByLevel(TopLevel)
}

func (s *defaultAreaService) AreaByLevel(level int) ([]Area, error) {
	err := s.checkLevel(level)
	if err != nil {
		return nil, err
	}
	ret := make([]Area, len(s.areas[level-1]))
	for i := range s.areas[level-1] {
		sub := Area{
			AreaData: s.areas[level-1][i],
		}
		_ = s.getChildren(&sub)
		ret[i] = sub
	}
	return ret, nil
}

func (s *defaultAreaService) AreaByName(name string) ([]Area, error) {
	var ret []Area
	for _, lv := range s.areas {
		for _, ad := range lv {
			if ad.Name == name {
				sub := Area{
					AreaData: ad,
				}
				_ = s.getChildren(&sub)
				ret = append(ret, sub)
			}
		}
	}
	return ret, nil
}

func (s *defaultAreaService) AreaByCode(code int) (Area, error) {
	for _, lv := range s.areas {
		for _, ad := range lv {
			if ad.Code == code {
				sub := Area{
					AreaData: ad,
				}
				_ = s.getChildren(&sub)
				return sub, nil
			}
		}
	}
	return Area{}, fmt.Errorf("Area with code %d not found. ", code)
}

func (s *defaultAreaService) checkLevel(level int) error {
	if level < TopLevel || level > len(s.areas) {
		return fmt.Errorf("Level %d out of range. ", level)
	}
	return nil
}

func (s *defaultAreaService) parse() error {
	d, err := s.Data()
	if err != nil {
		return err
	}
	s.areas = make([][]AreaData, 0, 3)
	for _, area := range d {
		if s.AreaLevel() < area.Level {
			lv := make([]AreaData, 0, 32)
			s.areas = append(s.areas, lv)
		}
		s.areas[area.Level-1] = append(s.areas[area.Level-1], area)
	}
	return nil
}

func (s *defaultAreaService) getChildren(area *Area) error {
	err := s.checkLevel(area.Level)
	if err != nil {
		return err
	}
	for _, a := range s.areas[area.Level] {
		if a.ParentCode == area.Code {
			sub := Area{
				AreaData: a,
			}
			area.Subareas = append(area.Subareas, sub)
			_ = s.getChildren(&sub)
		}
	}
	return nil
}
