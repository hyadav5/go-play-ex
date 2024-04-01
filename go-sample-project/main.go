package main

import (
	"context"
	"fmt"
)

type HealthRecordType string

type dataCollector interface {
	run()
	persistHealthRecord(recordType HealthRecordType, handlerFunc func(ctx context.Context))
}

//
//func (t *dataCollector) persistHealthRecord(healthRecordKey HealthRecordType, handlerFunc func(ctx context.Context)) {
//	ctx := context.Background()
//	handlerFunc(ctx)
//}

type systemHealth struct {
	localNodeName string
}

type dBusListener struct {
	sysHealth *systemHealth
	source    string
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *dBusListener) run() {
	fmt.Println("running dBusListener")
}

func (t *dBusListener) persistHealthRecord(healthRecordKey HealthRecordType, handlerFunc func(ctx context.Context)) {
	ctx := context.Background()
	println(t.sysHealth.localNodeName)
	handlerFunc(ctx)
}

func (t T) N() {
	fmt.Println(t.S)
}

func main() {
	var sh systemHealth
	sh.localNodeName = "localnode"
	var dBusListerObj dataCollector = &dBusListener{
		source:    "mySource",
		sysHealth: &sh,
	}
	dBusListerObj.run()
	dBusListerObj.persistHealthRecord("serviceName", func(ctx context.Context) {
		println("hey their")
	})
}

//import (
//	//aviutils "avi/utils"
//	"context"
//)
//

//
//type DataCollector interface {
//	Run()
//	PersistHealthRecord(HealthRecordType, handlerFunc func(ctx context.Context, recordKey string))
//}
//
//type SystemHealth struct {
//	// Keyed worker pool for persisting the Health Records from various sources into key-value Store
//	//wWorkerPool *aviutils.KeyedWorkerPool
//	// Local node name e.g. node1
//	localNodeName string
//}
//
//type dBusListener struct {
//	sysHealth *SystemHealth
//	source    string
//}
//
//func (sh *SystemHealth) Run() {
//	println("i am running")
//}
//
////func (sh SystemHealth) PersistHealthRecord(healthRecordKey HealthRecordType, handlerFunc func(ctx context.Context)) {
////	ctx := context.Background()
////	handlerFunc(ctx)
////}
//
////func (sh *dBusListener) run() {
////	println("i am running dbus")
////}
////
////func (sh *dBusListener) persistHealthRecord(healthRecordKey HealthRecordType, handlerFunc func(ctx context.Context)) {
////	ctx := context.Background()
////	handlerFunc(ctx)
////}
//
//func main() {
//	var sh SystemHealth
//	//sh.wWorkerPool = aviutils.MakeKeyedWorkerPool(runtime.GOMAXPROCS(0), 4*runtime.GOMAXPROCS(0))
//
//	var dbl dBusListener
//	dbl.source = "mysource"
//	dbl.sysHealth = &sh
//
//	//sh.run()
//	//sh.persistHealthRecord("key1", func(ctx context.Context) {
//	//	println("key1 handler:")
//	//})
//	//
//	//dbl.sysHealth.persistHealthRecord("key1", func(ctx2 context.Context) {
//	//	println("key1 handler:")
//	//	println(ctx2.Done())
//	//	println(dbl.source)
//	//})
//
//	//dbl.run()
//	//dbl.sysHealth.persistHealthRecord()
//
//	return
//}
