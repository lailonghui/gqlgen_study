// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Person struct {
	// 姓名
	Name *string `json:"name"`
	// 年龄
	Age *int `json:"age"`
}

type Student struct {
	// 姓名
	Name *string `json:"name"`
	// 年龄
	Age *int `json:"age"`
	// 班级
	Class *string `json:"class"`
}