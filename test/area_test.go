// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/carea"
	"testing"
)

func TestAreaAll(t *testing.T) {
	s := carea.NewAreaService()
	t.Log("level: ", s.AreaLevelNumber())
	v, err := s.Data()
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range v {
		t.Log(a)
	}

	t.Run("nosub", func(t *testing.T) {
		all, err := s.Areas(false)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})
	t.Run("one", func(t *testing.T) {
		all, err := s.Areas(true)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})
}

func TestAreaLevels(t *testing.T) {
	s := carea.NewAreaService()
	t.Log("level: ", s.AreaLevelNumber())
	for _, lv := range s.AreaLevels() {
		t.Log(lv)
	}

	t.Run("lv3", func(t *testing.T) {
		all, err := s.AreaByLevel("3", true)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})
}

func TestAreaCode(t *testing.T) {
	s := carea.NewAreaService()
	all, err := s.AreaByCode("110000", false)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("==========\n %s \n", all.String())

	lv2, err := s.SubareaByCode("110000", false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(lv2)

	all, err = s.AreaByCode("110000", true)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("==========\n %s \n", all.String())
}

func TestAreaSubareaByCode(t *testing.T) {
	s := carea.NewAreaService()
	t.Run("lv 1", func(t *testing.T) {
		all, err := s.Areas(false)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})

	t.Run("name", func(t *testing.T) {
		all, err := s.AreaByName("四川省", false)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})

	t.Run("lv 1", func(t *testing.T) {
		all, err := s.SubareaByCode("510000", false)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})

	t.Run("lv 2", func(t *testing.T) {
		all, err := s.SubareaByCode("510100", false)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})

	t.Run("lv 3", func(t *testing.T) {
		all, err := s.SubareaByCode("510181", false)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})
}

func TestAreaParentAreaByCode(t *testing.T) {
	s := carea.NewAreaService()
	t.Run("lv 1", func(t *testing.T) {
		all, err := s.Areas(false)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})

	t.Run("name", func(t *testing.T) {
		all, err := s.AreaByName("四川省", false)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range all {
			t.Logf("==========\n %s \n", a.String())
		}
	})

	t.Run("lv 1", func(t *testing.T) {
		all, err := s.ParentAreaByCode("510000", false)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(all)
	})

	t.Run("lv 2", func(t *testing.T) {
		all, err := s.ParentAreaByCode("510100", false)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(all)
	})

	t.Run("lv 3", func(t *testing.T) {
		all, err := s.ParentAreaByCode("510181", true)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(all)
	})
}

