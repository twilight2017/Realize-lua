//用数组和哈希表混合的方式来实现lua的表
package state

import "math"

import "luago/number"

type luaTable struct {
	arr  []luaValue
	_map map[luavalue]luaValue
}
