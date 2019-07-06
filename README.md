glog
====

在[golang/glog](https://github.com/golang/glog)的基础上做了一些修改。

## 修改的地方:
重新做了一层封装

##使用示例 
```
func main() {
    var logger = glog.Logger

    logger.Init(0, "/home/share/projects/my/go/src/projects/test/log", true, 10 * 1024)
    defer logger.Release()

    for i := 0; i < 10000000; i++ {
        logger.Debugln("111111111111111111111111")
        logger.Infoln("22222222222222222222222")
        logger.Warnln("333333333333333333333333333333333")
        logger.Errorln("444444444444444444444444")
        logger.Debugln("5555555555555555555555555555")
    }
 }
 ```
