package main

import (
	"fmt"

	examplepb "example.com/gazelle-example/myapp/proto/v1"

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
	v := &examplepb.Value{
		Value: &examplepb.Value_IntVal{
			IntVal: 100,
		},
		Sub: &examplepb.Message{
			Name: "foo",
			Id:   "bar",
		},
	}
	respAny, _ := ptypes.MarshalAny(v)

	pb := &examplepb.Thing{
		Payload: respAny,
	}
	fmt.Println(pb)
	tm := proto.TextMarshaler{
		ExpandAny: true,
	}
	b := tm.Text(pb)
	fmt.Println(string(b))
}
