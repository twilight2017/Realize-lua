package ch04

type LuaType = int

type luaStack struct {
	slots []luaValue //用于存放值
	top   int        //用于记录栈顶索引
}

//定义全部接口
type LuaState interface{
	/*basic stack manipulation*/
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx, toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx, n int)
	SetTop(idx int)
	/*access functions (stack->Go)*/
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInterger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	ToBoolean(idx int) bool
	ToInterger(idx int) int64
	ToIntergerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string bool)
	/*push functions (Go-> stack)*/
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)
}

//用于创建指定容量的栈
func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

//check方法检查栈的空闲空间是否还可以容纳至少n个值；不能时，调用append函数扩容
func (self *luaStack) check(n int) {
	free := len(self.slots) - self.top
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil)
	}
}

//pop方法从栈顶弹出一个值
func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("stack underflow!")
	}
	self.top--
	val := self.slots[self.top] //取栈顶元素传值给val变量
	self.slots[self.top] = nil
	return val
}

//该方法将索引值转换为绝对索引
func (self *luaStack) absIndex(idx int) int{
	if idx >= 0{
		return idx
	}
	else{
		return idx+self.top+1
	}
}

//判断索引值是否有效
func (self *luaStack) isValid(idx int) bool{
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx < = self.top
}

//根据索引从堆栈里取值，索引无效就用panic终止程序
func (self *luaStack) set(idx int, val luaValue){
	absIdx := self.absIndex(idx)
	if absIdx >=0 && absIdx < = self.top{
		self.slots[absIdx] = val
		return
	}
	panic("invalid index!")
}