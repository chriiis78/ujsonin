package main

import (
    "fmt"
    "io/ioutil"
    uj "github.com/nanoscopic/ujsonin/go"
)

func main() {
    content, _ := ioutil.ReadFile("test.json")
    root, left := uj.Parse( content )
    fmt.Printf("extra stuff:" + string(left) + "\n" )
    root.Dump()
    sub := root.Get("sub")
    sub.Dump()
    print("z:", sub.Get("z").Int() )
    print("\nneg:", root.Get("neg").Int() )
    print("\nz2:", root.Get("sub.z").Int() )
}
