package main

import (
	//"fmt"
    // "time"
	// "first/test"
	// "first/vgarcia"
	//  "os"
	//  "bufio"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
  // Create a Gin router with default middleware (logger and recovery)
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H {
        "message" : "def",
    })
  })
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  // port 8080 (default)
  r.Run()
}

// func main() {
//     fmt.Println("test")
//     for _, arg := range os.Args {
//         fmt.Printf(" %s\n", arg)
//     }
//     x := "test"
//     var yes bool
//     var flaot float64
//     fmt.Printf("%T\n", x)
//     fmt.Println(yes)
//     fmt.Println(flaot)
//         num := -4
//     if num > 0 {
//         fmt.Println("positif")
//     } else if num < 0 {
//         fmt.Println("négatif")
//     } else {
//         fmt.Println("zéro")
//     }
//     defer fmt.Println("defer")
//     var str string
//     var truc bool
//     test, testv2 := fmt.Scanln(&str, &truc)
//     println(str)
//     println(test, testv2)

//     reader := bufio.NewReader(os.Stdin)
//     fmt.Print("txt : ")
//     line, _ := reader.ReadString('\n')
//     fmt.Println("result :", line)
// }



// type rec struct {

//     width, height int
// }


// func sumAndDiff(a, b int) (int, int) {
//     return a +b, a - b
// }

// func (r *rec) area() int {

//     r.height = 56
//     return r.height * r.width
// }

// func main() {

//     names := []string{"ricko", "pierrick", "pied"}
//     fmt.Println("name:", names)

//     names = append(names, "test", "testv2", "testv3")
//     fmt.Println("after append", names)

//     names = append(names[:1], names[2:4]...)
//     fmt.Println("after removing", names)

//     ages := make(map[string]int)
//     ages["test"] = 2
//     ages["testv2"] = 454545
//     ages["testv3"] = 1111111
//     ages["testv4"] = -999

//     for name, age := range ages {
//         fmt.Printf("name: %s, age: %d\n", name, age)
//     }

//     x, y := sumAndDiff(5, 6)
//     fmt.Println(x, y)

//     rect := rec{width: 5, height: 6}
//     j := rect.area()

//     fmt.Println(j)
//     fmt.Printf("Height: %d, Width: %d\n", rect.height, rect.width)


//     for i := 0; i < 50; i++ {
//         fmt.Printf("%s\n", names[i])
//     }
//     // var val *int

//     // val = nil
//     // val = new(int)
//     // fmt.Printf("%d", val)
// }


// func counter() func() int {
//     count := 0
//     return func() int {
//         count++
//         return count
//     }
// }


// func test(fd chan int) {
//     defer close(fd)
//     defer fmt.Println("testv1")
//     fmt.Println("testv2")
// }

// func main() {

//     c := counter()

//     fmt.Println(c())
//     fmt.Println(c())
//     fmt.Println(c())
  //   names := []string{"ricko", "pierrick", "feet"}

  //   c := make(chan int, 2)
  //   defer func() {
  //   if r := recover(); r != nil {
  //     fmt.Println("Recovered:", r)
  //     test(c)
  //   }
  // }()
  // fmt.Println("test")
  //   for i := 0; i < 50; i++ {
  //       fmt.Printf("%s\n", names[i])
  //   }
  //panic("This is a panic")
// }


// func test(val []int, ch chan int) {

//     for i := range val {
//         ch <- val[i] * val[i]
//     }
//     close(ch)
// }

// func recieve(ch chan int) {

//     for num := range ch {
//         fmt.Println(num)
//     }
// }

// func main() {

//     val := []int{1,2,3,4,5}
//     numbersChannel := make(chan int)

//     go test(val, numbersChannel)
//     recieve(numbersChannel)
// }

// type bjr struct {
//     _jk int
// }

// type beta interface {
//     joke() int
// }

// type alpha interface {
//     Speak() string
// }

// type Test struct {
//     _name string
//     _age int
// }

// func (t bjr)Speak() string{
//     return "fhsdfgsfhsdf"
// }

// func (t Test)Speak() string{
//     return "fhsdfgsfhsdf"
// }

// func (t Test) Hello() {

//     fmt.Println("hello", t._age, "name:", t._name)
// }
// func main() {

//     yes := Test{_name : "non", _age : 56 }
//     yes.Hello()

//     var inter alpha
//     var ko alpha

//     ko = bjr{56}

//     ko = Test{"51", 23}
//     fmt.Println(ko)
//     inter = Test{_name: "testo", _age: 6}
//     fmt.Println(inter)

// }