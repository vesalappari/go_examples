package main

import "fmt"

func main() {
    var b int = 15
    var a int
    var t int = 10

    // Short cut for initialization of an array. Uo can initialize it partially

    numbers := [6]int{1,2,3,4}

    for a := 0; a < 10; a++ {
    fmt.Printf("Value of a:%d\n",a)
    }

    for a < b {
        a++
        fmt.Printf("Value of a:%d\n",a)
    }

    for i,x := range numbers {
    fmt.Printf("Value of x = %d at %d\n", x,i)
    }

    if t > 5 {
        fmt.Printf("%d is bigger than five \n", t)
    } else {
        fmt.Printf("%d is smaller or equal to five\n",t)
    }

    j := 2
    fmt.Print("Write ",j," as ")
    switch j {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("tThree")
        default:
            fmt.Println("no case for the value")
    }

    whatAmI := func (i interface{}) {
        switch t := i.(type) {
            case bool:
                fmt.Println("I am a boolean")
            case int:
                fmt.Println("I am an integer")
            default:
                fmt.Printf("Do not recognize type %T\n", t)
        }
    }

    whatAmI(true)
    whatAmI(5)
    whatAmI("hei")
}