package btree

type BTree struct {
	root uint64              // page number
	get  func(uint64) []byte // dereference pointer
	new  func([]byte) uint64 //allocate page
	del  func(uint64)        //deallocate page
}

type BNode []byte
