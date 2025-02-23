# WB-CLI



一个快速生成Ai-Thinker-WB2工程的cli工具



### 依赖

---

> 需要安装WB2的SDK到本机



### 安装

---

1. 构建

   - 下载

   ```shell
    git clone https://github.com/WildboarG/WB2-CLI.git
    cd WB2-CLI
   ```

   - 编译

   ```
   go build .
   ```

2.  或安装发行版

   - 下载

   

### 使用


- 查看帮助：

```SHELL
wb2-cli  -h
```

```sh
NAME:
   wb2-cli - wb2-cli [commands]

USAGE:
   wb2-cli [global options] command [command options]

VERSION:
   v0.0.1

COMMANDS:
   init     Initialize a  WB2 SDK PATH
   create   Create a new project
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```



- 初始化：

​	根据提示输入安装sdk的路径

​	建议将路径保存在系统环境变量下（2，3，4），这样就不用每次创建工程再次初始化

```shell
wb2-cli init
```



- 创建一个新工程:

​		根据提示输出工程名字即可

```shell
wb2-cli create
```



- 指定创建工程使用的sdk

```shell
wb2-cli create --sdk /path/to/yoursdk
```

