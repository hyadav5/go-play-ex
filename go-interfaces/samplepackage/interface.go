package samplepackage

type systemHealth struct {
	myName string
}

var sh systemHealth

type SystemHealthService interface {
	//AviServiceViewer
	GetMyName() string

	// Other Viewers interfaces
}

type systemHealthClientImpl struct {
	client *systemHealth
}

var SystemHealthClient systemHealthClientImpl

func Init() {
	sh.myName = "hemant yadav"
	SystemHealthClient.client = &sh
}

//var SystemHealthClient SystemHealthClientImpl

func (shc *systemHealthClientImpl) GetMyName() string {
	return shc.client.myName
}

//type Child1Interface interface {
//	Child1Method1()
//}
//
//type Child2Interface interface {
//	Child2Method1()
//}
//
//type MasterInterface interface {
//	Child1Interface
//	Child2Interface
//}
//
//type Master struct {
//}
//
//func (m *Master) Child1Method1() {
//	println("Child1Method1()")
//}
//
//func (m *Master) Child2Method1() {
//	println("Child2Method1()")
//}
