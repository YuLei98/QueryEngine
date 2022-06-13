package untils

import (
	"sort"
	"sync"
)

type Set struct {
	sync.RWMutex
	Doc_Id map[int]bool `json:"doc_Id"`
}

// 新建集合对象
func New(items ...int) (s *Set) {
	s = &Set{
		Doc_Id: make(map[int]bool, len(items)),
	}
	s.Add(items...)
	return s
}

func New_From_Map(fp map[int]bool) (s *Set) {
	s = &Set{
		Doc_Id: fp,
	}
	return s
}

// 添加元素
func (s *Set) Add(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		s.Doc_Id[v] = true
	}
}

// 删除元素
func (s *Set) Remove(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		delete(s.Doc_Id, v)
	}
}

// 判断元素是否存在
func (s *Set) Has(items ...int) bool {
	s.RLock()
	defer s.RUnlock()
	for _, v := range items {
		if _, ok := s.Doc_Id[v]; !ok {
			return false
		}
	}
	return true
}

// 元素个数
func (s *Set) Count() int {
	return len(s.Doc_Id)
}

// 清空集合
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.Doc_Id = map[int]bool{}
}

// 空集合判断
func (s *Set) Empty() bool {
	return len(s.Doc_Id) == 0
}

// 无序列表
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.Doc_Id))

	for item := range s.Doc_Id {
		list = append(list, item)
	}
	return list
}

// 排序列表
func (s *Set) SortList() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.Doc_Id))
	for item := range s.Doc_Id {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}

// 并集
func (s *Set) Union(sets ...*Set) *Set {
	r := New(s.List()...)
	for _, set := range sets {
		for e := range set.Doc_Id {
			r.Doc_Id[e] = true
		}
	}
	return r
}

// 差集
func (s *Set) Minus(sets ...*Set) *Set {
	r := New(s.List()...)
	for _, set := range sets {
		for e := range set.Doc_Id {
			if _, ok := s.Doc_Id[e]; ok {
				delete(r.Doc_Id, e)
			}
		}
	}
	return r
}

// 交集
func (s *Set) Intersect(sets ...*Set) *Set {
	r := New(s.List()...)
	for _, set := range sets {
		for e := range s.Doc_Id {
			if _, ok := set.Doc_Id[e]; !ok {
				delete(r.Doc_Id, e)
			}
		}
	}
	return r
}

// 补集
func (s *Set) Complement(full *Set) *Set {
	r := New()
	for e := range full.Doc_Id {
		if _, ok := s.Doc_Id[e]; !ok {
			r.Add(e)
		}
	}
	return r
}

