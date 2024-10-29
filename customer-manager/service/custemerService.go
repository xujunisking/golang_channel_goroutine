package service

import (
	"account/customer-manager/model"
)

// 该Customerservice,完成对customer的操作,包括增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

// 编写一个方法，可以返回 *customerservice
func NewCustomerservice() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

// 根据id删除客户(从切片中删除)
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	}

	//如何从切片中删除一个元素
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

// 根据id查找客户在切片中对应下标,如果没有该客户，返回-1
func (cs *CustomerService) FindById(id int) int {
	index := -1
	//遍历this.customers 切片 I
	for i, v := range cs.customers {
		if v.Id == id {
			index = i
			break
		}
	}
	return index
}

func (cs *CustomerService) Mod(customer model.Customer) bool {
	index := cs.FindById(customer.Id)
	if index == -1 {
		return false
	}
	cs.customers[index] = customer

	return true
}
