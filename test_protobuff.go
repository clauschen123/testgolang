package main

//"example"
//	"example/test2"
//	"fmt"
//	"log"
//	"github.com/golang/protobuf/proto"

func save_file(data []byte) {

}

func test_protobuf() {
	/*
		test := &example.Test{
			Label: proto.String("hello"),
			Type:  proto.Int32(17),
			Optionalgroup: &example.Test_OptionalGroup{
				RequiredField: proto.String("good bye"),
			},
		}

		//encode
		data, err := proto.Marshal(test)
		if err != nil {
			log.Fatal("Masharl error:", err)
		}
		fmt.Println(string(data))

		//decode
		newTest := &example.Test{}
		err = proto.Unmarshal(data, newTest)
		if err != nil {
			log.Fatal("Unmarshal error:", err)
		}

		if test.GetLabel() != newTest.GetLabel() {
			log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
		} else {
			fmt.Println("OK")
		}
	*/
}
