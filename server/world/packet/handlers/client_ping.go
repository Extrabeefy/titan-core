package handlers

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/packet"
	"github.com/Extrabeefy/titan_core/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientPing(pkt *packet.ClientPing, state *system.State) ([]interfaces.ServerPacket, error) {
	response := new(packet.ServerPong)
	response.Pong = pkt.Ping

	return []interfaces.ServerPacket{response}, nil
}
