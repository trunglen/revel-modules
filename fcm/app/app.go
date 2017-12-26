package app

import (
	"github.com/revel/revel"
)

var (
	Fcm *FcmClient
)

func InitFCM() {
	serverKey, _ := revel.Config.String("fcm.server-key")
	if Fcm == nil {
		Fcm = NewFCM(serverKey)
	}
}
func init() {
	revel.RegisterModuleInit(func(module *revel.Module) {
		InitFCM()
	})
}
