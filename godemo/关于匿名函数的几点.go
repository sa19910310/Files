//关于匿名函数的几点意见，有匿名函数的地方容易有闭包
// 关于闭包闭包里使用的母函数中的局部变量会一直存活在闭包中
//并且这个变量还是一直是引用变量，就是是用的地址，这就是闭包的特性之一。
package main

import (
	"fmt"
	"os"
)

func main() {
	lib()
}

func lib() {
	var rmdirs []func()
	dirs := tempDirs()
	for i := 0; i < len(dirs); i++ {
		j := i
		fmt.Println(&j, &i)
		//0xc000012078 0xc000012070----
		//0xc000012098 0xc000012070
		//0xc0000120a8 0xc000012070
		//0xc0000120b8 0xc000012070
		//0xc0000120c8 0xc000012070
		// 这里因为并不牵涉到闭包所以就是正常的每次循环的i
		// 这里之所以是使用一个局部变量是因为想记录每次的i值因为闭包中的i一直是引用值
		//但是如果使用j这个局部变量的话，那么每次使用的就是实际值了。
		os.MkdirAll(dirs[i], 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dirs[j]) // NOTE: incorrect!
		})
	}
	for _, v := range rmdirs {
		v()
	}
}

func tempDirs() []string {
	return []string{"1", "2", "3", "4", "5"}
}

//解释：ij：

//j 局部变量，每次都不是同一个地址，i循环变量每次都是同一个地址。
//那么如果是直接使用i也不会出现任何问题，因为它是直接每次循环中直接使用了这个变量，所以就算是一个地址也无所谓（举个例子，假如说不是
//这种匿名函数闭包的形式，是调用一个函数，例如fmt.Println(i),那么每次都是等这个函数执行完再进行下一个循环）但是匿名函数这种却不一样，
//或者说闭包就不一样，它因为只是定义并没有立即执行，所以它里面的值并不是每次都执行一次，导致在循环结束的时候它里面的i值是同一个也就是
//最后一个i的值，所以这个时候选择一个局部变量就很有必要了，因为局部变量并不是引用值，它每次都有不同的地址，导致每次声明函数的时候的j
//都不是一个j，所以这样就解决了循环变量是引用变量这个问题了（主要的原因是声明函数，并不是执行函数假如是匿名但是是 立即执行函数，那么
//也不会出现任何的问题）


//加入改成一下形式：
func test(){
	
for i := 0; i < len(dirs); i++ {
	j := &i//这里说明一下 i不是指针类型，它是引用类型。或者可以理解为 它是一个指向底层的指针，然后系统自动取值了。这里 加入j也赋予i同样的功能
	fmt.Println(j, &i)//那么跟i的功能就一样了变成了引用类型了，所以j之所以能用就是它必须是一个值类型，然后每次赋予给声明函数的都是一个值。
	os.MkdirAll(dirs[i], 0755)
	rmdirs = append(rmdirs, func() {
		os.RemoveAll(dirs[*j]) // NOTE: incorrect!
	})
}
