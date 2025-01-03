# gophers 训练营

## 第一期(12/16-12/22)

分别使用读写锁和channel实现一个线程安全的map.
写得比较好
<a href="https://github.com/DnullP/synmap" target="_blank">
DnullP</a>
<a href="https://github.com/EinoPlasma/gopher" target="_blank">
EinoPlasma</a>


## 第二期(1/4-1/10)

data.csv中有一列发布链接，**你需要检查该链接是否正常（可以访问），把异常的链接所在的行去除**。最终形成两个文件： 正常的文件`good.csv`, 异常的文件`bad.csv`。

用户要求： 

    - **生成一个可执行性文件给他，他需要定期运行。**
    - 执行时间越短越好，不能过长，不然不结尾款。囧
    - 执行过程要打印执行的进度。

提交格式:

    - 新建一个仓库，用于本项目。（建仓库又不要钱，哈哈 ^^）
    - **将处理后的两个文件，和可执行文件一起提交**
    - 在README.md 里写上**程序耗时**。
    - 在本仓库新建一个issue,并附上你的仓库的地址。

你将体会到：
    - go 天生并发的魅力
    - go 跨平台和直接编译成可执行文件的魅力

