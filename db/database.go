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
			var listLen uint32
			listLen, err = db.fileReader.ReadListLen()
			if err != nil {
				return
			}
			var i uint32 = 0
			var valuesSize []uint32 = make([]uint32, listLen)
			var values [][]byte = make([][]byte, listLen)
			for ; i < listLen; i++ {
				valuesSize[i], err = db.fileReader.ReadListValueSize()
				if err != nil {
					return
				}
				values[i], err = db.fileReader.ReadListValue(valuesSize[i])
				if err != nil {
					return
				}
			}
			db.listDict.RPush(string(baseEntry.Key), values)
		}
	}
}

func (db *Database) SaveData(path, name string) (err error) {
	writer, err := fr.NewFileWriter(path, name)
	if err != nil {
		return err
	}
	var offset int64 = 0
	var bytes []byte

	// write str entry
	bytes = db.strDict.ToBytes()
	_, err = writer.File.WriteAt(bytes, offset)
	offset += int64(len(bytes))
	// write list entry
	bytes = db.listDict.ToBytes()
	_, err = writer.File.WriteAt(bytes, offset)
	offset += int64(len(bytes))

	if err != nil {
		return err
	}
	return nil
}
