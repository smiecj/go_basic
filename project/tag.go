// package project 实际项目的一些库的用法
package project

const (
	tagName = "my_tag"
)

type withTagStruct struct {
	// 必须是公开对象才能获取
	Name string `my_tag:"name_"`
}
