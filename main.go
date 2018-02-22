package multidbfile

import (
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

type top struct {
	dbs      []*bolt.DB
	sequence int
}

const bucketkey = "blocks"

func dbname(sequence int) string {
	return "test" + strconv.Itoa(sequence) + ".db"
}

func (t *top) createDB() error {
	t.sequence++
	db, err := bolt.Open(dbname(t.sequence), 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		t.sequence--
		panic(err)
	}
	t.dbs = append(t.dbs, db)
	return nil
}

func (t *top) find(key []byte) []byte {
	var value []byte
	var err error
	for _, db := range t.dbs {
		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucketkey))
			value = b.Get(key)
			return nil
		})
		if value != nil {
			break
		}
	}
	if err != nil {
		return nil
	}
	return value
}

func (t *top) exist(key []byte) bool {
	return t.find(key) != nil
}

func (t *top) store(key []byte, value []byte) error {
	db := t.dbs[len(t.dbs)-1]
	err := db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(bucketkey))
		err := b.Put(key, value)
		return err
	})
	return err
}

func (t *top) update(key []byte, value []byte) error {
	var found bool
	var err error
	for _, db := range t.dbs {
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucketkey))
			v := b.Get(key)
			if v != nil {
				err := b.Put(key, value)
				found = true
				return err
			}
			return nil
		})
		if found {
			break
		}
	}
	return err
}

func (t *top) del(key []byte) error {
	var found bool
	var err error
	for _, db := range t.dbs {
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucketkey))
			v := b.Get(key)
			if v != nil {
				err := b.Delete(key)
				found = true
				return err
			}
			return nil
		})
		if found {
			break
		}
	}
	return err
}

func newTop() *top {
	return &top{
		dbs:      make([]*bolt.DB, 0),
		sequence: 0,
	}
}
