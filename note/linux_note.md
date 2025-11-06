# Linux系统学习


## 实用指令
| 序号 | 命令 | 功能说明 |
|-----|-----|------|
| 1 | `systemctl get-default` | 查询当前系统默认运行级别（systemd 目标） |
| 2 | `systemctl set-default <目标级别>.target` | 设置系统默认运行级别（例如：`multi-user.target` 或 `graphical.target`） |
| 3 | `init <数字>` | 临时切换运行级别（传统 SysV 方式，例如 `init 3` 进入多用户模式） |
| 4 | `pwd` | 显示当前工作目录的绝对路径 |
| 5 | `ls` | 列出当前目录下的文件和子目录 |
| 6 | `cd` | 切换当前工作目录（`cd /path` 进入指定路径；`cd ~` 返回主目录） |



## 常识
1. Linux下，以.开头的文件是隐藏文件。



## 目录结构详解

| 目录           | 全称                    | 主要作用                                                  |
| ---| --- | --- |
| `/`          | root（根目录）             | 整个文件系统的起点，所有目录都从这里分支。                                 |
| `/bin`       | binaries              | 存放最基本的用户命令（如 `ls`, `cp`, `mv`, `cat` 等），系统启动后立即可用。    |
| `/boot`      | boot loader files     | 存放系统引导所需文件（如 `vmlinuz`、`initrd.img`、`grub` 等）。        |
| `/dev`       | devices               | 存放设备文件（如硬盘、终端、USB 设备），以文件形式表示硬件。                      |
| `/etc`       | etcetera              | 系统配置文件目录，包含启动脚本、服务配置、用户信息等。                           |
| `/home`      | home directories      | 普通用户的家目录，每个用户一个（如 `/home/alice`）。                     |
| `/lib`       | libraries             | 存放系统运行所需的共享库（类似 Windows 的 DLL）。                       |
| `/lib64`     | 64-bit libraries      | 为 64 位系统提供的共享库文件。                                     |
| `/media`     | media                 | 可移动设备挂载点（如 U 盘、光盘）。                                   |
| `/mnt`       | mount                 | 临时挂载文件系统的位置（如临时挂载另一块硬盘）。                              |
| `/opt`       | optional packages     | 额外安装的软件（第三方应用）通常放在这里。                                 |
| `/proc`      | process               | 虚拟文件系统，提供内核和进程信息（如 `/proc/cpuinfo`, `/proc/meminfo`）。 |
| `/root`      | root user home        | 超级用户 root 的家目录。                                       |
| `/run`       | runtime data          | 存放系统运行时产生的临时数据（如 PID 文件、Socket）。                      |
| `/sbin`      | system binaries       | 系统管理命令（如 `reboot`, `ifconfig`, `fdisk`），通常只有 root 能用。 |
| `/srv`       | service data          | 存放某些系统服务的数据（如 Web 或 FTP 服务的数据）。                       |
| `/sys`       | system                | 内核虚拟文件系统，提供设备和驱动信息（配合 `/proc` 使用）。                    |
| `/tmp`       | temporary             | 临时文件目录，系统重启后会清空。                                      |
| `/usr`       | user system resources | 存放用户应用程序及文件，如命令、库、文档等。                                |
| `/usr/bin`   | user binaries         | 普通用户可用的命令（大部分应用程序都在这里）。                               |
| `/usr/sbin`  | system binaries       | 系统管理命令，通常由管理员使用。                                      |
| `/usr/lib`   | libraries             | 对应 `/usr/bin` 和 `/usr/sbin` 的库文件。                     |
| `/usr/local` | local software        | 用户自行编译或安装的软件通常放在这里，不影响系统默认软件。                         |
| `/var`       | variable              | 存放可变数据（如日志、缓存、邮件、数据库）。                                |
| `/var/log`   | logs                  | 各种系统日志文件（如 `/var/log/syslog`, `/var/log/auth.log`）。   |
| `/var/tmp`   | temporary             | 长期保存的临时文件，系统不会在重启时清除。                                 |

### 常见虚拟目录

| 目录      | 说明                                      |
| --- | --- |
| `/proc` | 显示内核和进程的实时信息（如 `/proc/1` 是 PID=1 的进程）   |
| `/sys`  | 展示系统硬件设备和内核接口（可动态修改内核参数）                |
| `/dev`  | 设备文件目录，例如 `/dev/sda`（硬盘）、`/dev/tty`（终端） |



## Vim编辑器

### Vim常用快捷键
| 模式       | 快捷键      | 说明                    |
| --- | --- | --- |
| **模式切换** | `i`      | 进入插入模式（Insert）        |
|          | `Esc`    | 返回普通模式（Normal）        |
|          | `v`      | 进入可视模式（Visual）        |
|          | `V`      | 进入行可视模式（Visual Line）  |
|          | `Ctrl+v` | 进入块可视模式（Visual Block） |
|          | `:`      | 进入命令模式（Command Line）  |
|          | `R`      | 进入替换模式（Replace）       |



### 删除与剪切

| 快捷键   | 功能                     |
| --- | ---|
| `x`   | 删除光标处字符                |
| `dd`  | 删除（剪切）当前行              |
| `ndd` | 删除 n 行，例如 `3dd` 删除 3 行 |
| `d$`  | 删除至行尾                  |
| `d0`  | 删除至行首                  |
| `dw`  | 删除至下一个单词开头             |
| `D`   | 删除至行尾（等价于 `d$`）        |

### 复制与粘贴

| 快捷键   | 功能              |
| --- | --- |
| `yy`  | 复制当前行（yank）     |
| `nyy` | 复制 n 行，例如 `3yy` |
| `p`   | 在光标后粘贴          |
| `P`   | 在光标前粘贴          |
| `yw`  | 复制一个单词          |
| `y$`  | 复制到行尾           |


### 文件操作

| 快捷键           | 功能               |
| --- | --- |
| `:w`          | 保存文件（write）      |
| `:w filename` | 另存为新文件           |
| `:q`          | 退出（quit）         |
| `:q!`         | 不保存退出            |
| `:wq` 或 `ZZ`  | 保存并退出            |
| `:e filename` | 打开文件             |
| `:r filename` | 读取文件内容插入到当前光标处   |
| `:x`          | 保存并退出（等价于 `:wq`） |



## 关机与重启命令

- 在执行关机与重启命令之前，最好执行sync命令来将内存中的数据存储到磁盘，避免数据丢失。


| 命令                              | 功能                          | 示例                    |
| --- | --- | --- |
| `shutdown -h now`               | **立即关机**（halt）              | 立刻关闭系统                |
| `shutdown -r now`               | **立即重启**（reboot）            | 立刻重启系统                |
| `shutdown -h +5`                | **5 分钟后关机**                 | 系统将在 5 分钟后自动关机        |
| `shutdown -r +10 "系统将于10分钟后重启"` | **带提示的定时重启**                | 给登录用户显示提示信息           |
| `shutdown -c`                   | **取消已安排的关机/重启**             | 必须在倒计时期间执行            |
| `reboot`                        | **立即重启系统**                  | 等价于 `shutdown -r now` |
| `halt`                          | **停止系统**，CPU 停止运行           | 不断电，只停止系统进程           |
| `poweroff`                      | **关闭电源**                    | 彻底关闭计算机（比 halt 更彻底）   |
| `init 0`                        | **进入运行级别 0（关机）**            | 传统 SysV 风格            |
| `init 6`                        | **进入运行级别 6（重启）**            | 传统 SysV 风格            |
| `systemctl poweroff`            | 使用 systemd **关闭系统**         | 推荐在较新的发行版中使用          |
| `systemctl reboot`              | 使用 systemd **重启系统**         | 现代 Linux 推荐方式         |
| `systemctl halt`                | 使用 systemd **停止所有进程但不切断电源** |                       |


## 用户登陆与注销

1. 登陆时，尽量少用root用户，避免操作失误，登陆后可使用`su - 用户名`来切换成管理员身份。

2. 在shell下输入logout即可注销用户。(在图形界面下无效)


## 用户管理

| 命令               | 功能                            | 示例                                      |
| --- | --- | --- |
| `useradd 用户名`    | 新建用户                          | `sudo useradd tom`                      |
| `adduser 用户名`    | 新建用户（交互式创建）                   | `sudo adduser tom`                      |
| `passwd 用户名`     | 设置或修改用户密码                     | `sudo passwd tom`                       |
| `usermod`        | 修改用户属性                        | `sudo usermod -aG sudo tom`（添加到 sudo 组） |
| `userdel 用户名`    | 删除用户但保留家目录                    | `sudo userdel tom`                      |
| `userdel -r 用户名` | 删除用户及家目录                      | `sudo userdel -r tom`                   |
| `id 用户名`         | 查看用户的 UID、GID、所属组             | `id tom`                                |
| `who`            | 查看当前登录的用户                     | `who`                                   |
| `whoami`         | 显示当前用户名                       | `whoami`                                |
| `w`              | 查看当前登录用户及运行情况                 | `w`                                     |
| `su 用户名`         | 切换到其他用户                       | `su tom`                                |
| `sudo 命令`        | 临时以 root 权限执行命令               | `sudo apt update`                       |
| `visudo`         | 编辑 sudo 权限配置文件 `/etc/sudoers` | 推荐用 `visudo` 修改                         |
| `chsh`           | 更改用户登录 Shell                  | `chsh -s /bin/bash tom`                 |
| `chfn`           | 更改用户信息（全名、电话等）                | `chfn tom`                              |
| `groups 用户名`     | 查看用户所属组                       | `groups tom`                            |
| `last`           | 显示登录历史记录                      | `last`                                  |
| `finger 用户名`     | 显示用户详细信息                      | `finger tom`（需安装 finger）                |

## Linux运行级别

| 运行级别 | 含义说明 | 对应 systemd 目标(target) |
|---|---|---|
| 0 | 关机（系统停止） | `poweroff.target` |
| 1 | 单用户模式（维护模式） | `rescue.target` |
| 2 | 多用户模式（无网络服务，部分发行版未使用） | `multi-user.target`（部分发行版） |
| 3 | 多用户模式（完整网络支持，纯命令行） | `multi-user.target` |
| 4 | 未定义/保留给用户自定义 | 可自定义目标 |
| 5 | 图形化多用户模式（含显示管理器） | `graphical.target` |
| 6 | 重启（重新启动系统） | `reboot.target` |
