package main

import "sync"

// WorkerPool: All methods of WorkerPool are goroutine-safe.
type WorkerPool struct {
	// work queue holds the tasks to be performed.
	work chan func()

	// wg supports the Wait() implementation
	wg sync.WaitGroup

	// mutex helps prevent race conditions.
	// For ex: A concurrent invocation of DoWork & Shutdown would result in predictable behaviour depending on which is performed first.
	// - DoWork before Shutdown would push the task into queue and it will be performed as Shutdown would wait for this task to complete.
	// - DoWork after Shutdown would cause panic.
	mutex sync.Mutex
}

// MakeWorkerPool: return an WorkerPool that performs tasks concurrently with a pool of go routines.
// * Spawns numWorkers count of goroutine that processes the tasks from queue.
// * The goroutines are stopped on invocation of Close() or Shutdown()
// * Queue accumulates tasks when all workers are in use.
func MakeWorkerPool(numWorkers, maxQueueLength int) *WorkerPool {
	m := WorkerPool{
		work: make(chan func(), maxQueueLength),
	}
	for i := 0; i < numWorkers; i++ {
		go func() {
			var f func()
			ok := true
			for f != nil || ok {
				f, ok = <-m.work
				if f != nil {
					f() // actually do the worker
					m.wg.Done()
				}
			}
		}()
	}
	return &m
}

func (m *WorkerPool) close() {
	close(m.work)
}

// Close closes the task queue and won't accept new tasks.
// * In-progress tasks continue to run.
// * Pending tasks are canceled.
// * Submitting new tasks will cause panic.
// * Use Shutdown() to wait for in-progress tasks and close.
func (wp *WorkerPool) Close() {
	wp.mutex.Lock()
	defer wp.mutex.Unlock()
	wp.close()
}

func (wp *WorkerPool) doWork(task func()) {
	wp.wg.Add(1)
	wp.work <- task
}

// DoWork pushes the task into queue.
// * When max queue size is reached, submitting new task is a blocking call until a previous task is completed.
func (m *WorkerPool) DoWork(task func()) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.doWork(task)
}

// Length returns the size of pending tasks.
func (m *WorkerPool) Length() int {
	return len(m.work)
}

func (m *WorkerPool) wait() {
	m.wg.Wait()
}

// Wait waits till all go routines are completed.
// * New tasks are still accepted.
func (m *WorkerPool) Wait() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.wait()
}

// Shutdown waits till all tasks are completed then closes the task queue.
// * Submitting a new task after invoking shutdown will cause a panic.
func (m *WorkerPool) Shutdown() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.wait()
	m.close()
}
