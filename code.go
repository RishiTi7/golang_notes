// package main

// // standard library fmt
// import "fmt"

// func main(){
// 	fmt.Println("hello Rishikar welcome to golang")
// }

// package main

// import (
// 	"fmt"
// )
//we have to declare variables in go the way we speak eg theres a var i with int .....
// func main(){
// 	var i int
// 	i = 21
// 	fmt.Println(i)
// }

// func main(){
// 	var i int = 23
// 	fmt.Println(i)
// }

// func main(){
// 	var i int
// 	i = 32
// 	var j int = 23
// 	k:= 23
// 	fmt.Printf("%v,%T",i,i)
// }

// var i int = 32
// func main (){
// 	fmt.Println(i)
// }

// one way

// var actorname string = "tom cruise"
// var sidechar string = "rdj"
// var doctornumber int = 22
// var season int = 32

// func main(){
// 	fmt.Printf("%v,%T", sidechar, sidechar)
// }

// var(
// 	actorname string = "tom cruise"
//     sidechar string = "rdj"
// 	doctornumber int = 22
// 	season int = 32
// )

// var(

// 	counter int = 21
// )

// func main(){
// 	fmt.Printf("%v,%T", sidechar, sidechar)
// }

// the var at the package level gets less precedence and the package at the function level gets more precedence this is called shadowing
// the print statement before and after displays accordingly

// naming conventions in go is of two types one is if the var is lowercase then the var is targeted to that package if the var if uppercase then the var is targeted to the global level visiblility if of global if lower then the visiblility is to that package only
//then the block scope
//the length of the variable is also imp if short then shorter lifespan if long then longer lifespan
//acronyms should be uppercase

//conversion in go

// func main(){
// 	var i int = 42
// 	fmt.Printf("%v,%T\n",i,i)

// 	var j float32
// 	j = float32(i) //here this float is acting like a function which converts
// 	fmt.Printf("%v,%T\n",j,j)
// }
// strconv package converts string
// cant redeclare var, but can shadow them

//primitive data types
//boolean

// func main(){
// 	n:= 1==1
// 	m:= 1==2
// 	fmt.Printf("%v,%T\n",n,n)
// 	fmt.Printf("%v,%T",m,m)
// }

//unassigned alues in bool is 0 so false

//ARITHMETIC
// func main(){
// 	a:=2
// 	b:=12
// 	fmt.Println(a+b)
// 	fmt.Println(a*b)
// 	fmt.Println(a-b)
// 	fmt.Println(a/b)
// 	fmt.Println(a%b)

// }

// 	func main(){
// 		var a int = 10
// 		var b int8 = 21
// 		// fmt.Println(a + b)
// 		//we have to do implicit data type converions
// 	//mismatched integers do it doesnt work
// 	fmt.Println(a + int(b))

// }

//BITWISE
// func main(){
// 	a:=10 // 1010
// 	b:=3 //  0011
// 	fmt.Println(a&b) // 0010
// 	fmt.Println(a|b) // 1011
// 	fmt.Println(a^b) // 1001
// 	fmt.Println(a&^b) // 0100
// }

//BIT SHIFTING
// func main(){
// 	a:=8  // 2^3
// 	fmt.Println(a<<3) // 2^3 * 2^3 = 2^6
// 	fmt.Println(a>>3) // 2^3 /2^3 = 2^0
// }

// go gives more precedence to complex numbers, because of this it gets used in data science

// func main(){
// 	var n complex128 = complex(5,12)
// 	fmt.Printf("%v,%T", n,n)
// }

//strings in go acts as arrays but with byte alias string8
//cant be immutable
// func main(){
// 	s:= "hello rishikar"
// 	fmt.Printf("%v,%T",string(s[2]),s[2])
// }

//concat
// func main(){
// 	s:= "hello rishikar "
// 	s1:= "this is go lang"
// 	fmt.Printf("%v,%T",s+s1,s+s1)
// }
// func main(){
// 	s:= "hello rishikar "
// 	b:= []byte(s)
// 	fmt.Printf("%v,%T",b,b)
// }

//rune is integer 32

// func main(){
// 	var r rune = 'a'
// 	fmt.Printf("%v,%T",r,r)
// }

//constant
//internal constant

//types consts
// func main(){
// 	const myCosnt int = 21
// 	// myConst = 21 we cant do that coz const
// 	fmt.Printf("%v,%T", myCosnt,myCosnt)
// }

//const and variable are the same types with same data types

//enumerated consts

// iota is used here it can act as counter **

// const(
// 	a = iota
// 	b
// 	c
// )
// const(
// 	a = iota
// 	b = iota
// 	c = iota
// )

// const (
// 	a2 = iota
// )
// func main(){
// 	fmt.Printf("%v\n",a)
// 	fmt.Printf("%v\n",b)
// 	fmt.Printf("%v\n",c)
// 	fmt.Printf("%v\n",a2)
// }

//error keyword is used for reassigning value to zero

//Arrays
// package main

// import(
// 	"fmt"
// )

// func main(){
// 	// grades := [...] int{1,2,3}
// 	grades := [3] int{1,2,3}
// 	fmt.Printf("%v", grades)
// }

// func main(){
// 	var students[3]string
// 	fmt.Printf("students:%v\n", students)
// 	students[0] = "lisa"
// 	students[1] = "rishi"
// 	students[2] = "kar"
// 	// fmt.Printf("students:%v", students[1])
// 	fmt.Printf("students:%v", len(students))
// }

//array inside array
// func main(){
// 	var identityMatrix [3][3]int
// 	identityMatrix[0] = [3]int{0,1,1}
// 	identityMatrix[0] = [3]int{0,0,1}
// 	identityMatrix[0] = [3]int{1,0,0}
// 	fmt.Println(identityMatrix)
// }

//arrays re generally considered as values but not in go
//in go it copies elements

// func main(){
// 	a:=[...]int{2,3,4}
// 	b:=a
// 	b[1]=2
// 	fmt.Println(a)
// 	fmt.Println(b)
// b is copies to a

// }

//for simple arrays its fine but for higher value arrays its tough so we use pointers

// pointers
// func main(){
// 	a:=[...]int{2,3,4}
// 	b:=&a
// 	b[1]=2
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// b is pointing to data that a has

//slice
// func main(){
// 	a:=[]int{2,3,4}
// 	fmt.Println(a)
// 	fmt.Println(len(a))
// 	fmt.Println(cap(a)) //capacity
// }

// func main(){
// 	a:=[]int{1,2,3,4,5,6,7,8,9,10}
// 	b:=a[:]
// 	c:=a[3:]
// 	d:=a[:6]
// 	e:=a[3:6]
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// }

// func main(){
// 	a:=make([]int,3,100)
// 	fmt.Println((a))
// 	fmt.Println(len(a))
// }

// func main(){
// 	a:=[]int{}
// 	// fmt.Println(a)
// 	// fmt.Println(len(a))
// 	// fmt.Println(cap(a))
// 	a = append(a, 1)
// 	// fmt.Println(a)
// 	// fmt.Println(len(a))
// 	// fmt.Println(cap(a))
// 	a = append(a, []int{2,3,4,5,6}...) // slice concatenation
// 	fmt.Println(a)
// 	fmt.Println(len(a))
// 	fmt.Println(cap(a))
// }

//putting stack

// func main(){
// 	a:=[]int{1,2,3,4,5}
// 	b:=append(a[:2], a[3:]...)
// 	fmt.Println(b)
// }

/// [....]this is best and robust

// underline of slice is an array

//MAPS

// there is no order in map unlike array, slice etc

// func main(){
// 	statePopulation := make(map[string]int,10)
// 	// statePopulation := map[string]int{ //another way
// 	statePopulation = map[string]int{
// 		"california": 12432335,
// 		"texas": 133215135,
// 		"florida": 978235238,
// 		"new ha": 325235325,
// 		"ohio": 798629700,
// 		"iowa": 464464646,

// 	}
// 	statePopulation["georgia"] = 10000022
// 	delete(statePopulation,"georgia")
// 	fmt.Println(statePopulation["georgia"])

// 	pop, ok := statePopulation["ohio"]
// 	fmt.Println(pop,ok) //this corrects error or mismatches
// }

//struct
//collection of multiple datatypes

// type Doctor struct{
// 	Number int
// 	ActorName string
// 	Companion []string
// 	Episodes []string
// }

// func main(){
// 	aDoctor := Doctor{
// 		Number: 3,
// 		ActorName: "rishi",
// 		Companion: []string{
// 			"kar",
// 			"had",
// 			"adad",
// 		},
// 		Episodes: []string{
// 			"1",
// 			"2",
// 			"3",
// 		},
// 	}
// 	fmt.Println(aDoctor)
// }

//another way

// func main(){
// 	aDoctor:=struct{name string}{name: "rishikar rishi"}
// 	fmt.Println(aDoctor)
// }

//GO tradionally doesnt have inheritance

//Go uses compostion through embedding similar to inheritance

// type Animal struct {
// 	Name string
// 	Origin string
// }
// type Bird struct{
// 	Animal
// 	Speedkph float32
// 	Canfly bool
// }

// func main(){
// 	b:= Bird{
// 	Animal: Animal{Name: "emu", Origin: "india"},
// 	// b.Name = "emu"
// 	// b.Origin = "india"
// 	Speedkph: 20.33,
// 	Canfly: true,
// 	}

// 	fmt.Println(b)
// }

//another way

// func main(){
// 	b:= Bird{}
// 	b.Name = "emu"
// 	b.Origin = "india"
// 	b.Speedkph = 20.33
// 	b.Canfly = true
// 	fmt.Println(b)
// 	}

//TAGS
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Animal struct {
// 	Name string `required max: "100"`
// 	Origin string
// }
// func main(){
// 	t:= reflect.TypeOf(Animal{})
// 	field, _:=t.FieldByName("Name")
// 	fmt.Println(field.Tag)
// }

// package main

// import(
// 	"fmt"
// )

// func main(){
// 	if true{
// 		fmt.Println("the test is true")
// 	}
// }

// func main(){
// 	    statePopulations := map[string]int{
// 			"california": 12432335,
// 			"texas": 133215135,
// 			"florida": 978235238,
// 			"new ha": 325235325,
// 			"ohio": 798629700,
// 			"iowa": 464464646,
// 		}
// 		if pop, ok := statePopulations["florida"]; ok{
// 			fmt.Println(pop)
// 		}
// }

//this will work only for int types not string types.

// func main(){
// 	number:= 21
// 	guess:= -10

// 	if guess < 1 || guess >100{
// 		fmt.Println("the guess must be in 1 to 100")
// 	}
// 	if guess>=1 && guess<=100{
// 	if guess > number{
// 		fmt.Println("too high")
// 	}
// 	if guess < number{
// 		fmt.Println("too low")
// 	}
// 	if guess == number{
// 		fmt.Println("equal")
// 	}
// 	fmt.Println(number<= guess, number>=guess, number!=guess)
//     }
// }

//else if

// func main(){
// 	number:= 21
// 	guess:= -10

// 	if guess < 1 {
// 		fmt.Println("the guess must be greater than 1")
// 	} else if guess > 100{
// 		fmt.Println("the guess imust be less than 100")
// 	}else{
// 		if guess>=1 && guess<=100{
// 		if guess > number{
// 			fmt.Println("too high")
// 		}
// 		if guess < number{
// 			fmt.Println("too low")
// 		}
// 		if guess == number{
// 			fmt.Println("equal")
// 		}
// 		fmt.Println(number<= guess, number>=guess, number!=guess)
// 		}
// 	}
// }

//Switch

//  we can use fallthrough if we use this we get the next statement executed but it is logic less no matter what it gets executed not often used generally tag is used
// func main(){
// 	switch 2 {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	default:
// 	    fmt.Println("1 or 2")
// 	}
// }
//break is used as short circuiting

//you can put any type in interface

// func main(){
// 	var i interface{} = []int{}
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("i is an int")
// 	case float64:
// 		fmt.Println("i is an float")
// 	case string:
// 		fmt.Println("i is an string")
// 	default:
// 		fmt.Println("i is an another type")
// 	}
// }

//Looping
//for loop

// func main(){
// 	for i:= 0; i <= 5; i++{
// 		fmt.Println(i)
// 	}
// }

// there is no comma operator in go so we cant use different logic in the same line
// func main(){
// 	for i:= 0; i <= 5; i = i+2{
// 		fmt.Println(i)
// 	}
// }
// func main(){
// 	for i, j:= 0,0 ; i <= 5; i,j = i+1, j+2{
// 		fmt.Println(i,j)
// 	}
// }

//continue breaks out of current iteration
// func main(){
// 	for i := 0 ; i<5 ;i++{
// 		if i%2 ==0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

//here loop is a tag if we dont attach loop here if just exits out of on loop or the closest loop but loop tag is attached to outer loop o it exits that as well

// func main(){
// Loop:
// for i:=1 ; i<=3; i++{
// 	for j:=1; j<= 3; j++{
// 		fmt.Println(i*j)
// 		if i*j >=3{
// 			break Loop
// 		}
// 	}
// }

// }

//in order to work for loop in collections we gotta use fo range keyword
// func main(){
// 	s:= []int{1,2,3}
// 	for k, v := range s{
// 		fmt.Println(k,v)
// 	}
// }
// this returns index and value

// _ is a write operator which tells the compiler that we dont need that thing right now

//  defer
//  panic
//  recover

//defer keyword works as it waits till the current function completes but executes before the main function

// func main(){
// 	fmt.Println("start")
// 	defer fmt.Println("middle")
// 	fmt.Println("end")
// }
// it works in LIFO
// func main(){
// 	defer fmt.Println("start")
// 	defer fmt.Println("middle")
// 	defer fmt.Println("end")
// }

// package main

// import(
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main(){
// 	res, err := http.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	robots, err := ioutil.ReadAll(res.Body)
// 	res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", robots)
// }

// panic
// there are no exceptions in GO because its considered as common errors
// we use panic instead of exceptions

// first deferr then panic gets executed
// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main(){
// 	fmt.Println("start")
// 	panic("fuck u")
// 	fmt.Println("bad")
// }

//pointers
// b is pointer to integer and it is pointing to address a

// func main(){
// 	var a int = 43
// 	var b *int = &a
// 	fmt.Println(&a,b)
// 	fmt.Println(a,*b) // derefencing it
// 	a = 21
// 	fmt.Println(a,*b)
// }
// we cant do traditional pointer arithmetic in go

// 	func main(){
// 	a:=[]int{1,2,3}
// 	b:=&a[0]
// 	// c:=&a[1] - 4 // we cant do that here
// 	c:=&a[1]

// 	fmt.Printf("%v,%p,%p\n", a,b,c)
// }

//unsafe package used for pointer arthmetic

// we can initialize variables with pointers withobjects
// func main(){
// 	var ms *myStruct
// 	ms = &myStruct{foo:42}
// 	fmt.Println(ms)

// }
// type myStruct struct{
// 	foo int
// }
// we cant intialize in object initialization in new keyowrd
// func main(){
// 	var ms *myStruct
// 	ms =new(myStruct)//returns zero
// 	fmt.Println(ms)

// }
// type myStruct struct{
// 	foo int
// }

//so

// func main(){
// 	var ms *myStruct
// 	ms =new(myStruct)
// 	ms.foo = 32  //using the underlying object
// 	fmt.Println(ms.foo)

// }
// type myStruct struct{
// 	foo int
// }

//another way
// (*ms).foo =42
// (*ms).foo

//slices and maps points to the underlying data. they are pointing to pointers to the underlying data. so the entire thing gets changed.

//functions

//pascal case or camel case.
// uppercase = global lowecase = that local

// passing parameters
// func main(){
// 	sayMessage("hello go")
// }

// func sayMessage(msg string){
// 	fmt.Println(msg)
// }

//variadic parameters
//but ht end should have intrinsic values
// func main(){
// 	sum( "the sum is ", 1,2,3,4,5)
// }
// func sum(msg string, values ...int){
// 	fmt.Println(values)
// 	result :=0
// 	for _, v := range values{
// 		result += v
// 	}
// 	fmt.Println(msg,result)

// }

// error parameter

// func main(){
// 	d,err := divide(5.0, 0.0)
// 	if err != nil{
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)
// }

// func divide(a,b float64)(float64, error){
// 	if b ==0.0{
// 		return 0.0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a/b, nil
// }

// Anonymous functions, here they are not traditionl functions but more of a type functions

// func main(){
// 	func(){
// 		fmt.Println("rishi")
// 	}() // invoke function
// }

// function defined as a variable

// func main(){
// 	var divide func(float64, float64)(float64, error)
// 	divide = func(a,b float64)(float64,error){
// 		if b ==0.0{
// 			return 0.0, fmt.Errorf("fuck you cannot divide by zero")
// 		}else{
// 			return a/b, nil
// 		}
// 	}
// 	d,err := divide(5.0,0.0)
// 	if err!= nil{
// 		fmt.Println(err)
// 	}
// 	fmt.Println(d)

// }

//greeter is an object greet is an method
//method also same kinda like a function
// func main(){
// 	g:= greeter{
// 		greeting: "hello!",
// 		Name: "go",
// 	}
// 	g.greet()
// }
// type greeter struct{
// 	greeting string
// 	Name string
// }
// func (g greeter) greet(){
// 	fmt.Println(g.greeting, g.Name)
// }

//Interfaces

//interfaces dont describe data they describe behaviors

// we gonna store method defintions
// func main() {
// 	// fmt.Println("hello")
// 	var w Writer = ConsoleWriter{}
// 	w.Write([]byte("hello rishi"))
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type ConsoleWriter struct{}

// func (cw ConsoleWriter) write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

// func main(){
// 	myInt := IntCounter(0)
// 	var inc Incrementer = &myInt
// 	for i :=0; i<10; i++{
// 		fmt.Println(inc.Increment())
// 	}
// }

// //increment is a method and interface implementing that method
// type Incrementer interface{
// 	Increment() int
// }
// type IntCounter int
// func (ic *IntCounter) Increment() int{
// 	*ic++
// 	return int(*ic)
// }

//an empty interface is something which doesnt have any methods

//usual thing is listing out different types with an empty interface used in type switches

//multiple interfaces implementing two interfaces in one

// Go routines for concurrent and parallel programming

//go is an abstraction of threads used by OS so we can run this small multiple high level threads instead of one low level thread
// func main(){
// 	go sayHello()
// 	time.Sleep(100* time.Millisecond)
// }

// func sayHello (){
// 	fmt.Println("say hello")
// }

//do not use realworls time to execute the go let it use its own application time

// var wg = sync.WaitGroup{}

// func main(){
// 	var msg = "hello"
// 	wg.Add(1)
// 	go func (msg string){
// 		fmt.Println(msg)
// 		wg.Done()
// 	}(msg)
// 	wg.Wait()
// }

//theres no sync in this eg every go is running differently every one is running fast as they can to complete the application
// var wg = sync.WaitGroup{}
// var counter = 0

// func main(){
// 	for i :=0; i <10; i++{
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }
// func sayHello(){
// 	fmt.Println("hello\n", counter)
// 	wg.Done()
// }
// func increment(){
// 	counter++
// 	wg.Done()
// }

// so we use mutex here the mutex allows multiple read actions but no write actions until one is finished so it outs a lock until one is done

// func main(){
// 	runtime.GOMAXPROCS(1)  //to set the threads if we set this to a higher number than the performance peaks and falls off coz the scheduler has to schedule continioulsy that looses performance
// 	fmt.Printf("threads: %v\n",runtime.GOMAXPROCS(-1))
// } 

// Channels

// package main

// import (
// 	"fmt"
// 	"sync"   
// )

// var wg = sync.WaitGroup{}
// func main(){
// 	ch := make(chan int)
// 	wg.Add(2)

// 	//data recieving from the channel
// 	go func (){
// 		i:= <- ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}()

// 	// putting data into the channel data flow to channel sender
// 	go func ()  {
// 		ch <-42
// 		wg.Done()
		
// 	}()
// 	wg.Wait()

// }

//multiple senders but only one reciever so it doesnt process all of them


// var wg = sync.WaitGroup{}


// func main(){
// 	ch:= make(chan int)
// 	go func(){
// 		i:= <-ch 
// 		fmt.Println(i)
// 		wg.Done()
// 	}()
// 	for j:=0; j<5; j++{
// 		wg.Add(2)

// 		go func(){
// 			ch<- 42
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()

// }

// func main(){
// 	ch:= make(chan int)
// 	wg.Add(2)
// 	go func(ch <-chan int){  // recieve only
// 		i:= <- ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}(ch)
// 	go func (ch chan<- int){ //send only
// 		ch <- 42
// 		wg.Done()
// 	}(ch)

// 	wg.Wait()
	
// }

// problem of one reciever and muliple senders

// func main(){
// 	ch:= make(chan int,50) // 50 here creates a buffer and stores the desired amount
// 	wg.Add(2)
// 	go func(ch <-chan int){  // recieve only
// 		i:= <- ch
// 		fmt.Println(i)
// 		i = <- ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}(ch)
// 	go func (ch chan<- int){ //send only
// 		ch <- 42
// 		ch <- 29
// 		wg.Done()
// 	}(ch)

// 	wg.Wait()	
// }


//looping channels

// func main(){
// 	ch:= make(chan int,50) // 50 here creates a buffer and stores the desired amount
// 	wg.Add(2)
// 	go func(ch <-chan int){
// 		for i := range ch {
// 			fmt.Println(i)
// 		}
// 		wg.Done()
// 	}(ch)
// 	go func (ch chan<- int){ //send only
// 		ch <- 42
// 		ch <- 29
// 		close(ch) // to close the channel so the for loop doesnt iterate
// 		wg.Done()
// 	}(ch)

// 	wg.Wait()	
// }
//but in this we solved the deadlock in sender but we didnt exit the for loop so we still waiting for senders but there isnt any so its deadlocking

// we can do this coming out of loop with OK syntax aswell


// func main(){
// 	ch:= make(chan int,50) // 50 here creates a buffer and stores the desired amount
// 	wg.Add(2)
// 	go func(ch <-chan int){
// 		for {
// 			if i, ok := <-ch;ok{  // here ok returns a boolean 
// 				fmt.Println(i)
// 			}else{
// 				break
// 			}
// 		}
// 		wg.Done()
// 	}(ch)
// 	go func (ch chan<- int){ //send only
// 		ch <- 42
// 		ch <- 29
// 		close(ch) // to close the channel so the for loop doesnt iterate
// 		wg.Done()
// 	}(ch)

// 	wg.Wait()	
// }


//Select statements

// var donech = make(chan struct{}) // signal only channel zero memory allocations but lets the people know what the message is

// select lets the messages to signal after the signaling it shuts down
// by passing case <-donech: break
//blocking select statement
// syntax for passing done is donech <- struct {}{}

