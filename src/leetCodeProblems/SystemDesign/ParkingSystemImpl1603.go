package main

/*
- LeetCode - https://leetcode.com/problems/design-parking-system/
- Time - O(1)
- Space - O(1)
*/

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

// func Constructor(big int, medium int, small int) ParkingSystem {
// 	return ParkingSystem{
// 		big:    big,
// 		medium: medium,
// 		small:  small,
// 	}
// }

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big <= 0 {
			return false
		}
		this.big--
	case 2:
		if this.medium <= 0 {
			return false
		}
		this.medium--
	case 3:
		if this.small <= 0 {
			return false
		}
		this.small--
	}
	return true
}

// func main() {
// 	// obj := Constructor(3)
// 	// obj.Put(1, 2)
// 	// obj.Put(3, 4)
// 	// obj.Put(5, 6)
// 	// log.Println(obj.Get(-1))
// 	// log.Println(obj.Get(1))
// 	// log.Println(obj.Get(3))
// 	// log.Println(obj.Get(5))

// 	obj := Constructor(10)
// 	obj.Put(10, 13)
// 	obj.Put(3, 17)
// 	obj.Put(6, 11)
// 	obj.Put(10, 5)
// 	obj.Put(9, 10)
// 	obj.Get(13)
// 	obj.Put(2, 19)
// 	obj.Get(2)
// 	obj.Get(3)
// 	obj.Put(5, 25)
// 	obj.Get(8)
// 	obj.Put(9, 22)
// 	obj.Put(5, 5)
// 	obj.Put(1, 30)
// 	obj.Get(11)
// 	obj.Put(9, 12)
// 	obj.Get(7)
// 	obj.Get(5)
// 	obj.Get(8)
// 	obj.Get(9)
// 	obj.Put(4, 30)
// 	obj.Put(9, 3)
// 	obj.Get(9)
// 	obj.Get(10)
// 	obj.Get(10)
// 	obj.Put(6, 14)
// 	obj.Put(3, 1)
// 	obj.Get(3)
// 	obj.Put(10, 11)
// 	obj.Get(8)
// 	obj.Put(2, 14)
// 	obj.Get(1)
// 	obj.Get(5)
// 	obj.Get(4)
// 	obj.Put(11, 4)
// 	obj.Put(12, 24)
// 	obj.Put(5, 18)
// 	obj.Get(13)
// 	obj.Put(7, 23)
// 	obj.Get(8)
// 	obj.Get(12)
// 	obj.Put(3, 27)
// 	obj.Put(2, 12)
// 	obj.Get(5)
// 	obj.Put(2, 9)
// 	obj.Put(13, 4)
// 	obj.Put(8, 18)
// 	obj.Put(1, 7)
// 	obj.Get(6)
// 	obj.Put(9, 29)
// 	obj.Put(8, 21)
// 	obj.Get(5)
// 	obj.Put(6, 30)
// 	obj.Put(1, 12)
// 	obj.Get(10)
// 	// obj.Put(9, 29)
// 	// obj.Put(8, 21)
// 	// obj.Get(5)
// 	// obj.Put(6, 30)
// 	// obj.Put(1, 12)
// 	// obj.Get(10)
// 	// obj.Put(4, 15)
// 	// obj.Put(7, 22)
// 	// obj.Put(11, 26)
// 	// obj.Put(8, 17)
// 	// obj.Put(9, 29)
// }
