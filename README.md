# gilix-lfs

[![Go Reference](https://pkg.go.dev/badge/github.com/lindorof/gilix-lfs.svg)](https://pkg.go.dev/github.com/lindorof/gilix-lfs)

基于 [gilix](https://github.com/lindorof/gilix) 平台的 LFS 协议实现。

### 从 C 到 Go 的数据转译规范

##### 宏

- 错误码、命令码使用 const 定义
- 命名、大小写规则保持与 LFS 一致

##### 参数规则

- 结构体
  - 保持 C 的结构体大写命名体系
  - 例如 ```LFSIDCCARDDATA```
- 数组
  - 定义新的 slice 类型
  - 例如 ```type xxx []LFSIDCCARDDATA``` 
  - 又如 ```type xxx []string```
- 简单类型
  - 定义新的结构体类型进行包装
  - 例如 ```type xxx struct { sth string }```

##### 新类型命名

- 如果是输入，则前缀 ```DefIn```
- 如果是输出，则前缀 ```DefOut```
- 如果是事件，则根据事件类型的不同，前缀：
  - ```DefExee``` 
  - ```DefUsre``` 
  - ```DefSrve``` 
  - ```DefSyse``` 
- 例如：
  - 命令 ```LFS_CMD_IDC_WRITE_RAW_DATA```
  - 命名 ```type DefInLFSIDCWRITERAWDATA []LFSIDCCARDDATA```
- 例如：
  - 命令 ```LFS_CMD_IDC_READ_RAW_DATA```
  - 命名 ```type DefOutLFSIDCREADRAWDATA []LFSIDCCARDDATA```
- 例如：
  - 事件 ```LFS_USRE_PTR_INKTHRESHOLD```
  - 命名 ```type DefUsreLFSPTRINKTHRESHOLD struct { InkThreshold uint }```

##### 变量类型

- 所有的无符号整型：```uint```
- 所有的有符号整型：```int```
- 普通字符串：```string```
- 特殊字符串：```[]byte```
- BOOL：```uint```
- BYTE/CHAR：```byte``` 或 ```uint```

##### 变量命名

- 整体延续 C 的名称
- 大写开头，驼峰式
- 但不使用下划线 ```_```
- 并去掉类型前缀
- 例如 C 的名称为 ```wData_Source``` ，则 Go 的名称为 ```DataSource```

##### 举例

例如 C 的定义如下：
``` C
typedef struct _LFS_IDC_card_data 
{
    WORD wData_Source; 
    WORD wStatus; 
    ULONG ulData_Length; 
    LPBYTE lpbData; 
    WORD wWrite_Method; 
} LFSIDCCARDDATA, *LPLFSIDCCARDDATA;
```
则转译到 Go 如下：
``` Go
type struct LFSIDCCARDDATA {
    DataSource uint
    Status uint
    DataLength uint
    Data []byte
    WriteMethod uint
}
```

### 从 Go 到 Json 的数据转译规范

##### 总述

- 所有与 Json 的数据转译，都是基于 LFSMsg 结构体
- 序列化时，gilix 自动调用 CBS 中的 LFSMsgEncode 接口
- 反序列化时，gilix 自动调用 CBS 中的 LFSMsgDecode 接口，得到 Msg

##### 序列化

- 序列化时，会自动识别 Lpara 中的参数类型
- 不管 Lpara 中的参数是 结构体、结构体指针、结构体数组，都能正确自动序列化

##### 反序列化

- 反序列化时，LFSMsg 中的数据会自动处理，但 Lpara 除外
- 对于 Lpara ，需基于反序列化后 LFSMsg 中的 type/code 等信息，通过 DecodeVal 接口进行二次反序列化
- 调用 DecodeVal 时，需传入参数 v ，而 v 标识了真实的需反序列化的数据类型

##### 如何传入 DecodeVal 的参数 v

- v 必须是指针，不管 struct 还是 []struct 
- 传入 v 时，推荐最简单的写法：```&xxx{}``` ，不需要为 v 中的数据赋值，也不需要为 v 中的数据提前分配内存
- 例如：反序列化的参数类型为 ``` type LFSIDCCARDDATA struct``` ，则传入 v 的写法为 ```Lpara.DecodeVal(&LFSIDCCARDDATA{})```
- 例如：反序列化的参数类型为 ``` type xxx []LFSIDCCARDDATA``` ，则传入 v 的写法为 ```Lpara.DecodeVal(&xxx{})```

##### DecodeVal 后得到的 Lpara.Val

- 在 DecodeVal 方法中，会通过 ```reflect.ValueOf.Elem``` 方法，将参数 v 也就是 ```&xxx{}``` 进行解引用
- 解引用后，得到非指针类型的值 ```xxx{}``` ，并赋值给 ```Lpara.Val ``` 这个 ```interface{}``` 类型的变量
- 因此，在读取 ```Lpara.Val``` 的值时，需首先进行一次类型断言，例如 ```v,ok := Lpara.Val.(xxx)``` ，或者先判断是否为 ```nil```

##### 为何 DecodeVal 要对 v 进行解引用

- go 有两种类型转换：
  1. 强制转换，例如 ```int64(i)``` ；前提：i 不是 interface{} 类型
  2. 类型断言，例如，```v.(*LFSMsg)``` ；前提：v 属于 interface{} 类型
- 当一个值是 ```nil``` 时，本身的类型信息已丢失，回到了 interface{} 类型，且 ```nil``` 进行任何类型断言都是失败的
- 针对 interface{} 类型，可以很方便的判断是否为 nil ，例如 ```if v == nil```
- 在 DecodeVal 中，Val 的类型是 interface{} ，因此只能进行类型断言
- 如果让 Val 保持指针类型，那么对于可能遇到的几种类型，处理情况如下：
  1. 对于结构体指针，例如 ```var v *LFSMsg{}``` ，访问成员时会自动进行解引用，例如 ```v.Lhs```
  2. 对于数组，例如 ```v := &[3]int{}``` ，也会自动解引用，例如 ```len(v)``` 、```v[1]```
  3. 但对于切片，例如 ```v := &[]int{}``` ，则 ```len(v)``` 、```v[1]``` 等操作都是不可以的
- 因此，DecodeVal 中自动进行解引用，以方便使用

### Go 与 Json 的数据类型对应

| Go | Json |
| ---- | ---- |
| bool | Boolean |
| 整型/浮点型 | Number |
| string(utf-8) | string(Unicode) |
| array/slice | array |
| []byte | string(Base64) |
| struct | Object |
| map[string]T | Object |
| nil | null |

### 配置文件规范

##### LFS 配置

- 格式为 ```.ini``` 
- 存储路径和配置项遵循 LFS 规范

##### GILIX 配置

- 格式为 ```.ini``` 
- 配置文件名称与 gilix 主程序名称一致
- 假设 gilix 主程序名称为 ```gilix_cbc``` ，则配置文件名称为 ```gilix_cbc.ini```

##### PHY 配置

- 格式为 ```.ini``` 
- 存储在当前路径，例如物理名为 ```xxx``` ，则存储为 ```./xxx.ini```
- 公共配置：
  - ```Class``` : 字符串，例如 ```IDCCRD```
  - ```Type``` : 字符串，例如 ```MTCR```
  - ```Polli``` : 整型，例如 ```1500```
- 其他配置自行定义

##### DRV 名称

- 物理名后加上 ```_DRV``` 表示文件名称
- 假设物理名为 ```xxx``` ，则 DRV 库名称为 ```xxx_DRV```

##### DRV 配置

- 存储在当前路径，名称与 DRV 库名称保持一致，格式为 ```.toml```
- 配置自行定义

##### 配置读写

- ini
  - 主要被 Go 使用，基于 [go-ini](github.com/go-ini/ini) 库，支持读/写
  - 对于 LFS 配置，需自行读取，但 kit 中提供了部分常用配置的读取方法
  - 对于 GILIX 配置，kit 中包装了 struct ，无需自行读取
  - 对于 PHY 配置，kit 中包装了 struct ，部分公共配置已读取在 struct 中，其它配置需自行读取
- toml
  - 主要被 C 使用，只支持读，平台提供了相关接口

---

> 其它未尽说明，请参考 ```lfs_test.go``` 