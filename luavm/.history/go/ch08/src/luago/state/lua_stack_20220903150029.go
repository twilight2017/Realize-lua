package state


type luaStack struct{
	slots []luaValue
	top int
	// 下面是新增加的字段
	prev *luaStack
	closure &luaClosure
	varargs []luaValue
	pc int
}