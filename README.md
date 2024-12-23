# gophers 训练营

## 第一期(12/16-12/22)

分别使用读写锁和channel实现一个线程安全的map.
写得比较好
<a href="https://github.com/DnullP/synmap" target="_blank">
DnullP</a>
<a href="https://github.com/EinoPlasma/gopher" target="_blank">
EinoPlasma</a>

## 第二期(12/23-12/29)
设计一个短链接生成器项目的后端
需求： 设计两个api接口

1. 接收长链接，返回短链接:

POST /api/v1/url 
    接收json: {"original_url": "..."}
    返回json: {"short_url": "..."}

2. 将短链接穷定向到长链接:

GET /:code 
返回重定向到长链接。