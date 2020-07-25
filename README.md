# WordConversion

#### 单词转换工具：基于Cobra和标准库strings,unicode实现多种模式的单词转换功能

查看帮助：

```go
/.WordConversion help word
```



该命令支持各种单词格式转换，模式如下：

1.全部单词转为大写

2.全部单词转为小写

3.下划线转大写驼峰

4.下划线转小写驼峰

5.驼峰转下划线单词

```go
/.WordConversion word -s=lee -m=1
/.WordConversion word -s=LEE -m=2
/.WordConversion word -s=mr_lee -m=1
/.WordConversion word -s=mr_lee -m=1
/.WordConversion word -s=MrLee -m=1

```

输出结果：LEE

输出结果：lee

输出结果：MrLee

输出结果：mrLee

输出结果：mr_lee

