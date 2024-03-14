package samplepackage

type Child1Interface interface {
	Child1Method1()
}

type Child2Interface interface {
	Child2Method1()
}

type MasterInterface interface {
	Child1Interface
	Child2Interface
}

type Master struct {
}

func (m *Master) Child1Method1() {
	println("Child1Method1()")
}

func (m *Master) Child2Method1() {
	println("Child2Method1()")
}
