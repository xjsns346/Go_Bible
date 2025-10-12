# 命令行使用技巧

1.标准输入在不同系统下的退出方式：
|系统|	退出标准输入的快捷键	|说明|
|:---:|:---:|:---:|
|Linux / macOS / WSL / Git Bash	|Ctrl + D	|表示 EOF（End Of File）结束输入|
|Windows (cmd / PowerShell)	|Ctrl + Z 然后按 Enter|	同样表示 EOF|


2.按下 Ctrl + C
向程序发送一个 中断信号 (SIGINT)
Go 程序会立即终止执行
控制权回到终端