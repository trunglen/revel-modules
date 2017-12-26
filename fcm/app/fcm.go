package app

import (
	fcm "github.com/NaySoftware/go-fcm"
)

const (
	RESPONSE_FAIL = "fail"
)

type FcmClient struct {
	*fcm.FcmClient
}

func NewFCM(serverKey string) *FcmClient {
	return &FcmClient{
		FcmClient: fcm.NewFcmClient(serverKey),
	}
}

func (f *FcmClient) SendToMany(ids []string, data interface{}) (error, string) {
	f.NewFcmRegIdsMsg(ids, data)
	status, err := f.Send()
	if err != nil {
		return err, RESPONSE_FAIL
	}
	return nil, status.Err
}

func (f *FcmClient) SendToOne(id string, data interface{}) (error, string) {
	return f.SendToMany([]string{id}, data)
}
