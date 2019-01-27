package queue

//TODO: for empty list skip size of theindex file

import (
	"io"
	"time"

	"github.com/mixflowtech/go-librt/logger"
)

// QueueItem is elementh of the queue
type QueueItem struct { // nolint
	idx     StorageIdx
	ID      StorageIdx
	Stream  io.ReadSeeker
	storage storageProcessing
}

//Queue is a base structure for managing of the messages
type Queue struct {
	name         string
	options      *Options
	log          logger.Logger
	newMessage   chan struct{}
	stopEvent    chan struct{}
	stopedHandle chan struct{}
	//storage      *fileStorage
	//memory       *queueMemory
	//factory      WorkerFactory
	//inProcess    *inProcessingPerWorker
	total        int32
	//totalWorkers int32
	lastTimeGC   time.Duration
}

type newMessageNotificator interface {
	newMessageNotification()
}

func CreateQueue(Name, StoragePath string, Log logger.Logger, Options *Options) (*Queue, error) {
	return nil, nil
}

// Queue as newMessageNotificator
func (q *Queue) newMessageNotification() {
	select {
	case q.newMessage <- struct{}{}:
	default:
	}
}

// Member function of Queue
//Count returns the count of the messages in the queue
func (q *Queue) Count() uint64 {
	return 0
	//return q.storage.Count() + q.memory.Count()
}

// Insert appends the message into the queue. In depends of the timeout's option either is trying
// to write message to the disk or is trying to process this message in the memory and writing to the
// disk only if timeout is expired shortly. Returns false if aren't processing / writing of the message
// in the during of the timeout or has some problems with  writing to disk
func (q *Queue) Insert(buf []byte) bool {
	return q.insert(buf, nil)
	// after timeout, then write to disk ? as archived/libqueue?
}

func (q *Queue) insert(buf []byte, ch chan bool) bool {
	if ch == nil {
		/*
		ID, err := q.storage.Put(buf)

		if err == nil {
			q.log.Trace("[Q:%s:%d] Stored to file storage", q.name, ID)
			q.newMessageNotification()
		} else {
			q.log.Error("[Q:%s:%d] Storing to storage with error result [%s] ", q.name, ID, err.Error())
		}
		return err == nil
		*/
	}
	/*
	if q.storage.Count() == 0 {
		if ID, err := q.memory.Put(buf, ch); err == nil {
			q.log.Trace("[Q:%s:%d] Stored to memory storage", q.name, ID)
			q.newMessageNotification()
			return true
		}
	}
	*/
	/*
	ID, err := q.storage.Put(buf)
	if err == nil {
		q.log.Trace("[Q:%s:%d] Stored to file storage", q.name, ID)
		q.newMessageNotification()
	} else {
		q.log.Error("[Q:%s:%d] Storing to storage with error result [%s] ", q.name, ID, err.Error())
	}
	ch <- err == nil
	return err == nil
	*/
	return false
}

func (q *Queue) close() {
	q.stopEvent <- struct{}{}
	<-q.stopedHandle
}

/*
func (q *Queue) info() {
	// q.storage.info()
}
*/

// Close stops the handler of the messages, saves the messages located in
// the memory into the disk, closes all opened files.
func (q *Queue) Close() {
	q.log.Info("[Q:%s] is closed...", q.name)
	q.close()
	q.log.Info("[Q:%s] was closed...", q.name)
}
