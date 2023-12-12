package buffer

import (
	"bulwark/config"
	"fmt"
	"time"

	bconst "github.com/jfcarter2358/bulwark/bulwarkmp/v1/constants"

	"github.com/jfcarter2358/go-logger"
)

type Buffer struct {
	Data     interface{} `json:"data"`
	LastPing int         `json:"last_ping"`
}

var buffers map[string]Buffer
var locks map[string]bool

const BUFFER_LOCK_WAIT = 500
const BUFFER_PRUNE_WAIT = 1000

func Init() error {
	buffers = make(map[string]Buffer)
	locks = make(map[string]bool)
	if config.Config.BufferLifetime > -1 {
		go pruneBuffer()
	}
	return nil
}

func pruneBuffer() {
	for {
		for name := range buffers {
			for locks[name] {
				time.Sleep(BUFFER_LOCK_WAIT * time.Millisecond)
			}
			locks[name] = true
			b := buffers[name]
			if b.LastPing > config.Config.BufferLifetime {
				logger.Debugf("", "Deleting buffer %s", name)
				delete(buffers, name)
				delete(locks, name)
				continue
			}
			b.LastPing += 1
			buffers[name] = b
			locks[name] = false
		}
		time.Sleep(BUFFER_PRUNE_WAIT * time.Millisecond)
	}
}

func Create(name string) error {
	buffers[name] = Buffer{
		Data:     make([]interface{}, 0),
		LastPing: 0,
	}
	locks[name] = false
	return nil
}

func Delete(name string) error {
	for locks[name] {
		time.Sleep(BUFFER_LOCK_WAIT * time.Millisecond)
	}
	if _, ok := locks[name]; !ok {
		return fmt.Errorf("cannot push to buffer %s: buffer has been deleted", name)
	}
	locks[name] = true
	delete(buffers, name)
	delete(locks, name)
	locks[name] = false
	return nil
}

func Set(name string, datum string) error {
	for locks[name] {
		time.Sleep(BUFFER_LOCK_WAIT * time.Millisecond)
	}
	if _, ok := locks[name]; !ok {
		return fmt.Errorf("cannot push to buffer %s: buffer has been deleted", name)
	}
	locks[name] = true
	b := buffers[name]
	b.Data = datum
	b.LastPing = 0
	buffers[name] = b
	locks[name] = false
	return nil
}

func Get(name string) (string, string, error) {
	for locks[name] {
		time.Sleep(BUFFER_LOCK_WAIT * time.Millisecond)
	}
	if _, ok := locks[name]; !ok {
		return nil, fmt.Errorf("cannot pop from buffer %s: buffer has been deleted", name)
	}
	locks[name] = true
	b := buffers[name]
	b.LastPing = 0
	buffers[name] = b
	out := b.Data.(string)
	locks[name] = false
	return bconst.CONTENT_TYPE_TEXT, out, nil
}
