// Author: magician
// File:   main.go
// Date:   2021/1/3
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Username string
}

type Admin struct {
	User
	title string
}

type User1 struct {
}

type Admin1 struct {
	User1
}

func (*User1) ToString() {
}

func (Admin1) test() {

}

type User2 struct {
	Username string
	age      int
}

type Admin2 struct {
	User2
	title string
}

type User3 struct {
	Name string `field:"username" type:"nvarchar(20)"`
	Age  int    `field:"age" type:"tinyint"`
}

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

type Data struct {
}

func (*Data) String() string {
	return ""
}

type Data1 struct {
	b byte
	x int32
}

type User4 struct {
	Username string
	age      int
}

type Admin3 struct {
	User4
	title string
}

type Data2 struct {
}

func (*Data2) Test(x, y int) (int, int) {
	return x + 100, y + 100
}

func (*Data2) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}
	return fmt.Sprintf(s, c)
}

func info(m reflect.Method) {
	t := m.Type
	fmt.Println(m.Name)
	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Printf("  in[%d] %v\n", i, t.In(i))
	}
	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Printf("  out[%d] %v\n", i, t.In(i))
	}
}

func Make(T reflect.Type, fptr interface{}) {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{
			reflect.MakeSlice(
				reflect.TypeOf(T),
				int(in[0].Int()),
				int(in[1].Int()),
			),
		}
	}

	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), swap)
	fn.Set(v)
}

func main() {
	var u Admin
	t := reflect.TypeOf(u)

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}

	u1 := new(Admin)
	t1 := reflect.TypeOf(u1)
	if t1.Kind() == reflect.Ptr {
		t1 = t1.Elem()
	}

	for i, n := 0, t1.NumField(); i < n; i++ {
		f := t1.Field(i)
		fmt.Println(f.Name, f.Type)
	}

	// value-interface, pinter-interface
	//var u2 Admin1
	//methods := func(t reflect.Type) {
	//	for i, n := 0, t.NumField(); i < n; i++ {
	//		m := t.Method(i)
	//		fmt.Println(m.Name)
	//	}
	//}
	//fmt.Println("---value interface---")
	//methods(reflect.TypeOf(u2))
	//fmt.Println("---pointer interface---")
	//methods(reflect.TypeOf(&u2))

	var u3 Admin2
	t2 := reflect.TypeOf(u3)
	f, _ := t2.FieldByName("title")
	fmt.Println(f.Name)
	f, _ = t2.FieldByName("User2")
	fmt.Println(f.Name)
	f, _ = t2.FieldByName("Username")
	fmt.Println(f.Name)
	f = t2.FieldByIndex([]int{0, 1})
	fmt.Println(f.Name)

	// ORM Model
	var u4 User3
	t3 := reflect.TypeOf(u4)
	f1, _ := t3.FieldByName("Name")
	fmt.Println(f1.Tag)
	fmt.Println(f1.Tag.Get("field"))
	fmt.Println(f1.Tag.Get("type"))

	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)
	m := reflect.MapOf(String, Int)
	fmt.Println(m)

	s := reflect.SliceOf(Int)
	fmt.Println(s)
	t4 := struct {
		Name string
	}{}
	p := reflect.PtrTo(reflect.TypeOf(t4))
	fmt.Println(p)

	t5 := reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(t5)

	// Implements AssignableTo, ConvertibleTo
	var d *Data
	t6 := reflect.TypeOf(d)
	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t6.Implements(it))

	var d1 Data1
	t7 := reflect.TypeOf(d1)
	fmt.Println(t7.Size(), t7.Align())
	f2, _ := t7.FieldByName("b")
	fmt.Println(f2.Type.FieldAlign())

	// Value
	u5 := &Admin3{User4{"Jack", 23}, "NT"}
	v := reflect.ValueOf(u5).Elem()
	fmt.Println(v.FieldByName("title").String())
	fmt.Println(v.FieldByName("age").Int())
	fmt.Println(v.FieldByIndex([]int{0, 1}).Int())

	// CanInterface
	u6 := User4{"Jack", 23}
	v1 := reflect.ValueOf(u6)
	fmt.Println(v1.FieldByName("Username").Interface())
	//fmt.Println(v1.FieldByName("age").Interface())
	f3 := v1.FieldByName("age")
	if f3.CanInterface() {
		fmt.Print(f3.Interface())
	} else {
		fmt.Println(f3.Int())
	}

	// reflect(array, slice, map)
	v2 := reflect.ValueOf([]int{1, 2, 3})
	for i, n := 0, v2.Len(); i < n; i++ {
		fmt.Println(v2.Index(i).Int())
	}
	fmt.Println("-------------------")
	v2 = reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	for _, k := range v2.MapKeys() {
		fmt.Println(k.String(), v2.MapIndex(k).Int())
	}

	var p1 *int
	var x interface{} = p1
	fmt.Println(x == nil)
	v3 := reflect.ValueOf(p1)
	fmt.Println(v3.Kind(), v3.IsNil())

	u7 := User4{"Jack", 23}
	v4 := reflect.ValueOf(u7)
	p2 := reflect.ValueOf(&u7)
	fmt.Println(v4.CanSet(), v4.FieldByName("Username").CanSet())
	fmt.Println(p2.CanSet(), p2.Elem().FieldByName("Username").CanSet())

	u8 := User4{"Jack", 23}
	p3 := reflect.ValueOf(&u8).Elem()
	p3.FieldByName("Username").SetString("Tom")
	f4 := p3.FieldByName("age")
	fmt.Println(f4.CanSet())
	if f4.CanAddr() {
		age := (*int)(unsafe.Pointer(f4.UnsafeAddr()))
		//age := (*int)(unsafe.Pointer(f4.Addr().Pointer()))
		*age = 88
	}
	fmt.Println(u8, p3.Interface().(User4))

	s1 := make([]int, 0, 10)
	v5 := reflect.ValueOf(&s1).Elem()
	v5.SetLen(2)
	v5.Index(0).SetInt(100)
	v5.Index(1).SetInt(200)
	fmt.Println(v5.Interface(), s1)
	v6 := reflect.Append(v5, reflect.ValueOf(300))
	v6 = reflect.AppendSlice(v6, reflect.ValueOf([]int{400, 500}))
	fmt.Println(v6.Interface())
	fmt.Println("---------------")
	m1 := map[string]int{"a": 1}
	v5 = reflect.ValueOf(&m1).Elem()
	v5.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) // update
	v5.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) // add
	fmt.Println(v5.Interface(), m1)

	// Method
	d2 := new(Data2)
	t8 := reflect.TypeOf(d2)
	test, _ := t8.MethodByName("Test")
	info(test)
	sum, _ := t8.MethodByName("Sum")
	info(sum)

	d3 := new(Data2)
	v7 := reflect.ValueOf(d3)

	exec := func(name string, in []reflect.Value) {
		m := v7.MethodByName(name)
		out := m.Call(in)

		for _, v7 := range out {
			fmt.Println(v7.Interface())
		}
	}

	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
	fmt.Println("-----------------")
	exec("Sum", []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})

	// CallSlice
	d4 := new(Data2)
	v8 := reflect.ValueOf(d4)
	m2 := v8.MethodByName("Sum")
	in := []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf([]int{1, 2}),
	}
	out := m2.CallSlice(in)
	for _, v8 := range out {
		fmt.Println(v8.Interface())
	}

	// Make
	//var makeints func(int, int) []int
	//var makestrings func(int, int) []string
	//Make(Int, &makeints)
	//Make(String, &makestrings)
	//x1 := makeints(5, 10)
	//fmt.Printf("%#v\n", x1)
	//s11 := makestrings(3, 10)
	//fmt.Printf("%#v\n", s11)
}
