package BytePacker

import (
	"math"
	"bytes"
	"encoding/binary"
)

// Packer is a binary packer helps you pack data into an io.Writer.
type Packer struct {
	buffer *bytes.Buffer
	order binary.ByteOrder
	err error
}

func (p *Packer) SetOrder (order binary.ByteOrder) {
	p.order = order
}

// NewPacker returns a *Packer hold an io.Writer.
func NewPacker(buffer *bytes.Buffer) *Packer {
	return &Packer{
		buffer: buffer,
		order: binary.BigEndian,
	}
}

// Error returns an error if any errors exists
func (p *Packer) Error() error {
	return p.err
}

func (p *Packer) PushBool(b bool) *Packer {
	return p.errFilter(func() {
		val := byte(0)

		if b {
			val = byte(1)
		}

		_, p.err = p.buffer.Write([]byte{val})
	})
}

// PushByte write a single byte into buffer.
func (p *Packer) PushByte(b byte) *Packer {
	return p.errFilter(func() {
		_, p.err = p.buffer.Write([]byte{b})
	})
}

// PushBytes write a bytes array into buffer.
func (p *Packer) PushBytes(bytes []byte) *Packer {
	return p.errFilter(func() {
		_, p.err = p.buffer.Write(bytes)
	})
}

// PushUint16 write a uint16 into buffer.
func (p *Packer) PushUint16(i uint16) *Packer {
	return p.errFilter(func() {
		buffer := make([]byte, 2)
		p.order.PutUint16(buffer, i)
		_, p.err = p.buffer.Write(buffer)
	})
}

// PushUint16 write a int16 into buffer.
func (p *Packer) PushInt16(i int16) *Packer {
	return p.PushUint16(uint16(i))
}

// PushUint32 write a uint32 into buffer.
func (p *Packer) PushUint32(i uint32) *Packer {
	return p.errFilter(func() {
		buffer := make([]byte, 4)
		p.order.PutUint32(buffer, i)
		_, p.err = p.buffer.Write(buffer)
	})
}

// PushInt32 write a int32 into buffer.
func (p *Packer) PushInt32(i int32) *Packer {
	return p.PushUint32(uint32(i))
}

// PushUint64 write a uint64 into buffer.
func (p *Packer) PushUint64(i uint64) *Packer {
	return p.errFilter(func() {
		buffer := make([]byte, 8)
		p.order.PutUint64(buffer, i)
		_, p.err = p.buffer.Write(buffer)
	})
}

// PushInt64 write a int64 into buffer.
func (p *Packer) PushInt64(i int64) *Packer {
	return p.PushUint64(uint64(i))
}

// PushFloat32 write a float32 into buffer.
func (p *Packer) PushFloat32(i float32) *Packer {
	return p.PushUint32(math.Float32bits(i))
}

// PushFloat64 write a float64 into buffer.
func (p *Packer) PushFloat64(i float64) *Packer {
	return p.PushUint64(math.Float64bits(i))
}

// PushString write a string into buffer.
func (p *Packer) PushString(s string) *Packer {
	return p.errFilter(func() {
		_, p.err = p.buffer.Write([]byte(s))
	})
}

func (p *Packer) Bytes() []byte {
	return p.buffer.Bytes()
}

func (p *Packer) errFilter(f func()) *Packer {
	if p.err == nil {
		f()
	}
	return p
}
