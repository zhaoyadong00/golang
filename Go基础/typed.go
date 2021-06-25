package main

import (
	. "log"
)

func add(a [3]*int) {
	p := 1
	a[1] = &p
}
func sliAdd(p []int) {
	p = append(p, 2)
	p[2] = 4
}
func modifyMap(m map[string]string)  {
	m["b"] = "asd"
}
func modifyChannel(c chan int)  {
	c <- 10
}
func addSli(p *[3]int)  {
	p[1] = 1
}
//全局array
var (
	Arr [5]int = [5]int{ 0 , 1 , 2 , 3 , 4}
	Arr1 = [...]int{1 , 2 , 3, 4,5}
	Arr2 = [5]int{}
)
func arrayOperation()  {
	//数组定义
	arr:= [5]int{}
	Println(arr)
	arr1 := [...]int{}
	Println(arr1)
}
//求所有元素之和
func sumArr(a [10]int) int {
	var sum = 0
	for _ , v := range a{
		sum += v
	}
	return sum
}
//找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）
func myTest(data [10]int , target int) {
	for index , v1 := range data{
		for i := len(data) - 1 ; i > 0 ; i-- {
			if data[i] + v1 == target {
				Println(index , i)
			}
		}
	}
}
//var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//var s1 []int = arr[2:8] //2, 3, 4, 5, 6, 7
//var s2 []int = arr[:6]//0, 1, 2, 3, 4, 5
//var s3 []int = arr[5:]//5, 6, 7, 8, 9
//var s4 []int = s2[2:4]//2, 3
//var s1 = make([]int , 10)
//var s2 = make([]int , 10 , 20)//make([]type , len , cap)
func main() {
	//元素为切片的map
	var m = make(map[string]interface{} , 10)
	var m1 map[string][]string
	m["1"] = []string{"123"}
	Println(m , m == nil , m1 , m1 == nil)

	//元素为map类型的切片
	//var sli []map[string]string
	//sli = make([]map[string]string , 10)
	//for i := 0 ; i < len(sli) ; i++ {
	//	m := make(map[string]string , 1)
	//	m["a"] = "bbb"
	//	sli[i] = m
	//}
	//Println(sli , sli == nil)
	//指针
	//var m = make(map[string]int , 100)
	//Println(len(m))
	//for i := 0 ; i < 100; i++ {
	//	key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
	//	value := rand.Intn(100)          //生成0~99的随机整数
	//	m[key] = value
	//}
	////取出map中的所有key存入切片keys
	//var keys = make([]string, 0, 200)
	//for key := range m {
	//	keys = append(keys, key)
	//}
	////对切片进行排序
	//sort.Strings(keys)
	////按照排序后的key遍历map
	//for _, key := range keys {
	//	fmt.Println(key, m[key])
	//}
	//Println(m , len(m))
	//var m = map[string]int{
	//	"aaa":11,
	//	"bbb":123,
	//	"ccc":22,
	//}
	////delete
	//delete(m , "bbb")//删除
	//for v , k := range m{
	//	Println(v , k)
	//}
	//v , ok := m["aaa"]
	//Println(v , ok)
	//if v , ok := m["aaa"]; ok {//map检查某个key是否存在
	//	Println(v)
	//}
	//t := fmt.Sprintf("%T" , m)//打印类型
	//Println(t == "map[string]int")
	//考察指针
	//var a = 10
	//var p *int
	//p = &a
	//*p = 20
	//Println(p , a) //a = ?
	//var p = new(int)
	//*p = 123
	//Println(*p , &p , p)//分别是对P取值 获取变量P的内存地址 变量P的值的内存地址

	//var p *string
	//p = new(string)
	//*p = "123"
	//Println(p , *p)
	//类型断言
	//a := 10
	//var b interface{}
	//b = &a
	//if _ , ok := b.(*int); ok{
	//	Println(b)
	//}
	//var p *int
	//var b = 1
	//p = &b
	//Println(p)
	//试题
	//p := 1
	//b := &p //取址
	//Println(b , *b)//取值
	//切片扩容策略？
	//var sli = make([]int , 10 , 20)
	//for i := 0 ; i < 10 ; i++ {
	//	if len(sli) == cap(sli) {
	//		sli = append(sli , 1)
	//		Println(cap(sli))
	//	}else{
	//		sli = append(sli , 1)
	//		i--
	//	}
	//}
	////面试题
	//s1 := make([]int , 0)
	//s2 := []int{}
	//var s3 []int
	//Println(s1 == nil , s2 == nil , s3 == nil)//false false true

	//var data = [10]int{1,2,3,4,5,6,7,8,9,10}
	//s2 := data[4:7]
	//Println(s2 , len(s2) , cap(s2))//len是切片可用长度 因为定义是 4-7 前包后不包原则 567 len为7-4 = 3  cap为s2底层指针指向下标为4 总长度是10 10-4=6 所以cap是6、
	//面试题
	//s1 := make([]int , 10 , 20)
	//s2 := s1[5:10]
	//Println(s2 , len(s2) , cap(s2))//结果：[0 0 0 0 0] 5 15

	//s1[0] = 1
	//s1 = append(s1 , 10)
	//Println(s1 , len(s1) , cap(s1) , s2 , len(s2) , cap(s2))
	//
	//data := [...]int{0, 1, 2, 3, 4, 5}
	//s := data[2:4]
	//s[0] += 100
	//s[1] += 200
	//s = append(s , 12)
	//Println(s)
	//Println(data)
	//arrayOperation()
	//切片操作
	//Println(s1 , s2 , len(s2) , cap(s2), s3 , s4 , len(s4) , cap(s4))
	//var arr = [5]int{1,2,3,4,5}
	//s1 := arr[:]//从下标1 到第4个元素 234
	//Println(s1 , len(s1) , cap(s1) , len(arr) , cap(arr))
	//s1 = append(s1 , 6)
	//Println(s1 , len(s1) , cap(s1) , len(arr) , cap(arr))
	//s1 = append(s1 , 6)
	//Println(s1 , len(s1) , cap(s1) , len(arr) , cap(arr))
	//s1 = append(s1 , 6)
	//Println(s1 , len(s1) , cap(s1) , len(arr) , cap(arr))
	//s1 = append(s1 , 6)
	//Println(s1 , len(s1) , cap(s1) , len(arr) , cap(arr))
	//var sli []int = make([]int , 0)
	//var s1 = []int{}
	////sli == nil 时 未初始化
	////sli[0] = 1//如果切片的可用长度不足 则爆panic
	//s1 = append(s1, 1)//len(s1) = 1 , cap(s1) = 1
	//Println(s1 == nil , len(s1) , cap(s1))
	//s1 = append(s1, 1)//len(s1) =2 , cap(s1) = 2
	//Println(s1 == nil , len(s1) , cap(s1))
	//s1 = append(s1, 1)//len(s1) =3 , cap(s1) = 4
	//Println(s1 == nil , len(s1) , cap(s1))
	//s1 = append(s1, 1)//len(s1) =4 , cap(s1) = 4
	//Println(s1 == nil , len(s1) , cap(s1))
	//s1 = append(s1, 1)//len(s1) =5 , cap(s1) = 8
	//Println(s1 == nil , len(s1) , cap(s1),sli == nil ,sli , len(sli) , cap(sli))
	//var d = [10]int{1,2,3,4,5,6,7,8,9,10}
	//myTest(d , 11)
	//sum := sumArr(d)
	//Println(sum)
	//数组遍历
	//var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	//for index , v := range f{
	//	for in , v1 := range v{
	//		Println(index , v , in , v1)
	//	}
	//}

	//a := [4]int{}
	//Println(len(a) , cap(a))
	//b := a[0:1]
	//Println(b , len(b) , cap(b)) // [0] 1 4？ 为什么cap是4 len和cap的区别 len的定义是数组的可用长度 cap 是底层数组的容量
	//指针数组 [5]*int & 数组指针*[5]int
	//a := [3]*int{}
	//定义未赋值
	//var a [3]int
	//Println(a)
	////定义并赋值
	//var b = [3]int{1,3,4}
	//Println(b)
	//var c [3]string
	//Println(c)
	//var d [3]struct{}
	//Print(d)
	//var e [3]map[string]string
	//Println(e)
	//add(a)
	//Println(a)
	//p := &[3]int{}
	//Println(p)
	//addSli(p)
	//Println(p)

	//channel 引用传递
	//c := make(chan int , 10)
	//modifyChannel(c)
	//log.Println(<-c)
	//map 引用传递
	//m := make(map[string]string , 10)
	//log.Println(m)
	//modifyMap(m)
	//log.Println(m)
	////引用类型实例
	//arr := [...]int{0 , 1 , 2, 3 , 4, 5,6,7,8,9}
	//sli := arr[:]
	//log.Println(arr , sli)
	//arr[9] = 10
	//log.Println(arr , sli)
	//sli[0] = 100
	//log.Println(arr , sli)
	//sliAdd(sli)
	//log.Println(arr , sli)
	//b := new(int)
	//a := new([]int)
	//*b = 12
	//log.Println(a , b , *b)
	//var sli = make([]int , 4)
	//log.Println(sli)
	//sliAdd(sli)
	//log.Println(sli)
}

