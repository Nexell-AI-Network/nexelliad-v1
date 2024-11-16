package standalone

import (
	"github.com/Nexell-AI-Network/nexelliad/infrastructure/logger"
	"github.com/Nexell-AI-Network/nexelliad/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
