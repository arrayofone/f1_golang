package utils

type ByteCursor struct {
	bytes []byte
	curs  int
	max   int
}

func NewByteCursor(b []byte) *ByteCursor {
	return &ByteCursor{
		bytes: b,
		curs:  0,
		max:   len(b),
	}
}

func (b *ByteCursor) b(length int) []byte {
	var r []byte

	l := b.curs + length

	if length == -1 {
		r = b.bytes[b.curs:]
		b.curs = 0
	} else {
		r = b.bytes[b.curs:l]
		b.curs = l
	}

	return r
}

func (b *ByteCursor) B(length int) []byte {
	return b.b(length)
}

func (b *ByteCursor) Uint8() uint8 {
	return D_uint8(b.B(1))
}

func (b *ByteCursor) Uint16() uint16 {
	return D_uint16(b.B(2))
}

func (b *ByteCursor) Uint32() uint32 {
	return D_uint32(b.B(4))
}

func (b *ByteCursor) Uint64() uint64 {
	return D_uint64(b.B(8))
}

func (b *ByteCursor) Float() float32 {
	return D_float(b.B(4))
}

func (b *ByteCursor) Int8() int8 {
	return D_int8(b.B(1))
}

func (b *ByteCursor) Int16() int16 {
	return D_int16(b.B(2))
}

func (b *ByteCursor) Int32() int32 {
	return D_int32(b.B(4))
}

func (b *ByteCursor) Int64() int64 {
	return D_int64(b.B(8))
}

func (b *ByteCursor) Bool() bool {
	return D_bool(b.B(1))
}

func (b *ByteCursor) String(len int) string {
	return D_string(b.B(len))
}
