package main

import (
	"fmt"
)

var stockcoststock, stockgavemoney, stockgavestock, stockaddmoney, stockprice float64
var costmoney, growth_rate, addmoney float64
var moneymode, age, bAge int

type Growthmoney struct {
	costMoney, growth_Rate, addMoney float64
	age, bAge                        int
}

type Stockmoney struct {
	stockcostStock, stockgaveMoney, stockgaveStock, stockaddMoney, stockPrice float64
	age, bAge                                                                 int
}

func (g Growthmoney) growthmoney() {

	g.costMoney *= 10000
	g.addMoney *= 10000
	g.growth_Rate /= 100

	for i := g.age; i < (g.bAge + 1); i++ {
		g.costMoney = g.costMoney + (g.costMoney * g.growth_Rate)
		fmt.Printf("第%v歲,可得: %v 萬\n", i, (g.costMoney / 10000))
		g.costMoney += g.addMoney
	}

}

func (s Stockmoney) stockmoney() {

	for i := s.age; i < (s.bAge + 1); i++ {
		s.stockcostStock = s.stockcostStock + ((s.stockcostStock * s.stockgaveMoney) / s.stockPrice) + ((s.stockcostStock * s.stockgaveStock) / 10)
		fmt.Printf("第%v歲,可得: %v 張\n", i, (s.stockcostStock / 1000))
		s.stockaddMoney /= s.stockPrice
		s.stockcostStock += s.stockaddMoney
	}

}

func sinio() (float64, float64, float64, float64, float64, int, int) {

	fmt.Println("請輸入成本(股): ")
	fmt.Scanln(&stockcoststock)
	fmt.Println("請輸入配多少錢(小數): ")
	fmt.Scanln(&stockgavemoney)
	fmt.Println("請輸入配多少股(小數): ")
	fmt.Scanln(&stockgavestock)
	fmt.Println("請輸入每年投入(元): ")
	fmt.Scanln(&stockaddmoney)
	fmt.Println("請輸入股票價格(一股幾元): ")
	fmt.Scanln(&stockprice)
	fmt.Println("請輸入年紀(歲): ")
	fmt.Scanln(&age)
	fmt.Println("請輸入算到幾歲: ")
	fmt.Scanln(&bAge)

	return stockcoststock, stockgavemoney, stockgavestock, stockaddmoney, stockprice, age, bAge

}

func ginio() (float64, float64, float64, int, int) {

	fmt.Println("請輸入成本(萬): ")
	fmt.Scanln(&costmoney)
	fmt.Println("請輸入成長率(雙位整數): ")
	fmt.Scanln(&growth_rate)
	fmt.Println("請輸入每年投入(萬): ")
	fmt.Scanln(&addmoney)
	fmt.Println("請輸入年紀(歲): ")
	fmt.Scanln(&age)
	fmt.Println("請輸入算到幾歲: ")
	fmt.Scanln(&bAge)

	return costmoney, growth_rate, addmoney, age, bAge
}

func main() {

	fmt.Println("********* 資產成長率計算機 **********")
	fmt.Println("計算金錢成長選1 計算股數成長選2 => ")
	fmt.Scanln(&moneymode)

	switch moneymode {
	case 1:
		ginio()
		gg := Growthmoney{costmoney, growth_rate, addmoney, age, bAge}
		gg.growthmoney()
	case 2:
		sinio()
		ss := Stockmoney{stockcoststock, stockgavemoney, stockgavestock, stockaddmoney, stockprice, age, bAge}
		ss.stockmoney()
	}

}
