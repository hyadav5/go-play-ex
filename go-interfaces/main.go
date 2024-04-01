package main

import (
	"context"
	"study/go-play-ex/go-interfaces/samplepackage"
)

type HealthRecordType string

type sampleStruct struct {
	name string
}

const (
	// PersistentVolumeClaimResizing - a user trigger resize of pvc has been started
	PersistentVolumeClaimResizing HealthRecordType = "Resizing"
	// PersistentVolumeClaimFileSystemResizePending - controller resize is finished and a file system resize is pending on node
	PersistentVolumeClaimFileSystemResizePending HealthRecordType = "FileSystemResizePending"
)

type object1 struct {
	name string
}

type object2 struct {
	address string
}

func checkType(in HealthRecordType) {
	println(in)
}

func persistHealthRecord(healthRecordKey HealthRecordType, handlerFunc func(ctx context.Context)) {
	ctx := context.Background()
	handlerFunc(ctx)
}

func takeInterfaceObject(workerkey string, obj interface{}, f func()) {
	switch obj.(type) {
	case *object1:
		println(obj.(*object1).name)
		f()
	case *object2:
		println(obj.(*object2).address)
		f()
	default:
		println("default")
	}
}

func main() {
	//samplepackage.Init()
	//samplepackage.SystemHealthClient.GetMyName()
	//print(samplepackage.SystemHealthClient.GetMyName())

	var obj1 object1
	obj1.name = "hemant"
	var obj2 object2
	obj2.address = "vidisha"
	//takeInterfaceObject("key", &obj1, func() {
	//	println("obj1 handler:", obj1.name)
	//})
	//takeInterfaceObject("key", &obj2, func() {
	//	println("obj2 handler:", obj2.address)
	//})

	persistHealthRecord("key1", func(ctx context.Context) {
		println("obj1 handler:", obj1.name)
	})
	persistHealthRecord("key2", func(ctx2 context.Context) {
		ctx2.Value("")
		println("obj2 handler:", obj2.address)
	})

	//checkType(PersistentVolumeClaimResizing)
	//name := HealthRecordType("hemant")
	//checkType(name)

	//caller := samplepackage.NewSystemHealthServiceClient()
	//caller.GetService()
	//caller.GetServices()

	samplepackage.SystemHealthClient.GetService()
	samplepackage.SystemHealthClient.GetServices()
}
