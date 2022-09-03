package api

func (self *luaState) Len(idx int) {
	val := self.stack.get(idx)
	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else if t, ok := val.(*luaTable); ok {
		self.stack.push(t.len())
	} else {
		panic("length error!")
	}
}

class A:
	a = 123


def foo():
    a = []
