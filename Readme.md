# 统一管理go协程

为了方便程序实现优雅重启，还有Hold住主进程，确保所有的协程都运行完成，程序才退出。

基础用法：
```golang

for i:=0;i< 10;i++{
    coroutine.Default.Run(func(a,b int){
        fmt.Println("a:", a, "b:", b)
        time.Sleep(10 * time.Second)
    })
}
coroutine.Wait() // 程序会在这儿等到10个协程都运行完成
```

也可自己定义新的group来管理

请查看`test`用例中的使用方式
