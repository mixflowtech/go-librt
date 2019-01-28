package queue

//TODO: for empty list skip size of theindex file

import (
	"encoding/binary"
	"errors"
	//"fmt"

	//"encoding/json"
	///"fmt"
	"io"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"

	"github.com/mixflowtech/go-librt/logger"
)

var startTime = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

// QueueItem is elementh of the queue
type QueueItem struct { // nolint
	//idx     StorageIdx
	//ID      StorageIdx
	Stream  io.ReadSeeker
	//storage storageProcessing
}

//Queue is a base structure for managing of the messages
type Queue struct {
	name         string
	options      *Options
	log          logger.Logger
	newMessage   chan struct{}
	stopEvent    chan struct{}
	stopedHandle chan struct{}
	db 			 *bolt.DB
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

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func uitob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func CreateQueue(Name, StoragePath string, Log logger.Logger, Options *Options) (*Queue, error) {
	db, err := bolt.Open(StoragePath, 0666, nil)
	if err != nil {
		return nil, err
	}

	tmp := &Queue{
		total:        0,
		stopedHandle: make(chan struct{}),
		newMessage:   make(chan struct{}, 1),
		log:          Log,
		options:      Options,
		name:         Name,
		stopEvent:    make(chan struct{}),
		lastTimeGC:   time.Since(startTime),
		db:			  db,
	}
	return tmp, nil
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
	var count uint64 = 0

	if err := q.db.View(func(tx *bolt.Tx) error {
		// Create a new bucket.
		b := tx.Bucket([]byte(q.name))
		if b == nil {
			return errors.New("Bucked not found.")
		}
		count = b.Sequence()
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return count
}

// Insert appends the message into the queue. In depends of the timeout's option either is trying
// to write message to the disk or is trying to process this message in the memory and writing to the
// disk only if timeout is expired shortly. Returns false if aren't processing / writing of the message
// in the during of the timeout or has some problems with  writing to disk
func (q *Queue) Insert(buf []byte) error {
	return q.insert(buf, nil)
	// after timeout, then write to disk ? as archived/libqueue?
}

func (q *Queue) insert(buf []byte, ch chan bool) error {
	// FIXME: add batch insert mode.
	// in bbolt , use batch update
	if err := q.db.Update(func(tx *bolt.Tx) error {
		// Create a new bucket.
		b, err := tx.CreateBucketIfNotExists([]byte(q.name))
		if err != nil {
			return err
		}

		id, _ := b.NextSequence()
		rec_no := int(id)
		// Persist bytes to users bucket.
		return b.Put(itob(rec_no), buf)
	}); err != nil {
		log.Fatal(err)
	}


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
	return nil
}

type FetchItemCb func(buf []byte) error

func (q *Queue) Fetch(offset uint64, count uint64, cb FetchItemCb) error {
	// FIXME:  unstable API
	if err := q.db.View(func(tx *bolt.Tx) error {
		// Create a new bucket.
		b := tx.Bucket([]byte(q.name))
		if b == nil {
			return errors.New("Bucked not found.")
		}
		// Create a cursor for iteration.
		c := b.Cursor()

		// Iterate over items in sorted key order. This starts from the
		// first key/value pair and updates the k/v variables to the
		// next key/value on each iteration.
		//
		// The loop finishes at the end of the cursor when a nil key is returned.
		for k, v := c.Seek(uitob(offset)); k != nil; k, v = c.Next() {
			cb_e := cb(v)
			if cb_e != nil {
				return cb_e
			}
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (q *Queue) close() {
	//q.stopEvent <- struct{}{}
	//<-q.stopedHandle
	q.db.Close()
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
