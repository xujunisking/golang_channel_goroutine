package utils

import (
	"fmt"
)

type FamilyAccount struct {
	key     string
	balance float32
	money   float32
	note    string
	flag    bool
	details string
	loop    bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说    明",
	}
}

// 将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("------------------当前收支明细记录------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧!")
	}
}

func (this *FamilyAccount) income() {
	this.flag = true
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	fmt.Println(this.details)
}

func (this *FamilyAccount) pay() {
	this.flag = true
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	if this.money > this.balance {
		fmt.Println("余额不足...")
		return
	}
	this.balance -= this.money // 修改账户余额
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	//将这个支出情况，拼接到details变量
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	fmt.Println(this.details)
}

func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗?y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "x" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n------------------家庭收支记账软件------------------")
		fmt.Println("                   1 收支明细")
		fmt.Println("                   2 登记收入")
		fmt.Println("                   3 登记支出")
		fmt.Println("                   4 退出软件")

		fmt.Print("请选择(1-4):")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项...")
		}

		if !this.loop {
			break
		}
	}
}
