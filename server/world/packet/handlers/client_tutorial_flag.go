package handlers

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/packet"
	"github.com/Extrabeefy/titan_core/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientTutorialFlag(pkt *packet.ClientTutorialFlag, state *system.State) ([]interfaces.ServerPacket, error) {
	state.Character.Tutorials[pkt.Flag] = true
	return nil, nil
}
