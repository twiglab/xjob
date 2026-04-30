package hik

import "github.com/twiglab/xjob/pfsdk/hik/cfas"

func GroupBySum(l []cfas.PassengerFlowOut) *GroupBy {
	x := NewGroupBy()
	for _, e := range l {
		x.Add(e.GroupID, e.FlowInNum, e.FlowOutNum)
	}
	return x
}

type GroupByItem struct {
	Key string
	In  int
	Out int
}

type GroupBy struct {
	m map[string]*GroupByItem
}

func NewGroupBy() *GroupBy {
	return &GroupBy{
		m: make(map[string]*GroupByItem),
	}
}

func (x *GroupBy) Get(k string) (y *GroupByItem) {
	var ok bool
	if y, ok = x.m[k]; ok {
		return
	}
	y = &GroupByItem{Key: k}
	x.m[k] = y
	return
}

func (x *GroupBy) Add(k string, in, out int) {
	y := x.Get(k)
	y.In += in
	y.Out += out
}
