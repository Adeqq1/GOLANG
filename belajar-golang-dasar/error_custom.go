package main

import "fmt"

type validationError struct {
    Message string
}

func (v *validationError) Error() string {
    return v.Message
}

type notFoundError struct {
    Message string
}

func (n *notFoundError) Error() string {
    return n.Message
}

func SaveData(id string, data any) error {
    if id == "" {
        return &validationError{"validation error"}
    }

    if id != "Ade" {
        return &notFoundError{"data not found"}
    }

    return nil
}

func main() {
    err := SaveData("Ade", nil)
    // if err != nil {
    //     // terjadi error
    //     if validationErr, ok := err.(*validationError); ok{
    //         fmt.Println("Validation Error:", validationErr.Error())
    //     }else if notFoundErr, ok := err.(*notFoundError); ok {
    //         fmt.Printnl ("not found error")
    //     }
    //     } else {
    //         // sukses
    //         fmt.Println("Sukses")
    //     }
.
    switch finalError := err.(type) {
        case *validationErrror
        fmt.Println("Validation eror"), finalError.Error()
        case *notFoundError
        fmt.Println("not found error")
    default:
        fmt.Println("Unkown error")
    }
    
    
    fmt.Pritnl("Validation eror"), finalError.Error()
}