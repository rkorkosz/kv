package common

import (
	"io"
	"log"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	bolt "go.etcd.io/bbolt"
)

// Store describes how kv handles database
type Store interface {
	io.Closer
	Get(key string) ([2]string, error)
	GetMany() ([][2]string, error)
	Set(env [2]string) error
	SetMany(envs [][2]string) error
}

// NewStore creates new data store
func NewStore() Store {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	dbPath := filepath.Join(home, ".kv", "kv.db")
	return newBoltStore(dbPath)
}

// BoltStore implements store
type BoltStore struct {
	db     *bolt.DB
	bucket []byte
}

func newBoltStore(dbPath string) Store {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	bucket, err := ProjectName()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	createProjectBucket(db)
	return &BoltStore{db: db, bucket: []byte(bucket)}
}

// Close closes store database
func (bs *BoltStore) Close() error {
	return bs.db.Close()
}

// Get retrieves value for a key
func (bs *BoltStore) Get(key string) ([2]string, error) {
	e := [2]string{key}
	err := bs.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bs.bucket)
		e[1] = string(b.Get([]byte(key)))
		return nil
	})
	return e, err
}

// GetMany fetches all key value pairs for project
func (bs *BoltStore) GetMany() ([][2]string, error) {
	var values [][2]string
	err := bs.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bs.bucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			values = append(values, [2]string{string(k), string(v)})
		}
		return nil
	})
	return values, err
}

// Set creates env in db
func (bs *BoltStore) Set(env [2]string) error {
	return bs.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bs.bucket)
		return b.Put([]byte(env[0]), []byte(env[1]))
	})
}

// SetMany adds multiple envs to db
func (bs *BoltStore) SetMany(envs [][2]string) error {
	return bs.db.Batch(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket(bs.bucket)
		for _, env := range envs {
			err = b.Put([]byte(env[0]), []byte(env[1]))
		}
		return err
	})
}

func createProjectBucket(db *bolt.DB) error {
	name, err := ProjectName()
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(name))
		return err
	})
}
