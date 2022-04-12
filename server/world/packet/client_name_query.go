package packet

import (
	"encoding/binary"
	"io"

	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/data/static"
)

// ClientNameQuery is sent from the client periodically.
type ClientNameQuery struct {
	GUID interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientNameQuery) FromBytes(buffer io.Reader) error {
	return binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
}

// OpCode gets the opcode of the packet.
func (*ClientNameQuery) OpCode() static.OpCode {
	return static.OpCodeClientNameQuery
}
