# Surfstore

This is the starter code for Module 3: Surfstore. 

Before you get started, make sure you understand the following 2 things about Go. (These will also be covered in class and in discussions)
1. Interfaces: They are named collections of method signatures. Here are some good resources to understand interfaces in Go:
    a. https://gobyexample.com/interfaces
    b. https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

2. RPC: You should know how to write RPC servers and clients in Go. The [online documentation](https://golang.org/pkg/net/rpc/) of the *rpc* package is a good resource. 

## Data Types

Recall from the module write-up the following things:

1. The SurfStore service is composed of two services: BlockStore and MetadataStore 
2. A file in SurfStore is broken into an ordered sequence of one or more blocks which are stored in the BlockStore.
3. The MetadataStore maintains the mapping of filenames to hashes of these blocks (and versions) in a map.

The starter code defines the following types for your usage in `SurfstoreInterfaces.go`:

```go
type Block struct {
	BlockData []byte
	BlockSize int
}

type FileMetaData struct {
	Filename      string
	Version       int
	BlockHashList []string
}
```

## Surfstore Interface

`SurfstoreInterfaces.go` also contains interfaces for the BlockStore and the MetadataStore:

```go
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
```

## Server
`BlockStore.go` provides a skeleton implementation of the `BlockStoreInterface` and `MetaStore.go` provides a skeleton implementation of the `MetaStoreInterface` 
**You must implement the methods in these 2 files which have `panic("todo")` as their body.**

`init/SurfstoreServerExec/main.go` also has a method `startServer` **which you must implement**. Depending on the service type specified, it should register a `MetaStore`, `BlockStore`, or `Both` and start listening for connections from clients.

## Client
`SurfstoreRPCClient.go` provides the rpc client stub for the surfstore rpc server. **You must implement the methods in this file which have `panic("todo")` as their body.** (Hint: one of them has been implemented for you) 

`SurfstoreClientUtils.go` also has the following method which **you need to implement** for the sync logic of clients:
```go
/*
Implement the logic for a client syncing with the server here.
*/
func ClientSync(client RPCClient) {
	panic("todo")
}
```
## Setup
You will need to setup your runtime environment variables so that you can build your code and also use the executables that will be generated.
1. If you are using a Mac, open `~/.bash_profile` or if you are using a unix/linux machine, open `~/.bashrc`. Then add the following:
```
export GOPATH=<path to starter code>
export PATH=$PATH:$GOPATH/bin
export GO111MODULE=off
```
2. Run `source ~/.bash_profile` or `source ~/.bashrc`
## Usage
1. Only after you have implemented all the methods and completed the `Setup` steps, run the `build.sh` script provided with the starter code. This should create 2 executables in the `bin` folder inside your starter code directory.
```shell
> ./build.sh
> ls bin
SurfstoreClientExec SurfstoreServerExec
```

2. Run your server using the script provided in the starter code.
```shell
./run-server.sh -s <service> -p <port> -l -d (BlockStoreAddr*)
```
Here, `service` should be one of three values: meta, block, or both. This is used to specify the service provided by the server. `port` defines the port number that the server listens to (default=8080). `-l` configures the server to only listen on localhost. `-d` configures the server to output log statements. Lastly, (BlockStoreAddr\*) is zero or more initial BlockStore addresses that the server is configured with. For module 3, the MetaStore should always start with 1 BlockStore address and if `service=both` then the BlockStoreAddr should be the `ip:port` of this server.

Examples:
```shell
./run-server.sh -s both -p 8070 -l localhost:8070
```
This starts a server that listens only to localhost on port 8070 and services both the BlockStore and MetaStore interface.

```shell
Run the commands below on separate terminals (or nodes)
> ./run-server.sh -s block -p 8081 -l
> ./run-server.sh -s meta -l localhost:8081
```
The first line starts a server that services only the BlockStore interface and listens only to localhost on port 8081. The second line starts a server that services only the MetaStore interface, listens only to localhost on port 8080, and references the BlockStore we created as the underlying BlockStore. (Note: if these are on separate nodes, then you should use the public ip address and remove `-l`)

3. From a new terminal (or a new node), run the client using the script provided in the starter code (if using a new node, build using step 1 first). Use a base directory with some files in it.
```shell
> mkdir dataA
> cp ~/pic.jpg dataA/ 
> ./run-client.sh server_addr:port dataA 4096
```
This would sync pic.jpg to the server hosted on `server_addr:port`, using `dataA` as the base directory, with a block size of 4096 bytes.

4. From another terminal (or a new node), run the client to sync with the server. (if using a new node, build using step 1 first)
```shell
> ls dataB/
> ./run-client.sh server_addr:port dataB 4096
> ls dataB/
pic.jpg index.txt
```
We observe that pic.jpg has been synced to this client.

## Testing 
On gradescope, only a subset of test cases will be visible, so we highly encourage you to come up with different scenarios like the one described above. You can then match the outcome of your implementation to the expected output based on the theory provided in the writeup.

To avoid autograder configuration issues, we also encourage you to test your solution on AWS which is likely configured with the same go version (1.14.6). This also means that your solution should avoid using methods and packages introduced after go 1.14.6.

