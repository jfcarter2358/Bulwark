package queue

import (
	"bulwark/config"
	"fmt"
	"time"
)

type Queue struct {
	Data     []interface{} `json:"data"`
	LastPing int           `json:"last_ping"`
}

var queues map[string]Queue
var locks map[string]bool

const QUEUE_LOCK_WAIT = 500
const QUEUE_PRUNE_WAIT = 1000

func Init() error {
	queues = make(map[string]Queue)
	locks = make(map[string]bool)
	if config.Config.QueueLifetime > -1 {
		go pruneQueue()
	}
	return nil
}

func pruneQueue() {
	for {
		for name := range queues {
			for locks[name] {
				time.Sleep(QUEUE_LOCK_WAIT * time.Millisecond)
			}
			locks[name] = true
			q := queues[name]
			if q.LastPing > config.Config.QueueLifetime {
				delete(queues, name)
				delete(locks, name)
				continue
			}
			q.LastPing += 1
			queues[name] = q
			locks[name] = false
		}
		time.Sleep(QUEUE_LOCK_WAIT * time.Millisecond)
	}
}

func Create(name string) error {
	queues[name] = Queue{
		Data:     make([]interface{}, 0),
		LastPing: 0,
	}
	locks[name] = false
	return nil
}

func Delete(name string) error {
	for locks[name] {
		time.Sleep(QUEUE_LOCK_WAIT * time.Millisecond)
	}
	if _, ok := locks[name]; !ok {
		return fmt.Errorf("cannot push to queue %s: queue has been deleted", name)
	}
	locks[name] = true
	delete(queues, name)
	delete(locks, name)
	locks[name] = false
	return nil
}

func Push(name string, datum interface{}) error {
	for locks[name] {
		time.Sleep(QUEUE_LOCK_WAIT * time.Millisecond)
	}
	if _, ok := locks[name]; !ok {
		return fmt.Errorf("cannot push to queue %s: queue has been deleted", name)
	}
	locks[name] = true
	q := queues[name]
	q.Data = append(q.Data, datum)
	queues[name] = q
	locks[name] = false
	return nil
}

func Pop(name string) (string, string, error) {
	var out interface{}
	var err error

	for locks[name] {
		time.Sleep(QUEUE_LOCK_WAIT * time.Millisecond)
	}
	if _, ok := locks[name]; !ok {
		return nil, fmt.Errorf("cannot pop from queue %s: queue has been deleted", name)
	}
	locks[name] = true
	q := queues[name]
	if len(q.Data) > 1 {
		out = q.Data[0].(string)
		q.Data = q.Data[1:]
		q.LastPing = 0
		queues[name] = q
	}
	locks[name] = false
	return bmp.constants.CONTENT_TYPE_TEXT, out, err
}
