# RCEmap

```
Author：
           ___     ___     __ 
          / _ \___<  /__  / /_
         / ___/ __/ / _ \/ __/
        /_/  /_/ /_/_//_/\__/ 
                      
.----.  .---. .----..-.   .-.  .--.  .----.  
| {}  }/  ___}| {_  |  `.'  | / {} \ | {}  }    
| .-. \\     }| {__ | |\ /| |/  /\  \| .--'     
`-' `-' `---' `----'`-' ` `-'`-'  `-'`-'         
```



## 简介 
这是一款类似sqlmap的RCE自动化工具

具有如下功能：
1. 自动绕过php语言黑名单过滤
2. 无数字字母RCE
3. 部分特殊字符被限制绕过
4. 限定字符数量RCE

目标：  
1. 支持多语言RCE题目，如php, java，的过滤绕过与自动执行指定命令
2. 各框架（如thinkPHP）RCE漏洞自动打poc

预计将发布四个大版本，分别是
第一阶段： 无GUI的php版RCEmap
第二阶段： 有GUI版本
第三阶段： 带Java的版本
最终阶段： 带各框架漏洞的自动化RCE工具

## 更新日志
3月16日发布第一版v0.1,实现了最容易的功能：在php7.x版本下，在能够使用`(` `)` `~` `;`的情况下，在题目是eval的情况下，实现rce

## 使用方法
#### v0.1使用方法：
`./rcemap.exe --help`,可以查看目前仅有的四个参数,其中-v参数无用,不是php7.x的版本的东西还没写,将会在v0.2以后更新出来,如果-c后的内容为`system(ipconfig)`之类的内容建议使用双引号包裹起来,不然有可能会报错或者达不到预期目标(注:--help中使用方法有误,**或者**后的内容应为./rcemap.exe --help)  



### 备注
部分特殊题目做不了，比如ctfshowRCE最后一道题，那种的需要使用限定的解题方法就不是我这种工具里能涵盖的了，包括n1ctf中zako那道题，那种考察是思维的东西，考察grep命令的用法，我这工具做不到解出来那种题，哪怕我发布的最终版也做不到这种效果，那得是接入人工智能的才行了