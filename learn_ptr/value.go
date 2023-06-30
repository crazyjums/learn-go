package learn_ptr

func PtrValue() {
	var i *int
	//*i = 10 // panic: runtime error: invalid memory address or nil pointer dereference
	//i = 10 // 语法错误，因为 i 是一个指针，不能直接赋值为 int 类型
	i = new(int)
	*i = 10
	b := 20
	i = &b
	println(*i)
}
