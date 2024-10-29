package main

import (
	"account/customer-manager/model"
	"account/customer-manager/service"
	"fmt"
)

type customerView struct {
	key             string //接收用户输入
	loop            bool   //表示是否循环的显示主菜单
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (cv *customerView) list() {
	//首先，获取到当前所有的客户信息(在切片中)
	customers := cv.customerService.List()
	fmt.Println("-------------------------客户列表-------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, v := range customers {
		fmt.Println(v.GetCustomerInfo())
	}
	fmt.Printf("\n-----------------------客户列表完成-----------------------\n\n")
}

func (cv *customerView) add() {
	var name string
	var gender string
	var age int
	var phone string
	var email string

	fmt.Println("-------------------------添加客户-------------------------")

	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("电话：")
	fmt.Scanln(&phone)
	fmt.Print("email：")
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("-------------------------添加完成-------------------------")
	} else {
		fmt.Println("-------------------------添加失败-------------------------")
	}
}

func (cv *customerView) delete() {
	fmt.Println("-------------------------删除客户-------------------------")
	fmt.Println("请选择待删除客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)

	if id == -1 {
		return //放弃删除操作
	}

	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if cv.customerService.Delete(id) {
			fmt.Println("-------------------------删除成功-------------------------")
		} else {
			fmt.Println("----------------删除失败，输入的id号不存在-----------------")
		}
	}
}

func (cv *customerView) mod() {
	fmt.Println("-------------------------修改客户-------------------------")
	fmt.Println("请选择待修改客户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)

	if id == -1 {
		return //放弃删除操作
	}

	index := cv.customerService.FindById(id)
	if index == -1 {
		fmt.Println("----------------修改失败，输入的id号不存在-----------------")
		return
	}

	var name string
	var gender string
	var age int
	var phone string
	var email string
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("电话：")
	fmt.Scanln(&phone)
	fmt.Print("email：")
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if cv.customerService.Mod(customer) {
		fmt.Println("-------------------------修改成功-------------------------")
	} else {
		fmt.Println("----------------修改失败，输入的id号不存在-----------------")
	}
}

// 退出软件
func (cv *customerView) exit() {
	fmt.Println("确认是否退出(Y/N):")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "y" || cv.key == "Y" || cv.key == "N" || cv.key == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N):")
	}
	if cv.key == "y" || cv.key == "Y" {
		cv.loop = false
	}
}

// 显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("------------------客户信息管理软件------------------")
		fmt.Println("                  1 添 加 客 户")
		fmt.Println("                  2 修 改 客 户")
		fmt.Println("                  3 删 除 客 户")
		fmt.Println("                  4 客 户 列 表")
		fmt.Println("                  5 退       出")
		fmt.Print("请选择(1-5):")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.mod()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !cv.loop {
			break
		}
	}

	fmt.Println("你退出了客户关系管理系统...")
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}

	customerView.customerService = service.NewCustomerservice()
	customerView.mainMenu()
}
