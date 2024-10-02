package findmedianfromdatastream

import "testing"

func TestFindMedianFromDataStream(t *testing.T) {
	obj := Constructor()
	obj.AddNum(1)
	if obj.FindMedian() != 1.0 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 1.0)
	}
	obj.AddNum(2)
	if obj.FindMedian() != 1.5 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 1.5)
	}
	obj.AddNum(3)
	if obj.FindMedian() != 2.0 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 2.0)
	}
	obj.AddNum(4)
	if obj.FindMedian() != 2.5 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 2.5)
	}
	obj.AddNum(5)
	if obj.FindMedian() != 3.0 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 3.0)
	}
	obj.AddNum(6)
	if obj.FindMedian() != 3.5 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 3.5)
	}
	obj.AddNum(7)
	if obj.FindMedian() != 4.0 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 4.0)
	}
	obj.AddNum(8)
	if obj.FindMedian() != 4.5 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 4.5)
	}
	obj.AddNum(9)
	if obj.FindMedian() != 5.0 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 5.0)
	}
	obj.AddNum(10)
	if obj.FindMedian() != 5.5 {
		t.Fatal("got: ", obj.FindMedian(), "want: ", 5.5)
	}
}
