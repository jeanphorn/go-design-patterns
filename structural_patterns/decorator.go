package main

import "fmt"

type Object func(string) string

func Decorate(fn Object) Object {
    return func(base string) string {

        ret := fn(base)

        ret = ret + " and Tshirt"
        return ret
    }
}

func Dressing(cloth string) string {
    return "dressing " + cloth
}

func main() {
    f := Decorate(Dressing)

    fmt.Println(f("shoes"))
}
