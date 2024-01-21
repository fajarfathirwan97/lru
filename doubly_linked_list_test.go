package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	dbl := &DoublyLinkedList{}
	assert.Equal(t, dbl.Len(), int64(0))
	dbl.Add("A")
	assert.Equal(t, dbl.Len(), int64(1))
	dbl.Add("B")
	assert.Equal(t, dbl.Len(), int64(2))
	dbl.Add("C")
	assert.Equal(t, dbl.Len(), int64(3))
	assert.Equal(t, dbl.Search(dbl.Head, "B"), "B")
	_, ok := dbl.Update("B", "S")
	assert.Equal(t, ok, true)
	_, ok = dbl.Update("D", "T")
	assert.Equal(t, ok, false)
	assert.Equal(t, dbl.Search(dbl.Head, "B"), "")
	dbl.Append("NEW")
	assert.Equal(t, dbl.Len(), int64(4))
	assert.Equal(t, dbl.Tail.Data, "NEW")
	dbl.Prepend("NEW HEAD")
	assert.Equal(t, dbl.Len(), int64(5))
	assert.Equal(t, dbl.Head.Data, "NEW HEAD")
	dbl.Remove("S")
	assert.Equal(t, dbl.Len(), int64(4))
	assert.Equal(t, dbl.Search(dbl.Head, "S"), "")
	dbl.Remove("NEW")
	assert.Equal(t, dbl.Len(), int64(3))
	dbl.Prepend("END")
	dbl.Remove("END")
	assert.Equal(t, dbl.Len(), int64(3))
	dbl.Append("TAIL")
	dbl.Prepend("NEW HEAD again")
	assert.Equal(t, dbl.Len(), int64(5))
	dbl.InsertAtIndex(3, "FOUR")
	assert.Equal(t, dbl.Len(), int64(6))
	dbl.Append("NEW TAIL")
	assert.Equal(t, dbl.Len(), int64(7))
	dbl.Append("FOURss at end")
	assert.Equal(t, dbl.Len(), int64(8))
	dbl.InsertAtIndex(3, "at idx 3")
	dbl.Debug()
}
