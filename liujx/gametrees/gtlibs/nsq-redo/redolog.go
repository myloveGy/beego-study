package nsqredo

import (
	"bytes"
	log "gametrees/gtlibs/nsq-logger"
	"gopkg.in/vmihailenco/msgpack.v2"
	"net/http"
	"os"
)

const (
	ENV_NSQD         = "NSQD_HOST"
	DEFAULT_PUB_ADDR = "http://172.17.42.1:4151/pub?topic=REDOLOG"
	MIME             = "application/octet-stream"
)

// a data change
type Change struct {
	Collection string      // represents document name
	SubDoc     string      // represents subdocument "a.b.c.1.d"
	Old        interface{} // value before change
	New        interface{} // value after change
}

// a redo record represents complete transaction
type RedoRecord struct {
	Api     string // the api name
	Uid     int32  // userid
	Changes []Change
	TS      uint64 // timestamp should get from snowflake
}

var (
	_pub_addr string
	_prefix   string
)

func init() {
	// get nsqd publish address
	_pub_addr = DEFAULT_PUB_ADDR
	if env := os.Getenv(ENV_NSQD); env != "" {
		_pub_addr = env + "/pub?topic=LOG"
	}
}

// add a change with o(old value) and n(new value)
func (r *RedoRecord) AddChange(collection, subdoc string, o interface{}, n interface{}) {
	r.Changes = append(r.Changes, Change{Collection: collection, SubDoc: subdoc, Old: o, New: n})
}

// set timestamp
func (r *RedoRecord) SetTS(ts uint64) {
	r.TS = ts
}

// set api
func (r *RedoRecord) SetApi(name string) {
	r.Api = name
}

// set user id
func (r *RedoRecord) SetUid(uid int32) {
	r.Uid = uid
}

func NewRedoRecord() *RedoRecord {
	return new(RedoRecord)
}

// publish to nsqd (localhost nsqd is suggested!)
func Publish(r *RedoRecord) {
	// pack message
	pack, err := msgpack.Marshal(r)
	if err != nil {
		log.Critical(err)
		return
	}

	// post to nsqd
	resp, err := http.Post(_pub_addr, MIME, bytes.NewReader(pack))
	if err != nil {
		log.Critical(err)
		return
	}
	defer resp.Body.Close()
}
