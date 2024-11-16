package prefixmanager

import (
	"github.com/Nexell-AI-Network/nexelliad/infrastructure/logger"
	"github.com/Nexell-AI-Network/nexelliad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
