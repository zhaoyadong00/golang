1 golang的优点
    自带GC
    静态编译
    简单的思想，没有继承多态类等
    丰富的库和详细的文档
    原生支持并发和拥有同步并发的channel类型，是并发开发非常容易
    简单的语法，提高开发效率和代码的可读性可维护性
    超级简单的交叉编译，仅需改变环境变量
2 go语言的主要特性
    自动立即回收
    更丰富的内置类型 int 系列 byte array slice map channel interface function bool rune string error
    函数多返回值
    错误处理
    匿名函数和闭包
    类型和接口 type interface
    并发编程
    反射
3 golang 值类型
    bool int系列 string byte array float32 float64
4 golang 引用类型
    channel slice map 
5 数组的特性
    golang数组是一个固定长度固定类型的序列
    定义 var a [len]type 数组长度len必须是常量并且是类型的一部分 [5]int [6]int 是不同的类型 一旦定义长度不变
    数组可以通过下标访问 下标是从0开始 最后一个元素下标是len-1
    遍历方式 
    for i:=0 ; i < len(a) ; i++{
    }
    for index , v := range a {
    }
    越界会造成panic
    数组是值类型 赋值和传参会造成复制 所以在数据量比较大的时候 使用指针数组
    指针数组[4]*int 是数组 具有数组的特性 是传值
    数组指针*[5]int 是一个指针 具有指针的特性 是传引用
    len和cap的区别 len的定义是数组的可用长度 cap 是底层数组的容量
6 切片的特性
    切片是数组的引用，因此切片是引用类型，但自身是结构体，值copy，这也就能说明在传参时是值传递而不是引用传递
    切片的长度可以改变，因此切片是一个可变的数组
    切片的遍历方式和数组是一样的 for  for range 可以用len获取可用长度 用cap获取底层数组容量 读写操作不能超过len的数量
    切片的定义 var a []int  var a = []int{}
    如果slice == nil 那么len 和 cap都等于0
    
7 指针
    指针地址 指针类型和指针取值
    &取址 * 取值
8 map是一种无序的基于key-value的数据结构 golang的map为引用类型 必须初始化才能使用 初始化使用make
    var m = make(map[keyType]valueType , int)
    
    Printf("T%" , m)