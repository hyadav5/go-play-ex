package main

import (
	"fmt"
	"strings"
	"time"
)

type restartHistory struct {
	Time         string
	Reason       string
	JournalEntry string
}

//
//type restartHistories struct {
//	Latest    int8
//	Histories []restartHistory
//}

type restartEvent struct {
	Time time.Time
}

//
//type restartEvents struct {
//	Latest int8
//	Events []restartEvent
//}

func getWriteIndexForLatestIdx(latestIdx int8, factor int8) int8 {
	writeIdx := latestIdx - 1
	if writeIdx < 0 {
		writeIdx = factor + writeIdx
	}
	return writeIdx
}

type CircularQueueHistory struct {
	MaxItems    int8
	Items       map[int8]restartHistory
	LatestIndex int8
}

// NewCircularQueue creates a new circular queue.
func NewCircularQueueHistory(max int8) *CircularQueueHistory {
	items := make(map[int8]restartHistory, max)
	q := CircularQueueHistory{MaxItems: max, Items: items, LatestIndex: 1}
	return &q
}

// Add adds a new message to the queue, removing the oldest if necessary.
func (cb *CircularQueueHistory) Add(message restartHistory) {
	cb.LatestIndex = getWriteIndexForLatestIdx(cb.LatestIndex, cb.MaxItems)
	cb.Items[cb.LatestIndex] = message
}

func (cb *CircularQueueHistory) GetData() []restartHistory {
	ret := []restartHistory{}
	readIdx := cb.LatestIndex
	for i := 0; i < int(cb.MaxItems); i++ {
		value, exists := cb.Items[readIdx]
		if exists {
			ret = append(ret, value)
		}
		readIdx = readIdx + 1
		if readIdx > cb.MaxItems-1 {
			readIdx = readIdx - cb.MaxItems
		}
	}
	return ret
}

type CircularQueueEvent struct {
	MaxItems    int8
	Items       map[int8]restartEvent
	LatestIndex int8
}

// NewCircularQueue creates a new circular queue.
func NewCircularQueueEvent(max int8) *CircularQueueEvent {
	items := make(map[int8]restartEvent, max)
	q := CircularQueueEvent{MaxItems: max, Items: items, LatestIndex: 1}
	return &q
}

// Add adds a new message to the queue, removing the oldest if necessary.
func (cb *CircularQueueEvent) Add(message restartEvent) {
	cb.LatestIndex = getWriteIndexForLatestIdx(cb.LatestIndex, cb.MaxItems)
	cb.Items[cb.LatestIndex] = message
}

func (cb *CircularQueueEvent) GetData() []restartEvent {
	ret := []restartEvent{}
	readIdx := cb.LatestIndex
	for i := 0; i < int(cb.MaxItems); i++ {
		value, exists := cb.Items[readIdx]
		if exists {
			ret = append(ret, value)
		}
		readIdx = readIdx + 1
		if readIdx > cb.MaxItems-1 {
			readIdx = readIdx - cb.MaxItems
		}
	}
	return ret
}

// Represent the single service information
type service struct {
	Enabled              bool      // The flag to represent the monitoring status
	Name                 string    // The primary unit name as string
	Description          string    // The human-readable description string
	LoadState            string    // The load state (i.e. whether the unit file has been loaded successfully)
	ActiveState          string    // The active state (i.e. whether the unit is currently started or not)
	SubState             string    // The substate (a more fine-grained version of the active state that is specific to the unit type, which the active state is not)
	LastRecordUpdatetime time.Time // The last time when the record was updated, used for deduplication
	// TODO: Keep either of the one out of below two
	ActiveEnterTimestamp time.Time             // The timestamp when the service entered the active state
	LastRestartTime      time.Time             // The timestamp when the service entered the active state
	RestartHistories     *CircularQueueHistory // The series of restart histories
	RestartEvents        *CircularQueueEvent   // The series of restart events
}
type RestartHistory struct {
	Time           *string `protobuf:"bytes,1,req,name=time" json:"time,omitempty"`
	Reason         *string `protobuf:"bytes,2,req,name=reason" json:"reason,omitempty"`
	JournalEntries *string `protobuf:"bytes,3,req,name=journalEntries" json:"journalEntries,omitempty"`
}

type RestartEvent struct {
	Time *string `protobuf:"bytes,1,req,name=time" json:"time,omitempty"`
}

type SysHealthGetServiceResponse struct {
	Name                 *string           `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	LoadState            *string           `protobuf:"bytes,2,req,name=loadState" json:"loadState,omitempty"`
	ActiveState          *string           `protobuf:"bytes,3,req,name=activeState" json:"activeState,omitempty"`
	SubState             *string           `protobuf:"bytes,4,req,name=subState" json:"subState,omitempty"`
	LastRecordUpdatetime *string           `protobuf:"bytes,5,req,name=lastRecordUpdatetime" json:"lastRecordUpdatetime,omitempty"`
	History              []*RestartHistory `protobuf:"bytes,6,rep,name=history" json:"history,omitempty"`
	Event                []*RestartEvent   `protobuf:"bytes,7,rep,name=event" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func main() {
	cq := NewCircularQueueHistory(3)
	//cqe := NewCircularQueueEvent(10)

	//cqe := CircularQueueEvent{
	//	Items: make(map[int8]restartEvent, 10),
	//}

	item1 := restartHistory{
		Reason: "item1",
	}
	item2 := restartHistory{
		Reason: "item2",
	}
	item3 := restartHistory{
		Reason: "item3",
	}
	item4 := restartHistory{
		Reason: "item4",
	}
	item5 := restartHistory{
		Reason: "item5",
	}

	//item1e := restartEvent{
	//	Time: time.Now().Add(1 * time.Duration(time.Hour)),
	//}
	//item2e := restartEvent{
	//	Time: time.Now().Add(2 * time.Duration(time.Hour)),
	//}
	//item3e := restartEvent{
	//	Time: time.Now().Add(3 * time.Duration(time.Hour)),
	//}
	//item4e := restartEvent{
	//	Time: time.Now().Add(4 * time.Duration(time.Hour)),
	//}
	//item5e := restartEvent{
	//	Time: time.Now().Add(5 * time.Duration(time.Hour)),
	//}

	cq.Add(item1)
	fmt.Printf("values are %v\n", cq.GetData())
	cq.Add(item2)
	fmt.Printf("values are %v\n", cq.GetData())
	cq.Add(item3)
	fmt.Printf("values are %v\n", cq.GetData())
	cq.Add(item4)
	fmt.Printf("values are %v\n", cq.GetData())
	cq.Add(item5)
	fmt.Printf("values are %v\n", cq.GetData())

	var resp SysHealthGetServiceResponse
	someValue := "somevalue"
	resp.Name = &someValue
	resp.LoadState = &someValue
	resp.ActiveState = &someValue
	resp.SubState = &someValue
	timeTmp := time.Now().String()
	resp.LastRecordUpdatetime = &timeTmp
	// TODO: We are looping 2 times, for copying the objects.
	// TODO: We should return an iterator or something and reduce one loop.
	histories := cq.GetData()
	for i := 0; i < len(histories); i++ {
		history := &RestartHistory{
			Time:           &histories[i].Time,
			Reason:         &histories[i].Reason,
			JournalEntries: &histories[i].JournalEntry,
		}
		resp.History = append(resp.History, history)
	}
	//for _, elem := range histories {
	//	history := &RestartHistory{
	//		Time:           &elem.Time,
	//		Reason:         &elem.Reason,
	//		JournalEntries: &elem.JournalEntry,
	//	}
	//	resp.History = append(resp.History, history)
	//}

	fmt.Printf("string format of bytes= %s\n", resp)
	//events := serviceInfo.RestartEvents.GetData()
	//for _, elem := range events {
	//	resp.Event = append(resp.Event, &aviproto.RestartEvent{
	//		Time: &elem.Time,
	//	})
	//}

	//cqe.Add(item1e)
	//fmt.Printf("values are %v\n", cqe.GetData())
	//cqe.Add(item2e)
	//fmt.Printf("values are %v\n", cqe.GetData())
	//cqe.Add(item3e)
	//fmt.Printf("values are %v\n", cqe.GetData())
	//cqe.Add(item4e)
	//fmt.Printf("values are %v\n", cqe.GetData())
	//cqe.Add(item5e)
	//fmt.Printf("values are %v\n", cqe.GetData())

	//service1 := service{
	//	Enabled:              true,
	//	Name:                 "svc name",
	//	Description:          "description",
	//	LastRecordUpdatetime: time.Now(),
	//}
	//
	//if service1.RestartEvents == nil {
	//	println("RestartEvents is null as of now")
	//}
	//
	//service1.RestartEvents = &cqe
	//if service1.RestartEvents.Items == nil {
	//	println("RestartEvents map is null as of now")
	//}

	//service1.RestartHistories.LatestIndex

	//fmt.Printf("service1= %v\n", service1.RestartHistories.LatestIndex)

	//var service2 service
	//
	//service1bytes, _ := json.Marshal(service1)
	//fmt.Printf("string format of bytes= %s\n", string(service1bytes))
	//
	//_ = json.Unmarshal(service1bytes, &service2)
	//
	//fmt.Printf("service2= %v\n", service2)

	name := "node1.controller.local"
	localNodeName := name[:strings.IndexByte(name, '.')]
	println(localNodeName)
}
