# BTsearch 一个用go语言实现的磁力搜索网站
<p align="center">
  <a href="https://pkg.go.dev/github.com/Bmixo/btSearch"><img src="https://pkg.go.dev/badge/github.com/Bmixo/btSearch" alt="PkgGoDev"></a>
  <a href="https://goreportcard.com/report/github.com/Bmixo/btSearch"><img src="https://goreportcard.com/badge/github.com/Bmixo/btSearch" alt="Go Report Card"></a>
  <a href="https://github.com/Bmixo/btSearch"><img src="https://img.shields.io/badge/License-Apache-green.svg" alt="License: Apache"></a>
</p>
<br>

<div align="center">
<strong>
<samp>

[English](README.md) · [简体中文](README_CN.md)

</samp>
</strong>
</div>

## 1.页面展示

![image](https://raw.githubusercontent.com/Bmixo/btSearch/master/example/index_old.png)
![image](https://raw.githubusercontent.com/Bmixo/btSearch/master/example/index.PNG)
![image](https://raw.githubusercontent.com/Bmixo/btSearch/master/example/detail.PNG)

## 2.程序架构

名称   |  用途
|------------:|-----------
server |  收集torrent数据
worker | 收集Hash信息
web    |  数据展示
Tool   | 工具

![image](https://raw.githubusercontent.com/Bmixo/btSearch/master/example/framework.png)

## 注意:
1.项目使用了reuseport系统特性来监听端口，请保持新的linux 内核版本（4.9以上）

2.docker一键安装仅供开发测试和展示程序功能使用，请勿应用于生产环境。

3.若要将本程序应用于生产环境，作者假设使用者都会使用Golang，请Fork主分支的任意一个版本代码开发，不要合并后续主分支代码，主分支代码不保证不进行不兼容的改动。

## 安装（docker一键安装）：
```
git clone https://github.com/Bmixo/btSearch.git && cd btSearch/docker && docker-compose up 
```
等待片刻系统初始化后，开始采集数据。程序网页界面请访问 http://127.0.0.1:8080

## 依赖环境 
* 确保运行程序的服务器能被互联网外部主机访问到
* 推荐配置 8核心 3.8Ghz 32G内存
```
apt update
apt install curl git docker.io -y
```



## 设置Elasticsearch默认分词器为ik分词器 (可选)

```
curl --user elastic:changeme -XPUT http://localhost:9200/bavbt -H 'Content-Type: application/json'
curl --user elastic:changeme -XPOST 'localhost:9200/bavbt/_close'
curl --user elastic:changeme -XPUT localhost:9200/bavbt/_settings?pretty -d '{
"index":{
"analysis" : {
            "analyzer" : {
                "default" : {
                    "type" : "ik_max_word"
                }
            },
			"search_analyzer" : {
                "default" : {
                    "type" : "ik_max_word"
                }
            }
        }
    }
}'
curl --user elastic:changeme -XPOST 'localhost:9200/bavbt/_open'
```

## TODO

- [ ] 后台数据展示
- [x] 打包docker镜像
- [ ] 提供k8s高可用部署方案（mongodb sharding + 无状态均衡负载master + etcd）
- [ ] gin迁移iris
- [x] 换用mongodb官方驱动
- [ ] 支持 postgres 和 mysql
