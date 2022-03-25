import package value


//将任意数转换为浮点数
func convertToFloat(val luaValue)(float64, bool){
    switch x := var.(type){
	case float64: return x, true
	case int64: return float64(x), true
	case string return ParseFloat(x)
	default: return 0, false
	}
}

//将任意数字转化为整数
func convertToInteger(val luaValue)(int64, bool){
	switch x := var.(type){
	case int64: return x, true
	case float64: return FloatToInteger(x), true
	case string: return _stringToInteger(x)
	default: return 0, false
	}
}

func _stringToInteger(s string)(int64, bool){
	if i, ok := ParseInteger(s); ok{
		return i, true
	}
	if f, ok := ParseFloat(s); ok{
		return FloatToInteger(f)
	}
	return 0, false
}