package storage

import "github.com/syndtr/goleveldb/leveldb"

type LevelDB struct {
	db   *leveldb.DB
	path string
}

func NewLevelDB(path string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDB{db, path}, nil
}

func (ldb *LevelDB) Get(key []byte) ([]byte, error) {
	return ldb.db.Get(key, nil)
}

func (ldb *LevelDB) Put(key, value []byte) error {
	return ldb.db.Put(key, value, nil)
}

func (ldb *LevelDB) Has(key []byte) bool {
	has, _ := ldb.db.Has(key, nil)
	return has
}

func (ldb *LevelDB) Delete(key []byte) error {
	return ldb.db.Delete(key, nil)
}

func (ldb *LevelDB) BatchPut(kvs [][2][]byte) error {
	batch := new(leveldb.Batch)
	for i := range kvs {
		batch.Put(kvs[i][0], kvs[i][1])
	}
	return ldb.db.Write(batch, nil)
}

func (ldb *LevelDB) Open() error {
	var err error
	ldb.db, err = leveldb.OpenFile(ldb.path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (ldb *LevelDB) Close() error {
	return ldb.db.Close()
}

func (ldb *LevelDB) GetIterator() IIterator {
	return ldb.db.NewIterator(nil, nil)
}
