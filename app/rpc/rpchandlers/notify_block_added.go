package rpchandlers

import (
	"github.com/Nexell-AI-Network/nexelliad/app/appmessage"
	"github.com/Nexell-AI-Network/nexelliad/app/rpc/rpccontext"
	"github.com/Nexell-AI-Network/nexelliad/infrastructure/network/netadapter/router"
)

// HandleNotifyBlockAdded handles the respectively named RPC command
func HandleNotifyBlockAdded(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateBlockAddedNotifications()

	response := appmessage.NewNotifyBlockAddedResponseMessage()
	return response, nil
}
