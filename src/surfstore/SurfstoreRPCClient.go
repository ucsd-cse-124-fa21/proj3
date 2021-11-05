package surfstore

import (
	"net/rpc"
)

type RPCClient struct {
	ServerAddr string
	BaseDir    string
	BlockSize  int
}

func (surfClient *RPCClient) GetBlock(blockHash string, blockStoreAddr string, block *Block) error {
	// connect to the server
	conn, e := rpc.DialHTTP("tcp", blockStoreAddr)
	if e != nil {
		return e
	}

	// perform the call
	e = conn.Call("BlockStore.GetBlock", blockHash, block)
	if e != nil {
		conn.Close()
		return e
	}

	// close the connection
	return conn.Close()
}

func (surfClient *RPCClient) PutBlock(block Block, blockStoreAddr string, succ *bool) error {
	panic("todo")
}

func (surfClient *RPCClient) HasBlocks(blockHashesIn []string, blockStoreAddr string, blockHashesOut *[]string) error {
	panic("todo")
}

func (surfClient *RPCClient) GetFileInfoMap(succ *bool, serverFileInfoMap *map[string]FileMetaData) error {
	panic("todo")
}

func (surfClient *RPCClient) UpdateFile(fileMetaData *FileMetaData, latestVersion *int) error {
	panic("todo")
}

func (surfClient *RPCClient) GetBlockStoreMap(blockHashesIn []string, blockStoreMap *map[string][]string) error {
	panic("todo")
}

var _ ClientInterface = new(RPCClient)

// Create an Surfstore RPC client
func NewSurfstoreRPCClient(hostPort, baseDir string, blockSize int) RPCClient {

	return RPCClient{
		ServerAddr: hostPort,
		BaseDir:    baseDir,
		BlockSize:  blockSize,
	}
}
