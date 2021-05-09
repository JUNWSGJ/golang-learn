package main

import (
	"fmt"
)

func main() {

	//  声明了一个[]int类型的变量s1, 指定该切片长度为5
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)

	// 可以看出，用make函数初始化切片时，如果不指明其容量，那么它就会和长度一致。

	//  声明了一个[]int类型的变量s2, 指定该切片长度为5， 容量为8
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)

	// 切片的容量实际上代表了它的底层数组的长度，这里是8

	//	当我们用make函数或切片值字面量（比如[]int{1, 2, 3}）初始化一个切片时，
	//	该窗口最左边的那个小格子总是会对应其底层数组中的第 1 个元素

	// 通过切片表达式基于某个数组或切片生成新切片
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6] //  [4 5 6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	// 在底层数组不变的情况下，切片代表的窗口可以向右扩展，直至其底层数组的末尾
	// s4的容量就是其底层数组的长度8, 减去上述切片表达式中的那个起始索引3，即5
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	s4 = append(s4, 10) // [4 5 6 10]
	// 此时s3变为 [1 2 3 4 5 6 10 8] 因为s4的底层数组是s3, 更新s4的时候同时会改变s3里的元素值
	fmt.Printf("The value of s3: %d\n", s3)
	fmt.Printf("The value of s4: %d\n", s4)

	// 切片的扩容，
	// 一旦一个切片无法容纳更多的元素，Go 语言就会想办法扩容。
	//但它并不会改变原来的切片，而是会生成一个容量更大的切片，然后将把原有的元素和新元素一并拷贝到新切片中。
	//在一般的情况下，你可以简单地认为新切片的容量（以下简称新容量）将会是原切片容量（以下简称原容量）的 2 倍。

	//s4  当前值为 [4 5 6 10]， 容量为5， 再添加两个值，使其扩容，看容量大小
	s4 = append(s4, 11)
	s4 = append(s4, 12)

	fmt.Printf("s4, len:[%d], capacity:[%d], value:%d\n", len(s4), cap(s4), s4)

	// 切片的底层数组什么时候会被替换？
	// 确切地说，一个切片的底层数组永远不会被替换。为什么？
	// 虽然在扩容的时候 Go 语言一定会生成新的底层数组，但是它也同时生成了新的切片
	// 它只是把新的切片作为了新底层数组的窗口，而没有对原切片，及其底层数组做任何改动。
	// 在无需扩容时，append函数返回的是指向原底层数组的原切片，
	// 而在需要扩容时，append函数返回的是指向新底层数组的新切片。

}
