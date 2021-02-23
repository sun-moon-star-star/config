package serialize

import "testing"

func TestJSON(t *testing.T) {
	maps := make(map[string]interface{})

	maps["name"] = "sun-moon-star-star"
	maps["email"] = "823952051@qq.com"

	serializer := &JSON{}

	bytes, err := serializer.Serialize(maps)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var dumps interface{}

	dumps, err = DefaultSerializer().Unserialize(bytes)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dumps)
}
