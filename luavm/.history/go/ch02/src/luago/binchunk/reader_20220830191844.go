package binchunk

import "encoding/binary"
import "math"

type reader struct {
	data []byte
}

func (self *reader) readByte() byte {
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

func (self *reader) readUint32() uint32 {
	i := binary.LittleEndian.Uint32(self.data)
	self.data = self.data[4:]
	return i
}

func (self *reader) readUint64() uint64 {
	i := binary.LittleEndian.Uint64(self.data)
	self.data = self.data[8:]
	return i
}

func (self *reader) readLuaInteger() int64 {
	return int64(self.readUint64())
}

func (self *reader) readLuaNumber() float64 {
	return math.Float64frombits(self.readUint64())
}

func (self *reader) readString() string {
	size := uint(self.readByte())
	if size == 0 {
		//NULL字符串
		return ""
	}
	if size == 0xFF {
		//长字符串
		size = uint(self.readUint64())
	}
	bytes := self.readBytes(size - 1)
	return string(bytes)
}

func (self *reader) readBytes(n uint) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

func (self *reader) checkHeader() {
	if string(self.readBytes(4)) != LUA_SIGNATURE {
		panic("not a precompiled chunk!")
	} else if self.readByte() != LUAC_VERSION {
		panic("version mismatch ! ")
	} else if self.readByte() != LUAC_FORMAT {
		panic("format mismatch!")
	} else if string(self.readBytes(6)) != LUAC_DATA {
		panic("corrupted!")
	} else if self.readByte() != CINT_SIZE {
		panic("int size mismatch!")
	} else if self.readByte() != CSZIET_SIZE {
		panic("size_t size mismatch!")
	} else if self.readByte() != INSTRUCTION_SIZE {
		panic("instruction size mismatch!")
	} else if self.readByte() != LUA_INTEGER_SIZE {
		panic("lua_Integer size mismatch!")
	} else if self.readByte() != LUA_NUMBER_SIZE {
		panic("lua_Number size mismatch!")
	} else if self.readLuaInteger() != LUAC_INT {
		panic("endianness mismatch!")
	} else if self.readLuaNumber() != LUAC_NUM {
		panic("float format mismatch!")
	}
}

//从字节流中读取函数原型
func (self *reader) readProto(parentSource string) *Prototype {
	source := self.readString()
	if source == "" {
		source == parentSource
	}
	return &Prototype{
		Source:          source,
		LineDefined:     self.readUint32(),
		LastLineDefined: self.readUint32(),
		NumParams:       self.readByte(),
		IsVararg:        self.readByte(),
		MaxStackSize:    self.readByte(),
		Code:            self.readByte(),
		Constants:       self.readConstats(),
		Upvalues:        self.readUpvalues(),
		Protos:          self.readProtos(source),
		LineInfo:        self.readLineInfo(),
		LocVars:         self.readLocVars(),
		UpvalueNames:    self.readUpvalueNames(),
	}
}
