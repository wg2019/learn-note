package template

// Menu 菜单
type Menu interface {
	Name()
	StapleFood()
	Dish()
	Drink()
}

// OrderMeal 订餐
func OrderMeal(m Menu) {
	m.Name()
	m.StapleFood()
	m.Dish()
	m.Drink()
}
