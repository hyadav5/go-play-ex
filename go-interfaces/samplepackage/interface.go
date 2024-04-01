package samplepackage

import (
	"context"
	"github.com/coreos/go-systemd/dbus"
)

type healthRecordType string

const (
	// SecureChannelStatus - record to define the secure channel status across the controller nodes
	secureChannelStatus healthRecordType = "SecureChannelStatus"
)

type systemHealth struct {
	localNodeName string
	dBusListener  dataCollector
}

// Interface which expects the implementor to implement atleast below 3 functions.
type dataCollector interface {
	init(sh *systemHealth) error
	run() error
	persistHealthRecord(recordType healthRecordType, handlerFuncForRecordType func(ctx context.Context)) error
}

// dBusListener implements the dataCollector interface.
// It listens to the system-d bus event for the AVI service events.
type dBusListener struct {
	sysHealth  *systemHealth
	dBusClient *dbus.Conn
}

func (dbl *dBusListener) init(sh *systemHealth) error {
	println("init()")
	return nil
}

func (dbl *dBusListener) run() error {
	println("run()")
	return nil
}

func (dbl *dBusListener) persistHealthRecord(recordType healthRecordType, handlerFuncForRecordType func(ctx context.Context)) error {
	println("persistHealthRecord()")
	return nil
}

// dBusListener implements the dataCollector interface.
// It listens to the system-d bus event for the AVI service events.
type newsListener struct {
	sysHealth  *systemHealth
	dBusClient *dbus.Conn
}

func (nl *newsListener) init(sh *systemHealth) error {
	println("init()")
	return nil
}

func (nl *newsListener) run() error {
	println("run()")
	return nil
}

func (nl *newsListener) persistHealthRecord(recordType healthRecordType, handlerFuncForRecordType func(ctx context.Context)) error {
	println("persistHealthRecord()")
	return nil
}

func NewImplementors() {
	sh := systemHealth{}

	var dBusListenerObj dataCollector = &dBusListener{}
	dBusListenerObj.init(&sh)
	dBusListenerObj.run()
	dBusListenerObj.persistHealthRecord("", func(ctx context.Context) {})

	var newsListenerObj dataCollector = &newsListener{}
	newsListenerObj.init(&sh)
	newsListenerObj.run()
	newsListenerObj.persistHealthRecord("", func(ctx context.Context) {})
}

// ============ start of CLIENT APIS ================ //
type systemHealthClientIf interface {
	GetServices()
	GetService()
}

type systemHealthClient struct {
	client *string
}

func (impl *systemHealthClient) GetServices() {
	println("called GetServices()")
}
func (impl *systemHealthClient) GetService() {
	println("called GetService()")
}

var SystemHealthClient systemHealthClientIf = &systemHealthClient{}

/*
If systemHealthClient have few memory maps or pointers which need to be freed.
Some data connection which needs to closed.
Then we should provide the New() method and return the object.
that Object can be used to close() the things when caller is done.

Here in this particular example.
*/

//func NewSystemHealthServiceClient() SystemHealthClientIf {
//	var systemHealthClientImplObj SystemHealthClientIf = &systemHealthClientImpl{}
//	return systemHealthClientImplObj
//}

// ============ end of CLIENT APIS ================ //
