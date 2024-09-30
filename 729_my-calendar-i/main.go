package mycalendari

type MyCalendar struct {
	start, end  int
	left, right *MyCalendar
}

func Constructor() MyCalendar {
	return MyCalendar{
		start: -1,
		end:   0,
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if start < this.end && end > this.start {
		return false
	}

	if end <= this.start {
		if this.left == nil {
			this.left = &MyCalendar{
				start: start,
				end:   end,
			}

			return true
		}

		return this.left.Book(start, end)
	}

	if this.right == nil {
		this.right = &MyCalendar{
			start: start,
			end:   end,
		}

		return true
	}

	return this.right.Book(start, end)
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
