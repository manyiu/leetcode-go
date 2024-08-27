package kthlargestelementinastream

import "testing"

func TestKthLargest(t *testing.T) {
	k := 7
	nums := []int{-10, 1, 3, 1, 4, 10, 3, 9, 4, 5, 1}
	kthLargest := Constructor(k, nums)

	if kthLargest.Add(3) != 3 {
		t.Fatal("error")
	}

	if kthLargest.Add(2) != 3 {
		t.Fatal("error")
	}

	if kthLargest.Add(3) != 3 {
		t.Fatal("error")
	}

	if kthLargest.Add(1) != 3 {
		t.Fatal("error")
	}

	if kthLargest.Add(2) != 3 {
		t.Fatal("error")
	}

	if kthLargest.Add(4) != 3 {
		t.Fatal("error")
	}

	if kthLargest.Add(5) != 4 {
		t.Fatal("error")
	}

	if kthLargest.Add(5) != 4 {
		t.Fatal("error")
	}

	if kthLargest.Add(6) != 4 {
		t.Fatal("error")
	}

	if kthLargest.Add(7) != 5 {
		t.Fatal("error")
	}

	if kthLargest.Add(7) != 5 {
		t.Fatal("error")
	}

	if kthLargest.Add(8) != 5 {
		t.Fatal("error")
	}

}
