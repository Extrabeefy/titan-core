package dynamic

import (
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"
	"github.com/Extrabeefy/titan_core/server/world/data/dynamic/messages"
)

func (u *Unit) HandleAttack(attacker interfaces.GUID) {
	// TODO(jeshua): more complex AI.
	u.SendUpdates([]interface{}{
		&messages.UnitAttack{Target: attacker},
	})
}

func (u *Unit) HandleAttackStop(attacker interfaces.GUID) {
}
