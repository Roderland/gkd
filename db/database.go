package db

import (
	"gkd/ds"
	"gkd/ds/impl"
	"gkd/entry"
	"gkd/fr"
	"io"
	"os"
)

type Database struct {
	fileReader *fr.FileReader
	strDict    ds.Str
	listDict   ds.List
}

func NewDatabase(fileReader *fr.FileReader) *Database {
	return &Database{
		fileReader: fileReader,
		strDict:    impl.NewStrMap(),
		listDict:   impl.NewListMap(),
	}
}

func OpenDatabase(path string, name string) (database *Database, err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return
		}
	}
	file, err := fr.NewFileReader(path, name)
	if err != nil {
		return
	}
	database = NewDatabase(file)
	err = database.loadData()
	return
}

func (db *Database) loadData() (err error) {
	if db.fileReader == nil {
		return
	}
	for {
		var mark uint16
		mark, err = db.fileReader.ReadMark()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
		var keySize uint32
		keySize, err = db.fileReader.ReadKeySize()
		if err != nil {
			return
		}
		var key []byte
		key, err = db.fileReader.ReadKey(keySize)
		if err != nil {
			return err
		}
		baseEntry := entry.NewBaseEntry(mark, keySize, key)
		switch mark {
		case entry.StrMark:
			var valueSize uint32
			valueSize, err = db.fileReader.ReadStrValueSize()
			if err != nil {
				return
			}
			var value []byte
			value, err = db.fileReader.ReadStrValue(valueSize)
			if err != nil {
				return
			}
			strEntry := entry.NewStrEntry(baseEntry, valueSize, value)
			db.strDict.Set(string(strEntry.Base.Key), string(strEntry.Value))
		case entry.ListMark:

		}
	}
}

func (db *Database) SaveData(path, name string) (err error) {
	reader, err := fr.NewFileReader(path, name)
	if err != nil {
		return err
	}
	var offset int64 = 0
	bytes := db.strDict.ToBytes()
	_, err = reader.File.WriteAt(bytes, offset)
	if err != nil {
		return err
	}
	return nil
}
