---
title: "Bash"
description: "Bash"
date: 2020-02-22
draft: false
toc: true
---

## 字符串

替换 `echo "ab/c/df" | sed 's/\//-/g'`

## 参数

|变量|说明|示例|
|---|---|---|
|`$#`|参数个数||
|`$0`|当前脚本文件名|
|`$1`|第一个参数，后面的参数则依次类推，例如：第二个参数为`$2`||
|`$_`|上一条名称的最后一个参数|mkdir temp | cd $_|

## 命令自动补全

我们可以使用 [brew](/blog/os/mac/brew) 来安装 bash 的自动补全：`brew install bash-completion`

## 目录

我们可以把特定命令的自动补全脚本保存到这个目录中 `/usr/local/etc/bash_completion.d`

## 循环

```
for i in {1..5}
do
  echo $i
done
```
