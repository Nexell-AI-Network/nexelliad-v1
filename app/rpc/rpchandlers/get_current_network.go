package rpchandlers

import (
	"github.com/Nexell-AI-Network/nexelliad/app/appmessage"
	"github.com/Nexell-AI-Network/nexelliad/app/rpc/rpccontext"
	"github.com/Nexell-AI-Network/nexelliad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
