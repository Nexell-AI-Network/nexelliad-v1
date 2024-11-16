package flowcontext

import (
	"github.com/Nexell-AI-Network/nexelliad/domain"
)

// Domain returns the Domain object associated to the flow context.
func (f *FlowContext) Domain() domain.Domain {
	return f.domain
}
