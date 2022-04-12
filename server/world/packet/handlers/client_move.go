package handlers

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/packet"
	"github.com/Extrabeefy/titan_core/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientMove(pkt *packet.ClientMove, state *system.State) ([]interfaces.ServerPacket, error) {
	state.Character.MovementInfo = pkt.MovementInfo

	location := state.Character.GetLocation()
	location.X = pkt.MovementInfo.Location.X
	location.Y = pkt.MovementInfo.Location.Y
	location.Z = pkt.MovementInfo.Location.Z
	location.O = pkt.MovementInfo.Location.O

	return nil, nil
}
