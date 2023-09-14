package model

import "open.chat/pkg/util"

func GetFirstValue(params ...interface{}) interface{} {
	if len(params) == 0 {
		return nil
	}
	return params[0]
}

type IDList []int32

func (m *IDList) AddIfNot(id ...int32) []int32 {
	for _, id2 := range id {
		if id2 != 0 && !GetFirstValue(util.Contains(id2, *m)).(bool) {
			*m = append(*m, id2)
		}
	}
	return m.ToIDList()
}

func (m *IDList) AddFrontIfNot(id ...int32) []int32 {
	for _, id2 := range id {
		if id2 != 0 && !GetFirstValue(util.Contains(id2, *m)).(bool) {
			*m = append([]int32{id2}, *m...)
		}
	}
	return m.ToIDList()
}

func (m IDList) ToIDList() []int32 {
	if m == nil {
		return []int32{}
	}
	return m
}

func AddID32ListIfNot(idList IDList, id ...int32) []int32 {
	return idList.AddIfNot(id...)
}

func AddID32ListFrontIfNot(idList IDList, id ...int32) []int32 {
	return idList.AddFrontIfNot(id...)
}

type ID64List []int64

func (m *ID64List) AddIfNot(id ...int64) []int64 {
	for _, id2 := range id {
		if id2 != 0 && !GetFirstValue(util.Contains(id2, *m)).(bool) {
			*m = append(*m, id2)
		}
	}
	return m.ToIDList()
}

func (m *ID64List) AddFrontIfNot(id ...int64) []int64 {
	for _, id2 := range id {
		if id2 != 0 && !GetFirstValue(util.Contains(id2, *m)).(bool) {
			*m = append([]int64{id2}, *m...)
		}
	}
	return m.ToIDList()
}

func (m ID64List) ToIDList() []int64 {
	if m == nil {
		return []int64{}
	}
	return m
}

func AddID64ListIfNot(idList ID64List, id ...int64) []int64 {
	return idList.AddIfNot(id...)
}

func AddID64ListFrontIfNot(idList ID64List, id ...int64) []int64 {
	return idList.AddFrontIfNot(id...)
}
