// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/carea"
	"testing"
)

func TestArea(t *testing.T) {
	s := carea.NewAreaService()
	t.Log("level: ", s.AreaLevel())
	v, err := s.Data()
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range v {
		t.Log(a)
	}

	all, err := s.Areas()
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range all {
		t.Logf("==========\n %s \n", a.String())
	}
}

func TestAreaCode(t *testing.T) {
	s := carea.NewAreaService()
	all, err := s.AreaByCode("110000")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("==========\n %s \n", all.String())
}
