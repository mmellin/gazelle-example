package main

import (
	"fmt"

	"example.com/example/myapp/proto/example"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

// Example output
/*
payload: <
  [type.googleapis.com/example.Value]: <
    int_val: 100
    sub: <
      name: "foo"
      id: "bar"
    >
  >
>
*/

func main() {
	v := &example.Value{
		Value: &example.Value_IntVal{
			IntVal: 100,
		},
		Sub: &example.Message{
			Name: "foo",
			Id:   "bar",
		},
	}
	respAny, _ := ptypes.MarshalAny(v)

	pb := &example.Thing{
		Payload: respAny,
	}
	fmt.Println(pb)
	tm := proto.TextMarshaler{
		ExpandAny: true,
	}
	b := tm.Text(pb)
	fmt.Println(string(b))
}
