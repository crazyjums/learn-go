package learn_map

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

// InitMapByLiteral 通过字面量初始化
func InitMapByLiteral() {
	//通过字面量初始化
	strMap := map[string]string{
		"name": "zs",
		"age":  "18",
	}

	fmt.Println(strMap)

	hash := map[string]string{}
	hash["a"] = "aa"
	fmt.Println(hash, len(hash))
}

func InitMapByMake() {
	m := make(map[string]string, 6)
	fmt.Println(len(m))
}

func ReadValue() {
	hash := map[string]string{
		"name": "zs",
		"age":  "18",
	}

	//v := hash["name"]
	//fmt.Println(v)
	_, ok := hash["name"]
	fmt.Println(ok)
	hash["hh"] = ""
	_, ok = hash["h1h"]
	fmt.Println("ok=", ok)
}

type any = interface{}
type H map[string]any

func ValueIsAny() {
	var m H = make(H)
	m["name"] = "zs"
	m["age"] = 17

	m1 := H{
		"errno":  1,
		"errmsg": "success",
		"md5":    md5.Sum(([]byte)("aaa")),
		"data": H{
			"name": "jums",
			"age":  19,
		},
	}

	b, _ := json.Marshal(m1)

	fmt.Println(m, m1, string(b))
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//InitMapByLiteral()
	//ReadValue()
	//InitMapByMake()
	//ValueIsAny()

}
