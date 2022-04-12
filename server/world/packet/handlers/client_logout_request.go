package handlers

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/packet"
	"github.com/Extrabeefy/titan_core/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientLogoutRequest(pkt *packet.ClientLogoutRequest, state *system.State) ([]interfaces.ServerPacket, error) {
	response := new(packet.ServerLogoutResponse)

	// TODO: Actually implement this!
	response.Reason = 0
	response.InstantLogout = true

	if state.Character != nil {
		state.Updater.Logout(state.Character.GUID())
	}

	return []interfaces.ServerPacket{
		response, new(packet.ServerLogoutComplete),
	}, nil
}
