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

|     **类型**     |       **名称**        |         **数量**          |     说明     |
| :------------: | :-----------------: | :---------------------: | :--------: |
|       u4       |        magic        |            1            | 0xCAFEBABE |
|       u2       |    minor_version    |            1            |    次版本     |
|       u2       |    major_version    |            1            |    主版本     |
|       u2       | constant_pool_count |            1            |  常量池容量计数值  |
|    cp_info     |    constant_pool    | constant_pool_count - 1 |    常量池     |
|       u2       |    access_flags     |            1            |            |
|       u2       |     this_class      |            1            |            |
|       u2       |     super_class     |            1            |            |
|       u2       |  interfaces_count   |            1            |            |
|       u2       |     interfaces      |    interfaces_count     |            |
|       u2       |    fields_count     |            1            |            |
|   field_info   |       fields        |      fields_count       |            |
|       u2       |    methods_count    |            1            |            |
|  method_info   |       methods       |      methods_count      |            |
|       u2       |   attribute_count   |            1            |            |
| attribute_info |     attributes      |    attributes_count     |            |

- 常连池：

  - 每项数据都是一个表，公有14种表类型结构

  - 所有类型结构第一位是一个u1类型的标示，用于标示具体的常量类型

  - 常量池的项目类型汇总如下：

    ![常量池的项目类型.png](https://ooo.0o0.ooo/2017/06/30/5955ee231e3fd.png)

  - 常量池的项目类型实现：

    - `classfile/constant_info.go`
    - `classfile/cp_utf8.go`
    - `classfile/cp_numeric.go`
    - `classfile/cp_class.go`
    - `classfile/cp_string.go`
    - `classfile/cp_member_ref.go`
    - `classfile/cp_name_and_type.go`