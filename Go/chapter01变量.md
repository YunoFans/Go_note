# 一、声明与赋值

```go
(1)指定变量类型
var i int//不赋值就使用默认值


(2)自动推导类型
var i  = 100//根据赋予的值进行类型推导


(3)声明于赋值同时进行
name := "Lucy"


(4)多变量声明
var name, sex, age int
fmt.Println(name, sex, age)

var name, sex, age = "Lucy", true, 18
fmt.Println(name, sex, age)

name, sex, age := "Lucy", "男", 19
fmt.Println(name, sex, age)

```



**声明全局变量：函数外**

```go


（1）
var (
	name = "Lisi"
	age  = 90
)

（2）
var name = "Lucy"
var sex = "男"
var age = 90

（3）
var name, sex, age = "Lucy", true, 18


```



# 二、使用注意事项

- 变量在同一个数据类型中可以不断变换
- 变量在同一个作用域内不允许重复定义
- **只有类型一样才能做`+`运算，这点与java不同**



