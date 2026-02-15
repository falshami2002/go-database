package btree

// Split an oversized node into 2, 2nd always fits into one page
func nodeSplit2(left BNode, right BNode, old BNode) {
	nkeys := old.nkeys()
	assert(nkeys >= 2)

	nbytes := old.nbytes()

	mid := uint16(0)

	for i := uint16(1); i < nkeys; i++ {
		rightBytes := nbytes - old.kvPos(i)
		if rightBytes <= BTREE_PAGE_SIZE {
			mid = i
			break
		}
	}

	assert(mid > 0 && mid < nkeys)

	leftCount := mid
	rightCount := nkeys - mid

	btype := old.btype()

	left.setHeader(btype, leftCount)
	nodeAppendRange(left, old, 0, 0, leftCount)

	right.setHeader(btype, rightCount)
	nodeAppendRange(right, old, 0, mid, rightCount)
}

// split a node if it's too big. the results are 1~3 nodes.
func nodeSplit3(old BNode) (uint16, [3]BNode) {
	if old.nbytes() <= BTREE_PAGE_SIZE {
		old = old[:BTREE_PAGE_SIZE]
		return 1, [3]BNode{old} // not split
	}
	left := BNode(make([]byte, 2*BTREE_PAGE_SIZE)) // might be split later
	right := BNode(make([]byte, BTREE_PAGE_SIZE))
	nodeSplit2(left, right, old)
	if left.nbytes() <= BTREE_PAGE_SIZE {
		left = left[:BTREE_PAGE_SIZE]
		return 2, [3]BNode{left, right} // 2 nodes
	}
	leftleft := BNode(make([]byte, BTREE_PAGE_SIZE))
	middle := BNode(make([]byte, BTREE_PAGE_SIZE))
	nodeSplit2(leftleft, middle, left)
	assert(leftleft.nbytes() <= BTREE_PAGE_SIZE)
	return 3, [3]BNode{leftleft, middle, right} // 3 nodes
}
