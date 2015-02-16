package class

var (
	_shimClass  = &Class{name: "~shim"}
	_returnCode = []byte{0xb1} // return
	_athrowCode = []byte{0xbf} // athrow

	_shimMethod = &Method{
		ClassMember: ClassMember{
			AccessFlags: AccessFlags{ACC_STATIC},
			name:        "<return>",
			class:       _shimClass,
		},
		code: _returnCode,
	}

	_athrowMethod = &Method{
		ClassMember: ClassMember{
			AccessFlags: AccessFlags{ACC_STATIC},
			name:        "<athrow>",
			class:       _shimClass,
		},
		code: _athrowCode,
	}
)

func ShimMethod() *Method {
	return _shimMethod
}

func NewAthrowMethod(maxStack, maxLocals uint) *Method {
	return _athrowMethod
}