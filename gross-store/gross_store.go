package gross

// Units store the Gross Store unit measurement
// units := Units()
// fmt.Println(units)
// Output: map[...] with entries like ("dozen": 12)
func Units() map[string]int {
	m := make(map[string]int)
	m["quarter_of_a_dozen"] = 3
	m["half_of_a_dozen"] = 6
	m["dozen"] = 12
	m["small_gross"] = 120
	m["gross"] = 144
	m["great_gross"] = 1728

	return m
}

// NewBill create a new bill
// bill := NewBill()
// fmt.Println(bill)
// Output: map[]
func NewBill() map[string]int {
	m := make(map[string]int)
	return m
}

// AddItem add item to customer bill
// bill := NewBill()
// units := Units()
// ok := AddItem(bill, units, "carrot", "dozen")
// fmt.Println(ok)
// Output: true (since dozen is a valid unit)
func AddItem(bill, units map[string]int, item, unit string) bool {
	// check if the unit is in map
	quantity, ok := units[unit]
	if ok == false {
		return false
	}

	// add the item and number of units to the bill.
	bill[item] = quantity

	return true
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// return false if item is in the bill
	currentQuantity, ok := bill[item]
	if ok == false {
		return false
	}

	// return false if the unit is in the units map
	removeQuantity, ok := units[unit]
	if ok == false {
		return false
	}

	// doing if/else using switch without an expression
	switch {
	// return false if the new quantity would be less than zero
	case currentQuantity-removeQuantity < 0:
		return false

	// if the new quantity is zero, remove the item from the bill and return true
	case currentQuantity-removeQuantity == 0:
		delete(bill, item)
		return true
	default:
		// lastly, reduce quantity and return true
		bill[item] = currentQuantity - removeQuantity
		return true
	}
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	// return 0, false if the item is not on the bill
	amount, ok := bill[item]
	if ok == false {
		return 0, false
	}

	// return the quantity of the item and true
	return amount, true
}
