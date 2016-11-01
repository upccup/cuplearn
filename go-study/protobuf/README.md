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

* casttype (beta), Changes the generated fieldtype.  All generated code assumes that this type is castable to the protocol buffer field type.  It does not work for structs or enums.
* castkey (beta), Changes the generated fieldtype for a map key.  All generated code assumes that this type is castable to the protocol buffer field type.  Only supported on maps.
* castvalue (beta), Changes the generated fieldtype for a map value.  All generated code assumes that this type is castable to the protocol buffer field type.  Only supported on maps.

[gogoproto详细文档](https://github.com/gogo/protobuf/blob/master/gogoproto/doc.go)
