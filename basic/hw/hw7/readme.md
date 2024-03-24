作业：优化 Web 中打印日志的部分
在 Web 的 Handler 部分，有很多 if-else 分支，基本上都是在判定 err !=nil。如下图，每一个 if 里面都要打印日志。
![img.png](img.png)
现在要求你优化这些打印日志的逻辑，避免每一处 err !=nil 的时候，都得手动打一个日志。
