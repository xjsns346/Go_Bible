# 有关一些名词注解

## GCC
GCC（GNU Compiler Collection） 是 GNU 项目推出的一套 开源编译器套件，支持多种编程语言，包括：

C

C++

Objective-C

Fortran

Ada

Go

D 等

GCC 的核心作用是：

将源代码编译成可执行文件或目标文件（machine code），并提供优化和调试信息。

它是 类 Unix 系统（Linux、macOS）默认的主要编译器，同时也可在 Windows（通过 MinGW / Cygwin）使用。

⚙️ 二、GCC 的特点
特性	说明
🔧 多语言支持	C、C++、Fortran、Go 等多语言编译
🌍 跨平台	Linux、Windows、macOS 都可使用
🪄 优化能力	提供多种优化级别：-O1, -O2, -O3, -Os, -Ofast
🐞 调试支持	生成调试信息：-g 标志
📜 标准兼容	支持 ISO C/C++ 标准及 GNU 扩展
📚 开源	完全免费，可自由使用和修改
🧱 三、GCC 的组成

GCC 并不是单一程序，而是 一套工具链：

预处理器（CPP）

处理宏定义、头文件包含等

命令：gcc -E file.c → 输出预处理后的代码

编译器（cc1 / cc1plus）

将 C / C++ 代码编译成汇编代码

命令：gcc -S file.c → 输出汇编文件

汇编器（as）

将汇编代码转成目标文件（.o）

链接器（ld）

将多个目标文件和库链接成可执行文件

GCC 实际上把这些步骤封装成一个命令行工具 gcc 或 g++，用户只需要一条命令即可完成编译链接。

## Clion