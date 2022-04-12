package messages

import "github.com/Extrabeefy/titan_core/server/world/data/dynamic/interfaces"

type (
	UnitAttack struct {
		Target interfaces.GUID
	}

	UnitStopAttack struct{}

	UnitRegisterAttack struct {
		Attacker interfaces.GUID
	}

	UnitDeregisterAttacker struct {
		Attacker interfaces.GUID
	}

	UnitDied struct {
		DeadUnit interfaces.GUID
	}
)
