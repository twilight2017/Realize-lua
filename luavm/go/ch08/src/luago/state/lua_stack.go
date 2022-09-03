package state


type luaStack struct{
	slots []luaValue
	top int
	// 下面是新增加的字段
	prev *luaStack
	closure &luaClosure //把函数原型替换成了闭包
	varargs []luaValue
	pc int
}