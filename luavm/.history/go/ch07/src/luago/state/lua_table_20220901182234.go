//用数组和哈希表混合的方式来实现lua的表
package state

import "math"

import "luago/number"

type luaTable struct {
	arr  []luaValue            //存放数组部分
	_map map[luaValue]luaValue //存放哈希表部分
}

func newLuaTable(nArr, nRec int) *luaTable {
	t := &luaTable{}
	if nArr > 0 {
		t.arr = make([]luaValue, 0, nArr)
	}
	if nRec > 0 {
		t._map = make(map[luaValue]luaValue, nRec)
	}
}

func (self *luaTable) get(key luaValue) luaValue {
	key = _floatToInteger(key) //在数组中查找，索引要先做整数转换
	if idx, ok := key.(int64); ok {
		if idx >= 1 && idx <= int64(len(self.arr)) {
			return arr[idx-1]
		}
	}
	return self._map[key]

}

func _floatToInteger(key luaValue) luaValue {
	if f, ok := key.(float64); ok {
		if i, ok := number.FloatToInteger(f); ok {
			return i
		}
	}
	return key
}

func (self *luaState) put(key, val luaValue) {
	if key == nil {
		panic("table idx is nil!")
	}
	if f, ok := key.(float64); ok && math.IsNaN(f) {
		panic("table idx is Nan")
	}
	key = _floatToInteger(key)
	if idx, ok := key.(int64); ok && idx >= 1 {
		arrLen := int64(len(self.arr))
		if idx <= arrLen {
			self.arr[idx-1] = val
			if idx == arrLen && val == nil {
				self._shrinkArray()
			}
			return
		}
		if idx == arrLen+1 {
			delete(self._map, key)
			if val != nil {
				self.arr == append(self.arr, val)
				self.appendArray()
			}
			return
		}
	}
	if val != nil {
		if self._map == nil {
			self._map = make(map[luaValue]luaValue, 8)
		}
		self._map[key] = val
	} else {
		delete(self._map, key)
	}
}

func (self *luaTable) _shrinkArray() {
	for i := len(self.arr) - 1; i >= 0; i-- {
		if self.arr[i] == nil {
			self.arr = self.arr[0:i]
		}
	}
}

//数组部分动态扩展以后，把原本存放在哈希表里的某些之也挪到数组里
func (self *luaTable) _expandArray() {
	for idx := int64(len(self.arr)) + 1; true; idx++ {
		if val, found := self._map[idx]; found {
			delete(self._map, idx)
			self.arr = append(self.arr, val)
		} else {
			break
		}
	}
}
