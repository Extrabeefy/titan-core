package handlers

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/packet"
	"github.com/Extrabeefy/titan_core/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientQueryTime(pkt *packet.ClientQueryTime, state *system.State) ([]interfaces.ServerPacket, error) {
	return []interfaces.ServerPacket{new(packet.ServerQueryTimeResponse)}, nil
}
