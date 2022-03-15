// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package carea

import (
	"encoding/json"
	"fmt"
	"github.com/xfali/carea/static"
	"io/ioutil"
)

const (
	TopLevel    = AreaLevel("1")
	TopLevelInt = 1
)

type AreaService interface {
	// 获得原始区域数据
	Data() ([]AreaData, error)

	// 获得区域层级
	AreaLevelNumber() int

	// 获得区域层级列表
	AreaLevels() []AreaLevel

	// 从顶级层级获得区域信息
	// withSub： 是否遍历子区域
	Areas(withSub bool) ([]Area, error)

	// 获得指定层级区域信息
	// level：指定区域层级
	// withSub： 是否遍历子区域
	AreaByLevel(level AreaLevel, withSub bool) ([]Area, error)

	// 获得指定区域名称的区域信息
	// name：指定区域层级
	// withSub： 是否遍历子区域
	AreaByName(name string, withSub bool) ([]Area, error)

	// 获得指定区域Code的区域信息
	// code：指定区域Code
	// withSub： 是否遍历子区域
	AreaByCode(code AreaCode, withSub bool) (Area, error)

	// 获得指定区域Code的子区域信息
	// code：指定区域Code
	// recursion： 是否遍历所有子区域
	SubareaByCode(code AreaCode, recursion bool) ([]Area, error)

	// 获得指定区域Code的父区域信息
	// code：指定区域Code
	// recursion： 是否遍历所有父区域
	ParentAreaByCode(code AreaCode, recursion bool) (Area, error)
}

type defaultAreaService struct {
	ds     DataSource
	areas  [][]AreaData
	levels []AreaLevel
}

type Opt func(s *defaultAreaService)

func NewAreaService(opts ...Opt) *defaultAreaService {
	ret := &defaultAreaService{
		ds: buildinDataSource,
	}
	for _, opt := range opts {
		opt(ret)
	}
	err := ret.parse()
	if err != nil {
		return nil
	}
	return ret
}

func NewAreaServiceFromFile(path string) *defaultAreaService {
	return NewAreaService(DefaultOpt.LoadFromFile(path))
}

func (s *defaultAreaService) Data() ([]AreaData, error) {
	return s.ds()
}

func (s *defaultAreaService) AreaLevelNumber() int {
	return len(s.areas)
}

func (s *defaultAreaService) AreaLevels() []AreaLevel {
	return s.levels
}

func (s *defaultAreaService) Areas(withSub bool) ([]Area, error) {
	return s.AreaByLevel(TopLevel, withSub)
}

func (s *defaultAreaService) AreaByLevel(areaLv AreaLevel, withSub bool) ([]Area, error) {
	level := areaLv.Int()
	err := s.checkLevel(level)
	if err != nil {
		return nil, err
	}
	ret := make([]Area, len(s.areas[level-1]))
	for i := range s.areas[level-1] {
		sub := Area{
			AreaData: s.areas[level-1][i],
		}
		if withSub {
			_ = s.getChildren(&sub, true)
		}
		ret[i] = sub
	}
	return ret, nil
}

func (s *defaultAreaService) AreaByName(name string, withSub bool) ([]Area, error) {
	var ret []Area
	for _, lv := range s.areas {
		for _, ad := range lv {
			if ad.Name == name {
				sub := Area{
					AreaData: ad,
				}
				if withSub {
					_ = s.getChildren(&sub, true)
				}
				ret = append(ret, sub)
			}
		}
	}
	return ret, nil
}

func (s *defaultAreaService) AreaByCode(code AreaCode, withSub bool) (Area, error) {
	for _, lv := range s.areas {
		for _, ad := range lv {
			if ad.Code == code {
				sub := Area{
					AreaData: ad,
				}
				if withSub {
					_ = s.getChildren(&sub, true)
				}
				return sub, nil
			}
		}
	}
	return Area{}, fmt.Errorf("Area with code %v not found. ", code)
}

func (s *defaultAreaService) checkLevel(level int) error {
	if level < TopLevelInt || level > len(s.areas) {
		return fmt.Errorf("Level %d out of range. ", level)
	}
	return nil
}

func (s *defaultAreaService) SubareaByCode(code AreaCode, recursion bool) ([]Area, error) {
	for _, lv := range s.areas {
		for _, ad := range lv {
			if ad.Code == code {
				sub := Area{
					AreaData: ad,
				}
				err := s.getChildren(&sub, recursion)
				return sub.Subareas, err
			}
		}
	}
	return nil, fmt.Errorf("Area with code %v not found. ", code)
}

func (s *defaultAreaService) ParentAreaByCode(code AreaCode, recursion bool) (Area, error) {
	for i := len(s.areas) - 1; i >= 0; i-- {
		lv := s.areas[i]
		for _, ad := range lv {
			if ad.Code == code {
				parent := &Area{
					AreaData: ad,
				}
				p, err := s.getParent(parent, recursion)
				if p == nil {
					return Area{}, err
				}
				return *p, err
			}
		}
	}
	return Area{}, fmt.Errorf("Area with code %v not found. ", code)
}

func (s *defaultAreaService) parse() error {
	d, err := s.Data()
	if err != nil {
		return err
	}
	s.areas = make([][]AreaData, 0, 3)
	for _, area := range d {
		lv := area.Level.Int()
		if s.AreaLevelNumber() < lv {
			lv := make([]AreaData, 0, 32)
			s.areas = append(s.areas, lv)
			s.levels = append(s.levels, area.Level)
		}
		s.areas[lv-1] = append(s.areas[lv-1], area)
	}
	return nil
}

func (s *defaultAreaService) getChildren(area *Area, recursion bool) error {
	lv := area.Level.Int()
	err := s.checkLevel(lv)
	if err != nil {
		return err
	}
	if lv == len(s.areas) {
		return nil
	}
	for _, a := range s.areas[lv] {
		if a.ParentCode == area.Code {
			sub := Area{
				AreaData: a,
			}
			if recursion {
				_ = s.getChildren(&sub, recursion)
			}
			area.Subareas = append(area.Subareas, sub)
		}
	}
	return nil
}

func (s *defaultAreaService) getParent(area *Area, recursion bool) (*Area, error) {
	lv := area.Level.Int()
	err := s.checkLevel(lv)
	if err != nil {
		return nil, err
	}
	if area.Level == TopLevel {
		return area, nil
	}
	lv -= 2
	parent := area
	for _, a := range s.areas[lv] {
		if a.Code == area.ParentCode {
			parent = &Area{
				AreaData: a,
				Subareas: []Area{*area},
			}
			if recursion {
				parent, _ = s.getParent(parent, recursion)
			}
			return parent, nil
		}
	}
	return parent, nil
}

type defaultOption struct{}

var DefaultOpt defaultOption

type DataSource func() ([]AreaData, error)

func buildinDataSource() ([]AreaData, error) {
	return loadFromData([]byte(static.Areas))
}

func loadFromData(data []byte) ([]AreaData, error) {
	var ret []AreaData
	err := json.Unmarshal(data, &ret)
	return ret, err
}

func (opt defaultOption) SetDataSource(ds DataSource) Opt {
	return func(s *defaultAreaService) {
		s.ds = ds
	}
}

func (opt defaultOption) LoadFromFile(path string) Opt {
	return func(s *defaultAreaService) {
		s.ds = func() (data []AreaData, e error) {
			d, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, err
			}
			return loadFromData(d)
		}
	}
}
