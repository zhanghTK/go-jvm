# GJvm
根据《自己动手写 Java 虚拟机》一书实现的一个简单 JVM，用以学习 JVM & Go

## 命令行解析
使用 flag 包下内容，flag 使用参照[命令行标志](http://www.ctolib.com/docs/sfile/gobyexample/command-line-flags/)

## 类加载
- 类路径加载策略接口：`classpath/entry.go`

  实现根据不同路径参数特征，使用不同策略类加载策略

- 不同类路径加载策略的实现：
  1. 目录形式的类路经：`classpath/entry_dir.go`
  2. 包形式的类路径：`classpath/entry_zip.go`
  3. 组合形式的类路径：`classpath/entry_composite.go`
  4. 通配符形式的类路径：`classpath/entry_wildcard.go`

- 类加载：`classpath/classpath.go`

  根据命令行参数解析启动/扩展类加载器对应路径和用户类加载对应路径;

  依次加载启动类加载器，扩展类加载器，用户类加载器对应路径下的类;

  具体加载策略根据路径形式的不同,使用 `entry` 的不同实现加载;

## 类解析

Class文件是一组以8位字节为基础单位的二进制流

Class文件格式只包含两种数据类型：

- 无符号数：以u1，u2，u4，u8代表1字节，2字节，4字节，8字节和无符号数。无符号数用来描述数字、索引引用、数量值或按照UTF-8编码构成字符串值
- 表：由多个无符号数或其他表作为数据项构成的符合数据类型，以“_info”结尾

整个Class文件本质也是一个表，其数据项如下：
|     **类型**     |       **名称**        |         **数量**          |     描述     |
| :------------: | :-----------------: | :---------------------: | :--------: |
|       u4       |        magic        |            1            | 0xCAFEBABE |
|       u2       |    minor_version    |            1            |    次版本     |
|       u2       |    major_version    |            1            |    主版本     |
|       u2       | constant_pool_count |            1            |  常量池容量计数值  |
|    cp_info     |    constant_pool    | constant_pool_count - 1 |    常量池     |
|       u2       |    access_flags     |            1            |    访问标志    |
|       u2       |     this_class      |            1            |    类索引     |
|       u2       |     super_class     |            1            |    父类索引    |
|       u2       |  interfaces_count   |            1            |   接口计数值    |
|       u2       |     interfaces      |    interfaces_count     |   接口索引集合   |
|       u2       |    fields_count     |            1            |  字段表集合计数值  |
|   field_info   |       fields        |      fields_count       |   字段表集合    |
|       u2       |    methods_count    |            1            |  方法表集合计数值  |
|  method_info   |       methods       |      methods_count      |   方法表索引    |
|       u2       |   attribute_count   |            1            |   属性表计数值   |
| attribute_info |     attributes      |    attributes_count     |    属性表     |

- 常连池：

  - Class文件中的仓库资源，存储自字面量和符号引用
  - 每项数据都是一个表，公有14种表类型结构
  - 所有类型结构第一位是一个u1类型的标示，用于标示具体的常量类型
  - 类型可以大致分为，（详细可以参见[常量池的项目类型.png](https://ooo.0o0.ooo/2017/06/30/5955ee231e3fd.png)）：
    - UTF8编码字符串
    - 数字字面量
    - 字符串字面量
    - 类和接口符号引用
    - 字段、（类/接口）方法的符号引用
    - 字段、方法的部分符号引用
    - 动态语言调用支持
  - 常量池的项目类型实现：`classfile/cp_*.go`

- 字段表

  - 用于描述接口或者类中声明的变量，其结构：

    |       类型       |        名称        |       数量        |      描述      |
    | :------------: | :--------------: | :-------------: | :----------: |
    |       u2       |   access_flags   |        1        |    字段访问标志    |
    |       u2       |    name_index    |        1        |    字段简单名称    |
    |       u2       | descriptor_index |        1        |    字段描述符     |
    |       u2       | attribute_count  |        1        |    属性表计数值    |
    | attribute_info |    attributes    | attribute_count | 属性表列表，记录额外信息 |

  - 字段表集合不会从父类继承任何字段

  - Java语言中字段无法重载，字节码中描述符不一致即可重载

- 方法表

  - 用于描述接口或者类中的方法，其结构与字段表结构完全一致
  - 方法内部的实现等其他信息会在属性表中记录
  - 如果没有重写父类方法，方法表不会出现来自父类的方法信息
  - Java方法签名：方法名，参数类型及顺序；字节码方法签名：方法名，参数类型及顺序，返回值，异常表

- 属性表

  - 用于描述某些场景专有信息

  - 属性表结构定义较为松散，满足基本定义：

    |  类型  |          名称          |        数量        |                  描述                  |
    | :--: | :------------------: | :--------------: | :----------------------------------: |
    |  u2  | attribute_name_index |        1         | 常量池中CONSTANT_Utf8_info类型常量，表示属性表具体类型 |
    |  u4  |   attribute_length   |        1         |               属性值占用位数                |
    |  u1  |         info         | attribute_length |             属性值，根据不用属性定义             |

  - 部分属性说明：

    - Code：记录方法体的代码
    - Exceptions：方法中可能抛出的受查异常
    - LineNumberTable：源码行号与字节码偏移量之间对应关系
    - LocalVariableTable：栈帧中局部变量表中的变量与源码中定义的变量之间的对应关系
    - SourceFile：Class文件的源码文件名称
    - ConstantValue：通知虚拟机自动为静态变量赋值
    - InnerClass：记录内部类与宿主类之间的关联
    - Deprecated：表示一个类/字段/方法不再推荐使用
    - Synthetic：表示字段或方法不是源码直接产生的
