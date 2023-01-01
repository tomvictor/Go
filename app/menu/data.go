package menu

// this variable will be vissible in the locgic file
// its go's way of local packaging
// Also menuItem type is fetched from the logic file
// tht means, every files in a package is a single file
var menu = []menuItem{
	{
		name: "Coffe",
		prices: map[string]float64{
			"regular": 12.5,
			"large":   20,
		},
	},
	{
		name: "tea",
		prices: map[string]float64{
			"single": 10,
			"double": 15,
			"triple": 20,
		},
	},
}
