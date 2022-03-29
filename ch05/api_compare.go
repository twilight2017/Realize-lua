//Compare()方法对指定索引处的两个值进行比较，返回结果
package ch05

func (self *luaState) Compare(idx1, idx2 int, op CompareOp) bool {
	a := self.stack.get(idx1)
	b := self.stack.get(idx2)
	switch op {
	case LUA_OPEQ:
		return _eq(a, b)
	case LUA_OPLT:
		return _lt(a, b)
	case LUA_OPLE:
		return _le(a, b)
	default:
		panic("invalid compare op!")
	}
}

func _eq(a, b luaValue) bool {
	switch x := a.(type) {
	case nil: //a的类型为nil
		return b == nil
	case bool:
		y, ok := b.(bool) //ok表示b能转化为bool型
		return ok && x == y
	case string:
		y, ok = b.(string)
		return ok && x == y
	case int64:
		switch y := b.(type); {
		case int64:
			return x == y
		case float64:
			return float64(x) == y
		default:
			return false
		}
	case float64:
		switch y := b.(type) {
		case int64:
			return x == float64(y)
		case float64:
			return x == y
		default:
			return false
		}
	default:
		return a == b
	}
}
