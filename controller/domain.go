package controller

type Person struct {
	Pid string
	Name string
	Password string
	Age int
	Gender string
}

var Persons []Person

func init() {
	//创建5个用户，数组模拟数据库存储
	p := Person{Pid:"111", Name:"jack", Password:"123456", Age:21, Gender:"m"}
	p2 := Person{Pid:"111", Name:"Jerry", Password:"123", Age:23, Gender:"m"}
	p3 := Person{Pid:"111", Name:"alice", Password:"123456", Age:26, Gender:"f"}
	p4 := Person{Pid:"111", Name:"bob", Password:"qazx", Age:29, Gender:"m"}
	p5 := Person{Pid:"111", Name:"candy", Password:"123456", Age:22, Gender:"f"}

	Persons = append(Persons, p)
	Persons = append(Persons, p2)
	Persons = append(Persons, p3)
	Persons = append(Persons, p4)
	Persons = append(Persons, p5)
}

