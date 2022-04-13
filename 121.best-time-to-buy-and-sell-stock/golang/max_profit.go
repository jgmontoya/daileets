package max_profit

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	if len(prices) == 2 {
		return max(0, prices[1]-prices[0])
	}

	max_profit := 0
	var selected_buy_price int
	var selected_sell_price int

	if prices[0] <= prices[1] {
		selected_buy_price = prices[0]
		selected_sell_price = prices[1]
	} else {
		selected_buy_price = prices[1]
		selected_sell_price = prices[2]
	}

	max_profit = selected_sell_price - selected_buy_price
	for index, value := range prices[2:] {
		index = index + 2
		if value > selected_sell_price {
			selected_sell_price = value
		} else if value < selected_buy_price {
			selected_buy_price = value
			if index < len(prices)-1 {
				selected_sell_price = prices[index+1]
			} else {
				selected_sell_price = -1
			}
		}
		max_profit = max(max_profit, selected_sell_price-selected_buy_price)
	}

	return max(0, max_profit)
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
