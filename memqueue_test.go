package gomemqueue

import (
	"log"
	"testing"
)


func TestQueue_Push(t *testing.T) {
	q := New()

	q.Push("123")
	q.Push("456")

	result := q.BatchPop(2)

	log.Println(result)
}