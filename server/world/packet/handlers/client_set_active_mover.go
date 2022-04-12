package handlers

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/packet"
	"github.com/Extrabeefy/titan_core/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientSetActiveMover(pkt *packet.ClientSetActiveMover, state *system.State) ([]interfaces.ServerPacket, error) {
	if pkt.GUID != state.Character.GUID() {
		state.Log.Errorf("Incorrect mover GUID: it is %v, but should be %v", pkt.GUID, state.Character.GUID())
	}

	return nil, nil
}
