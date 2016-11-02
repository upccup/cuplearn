## protobuf 使用简介

#### mac 安装protobuf
        brew tap homebrew/versions 
      
安装指定版本:    
```
        brew install protobuf241
        brew link --force --overwrite protobuf241
```

安装 gogoproto 扩展插件
```
        go get github.com/gogo/protobuf/proto
        go get github.com/gogo/protobuf/{binary}
        go get github.com/gogo/protobuf/gogoproto
```

#### 编写 xxx.proto 文件 生成相应的go文件
protobuf 语法参考 [Go Generated Code](https://developers.google.com/protocol-buffers/docs/reference/go-generated)

自动生成 xxx.proto.go 的命令
```
protoc --proto_path=../../../../../:.  --gogo_out=. prototest.proto
```

* --proto_path 指定 proto 文件中 import 文件的位置
* --gogo_out 指定生成文件的语言

####  关于 gogoproto 选项说明

* nullable, if false, a field is generated without a pointer (see warning below).
```
message A {
		optional string Description = 1 [(gogoproto.nullable) = false];
		optional int64 Number = 2 ;
	}
  
type A struct {
		Description string
		Number      *int64
	} 
```

* embed, if true, the field is generated as an embedded field.
```
	message B {
		optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
		repeated bytes G = 2 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uint128", (gogoproto.nullable) = false];
	}

	type B struct {
		A
		G []github_com_gogo_protobuf_test_custom.Uint128
	}
```

* customtype, It works with the Marshal and Unmarshal methods, to allow you to have your own types in your struct, but marshal to bytes. For example, custom.Uuid or custom.Fixed128
```
	message A {
		optional string Description = 1 [(gogoproto.nullable) = false];
		optional int64 Number = 2 [(gogoproto.nullable) = false];
		optional bytes Id = 3 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uuid", (gogoproto.nullable) = false];
	}
  
  type A struct {
		Description string
		Number      int64
		Id          github_com_gogo_protobuf_test_custom.Uuid
	}
```

* customname (beta), Changes the generated fieldname. This is especially useful when generated methods conflict with fieldnames.
```
 message C {
		optional int64 size = 1 [(gogoproto.customname) = "MySize"];
	}
  
  type C struct {
		MySize		*int64
	}

```

* goproto_enum_prefix, if false, generates the enum constant names without the messagetype prefix
```
enum E {
    A = 0;
    B = 2;
}
const (
	E_A E = 0
	E_B E = 2
)


enum E {
    option (gogoproto.goproto_enum_prefix) = false;

    A = 0;
    B = 2;
}
const (
	A E = 0
	B E = 2
)
```

* goproto_getters, if false, the message is generated without get methods, this is useful when you would rather want to use face
```
b code:
message test {
    // option (gogoproto.goproto_getters) = false;
    E e = 1;
}
go code:
type Test struct {
	E                *E     `protobuf:"varint,1,opt,name=e,enum=test.E" json:"e,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Test) GetE() E {
	if m != nil && m.E != nil {
		return *m.E
	}
	return A
}

or

pb code:
message test {
    option (gogoproto.goproto_getters) = false;
    E e = 1;
}

go code:
type Test struct {
	E                *E     `protobuf:"varint,1,opt,name=e,enum=test.E" json:"e,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

```

* gogoproto.face (gogoproto.face_all) The face plugin generates a function will be generated which can convert a structure which satisfies an interface (face) to the specified structure.
```
The face plugin generates a function will be generated which can convert a structure which satisfies an interface (face) to the specified structure.
This interface contains getters for each of the fields in the struct.
The specified struct is also generated with the getters.
This means that getters should be turned off so as not to conflict with face getters.
This allows it to satisfy its own face.
It is enabled by the following extensions:
  - face
  - face_all
Turn off getters by using the following extensions:
  - getters
  - getters_all
The face plugin also generates a test given it is enabled using one of the following extensions:
  - testgen
  - testgen_all
Let us look at:
  github.com/gogo/protobuf/test/example/example.proto
Btw all the output can be seen at:
  github.com/gogo/protobuf/test/example/*
The following message:
  message A {
	option (gogoproto.face) = true;
	option (gogoproto.goproto_getters) = false;
	optional string Description = 1 [(gogoproto.nullable) = false];
	optional int64 Number = 2 [(gogoproto.nullable) = false];
	optional bytes Id = 3 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uuid", (gogoproto.nullable) = false];
  }
given to the face plugin, will generate the following code:
	type AFace interface {
		Proto() github_com_gogo_protobuf_proto.Message
		GetDescription() string
		GetNumber() int64
		GetId() github_com_gogo_protobuf_test_custom.Uuid
	}
	func (this *A) Proto() github_com_gogo_protobuf_proto.Message {
		return this
	}
	func (this *A) TestProto() github_com_gogo_protobuf_proto.Message {
		return NewAFromFace(this)
	}
	func (this *A) GetDescription() string {
		return this.Description
	}
	func (this *A) GetNumber() int64 {
		return this.Number
	}
	func (this *A) GetId() github_com_gogo_protobuf_test_custom.Uuid {
		return this.Id
	}
	func NewAFromFace(that AFace) *A {
		this := &A{}
		this.Description = that.GetDescription()
		this.Number = that.GetNumber()
		this.Id = that.GetId()
		return this
	}
and the following test code:
	func TestAFace(t *testing7.T) {
		popr := math_rand7.New(math_rand7.NewSource(time7.Now().UnixNano()))
		p := NewPopulatedA(popr, true)
		msg := p.TestProto()
		if !p.Equal(msg) {
			t.Fatalf("%#v !Face Equal %#v", msg, p)
		}
	}
The struct A, representing the message, will also be generated just like always.
As you can see A satisfies its own Face, AFace.
Creating another struct which satisfies AFace is very easy.
Simply create all these methods specified in AFace.
Implementing The Proto method is done with the helper function NewAFromFace:
	func (this *MyStruct) Proto() proto.Message {
	  return NewAFromFace(this)
	}
just the like TestProto method which is used to test the NewAFromFace function.
```

* goproto_stringer, if false, the message is generated without the default string method, this is useful for rather using stringer, or allowing you to write your own string method.
```
pb code:
message test {
	option (gogoproto.goproto_stringer) = true;
	string msg = 1 [(gogoproto.nullable) = false, (gogoproto.customname) = "MyMsg"];
}
go code:
type Test struct {
	MyMsg            string `protobuf:"bytes,1,opt,name=msg" json:"msg"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}

or
pb code:
option (gogoproto.goproto_getters_all) = false;

message test {
	option (gogoproto.goproto_stringer) = false;
	string msg = 1 [(gogoproto.nullable) = false, (gogoproto.customname) = "MyMsg"];
}
go code:
type Test struct {
	MyMsg            string `protobuf:"bytes,1,opt,name=msg" json:"msg"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Test) Reset()      { *m = Test{} }
func (*Test) ProtoMessage() {}
```

* The gostring plugin generates a GoString method for each message.
```
The GoString method is called whenever you use a fmt.Printf as such:
  fmt.Printf("%#v", mymessage)
or whenever you actually call GoString()
The output produced by the GoString method can be copied from the output into code and used to set a variable.
It is totally valid Go Code and is populated exactly as the struct that was printed out.
It is enabled by the following extensions:
  - gostring
  - gostring_all
The gostring plugin also generates a test given it is enabled using one of the following extensions:
  - testgen
  - testgen_all
Let us look at:
  github.com/gogo/protobuf/test/example/example.proto
Btw all the output can be seen at:
  github.com/gogo/protobuf/test/example/*
The following message:
  option (gogoproto.gostring_all) = true;
  message A {
	optional string Description = 1 [(gogoproto.nullable) = false];
	optional int64 Number = 2 [(gogoproto.nullable) = false];
	optional bytes Id = 3 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uuid", (gogoproto.nullable) = false];
  }
given to the gostring plugin, will generate the following code:
  func (this *A) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&test.A{` + `Description:` + fmt1.Sprintf("%#v", this.Description), `Number:` + fmt1.Sprintf("%#v", this.Number), `Id:` + fmt1.Sprintf("%#v", this.Id), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
  }
and the following test code:
	func TestAGoString(t *testing6.T) {
		popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
		p := NewPopulatedA(popr, false)
		s1 := p.GoString()
		s2 := fmt2.Sprintf("%#v", p)
		if s1 != s2 {
			t.Fatalf("GoString want %v got %v", s1, s2)
		}
		_, err := go_parser.ParseExpr(s1)
		if err != nil {
			panic(err)
		}
	}
Typically fmt.Printf("%#v") will stop to print when it reaches a pointer and
not print their values, while the generated GoString method will always print all values, recursively.
```

* The populate plugin generates a NewPopulated function.
```
This function returns a newly populated structure.
It is enabled by the following extensions:
  - populate
  - populate_all
Let us look at:
  github.com/gogo/protobuf/test/example/example.proto
Btw all the output can be seen at:
  github.com/gogo/protobuf/test/example/*
The following message:
  option (gogoproto.populate_all) = true;
  message B {
	optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
	repeated bytes G = 2 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uint128", (gogoproto.nullable) = false];
  }
given to the populate plugin, will generate code the following code:
  func NewPopulatedB(r randyExample, easy bool) *B {
	this := &B{}
	v2 := NewPopulatedA(r, easy)
	this.A = *v2
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.G = make([]github_com_gogo_protobuf_test_custom.Uint128, v3)
		for i := 0; i < v3; i++ {
			v4 := github_com_gogo_protobuf_test_custom.NewPopulatedUint128(r)
			this.G[i] = *v4
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedExample(r, 3)
	}
	return this
  }
The idea that is useful for testing.
Most of the other plugins' generated test code uses it.
You will still be able to use the generated test code of other packages
if you turn off the popluate plugin and write your own custom NewPopulated function.
If the easy flag is not set the XXX_unrecognized and XXX_extensions fields are also populated.
These have caused problems with JSON marshalling and unmarshalling tests.
```

* The testgen plugin generates Test and Benchmark functions for each message.
```
Tests are enabled using the following extensions:
  - testgen
  - testgen_all
Benchmarks are enabled using the following extensions:
  - benchgen
  - benchgen_all
Let us look at:
  github.com/gogo/protobuf/test/example/example.proto
Btw all the output can be seen at:
  github.com/gogo/protobuf/test/example/*
The following message:
  option (gogoproto.testgen_all) = true;
  option (gogoproto.benchgen_all) = true;
  message A {
	optional string Description = 1 [(gogoproto.nullable) = false];
	optional int64 Number = 2 [(gogoproto.nullable) = false];
	optional bytes Id = 3 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uuid", (gogoproto.nullable) = false];
  }
given to the testgen plugin, will generate the following test code:
	func TestAProto(t *testing.T) {
		popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
		p := NewPopulatedA(popr, false)
		dAtA, err := github_com_gogo_protobuf_proto.Marshal(p)
		if err != nil {
			panic(err)
		}
		msg := &A{}
		if err := github_com_gogo_protobuf_proto.Unmarshal(dAtA, msg); err != nil {
			panic(err)
		}
		for i := range dAtA {
			dAtA[i] = byte(popr.Intn(256))
		}
		if err := p.VerboseEqual(msg); err != nil {
			t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
		}
		if !p.Equal(msg) {
			t.Fatalf("%#v !Proto %#v", msg, p)
		}
	}
	func BenchmarkAProtoMarshal(b *testing.B) {
		popr := math_rand.New(math_rand.NewSource(616))
		total := 0
		pops := make([]*A, 10000)
		for i := 0; i < 10000; i++ {
			pops[i] = NewPopulatedA(popr, false)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dAtA, err := github_com_gogo_protobuf_proto.Marshal(pops[i%10000])
			if err != nil {
				panic(err)
			}
			total += len(dAtA)
		}
		b.SetBytes(int64(total / b.N))
	}
	func BenchmarkAProtoUnmarshal(b *testing.B) {
		popr := math_rand.New(math_rand.NewSource(616))
		total := 0
		datas := make([][]byte, 10000)
		for i := 0; i < 10000; i++ {
			dAtA, err := github_com_gogo_protobuf_proto.Marshal(NewPopulatedA(popr, false))
			if err != nil {
				panic(err)
			}
			datas[i] = dAtA
		}
		msg := &A{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			total += len(datas[i%10000])
			if err := github_com_gogo_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
				panic(err)
			}
		}
		b.SetBytes(int64(total / b.N))
	}
	func TestAJSON(t *testing1.T) {
		popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
		p := NewPopulatedA(popr, true)
		jsondata, err := encoding_json.Marshal(p)
		if err != nil {
			panic(err)
		}
		msg := &A{}
		err = encoding_json.Unmarshal(jsondata, msg)
		if err != nil {
			panic(err)
		}
		if err := p.VerboseEqual(msg); err != nil {
			t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
		}
		if !p.Equal(msg) {
			t.Fatalf("%#v !Json Equal %#v", msg, p)
		}
	}
	func TestAProtoText(t *testing2.T) {
		popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
		p := NewPopulatedA(popr, true)
		dAtA := github_com_gogo_protobuf_proto1.MarshalTextString(p)
		msg := &A{}
		if err := github_com_gogo_protobuf_proto1.UnmarshalText(dAtA, msg); err != nil {
			panic(err)
		}
		if err := p.VerboseEqual(msg); err != nil {
			t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
		}
		if !p.Equal(msg) {
			t.Fatalf("%#v !Proto %#v", msg, p)
		}
	}
	func TestAProtoCompactText(t *testing2.T) {
		popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
		p := NewPopulatedA(popr, true)
		dAtA := github_com_gogo_protobuf_proto1.CompactTextString(p)
		msg := &A{}
		if err := github_com_gogo_protobuf_proto1.UnmarshalText(dAtA, msg); err != nil {
			panic(err)
		}
		if err := p.VerboseEqual(msg); err != nil {
			t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
		}
		if !p.Equal(msg) {
			t.Fatalf("%#v !Proto %#v", msg, p)
		}
	}
Other registered tests are also generated.
Tests are registered to this test plugin by calling the following function.
  func RegisterTestPlugin(newFunc NewTestPlugin)
where NewTestPlugin is:
  type NewTestPlugin func(g *generator.Generator) TestPlugin
and TestPlugin is an interface:
  type TestPlugin interface {
	Generate(imports generator.PluginImports, file *generator.FileDescriptor) (used bool)
  }
Plugins that use this interface include:
  - populate
  - gostring
  - equal
  - union
  - and more
Please look at these plugins as examples of how to create your own.
A good idea is to let each plugin generate its own tests.
```

* The marshalto plugin generates a Marshal and MarshalTo method for each message.
```
The `Marshal() ([]byte, error)` method results in the fact that the message
implements the Marshaler interface.
This allows proto.Marshal to be faster by calling the generated Marshal method rather than using reflect to Marshal the struct.
If is enabled by the following extensions:
  - marshaler
  - marshaler_all
Or the following extensions:
  - unsafe_marshaler
  - unsafe_marshaler_all
That is if you want to use the unsafe package in your generated code.
The speed up using the unsafe package is not very significant.
The generation of marshalling tests are enabled using one of the following extensions:
  - testgen
  - testgen_all
And benchmarks given it is enabled using one of the following extensions:
  - benchgen
  - benchgen_all
Let us look at:
  github.com/gogo/protobuf/test/example/example.proto
Btw all the output can be seen at:
  github.com/gogo/protobuf/test/example/*
The following message:
option (gogoproto.marshaler_all) = true;
message B {
	option (gogoproto.description) = true;
	optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
	repeated bytes G = 2 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uint128", (gogoproto.nullable) = false];
}
given to the marshalto plugin, will generate the following code:
  func (m *B) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
  }
  func (m *B) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintExample(dAtA, i, uint64(m.A.Size()))
	n2, err := m.A.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.G) > 0 {
		for _, msg := range m.G {
			dAtA[i] = 0x12
			i++
			i = encodeVarintExample(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
  }
As shown above Marshal calculates the size of the not yet marshalled message
and allocates the appropriate buffer.
This is followed by calling the MarshalTo method which requires a preallocated buffer.
The MarshalTo method allows a user to rather preallocated a reusable buffer.
The Size method is generated using the size plugin and the gogoproto.sizer, gogoproto.sizer_all extensions.
The user can also using the generated Size method to check that his reusable buffer is still big enough.
The generated tests and benchmarks will keep you safe and show that this is really a significant speed improvement.
An additional message-level option `stable_marshaler` (and the file-level
option `stable_marshaler_all`) exists which causes the generated marshalling
code to behave deterministically. Today, this only changes the serialization of
maps; they are serialized in sort order.
```

* The size plugin generates a Size or ProtoSize method for each message.
```
This is useful with the MarshalTo method generated by the marshalto plugin and the
gogoproto.marshaler and gogoproto.marshaler_all extensions.
It is enabled by the following extensions:
  - sizer
  - sizer_all
  - protosizer
  - protosizer_all
The size plugin also generates a test given it is enabled using one of the following extensions:
  - testgen
  - testgen_all
And a benchmark given it is enabled using one of the following extensions:
  - benchgen
  - benchgen_all
Let us look at:
  github.com/gogo/protobuf/test/example/example.proto
Btw all the output can be seen at:
  github.com/gogo/protobuf/test/example/*
The following message:
  option (gogoproto.sizer_all) = true;
  message B {
	option (gogoproto.description) = true;
	optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
	repeated bytes G = 2 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uint128", (gogoproto.nullable) = false];
  }
given to the size plugin, will generate the following code:
  func (m *B) Size() (n int) {
	var l int
	_ = l
	l = m.A.Size()
	n += 1 + l + sovExample(uint64(l))
	if len(m.G) > 0 {
		for _, e := range m.G {
			l = e.Size()
			n += 1 + l + sovExample(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
  }
and the following test code:
	func TestBSize(t *testing5.T) {
		popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
		p := NewPopulatedB(popr, true)
		dAtA, err := github_com_gogo_protobuf_proto2.Marshal(p)
		if err != nil {
			panic(err)
		}
		size := p.Size()
		if len(dAtA) != size {
			t.Fatalf("size %v != marshalled size %v", size, len(dAtA))
		}
	}
	func BenchmarkBSize(b *testing5.B) {
		popr := math_rand5.New(math_rand5.NewSource(616))
		total := 0
		pops := make([]*B, 1000)
		for i := 0; i < 1000; i++ {
			pops[i] = NewPopulatedB(popr, false)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			total += pops[i%1000].Size()
		}
		b.SetBytes(int64(total / b.N))
	}
The sovExample function is a size of varint function for the example.pb.go file.
```

* The equal plugin generates an Equal and a VerboseEqual method for each message.
```
These equal methods are quite obvious.
The only difference is that VerboseEqual returns a non nil error if it is not equal.
This error contains more detail on exactly which part of the message was not equal to the other message.
The idea is that this is useful for debugging.
Equal is enabled using the following extensions:
  - equal
  - equal_all
While VerboseEqual is enable dusing the following extensions:
  - verbose_equal
  - verbose_equal_all
The equal plugin also generates a test given it is enabled using one of the following extensions:
  - testgen
  - testgen_all
Let us look at:
  github.com/gogo/protobuf/test/example/example.proto
Btw all the output can be seen at:
  github.com/gogo/protobuf/test/example/*
The following message:
  option (gogoproto.equal_all) = true;
  option (gogoproto.verbose_equal_all) = true;
  message B {
	optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
	repeated bytes G = 2 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uint128", (gogoproto.nullable) = false];
  }
given to the equal plugin, will generate the following code:
	func (this *B) VerboseEqual(that interface{}) error {
		if that == nil {
			if this == nil {
				return nil
			}
			return fmt2.Errorf("that == nil && this != nil")
		}
		that1, ok := that.(*B)
		if !ok {
			return fmt2.Errorf("that is not of type *B")
		}
		if that1 == nil {
			if this == nil {
				return nil
			}
			return fmt2.Errorf("that is type *B but is nil && this != nil")
		} else if this == nil {
			return fmt2.Errorf("that is type *B but is not nil && this == nil")
		}
		if !this.A.Equal(&that1.A) {
			return fmt2.Errorf("A this(%v) Not Equal that(%v)", this.A, that1.A)
		}
		if len(this.G) != len(that1.G) {
			return fmt2.Errorf("G this(%v) Not Equal that(%v)", len(this.G), len(that1.G))
		}
		for i := range this.G {
			if !this.G[i].Equal(that1.G[i]) {
				return fmt2.Errorf("G this[%v](%v) Not Equal that[%v](%v)", i, this.G[i], i, that1.G[i])
			}
		}
		if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
			return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
		}
		return nil
	}
	func (this *B) Equal(that interface{}) bool {
		if that == nil {
			if this == nil {
				return true
			}
			return false
		}
		that1, ok := that.(*B)
		if !ok {
			return false
		}
		if that1 == nil {
			if this == nil {
				return true
			}
			return false
		} else if this == nil {
			return false
		}
		if !this.A.Equal(&that1.A) {
			return false
		}
		if len(this.G) != len(that1.G) {
			return false
		}
		for i := range this.G {
			if !this.G[i].Equal(that1.G[i]) {
				return false
			}
		}
		if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
			return false
		}
		return true
	}
and the following test code:
	func TestBVerboseEqual(t *testing8.T) {
		popr := math_rand8.New(math_rand8.NewSource(time8.Now().UnixNano()))
		p := NewPopulatedB(popr, false)
		dAtA, err := github_com_gogo_protobuf_proto2.Marshal(p)
		if err != nil {
			panic(err)
		}
		msg := &B{}
		if err := github_com_gogo_protobuf_proto2.Unmarshal(dAtA, msg); err != nil {
			panic(err)
		}
		if err := p.VerboseEqual(msg); err != nil {
			t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}

```

* The description (experimental) plugin generates a Description method for each message.
```
The Description method returns a populated google_protobuf.FileDescriptorSet struct.
This contains the description of the files used to generate this message.
It is enabled by the following extensions:
  - description
  - description_all
The description plugin also generates a test given it is enabled using one of the following extensions:
  - testgen
  - testgen_all
Let us look at:
  github.com/gogo/protobuf/test/example/example.proto
Btw all the output can be seen at:
  github.com/gogo/protobuf/test/example/*
The following message:
  message B {
	option (gogoproto.description) = true;
	optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
	repeated bytes G = 2 [(gogoproto.customtype) = "github.com/gogo/protobuf/test/custom.Uint128", (gogoproto.nullable) = false];
  }
given to the description plugin, will generate the following code:
  func (this *B) Description() (desc *google_protobuf.FileDescriptorSet) {
	return ExampleDescription()
  }
and the following test code:
  func TestDescription(t *testing9.T) {
	ExampleDescription()
  }
The hope is to use this struct in some way instead of reflect.
This package is subject to change, since a use has not been figured out yet.
```

* casttype (beta), Changes the generated fieldtype.  All generated code assumes that this type is castable to the protocol buffer field type.  It does not work for structs or enums.
* castkey (beta), Changes the generated fieldtype for a map key.  All generated code assumes that this type is castable to the protocol buffer field type.  Only supported on maps.
* castvalue (beta), Changes the generated fieldtype for a map value.  All generated code assumes that this type is castable to the protocol buffer field type.  Only supported on maps.

[gogoproto详细文档](https://github.com/gogo/protobuf/blob/master/gogoproto/doc.go) <br>
[gogoprotobuf使用](https://my.oschina.net/alexstocks/blog/387031)
