package packet

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"math/big"

	"github.com/Extrabeefy/titan-core/packet"
	"github.com/Extrabeefy/titan-core/server/auth/srp"
)

// OpCodes used by the AuthServer.
const (
	ClientLoginChallengeOpCode = 0
	ServerLoginChallengeOpCode = 0
)

// ClientLoginChallenge encodes information about a new connection to the
// login server.
type ClientLoginChallenge struct {
	Error          uint8
	GameName       [4]byte
	Version        [3]uint8
	Build          uint16
	Platform       [4]byte
	OS             [4]byte
	Locale         [4]byte
	TimezoneOffset uint32
	IPAddress      uint32
	AccountName    []byte
}

// Read will load a ClientLoginChallenge packet from a buffer.
// An error will be returned if at least one of the fields didn't load correctly.
func (pkt *ClientLoginChallenge) Read(buffer *bufio.Reader) error {
	binary.Read(buffer, binary.LittleEndian, &pkt.Error)
	buffer.Read(make([]byte, 2)) // unused: packet length
	binary.Read(buffer, binary.LittleEndian, &pkt.GameName)
	binary.Read(buffer, binary.LittleEndian, &pkt.Version)
	binary.Read(buffer, binary.LittleEndian, &pkt.Build)
	binary.Read(buffer, binary.LittleEndian, &pkt.Platform)
	binary.Read(buffer, binary.LittleEndian, &pkt.OS)
	binary.Read(buffer, binary.LittleEndian, &pkt.Locale)
	binary.Read(buffer, binary.LittleEndian, &pkt.TimezoneOffset)
	binary.Read(buffer, binary.BigEndian, &pkt.IPAddress)

	var accountNameLen uint8
	binary.Read(buffer, binary.LittleEndian, &accountNameLen)

	pkt.AccountName = make([]byte, accountNameLen)
	return binary.Read(buffer, binary.LittleEndian, &pkt.AccountName)
}

// ServerLoginChallenge is the server's response to a client's challenge. It contains
// some SRP information used for handshaking.
type ServerLoginChallenge struct {
	Error   uint8
	B       big.Int
	Salt    big.Int
	SaltCRC big.Int
}

// Bytes writes out the packet to an array of bytes.
func (pkt *ServerLoginChallenge) Bytes() []byte {
	buffer := bytes.NewBufferString("")

	buffer.WriteByte(ServerLoginChallengeOpCode)
	buffer.WriteByte(0) // unk1
	buffer.WriteByte(pkt.Error)

	if pkt.Error == 0 {
		buffer.Write(padBigIntBytes(reverse(pkt.B.Bytes()), 32))
		buffer.WriteByte(1)
		buffer.WriteByte(srp.G)
		buffer.WriteByte(32)
		buffer.Write(reverse(srp.N().Bytes()))
		buffer.Write(padBigIntBytes(reverse(pkt.Salt.Bytes()), 32))
		buffer.Write(padBigIntBytes(reverse(pkt.SaltCRC.Bytes()), 16))
		buffer.WriteByte(0) // unk2
	}

	return buffer.Bytes()
}

// Handle will check the database for the account and send an appropriate response.
func (pkt *ClientLoginChallenge) Handle() ([]packet.ServerPacket, error) {
	response := new(ServerLoginChallenge)

	// TODO(jeshua): Read this data from a database instead.
	salt := srp.GenerateSalt()
	v := srp.GenerateVerifier("JESHUA", "JESHUA", salt)

	// TODO(jeshua): Save this information with the session.
	_, B := srp.GenerateEphemeral(v)

	response.Error = 0
	response.B.Set(B)
	response.Salt.Set(salt)
	response.SaltCRC.SetInt64(0)

	return []packet.ServerPacket{response}, nil
}
