package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkLRU(b *testing.B) {
	data := `{"asds" : "sadsada", "Adsada":"SADasdasd" , "Asdsadadsadwqew":1,"dksaodkwqeokasl;dkaskdlasd":"sadsadadsaad"}`
	lru := NewLRU(8192)
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("%v", i)
		lru.Add(key, data)
		lru.Get(key)
	}
}
func TestLRU(t *testing.T) {
	lru := NewLRU(10)
	lru.Add("1", "1st")
	lru.Add("2", "2nd")
	lru.Add("3", "3rd")
	lru.Add("4", "4th")
	lru.Add("5", "5th")
	lru.Add("6", "6th")
	lru.Add("7", "7th")
	lru.Add("8", "8th")
	lru.Add("9", "9th")
	lru.Add("10", "10th")
	assert.Equal(t, "10th", lru.Get("10"))
	assert.Equal(t, "10th", lru.list.Head.Data)
	assert.Equal(t, "1st", lru.list.Tail.Data)
	assert.Equal(t, "1st", lru.Get("1"))
	assert.Equal(t, "1st", lru.list.Head.Data)
	assert.Equal(t, "2nd", lru.list.Tail.Data)
	lru.Add("2", "UPDATED")
	assert.Equal(t, "UPDATED", lru.list.Tail.Data)
	lru.Add("11", "11th")
	assert.Equal(t, "3rd", lru.list.Tail.Data)
	lru.list.Debug()
}
