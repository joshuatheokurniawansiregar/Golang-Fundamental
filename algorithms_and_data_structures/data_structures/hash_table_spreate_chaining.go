package data_structures

// Hash Table Part(Arry):

// ArraySize is the size of the hashtable
const ArraySize = 7

// HashTable structure will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

func NewHashTable() *HashTable {
	return new(HashTable)
}

//Init: Initialization
func Init() *HashTable {
	var result *HashTable = NewHashTable()
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key nd return true if tht key is stored in the hash table
func (h *HashTable) Searh(key string) bool {
	return false
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) bool {
	return false
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Bucket Part(Linked List)

// bucketNode structure is a linked list node that holds another bucketNode
type bucketNode struct {
	key  string
	next *bucketNode
}

// bucket structure is a linked list in each slot of the hash table
type bucket struct {
	head *bucketNode
	port string
}

var Bucket bucket

func NewBucket() *bucket {
	return &bucket{}
}

// insert() will take in a key, create a node with key and insert the node in the bucket
func (b *bucket) insert(k string) {
	var newBucketNode *bucketNode = new(bucketNode)
	newBucketNode.key = k
	newBucketNode.next = b.head
	b.head = newBucketNode
}

// delete()
// search()
