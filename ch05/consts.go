const (
	LUA_OPADD  = iota //+
	LUA_OPSUB         //-
	LUA_OPMUL         //*
	LUA_OPMOD         //%
	LUA_OPPOW         //^
	LUA_OPDIV         // /
	LUA_OPIDIV        // //
	LUA_OPBAND        //&
	LUA_OPBOR         //|
	LUA_OPBXOR        //-
	LUA_OPSHL         //<<
	LUA_OPSHR         //>>
	LUA_OPUNM         //-
	LUA_OPBNOT        //-
)

const (
	LUA_OPEQ = iota //==
	LUA_OPLT        //<
	LUA_OPLE        // <=
)