package handlers

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/messages"
	"github.com/Extrabeefy/titan_core/server/world/packet"
	"github.com/Extrabeefy/titan_core/server/world/system"
)

func HandleClientAttackStop(pkt *packet.ClientAttackStop, state *system.State) ([]interfaces.ServerPacket, error) {
	state.Character.SendUpdates([]interface{}{
		&messages.UnitStopAttack{},
	})

	return []interfaces.ServerPacket{
		&packet.ServerAttackStop{
			Attacker: state.Character.GUID(),
			Target:   state.Character.Target,
		},
	}, nil
}
