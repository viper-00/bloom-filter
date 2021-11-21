package bloom

import "testing"

func ExampleCBF() {

}

func TestBasicCBF(t *testing.T) {
	cbf := NewCountingBloomFilter(100, 0.01)
	cbf.Add([]byte("abc"))
	cbf.Add([]byte("efg"))
	cbf.Add([]byte("hijk"))

	if cbf.Test([]byte("abc")) == false {
		t.Errorf("CBF: test failed\n")
	}

	cbf.Remove([]byte("efg"))
	if cbf.Test([]byte("efg")) == true {
		t.Errorf("CBF: test remove failed\n")
	}
}
