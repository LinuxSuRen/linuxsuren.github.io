---
title: "Jenkins 核心"
description: "Jenkins 核心的 API"
date: 2019-01-03T16:29:11+08:00
draft: false
keywords:
- curl
- api
toc: true
---

你可以通过如下的方式来访问 Jenkins：

## 单个任务

`curl http://localhost:8080/jenkins/job/jobName/api/json?pretty=true` 获取任务信息

`curl -X POST http://localhost:8080/jenkins/job/jobName/build` 触发构建，API 返回码是 201

``

`curl -X POST http://localhost:8080/jenkins/job/jobName/submitDescription?description=hello` 添加描述信息

`curl -X POST http://localhost:8080/jenkins/job/jobName/disable` 禁用任务

`curl -X POST http://localhost:8080/jenkins/job/jobName/enable` 启用任务

`curl -X POST http://localhost:8080/jenkins/job/jobName/doDelete` 删除任务

## 搜索任务

`curl https://localhost:8080/jenkins/search/suggest?query=test` 通过关键字搜索任务

## 构建历史

`curl http://localhost:8080/jenkins/job/jobName/1/consoleText` 获取构建历史的控制台日志

`curl http://localhost:8080/jenkins/job/jobName/1/logText/progressiveText?start=0` 获取部分日志，并返回是否还有未输出的日志。协议头中的 `X-Text-Size` 为日志内容的大小，`X-More-Data` 的值如果为 `true` 的话则标示还有未输出的日志。

`curl -X POST http://localhost:8080/jenkins/job/jobName/1/doDelete` 删除构建历史

`curl -X POST http://localhost:8080/jenkins/job/jobName/1/stop` 停止构建

下面是特定情况的构建：

`curl http://localhost:8080/jenkins/job/jobName/lastBuild/consoleText` 最近一次构建的日志

`curl http://localhost:8080/jenkins/job/jobName/lastFailedBuild/consoleText` 最近一次失败构建的日志

`curl http://localhost:8080/jenkins/job/jobName/lastUnsuccessfulBuild/consoleText` 最近一次未成功的构建日志

`curl http://localhost:8080/jenkins/job/jobName/lastCompletedBuild/consoleText` 最近一次完成的构建日志

## 插件管理

`curl -u admin:token http://localhost:8080/jenkins/pluginManager/uploadPlugin -F "name=@target/demo.hpi"` 上传插件

### 插件列表

/pluginManager/plugins

## 重启

`curl -X POST http://localhost:8080/jenkins/safeRestart` 必须使用 `POST` 请求

## 认证

这里既支持密码也支持 `Token` 的方式

`curl http://localhost:8080/jenkins/job/jobName/api/json -u user:passwd`

在 HTTP 协议头中添加认证信息：

```
curl http://localhost:8080/jenkins/job/jobName/api/json --header "Authorization: Basic `echo -n user:passwd | base64`"
```


https://129.28.182.197/jenkins/asynchPeople/api/json
//user/admin/submitDescription
//user/admin/delete
//%s/user/%s/descriptorByName/jenkins.security.ApiTokenProperty/rename
//%s/user/%s/descriptorByName/jenkins.security.ApiTokenProperty/Revoke
