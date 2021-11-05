package surfstore

type Block struct {
	BlockData []byte
	BlockSize int
}

type FileMetaData struct {
	Filename      string
	Version       int
	BlockHashList []string
}

type MetaStoreInterface interface {
	// Retrieves the server's FileInfoMap
	GetFileInfoMap(_ignore *bool, serverFileInfoMap *map[string]FileMetaData) error

	// Update a file's fileinfo entry
	UpdateFile(fileMetaData *FileMetaData, latestVersion *int) (err error)

	// Retrieve the mapping of BlockStore addresses to block hashes
	GetBlockStoreMap(blockHashesIn []string, blockStoreMap *map[string][]string) error
}

type BlockStoreInterface interface {

	// Get a block based on its hash
	GetBlock(blockHash string, block *Block) error

	// Put a block
	PutBlock(block Block, succ *bool) error

	// Check if certain blocks are alredy present on the server
	HasBlocks(blockHashesIn []string, blockHashesOut *[]string) error
}

type ClientInterface interface {
	// MetaStore
	GetFileInfoMap(_ignore *bool, serverFileInfoMap *map[string]FileMetaData) error
	UpdateFile(fileMetaData *FileMetaData, latestVersion *int) (err error)
	GetBlockStoreMap(blockHashesIn []string, blockStoreMap *map[string][]string) error

	// BlockStore
	GetBlock(blockHash string, blockStoreAddr string, block *Block) error
	PutBlock(block Block, blockStoreAddr string, succ *bool) error
	HasBlocks(blockHashesIn []string, blockStoreAddr string, blockHashesOut *[]string) error
}
