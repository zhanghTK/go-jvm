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
