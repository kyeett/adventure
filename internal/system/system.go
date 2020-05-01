package system

import "github.com/kyeett/single-player-game/internal/comp"

// A RenderSystem is updated every iteration, and draws to a screen
type System interface {
	Add(v interface{})
	Remove(id comp.ID)
}
