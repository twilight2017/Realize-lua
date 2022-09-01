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
