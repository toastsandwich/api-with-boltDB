package repository

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type UserRepository struct {
	bucketName []byte
	DB         *bolt.DB
}

func NewUserRepository(db *bolt.DB, bucketName string) (*UserRepository, error) {
	// this looks rediculous
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &UserRepository{
		bucketName: []byte(bucketName),
		DB:         db,
	}, nil
}

func (r *UserRepository) Set(key, value []byte) error {
	return r.DB.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(r.bucketName)
		if bkt == nil {
			return fmt.Errorf("bucket not found: %s", r.bucketName)
		}
		if err := bkt.Put(key, value); err != nil {
			return fmt.Errorf("error in putting values in bucket: %s", err.Error())
		}
		return nil
	})
}

func (r *UserRepository) Get(key []byte) ([]byte, error) {
	var value []byte
	err := r.DB.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(r.bucketName)
		if bkt == nil {
			return fmt.Errorf("bucket not found: %s", r.bucketName)
		}
		value = bkt.Get(key)
		if value == nil {
			return fmt.Errorf("bucket has no user with key: %s", string(key))
		}
		return nil
	})
	return value, err
}
