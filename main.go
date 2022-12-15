package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

type GetUserRoles struct {
	code    int32
	user_id int32
}

type Pack struct {
	opcode  int32
	payload interface{}
}

func (p *Pack) Write() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, compact_int(int(p.opcode)))
	binary.Write(&buf, binary.BigEndian, compact_int(int(unsafe.Sizeof(&p.payload))))
	binary.Write(&buf, binary.BigEndian, p.payload)
	return buf.Bytes()
}

func main() {
	userRoles := GetUserRoles{
		code:    -1,
		user_id: 32,
	}

	pack := Pack{
		opcode:  3401,
		payload: userRoles,
	}

	encodedMessage := pack.Write()

	fmt.Println(encodedMessage)
}
func compact_int(data int) []byte {
	if data <= 0x7F {
		return []byte{byte(data)}
	} else if data <= 16384 {
		encoded := make([]byte, 2)
		binary.BigEndian.PutUint16(encoded, uint16(data|0x8000))
		return encoded
	} else if data <= 536870912 {
		encoded := make([]byte, 4)
		binary.BigEndian.PutUint32(encoded, uint32(data|0xC0000000))
		return encoded
	}
	return []byte{224, byte(data)}
}
