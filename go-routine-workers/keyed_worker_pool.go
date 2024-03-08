package main

import "hash/fnv"

/* A keyed worker pool is a worker pool that guarantees that tasks with the same key are run in order*/
type KeyedWorkerPool struct {
	workers []chan func()
}

// Note that true "Length" in this case is numWorkers * maxQueueLength, as we make one channel of size maxQueueLength
// per Worker
func MakeKeyedWorkerPool(numWorkers int, maxQueueLength int) *KeyedWorkerPool {
	var workers []chan func()
	for i := 0; i < numWorkers; i++ {
		c := make(chan func(), maxQueueLength)
		go func(c chan func()) {
			var f func()
			ok := true
			for f != nil || ok {
				f, ok = <-c
				if f != nil {
					f() // actually do the worker
				}
			}
		}(c)
		workers = append(workers, c)
	}
	return &KeyedWorkerPool{workers}
}

func (kp *KeyedWorkerPool) Close() {
	for _, c := range kp.workers {
		close(c)
	}
}

func (kp *KeyedWorkerPool) DoWork(key string, f func()) {
	kp.workers[kp.getIDForKey(key)] <- f
}

func (kp *KeyedWorkerPool) getIDForKey(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32() % uint32(len(kp.workers))
}

func (kp *KeyedWorkerPool) Length() int {
	total := 0
	for _, c := range kp.workers {
		total += len(c)
	}
	return total
}

func (kp *KeyedWorkerPool) GetWorkerID(key string) uint32 {
	return kp.getIDForKey(key)
}

func (kp *KeyedWorkerPool) WorkerChannelLength(workerID uint32) int {
	return len(kp.workers[workerID])
}
