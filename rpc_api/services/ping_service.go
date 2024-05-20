package apiServices

// 以一种类似interface的方式定义了一个struct
type Test struct {
	Ping  func() (string, error)
	Hello func() (string, error)
}
