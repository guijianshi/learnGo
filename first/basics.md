# Go基础知识

## 关键字
1. 下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var

2. 除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：
append 	bool 	byte 	cap 	close 	complex 	complex64 	complex128 	uint16
copy 	false 	float32 	float64 	imag 	int 	int8 	int16 	uint32
int32 	int64 	iota 	len 	make 	new 	nil 	panic 	uint64
print 	println 	real 	recover 	string 	true 	uint 	uint8 	uintptr

## 类型

序号 | 类型 | 描述
- | :-:        | -:
1  | 布尔型     | true, false
2  | 数字类型   | int, float32, float64
3  | 字符串类型 | 使用UTF-8编码标识Unicode文本
4  | 派生类型   |      (a) 指针类型（Pointer）(b) 数组类型(c) 结构化类型(struct)(d) Channel 类型(e) 函数类型(f) 切片类型(g) 接口类型（interface）(h) Map 类型



### 整数型

序号| 类型 |描述
- | :-:   | -:
1 |	uint8 |无符号 8 位整型 (0 到 255)
2 |	uint16|无符号 16 位整型 (0 到 65535)
3 |	uint32|无符号 32 位整型 (0 到 4294967295)
4 |	uint64|无符号 64 位整型 (0 到 18446744073709551615)
5 |	int8  |有符号 8 位整型 (-128 到 127)
6 |	int16 |有符号 16 位整型 (-32768 到 32767)
7 |	int32 |有符号 32 位整型 (-2147483648 到 2147483647)
8 |	int64 |有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

### 浮点型
序号| 类型 |描述
- | :-:   | -:
1 |	float32 |32位浮点数
2 |	float64|64位浮点数
3 |	float64|32 位实数和虚数
4 |	complex128| 64 位实数和虚数

### 其他数字类型

序号| 类型 |描述
- | :-:   | -:
1 |	byte |类似于uint8
2 |	rune|类似于int32
3 |	uint|32位或64位
4 |	int| 与uint一样大小
4 |	uintptr| 无符号类型,用于存放一个指针
