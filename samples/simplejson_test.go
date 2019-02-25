package samples

import (
	"bufio"
	"io"
	"os"

	//"strconv"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
)

func TestGjson(t *testing.T) {
	const body = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		panic(err.Error())
	}
	first0 := js.GetPath("name", "first").MustString()
	t.Log(first0)
	first1 := js.Get("name").Get("first").MustString()
	t.Log(first1)
	age := js.Get("age").MustInt()
	t.Log(age)

	//strconv.FormatInt()
}

func TestReadLine(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		t.Log(line)
	}
}
