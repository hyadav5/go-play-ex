package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type restartHistory struct {
	Time         string
	Reason       string
	JournalEntry string
}

type restartHistories struct {
	Latest    int8
	Histories [3]restartHistory
}

type restartEvent struct {
	Time time.Time
}

type restartEvents struct {
	Latest int8
	Events [10]restartEvent
}

func getWriteIndexForLatestIdx(latestIdx int8, factor int8) int8 {
	//latestIndexInt, _ := strconv.Atoi(latestIdx)
	writeIndexInt := latestIdx - 1
	if writeIndexInt < 0 {
		writeIndexInt = factor + writeIndexInt
	}
	return writeIndexInt
	//return strconv.Itoa(writeIndexInt)
}

func (s *service) storeRestartHistories(rh restartHistory) {
	writeIdx := getWriteIndexForLatestIdx(s.RestartHistories.Latest, 3)
	s.RestartHistories.Histories[writeIdx] = rh
	s.RestartHistories.Latest = writeIdx
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
	ActiveEnterTimestamp time.Time        // The timestamp when the service entered the active state
	LastRestartTime      time.Time        // The timestamp when the service entered the active state
	RestartHistories     restartHistories // The series of restart histories
	RestartEvents        restartEvents    // The series of restart events
}

func main() {
	c := make(chan bool, 10)
	go func(c chan bool) {
		var f bool
		ok := true
		for f != true || ok {
			f, ok = <-c
			//if f != nil {
			//	f() // actually do the worker
			//}
		}
	}(c)

	service1 := service{
		Enabled:              true,
		Name:                 "svc name",
		Description:          "description",
		LastRecordUpdatetime: time.Now(),
		RestartHistories: restartHistories{
			Latest: 1,
			Histories: [3]restartHistory{
				{Time: time.Now().String(), Reason: "manual"},
			},
		},
		//RestartEvents: restartEvents{
		//	Latest: 1,
		//	Events: []restartEvent{},
		//},
	}

	fmt.Printf("service1= %v\n", service1)

	var service2 service

	service1bytes, _ := json.Marshal(service1)
	fmt.Printf("string format of bytes= %s\n", string(service1bytes))

	_ = json.Unmarshal(service1bytes, &service2)

	fmt.Printf("service2= %v\n", service2)

	service2.storeRestartHistories(restartHistory{Reason: "something 1"})
	service2.storeRestartHistories(restartHistory{Reason: "something 2"})
	fmt.Printf("service2= %v\n", service2)
}
