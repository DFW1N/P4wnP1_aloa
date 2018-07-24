package main

import (
	"../common"
	pb "../proto/gopherjs"
	"errors"
	"github.com/gopherjs/gopherjs/js"
)

type EventLog struct {
	*js.Object
	EvLogSource  string `js:"source"`
	EvLogLevel   int  `js:"level"`
	EvLogMessage string `js:"message"`
	EvLogTime string `js:"time"`
}

func DeconstructEventLog(gRPCEv *pb.Event) (res *EventLog, err error) {
	if gRPCEv.Type != common.EVT_LOG { return nil,errors.New("No log event")}

	res = &EventLog{Object:O()}
	switch vT := gRPCEv.Values[0].Val.(type) {
	case *pb.EventValue_Tstring:
		res.EvLogSource = vT.Tstring
	default:
		return nil, errors.New("Value at position 0 has wrong type for a log event")
	}
	switch vT := gRPCEv.Values[1].Val.(type) {
	case *pb.EventValue_Tint64:
		res.EvLogLevel = int(vT.Tint64)
	default:
		return nil, errors.New("Value at position 1 has wrong type for a log event")
	}
	switch vT := gRPCEv.Values[2].Val.(type) {
	case *pb.EventValue_Tstring:
		res.EvLogMessage = vT.Tstring
	default:
		return nil, errors.New("Value at position 2 has wrong type for a log event")
	}
	switch vT := gRPCEv.Values[3].Val.(type) {
	case *pb.EventValue_Tstring:
		res.EvLogTime = vT.Tstring
	default:
		return nil, errors.New("Value at position 3 has wrong type for a log event")
	}

	return res, nil
}
