package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var fod food
	if order == "burger" {
		fod.preptime = 15
		return fod.preptime
	} else if order == "chips" {
		fod.preptime = 10
		return fod.preptime
	} else if order == "nuggets" {
		fod.preptime = 12
		return fod.preptime
	} else {
		return 404
	}
}
