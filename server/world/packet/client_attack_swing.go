package packet

import (
	"encoding/binary"
	"io"

	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/data/static"
)

// ClientAttackSwing is sent from the client periodically.
type ClientAttackSwing struct {
	Target interfaces.GUID
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientAttackSwing) FromBytes(buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Target)
	return nil
}

// OpCode gets the opcode of the packet.
func (*ClientAttackSwing) OpCode() static.OpCode {
	return static.OpCodeClientAttackswing
}
