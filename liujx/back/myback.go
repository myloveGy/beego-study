package back // 我的一个包

import "fmt" // 引入文件包

func Mprintf(age ...int) {
	length := len(age)
	for i := 0; i < length; i++ {
		fmt.Printf("age i = %T, value = %T \n", i, age[i])
	}
}
