---
title: "多分支流水线"
description: "多分支流水线"
date: 2020-02-24
draft: false
toc: true
---

## 背景
通常的 Git 分支策略中，会有固定分支和临时分支。对于固定的，我们可以给每个分支都创建一条流水线。
但对于临时分支，有些团队的做法，会把分支作为参数。

但实际上，单个代码仓库完全可以使用同一个流水线，也就是多分支流水线

## 特点
支持基于一定的策略，扫描分支、PR、标签（tag），对符合过滤条件的会动态创建流水线。
当分支等被删除后，多分支流水线可以把之前创建出来的流水线删除。

多分支流水线，本质上是存放流水线的容器（类似于文件夹）。

## 分支扫描
默认支持的分支并发扫描可能是不够的，这时候，我们可以通过增加系统配置来修改：

`-Dcom.cloudbees.hudson.plugins.folder.computed.ThrottleComputationQueueTaskDispatcher.LIMIT=12`

要注意的是，插件会选择上面的限制值和当前系统的核心数中的较小值。

## Webhook
多分支流水线，可以自动添加 webhook 到 git 服务上，并基于 webhook 推送的事件负责分支流水线的生命流程管理。下面，是部分 git 服务中
所对应的 webhook 地址：  

| git 服务 | 地址 | ContentType | 限制 |
|---|---|---|---|
| Bitbucket | `http://localhost:8080/jenkins/bitbucket-scmsource-hook/notify` | | 每个仓库最多有50个 ｜
| Gitlab | `http://localhost:8080/jenkins/gitlab-webhook/post` | `application/json` | ｜
| GitHub | `http://localhost:8080/jenkins/github-webhook/` | `application/json` | ｜

Jenkins 中的多分支流水线删除后，也会同时删除 Bitbucket 中的 webhook 记录，但是 GitHub 不会。

## 缺点
对于 PR 的动态发现，目前只有 `GitHub`、`Gitlab` 和 `Bitbucket` 支持。
