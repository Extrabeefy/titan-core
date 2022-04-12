package handlers

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/data/static"
	"github.com/Extrabeefy/titan_core/server/world/packet"
	"github.com/Extrabeefy/titan_core/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientCharDelete(pkt *packet.ClientCharDelete, state *system.State) ([]interfaces.ServerPacket, error) {
	response := new(packet.ServerCharDelete)
	response.Error = static.CharErrorCodeDeleteFailed
	return []interfaces.ServerPacket{response}, nil
}
