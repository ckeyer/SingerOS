package leros

import (
	"github.com/insmtx/Leros/backend/events"
	"github.com/insmtx/Leros/backend/internal/agent"
	einoadapter "github.com/insmtx/Leros/backend/internal/agent/eino"
)

type runState struct {
	req          *agent.RequestContext
	emitter      *events.Emitter
	userInput    string
	systemPrompt string
	toolBinding  einoadapter.ToolBinding
	maxStep      int
}
