package surfstore

type MetaStore struct {
	FileMetaMap map[string]FileMetaData
	// Add additional data structure(s) to maintain BlockStore addresses
}

func (m *MetaStore) GetFileInfoMap(_ignore *bool, serverFileInfoMap *map[string]FileMetaData) error {
	panic("todo")
}

func (m *MetaStore) UpdateFile(fileMetaData *FileMetaData, latestVersion *int) (err error) {
	panic("todo")
}

func (m *MetaStore) GetBlockStoreMap(blockHashesIn []string, blockStoreMap *map[string][]string) error {
	panic("todo")
}

var _ MetaStoreInterface = new(MetaStore)

func NewMetaStore(blockStoreList []string) MetaStore {
	panic("todo")
}
