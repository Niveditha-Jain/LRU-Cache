package main

// DO NOT CHANGE THIS CACHE SIZE VALUE
const CACHE_SIZE int = 3

var lruMap map[int]node = make(map[int]node)
var head *node
var end *node

type node struct {
	key      int
	value    int
	previous *node
	next     *node
}

func Set(key int, value int) {

	if setNode, ok := lruMap[key]; ok {
		setNode.value = value
		lruMap[setNode.key] = setNode
		Get(setNode.key)
	} else {

		createdNode := node{key: key, value: value}
		if CACHE_SIZE == len(lruMap) {
			removeLast()
		}

		//fmt.Println("Cache size:", len(cache.lruMap))
		if len(lruMap) == 0 {
			createdNode.next = nil
			createdNode.previous = nil
			lruMap[key] = createdNode
			head = &createdNode
			end = &createdNode
		} else {
			var tempHead *node = head
			var previousNode = lruMap[head.key]
			previousNode.previous = &createdNode
			lruMap[previousNode.key] = previousNode
			createdNode.next = tempHead
			createdNode.previous = nil
			lruMap[key] = createdNode
			head = &createdNode
		}
	}
}

func Get(key int) int {

	if getNode, ok := lruMap[key]; ok {

		//fmt.Println("Get Node",getNode.key)
		// if key is not at head
		if head.key != key {

			// if key is at the end
			if end.key == key {
				lastNode := lruMap[end.key]
				end = lastNode.previous

				var secondLast = lruMap[lastNode.previous.key]
				secondLast.next = nil
				lruMap[secondLast.key] = secondLast
				end.next = nil
			}

			if getNode.next != nil {
				nextNode := lruMap[getNode.next.key]
				nextNode.previous = getNode.previous
				lruMap[nextNode.key] = nextNode
			}
			if getNode.previous != nil {
				prevNode := lruMap[getNode.previous.key]
				prevNode.next = getNode.next
				lruMap[prevNode.key] = prevNode
			}

			firstNode := lruMap[head.key]
			firstNode.previous = &getNode
			lruMap[firstNode.key] = firstNode

			getNode.next = &firstNode
			getNode.previous = nil
			lruMap[getNode.key] = getNode

			head = nil
			head = &getNode
		}
		return getNode.value
	} else {
		return -1
	}

}

func removeLast() {

	if len(lruMap) < 2 || lruMap == nil {
		lruMap = make(map[int]node)
		//fmt.Printf("corner")
	} else {

		lastNode := lruMap[end.key]
		end = lastNode.previous

		var secondLast = lruMap[lastNode.previous.key]
		secondLast.next = nil
		lruMap[secondLast.key] = secondLast
		end.next = nil

		delete(lruMap, lastNode.key)
	}
}
