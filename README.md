# auto create aria2 config file for linux
This tool will download the best trackers list, and auto update the config file of aria2 "~/.aria2/aria2.conf".

The default url is "https://ngosang.github.io/trackerslist/trackers_best.txt", you can change it by argument.

这个小工具用于下载最佳tracker列表，并且自动更新aria2配置文件"~/.aria2/aria2.conf"。

默认列表文件的下载地址是："https://ngosang.github.io/trackerslist/trackers_best.txt"，用户可以用参数改变下载地址。

#### compile from source(从源代码编译):
    go get -v github.com/rocket049/aria2-trackers

notice: You must install golang compiler first. 你必须先安装go语言编译器。

#### download binary(下载编译版本)：
[https://github.com/rocket049/aria2-trackers/releases/download/v0.1.0/aria2-trackers.gz](https://github.com/rocket049/aria2-trackers/releases/download/v0.1.0/aria2-trackers.gz)

#### Usage（用法）:
```
    //unzip
    gunzip aria2-trackers.gz

    //update default best list，下载默认最佳列表。
    ./aria2-trackers

    //update special list，下载指定列表
   ./aria2-trackers [url of list]
```
