

package main
import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"

)

type Foo struct{
	info data `json:"data"`
}

type data struct {
	id string `json:"id"`
	type1 string `json:"type"`
	attributes FirstStruc `json:"attributes"`
}
type FirstStruc struct{
	value SecondStruc `json:"value"`
	u32 int `json:"u32"`
}
type SecondStruc struct{
	type2 ThirdStruc `json:"type"`
}
type ThirdStruc struct{
	value int `json:"value"`
	name string `json:"name"`
}

type Item struct {
	foos [10] Foo
}
func main() {
	response, err := http.Get("http://localhost:8000/_/api/v3/key_values")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()



	dataInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return
	}


	array:=Item{}
    Datan:=Foo{}


	json.Unmarshal(dataInBytes, &Datan)
	fmt.Println(Datan.info.id)

	err0:= json.Unmarshal([]byte(dataInBytes), &Datan)
	if err0 != nil {
		fmt.Println(err0)
		return
	}
	fmt.Println(Datan)

	err1:= json.Unmarshal([]byte(dataInBytes), &array)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(array.foos[0])
	//for _, t := range array {
	//	fmt.Println(t, "-")
	//}



}
