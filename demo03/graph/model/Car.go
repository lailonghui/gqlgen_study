/*
@Time : 2020/11/11 10:42
@Author : lai
@Description :
@File : Car
*/
package model

type Person struct {
	Name string
}

type Car struct {
	Make            string
	Model           string
	Color           string
	OwnerID         *string
	OdometerReading int
}

func (c *Car) Owner() *Person {
	// get the car owner
	//....
	owner := &Person{
		Name: "赖龙辉",
	}
	return owner
}
