package main

import (
	"encoding/json"
	"fmt"
)

// type Number interface {
//     int | int8 | int16 | int32 | int64 |
//     uint | uint8 | uint16 | uint32 | uint64 |
//     float32 | float64 | string
// }

// func Max[T Number](values ...T) T {
//     if len(values) == 0 {
//         var zero T
//         return zero
//     }
//     max := values[0]
//     for _, v := range values[1:] {
//         if v > max {
//             max = v
//         }
//     }
//     return max
// }
type User struct{
    Name string `json:"name"`
}

type UserInfo struct{
    Name string `json:"name"`
    Age int `json:"age"`
}


type Response[T any] struct{
    Code int `json:"code"`
    Msg string `json:"msg"`

    Data T `json:"data"`
}


func main() {
    // user := Response{
    //     Code: 200,
    //     Msg: "success",
    //     Data: User{
    //         Name: "小明",
    //     },
    // }
    
    // userinfo := Response{
    //     Code: 200,
    //     Msg: "success",
    //     Data: UserInfo{
    //         Name: "小明",
    //         Age: 18,
    //     },
    // }

    // bytedata, _:= json.Marshal(userinfo)
    // fmt.Println(string(bytedata))
    var userResponse Response[User]

    json.Unmarshal([]byte(`{"code":200,"msg":"success","data":{"name":"小明"}}`),&userResponse)
    fmt.Println(userResponse.Data)
}
