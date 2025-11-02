1. 有关文件路径处理的常用函数:
| 函数                          | 作用                   |
| :---| :---|
| `filepath.Dir(path)`        | 获取路径的父目录             |
| `filepath.Base(path)`       | 获取路径的最后一部分（文件名或最后目录） |
| `filepath.Ext(path)`        | 获取文件扩展名              |
| `filepath.Join(a, b, c...)` | 拼接路径，自动加分隔符          |
