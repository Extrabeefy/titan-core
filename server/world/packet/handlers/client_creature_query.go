package handlers

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/data/static"
	"github.com/Extrabeefy/titan_core/server/world/packet"
	"github.com/Extrabeefy/titan_core/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientCreatureQuery(pkt *packet.ClientCreatureQuery, state *system.State) ([]interfaces.ServerPacket, error) {
	response := new(packet.ServerCreatureQueryResponse)

	response.Unit = nil
	if unit, ok := static.Units[int(pkt.Entry)]; ok {
		response.Unit = unit
		response.Entry = pkt.Entry
	} else if pkt.GUID != 0 && state.OM.Exists(pkt.GUID) {
		response.Unit = state.OM.GetUnit(pkt.GUID).Template()
		response.Entry = uint32(response.Unit.Entry)
	}

	return []interfaces.ServerPacket{response}, nil
}
