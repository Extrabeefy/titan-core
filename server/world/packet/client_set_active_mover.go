package packet

import (
	"encoding/binary"
	"io"

	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/data/static"
)

// ClientSetActiveMover is sent from the client periodically.
type ClientSetActiveMover struct {
	GUID interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientSetActiveMover) FromBytes(buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.GUID)
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientSetActiveMover) OpCode() static.OpCode {
	return static.OpCodeClientSetActiveMover
}
