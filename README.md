# RCEmap

```
Author：
           ___     ___     __ 
          / _ \___<  /__  / /_
         / ___/ __/ / _ \/ __/
        /_/  /_/ /_/_//_/\__/ 
                      
██████╗  ██████╗███████╗███╗   ███╗ █████╗ ██████╗ 
██╔══██╗██╔════╝██╔════╝████╗ ████║██╔══██╗██╔══██╗
██████╔╝██║     █████╗  ██╔████╔██║███████║██████╔╝
██╔══██╗██║     ██╔══╝  ██║╚██╔╝██║██╔══██║██╔═══╝ 
██║  ██║╚██████╗███████╗██║ ╚═╝ ██║██║  ██║██║     
╚═╝  ╚═╝ ╚═════╝╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝      
```

## 免责声明
本工具仅作为ctf题目中节省rce手测时间的自动化工具，禁止用于非法用途

## 简介 
这是一款类似sqlmap的RCE自动化工具

> ~~既然无法接入人工智能,为什么不直接接入人工呢~~  (这句话是在突然想到能让用户自己编写脚本来跑之后想到的)(奠定了半自动的基础x)
> 24/5/19更新：将开发GUI，并将全自动更改为半自动

v1.0将具有如下功能：
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

## 现在的功能

- [x] 检测源代码
- [x] 无数字字母eval
- [x] 无数字字母system
- [x] 无数字字母或(注意!现在不太完善,会不断遍历导致卡死!!!!)
- [ ] 黑名单(根据用户传入的命令进行执行)
- [x] 黑名单(执行固定逻辑)
- [x] 黑名单(根据过滤回显出可用函数用户自行执行)
- [x] 利用环境变量拼接rce
- [x] preg_replace/e的利用

## TODO
这个部分是后加的，本来没想写ToDo的，但是随着学习的深入和更多比赛的参加发现rce类型实在是太多了，于是这里写了个ToDo


- [x] 无数字字母异或plus(限制字符种类数量,还缺自定义)
- [x] 伪协议的利用（已完成php部分）
- [ ] 少量字符rce，例如7字符

- [ ] 重构一部分代码，新增探针
- [ ] 识别注释并用换行绕过
- [ ] fuzz出可用字符并进行利用

- [ ] 允许用户自己编写脚本

- [ ] rce-libs_AK

## 更新日志

3月16日发布第一版v0.1,实现了最容易的功能：在php7.x版本下，在能够使用`(` `)` `~` `;`的情况下，在题目是eval的情况下，实现rce  

3月17日v0.2更新说明：在system的情况下实现无数字字母，新增二开后的bashfuck工具，优化代码结构，增加彩色INFO和WARNING信息，不用额外的qufan.py，并新增fuzz函数(后续会用到)  

3月21日v0.3更新说明：新增在php5.x版本下，使用`><?=. /[-[];`这些字符进行rce，新增异或方法

3月27日v0.4更新说明：更新自增和%ff方法，彻底结束无数字字母RCE部分，即将开始黑名单绕过与无源码fuzz阶段，离v1.0越来越近了喵🥰

4月17日v0.5更新说明：新增无源码下fuzzpro出可用字符~~然后加以利用~~，增加代码可读性，优化代码结构  

隔了好久好不容易更新了一次估计又得停更好久，v1.0的发布也会延后好久，虽然现在能通用的脚本只差一个pwd，但是要做到当初承诺的能akphp版rce题目还有很长一段路，而且要去挖证书了，所以很少有时间会更新了  ~~我尽量这学期发布v1.0~~
要看开发笔记的话欢迎访问我的博客(虽然现在还没什么东西,不过过几天就会把一些真是笔记的东西写上去)[仙人指路](https://w0r1d-pr1nt.github.io/2024/03/04/RCEmap/)

9月1日v0.6更新说明:不说明了自己看现在的功能一栏和todo了解吧,太多了

## 使用方法
#### v0.1使用方法：
`./rcemap.exe --help`,可以查看目前仅有的四个参数,其中-v参数无用,不是php7.x的版本的东西还没写,将会在v0.2以后更新出来,如果-c后的内容为`system(ipconfig)`之类的内容建议使用双引号包裹起来,不然有可能会报错或者达不到预期目标(注:--help中使用方法有误,**或者**后的内容应为./rcemap.exe --help)（欸对，得搭配qufan.py一起使用，不然不好使，记得下下来，放在同一目录）  
#### v0.2使用方法：
`./rcemap.exe --help`，可以查看参数，v0.1中说的7.x版本内容还没写，等我后续继续更新，这个v0.2不用搭配qufan.py了，内置了
#### v0.3使用方法：
还是没变，最主要的是建议设置v
#### v0.4使用方法：
建议把这几个参数都设置上吧，反正也不多，然后自增方法在php5.x好像不好使，然后c里面的内容最好用双引号括起来，不然有可能被powershell视为其他东西导致不执行
#### v0.5使用说明：
~~离世界级项目越来越近了~~  额外写了一个fuzzpro脚本负责进行fuzz (x),优化代码结构,便于二开,以后也许会加上例如sqlmap一样可以自己编写脚本的功能
#### v0.6:
新增一大堆东西,更新日志没法详细写了记不住了,也不能再乱七八糟随便写个使用说明了,这次得写个详细点的说明书不然可能不会用

### 说明书

> 目前的中文语言只能在win上使用,linux编译不出来也没法用我编译好的,我在linux平台编译完也是乱码  
> 具体工具的画面不够好看字体不好看什么的这些东西后期会慢慢加,现在先加功能  
> 我会GUI主功能部分从上到下开始说明  
> 声明: 作者很菜,软件开发着玩的,真有做不出来的题别骂我,发issue里,看到了会加功能上去,也请不要抱太大期待,真别以为靠一个工具就能解决所有的rce题目了,就顶级项目sqlmap呗也有梭不了的时候,主要还是要以提升自身水平自己学习为主,真别太依赖工具  

首先打开软件之后,默认是自动模式,很不完善功能很少,建议直接看手动,自动模式现在具体有哪些功能我也不知道,后续会加上来.

手动模式下面就是两个参数点传入框,第二个一般是给伪协议情况下例如file_put-content这种需要传两个参数这种情况用的,其他的情况没试过不知道,然后是需要执行的命令,这个在下面情况下是不用写的,没有用:

1. 无数字字母下面你可能会看到一个叫`限制字符种类(固定)`的东西,自定义那个没写完,自定义那里面传执行的命令是有用的,固定这个没有用,这个题型主要是给`strlen(count_chars(strtolower($_), 0x3)) > 0xd `这种东西用的,你可能见过可能没见过,无所谓,当你见到的时候自然会知道我这个选项的作用
2. 黑名单绕过中,一个`执行固定逻辑`一个`根据输入的过滤回显出可用的函数自行执行`,这两个命令部分是没有用的,等到`用户传命令执行`部分实装了之后会有用
3. 环境变量构造里面没有用
4. 伪协议里面,当题目能使用php://input的时候有用,但是条件很严苛,其余情况,读取flag文件那里我写的很详细,包含了我能想到的各种flag名称,真有那逆天比赛flag名不叫flag的我也没办法,什么ffffflllllaaaaggg这玩意真没办法,其他的一个是写入,那里我选择写在a.php里面一个马,你访问一下就能看到了,至于保底的那个"前两个都不行"那个也只是把`rot13`改成了两次编码的`%7%32ot13`,其他的和写入是一样的

然后是GET&POST选项,这个还是建议什么时候都选一下的,至于一个参数是GET一个参数是POST那种我没考虑,你可以把文件本地弄一个然后跑一遍工具输出一个payload然后手动输进去(这个我真没办法,真有办法也得以后了)

PHP版本那里大部分时候无用,像preg_replace/e那种你选上了也没用,一般来讲我应该是给无数字字母那里准备的,其他时候选不选随意,反正也就点一下的事,如果有的时候不知道是5还是7的时候你两个都试一下,万一出结果了呢

过滤部分,这里其实是必填的,不填过滤后续代码全是错的,除非题目就没有过滤,不然不管滤的什么都请输进去,直接复制题目中两个引号中间的正则就对了,我的过滤匹配方式也是正则匹配

然后就是最主要的题型选择部分:
1. 无数字字母RCE,当你点进去可能只有一个选择,那就对了,你先选,再看后面的东西,大部分都塞eval下面了,自增和ff什么的我都是直接选的一系列题型中最后一道过滤最多的当payload的,***'或'那里有bug!!!!!注意!!!!!***,或这里我写屎山代码在不断遍历,现在不建议使用,其他的没什么说的,你要是选了哪个结果不好使的话我的建议是全选一遍,毕竟总共就这么几个选项都选一下也不会浪费你很多时间,万一对了呢,system那里我只塞了一个bashfuck的方法,其他的例如什么环境变量什么的在环境变量的部分
2. 黑名单,eval和system的选择不管了,你自己选,下面有三个选项,第一个现在没做,做好了说明书会更新。第二个执行固定逻辑那个就是能ls就先ls一下,ls不了就遍历一下然后将带有flag且被过滤的话就加单引号绕过或者反斜线绕过什么的,如果说他过滤了一个很诡异的东西说那个是flag的话我还是那句话,我真没办法,这绕一下又不难你自己手动一下吧。第三个就相当于一个简易的遍历了，能遍历出来可以使用的函数和可用的字符，具体函数作用啥的你自己百度一下吧，线下赛的（你都打线下了，怎么可能会遇到这么简单的题）我没办法了。
3. 少量字符RCE，未完待续
4. 环境变量构造，如题，虽然是半固定，最后的payload差不多就是执行`/???/??t ????`这种的了,具体情况你自己试吧（传入命令使用`cat` 或者 `tac` 即可）
5. preg_replace/e,这个就如题,但是太简单的我没写进去,一般考这个也都搭配`strtolower`一起用,太简单的情况你还是百度吧
6. 伪协议,其实一开始没打算写整个伪协议而是打算写file_put_content的利用的,但是后来发现单写一个还想通杀的难度不如直接就写伪协议了,然后就写的这个

最底下是回显区,分界线可拉动,字太多了可以滚动,字体大小可以在左边列表底部调节,最右边是开始按钮和复制按钮，能将poc和回显之类的进行复制

## 备注

部分特殊题目做不了，比如ctfshowRCE最后一道题，那种的需要使用限定的解题方法就不是我这种工具里能涵盖的了，包括n1ctf中zako那道题，那种考察是思维的东西，考察grep命令的用法，我这工具做不到解出来那种题，哪怕我发布的最终版也做不到这种效果，那得是接入人工智能的才行了  
再开发几个版本，后期可能会多加一些功能
