package multidbfile

import (
	"bytes"
	"testing"
)

func TestMultiDBfile(t *testing.T) {
	up := newTop()

	// level 1, one item
	up.createDB()
	up.store([]byte("hello world"), []byte("hello world"))
	value := up.find([]byte("hello world"))
	if value == nil {
		t.Error("there should have been stored the first item")
	}

	// level 2, two items
	up.createDB()
	up.store([]byte("hello golang"), []byte("hello golang"))
	up.store([]byte("hello php"), []byte("hello php"))

	// level 3, two items
	up.createDB()
	up.store([]byte("hello c++"), []byte("hello c++"))
	up.store([]byte("hello java"), []byte("hello java"))

	// level 4, one item
	up.createDB()
	up.store([]byte("hello ruby"), []byte("hello ruby"))

	// level 5, four items
	up.createDB()
	up.store([]byte("hello c"), []byte("hello c"))
	up.store([]byte("hello c#"), []byte("hello c#"))
	up.store([]byte("hello rust"), []byte("hello rust"))
	up.store([]byte("hello js"), []byte("hello js"))

	value = up.find([]byte([]byte("hello world")))
	if bytes.Compare(value, []byte("hello world")) != 0 {
		t.Error("total:level 4 find in level 1 error")
	}

	value = up.find([]byte([]byte("hello php")))
	if bytes.Compare(value, []byte("hello php")) != 0 {
		t.Error("total:level 4 find in level 2 error")
	}

	value = up.find([]byte([]byte("hello js")))
	if bytes.Compare(value, []byte("hello js")) != 0 {
		t.Error("total:level 4 find in level 5 error")
	}

	up.update([]byte("hello ruby"), []byte("hello modified ruby"))
	value = up.find([]byte([]byte("hello ruby")))
	if bytes.Compare(value, []byte("hello modified ruby")) != 0 {
		t.Error("total:level 4 find in level 4 error")
	}

	up.del([]byte("hello golang"))
	value = up.find([]byte([]byte("hello golang")))
	if value != nil {
		t.Error("this item should have been removed")
	}
}
