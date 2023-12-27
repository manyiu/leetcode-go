package timebasedkeyvaluestore

type TimeMap struct {
	item map[string][]Item
}

type Item struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{
		item: make(map[string][]Item),
	}
}

func (timeMap *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := timeMap.item[key]; !ok {
		timeMap.item[key] = []Item{}
	}

	timeMap.item[key] = append(timeMap.item[key], Item{
		value:     value,
		timestamp: timestamp,
	})
}

func (timeMap *TimeMap) Get(key string, timestamp int) string {
	if _, ok := timeMap.item[key]; !ok {
		return ""
	}

	item := timeMap.item[key]

	if item[0].timestamp > timestamp {
		return ""
	}

	l := 0
	r := len(item) - 1

	for l < r {
		m := (l + r) / 2

		if item[m].timestamp == timestamp {
			return item[m].value
		} else if item[m].timestamp < timestamp {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if item[l].timestamp <= timestamp {
		return item[l].value
	} else {
		return item[l-1].value
	}
}
