package utils

//target = 2880
//map[1:1000 2:2000 3:2200 4:2500 5:3500 6:4500 7:4600 8:5100 9:5500]
//map[1:10 2:11 3:12 4:13 5:15 6:16 7:17 8:18 9:19]
func Binary(target int, dataCount map[int]int, dataInfo map[int]string) string {
	var begin, end int
	begin = 1
	end = len(dataCount)

	for begin <= end {
		if target < dataCount[begin] {
			return dataInfo[begin]
		}
		if target > dataCount[end] {
			return dataInfo[end+1]
		}
		if begin == end-1 {
			return dataInfo[end]
		}

		middle := (begin + end) / 2
		if dataCount[middle] > target {
			end = middle
		} else if dataCount[middle] < target {
			begin = middle
		} else {
			return dataInfo[middle]
		}
	}

	return dataInfo[end]
}
