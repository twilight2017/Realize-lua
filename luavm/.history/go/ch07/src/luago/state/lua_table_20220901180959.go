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

func(self *luaState) put(key, val luaValue){
	if key == nil{
		panic("table idx is nil!")
	}
	if f, ok := key.(float64);ok && math.IsNaN(f)
}
