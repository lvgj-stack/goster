# 以下内容有些链接可能已经过时

**折腾完先说结论：还是直接买转发更香！** :xhj016:
# 不折腾会死
最近对转发很感兴趣，感觉挺有意思，可玩性也挺高，也都试了下`Realm`、`Gost`、`ZeroTier`这些，每个使用起来都不难，用下来用`Gost`多一点，给大家分享下，说白了就是折腾（直接nnr、鸽子、私家车不香吗？）

```
goster是一个小工具，可以看一些目前简单的规则，以及一些简单的转发、隧道

这个工具的优点是：可以在入口端管理多条隧道，每一个转发都可以选择想用哪条隧道转发，和那些转发面板有点像
```

# 基础Gost使用场景
## 基础中转
1. 一般玩中转需要两到三台机器，分别为：
- 入口鸡：一般指的国内机，和本地网络链接良好
- 线路鸡：一般指的和国内鸡网络链接较好，比如CMI（广移拉万物）
- 落地鸡：一般指的解锁、国际互联比较好，就不比如了容易被D
2. 基本的逻辑图大概是
   ![xx](https://img.181000.xyz/file/4fe33a3ebc1f345bf62bb.png)

## 内网穿透
1. 内网没有公网IP没办法直连
2. 想要在公司访问家里网络
3. 大概长这样
   ![xxx](https://img.181000.xyz/file/832f3e0d04b67703241f3.png)


# 基础中转操作

## 介绍小鸡
这里那三台机器来操作下：
- 国内鸡：是一个比较垃圾的广移NAT鸡，三天两头炸，后面就叫入口鸡
- 中转：示例为一台`AWS EC2 HK`
- 落地：示例为`BestVM HK Lite`

`AWS`延迟：
![aws](https://img.181000.xyz/file/e1fecb93e50854a92d46d.png)

`BestVM`延迟：
![bestvm](https://img.181000.xyz/file/4d0bdc7e5223aae9218d2.png)

广移到`AWS`：
![ping1](https://img.181000.xyz/file/8a71ca22c090b91579a70.png)

`AWS`到`BestVM`：
![ping2](https://img.181000.xyz/file/a1cece024c9efe035b260.png)

## 基础操作
goster是一个小工具，可以看一些目前简单的规则，以及一些简单的转发、隧道

遇到某一步不会用了就`-h`

1. goster安装
```
wget http://wap-tw.181000.xyz:39970/goster && chmod +x goster && mv goster /usr/bin

# ipv6 only
wget http://[2001:b030:a42d:5dcd::c8]:39970/goster && chmod +x goster && mv goster /usr/bin
```
2. goster基础使用
   ![intro](https://img.181000.xyz/file/d817a043781c9b192f284.png)
3. 安装gost
```
# Github能连通
goster install

# Github连不通
goster install http://wap-tw.181000.xyz:39970/gost.tar.gz

# ipv6 only
goster install http://[2001:b030:a42d:5dcd::c8]:39970/gost.tar.gz
```
![install](https://img.181000.xyz/file/4b7055da54393d937dc20.png)
4. 建立隧道
   命令：
```
#   goster chain add [隧道名] [隧道协议] [隧道对端ip+端口] [可选：账号] [可选：密码] [flags]
goster chain add aws-hk relay+tls 43.198.00.00:48877 lvgj lvpassword
```
![xx3](https://img.181000.xyz/file/330084e0d12e7385389b2.png)

5. 运行后会生成两条命令，分别为对端未安装gost和安装了gost
> **⚠️注意：如果对端已经运行了一些转发，可能需要自己拼一下json**

6. 根据情况在对端执行命令

对端运行完后会看到正在监听上面的端口

![duiduan](https://img.181000.xyz/file/707e309261b814d284aa0.png)

7. 添加转发规则
```
Usage:
  goster service add [转发名] [协议(tcp,udp,tcp+udp,rtcp,rudp)] [本地端口] [目标IP+端口] [隧道名(可选)] [flags]

Examples:

新增转发时：需要填目标IP+端口
goster service add bestvm-hk tcp :40005 234.23.234.87:40098 chain-0
新增内网穿透时：只需要公网入口的端口，不需要ip；本地可以填ip
goster service add bestvm-hk tcp :40005 :40098 chain-0


Flags:
  -h, --help   help for add

Global Flags:
  -c, --configPath string   配置文件路径 (default "/etc/gost/gost.json")

requires at least 4 arg(s), only received 0
```
```
# goster service add [转发名] [协议(tcp,udp,tcp+udp,rtcp,rudp)] [本地端口] [目标IP+端口] [隧道名(可选)] [flags]
goster service add bestvm-hk tcp+udp 40004 163.53.0.0:62188 aws-hk
```
**上面的40004为广移的端口，是入口端口，`163.53.0.0:62188`是BestVM，落地鸡的ip端口，这里访问BestVM的是隧道的对端**

8. 查看状态，此时已经完成转发
```
root@VM-V7H9Wv28aoRb:~# goster ps
转发:
+---------------+------+----------+---------------------+--------+
|   备注/名称   | 协议 | 本地端口 |      目的端口       |  隧道  |
+---------------+------+----------+---------------------+--------+
| bestvm-hk-tcp | tcp  |  :40004  | 163.53.0.0:62188 | aws-hk |
| bestvm-hk-udp | udp  |  :40004  | 163.53.0.0:62188 | aws-hk |
+---------------+------+----------+---------------------+--------+

内网穿透:
+-----------+------+--------------------+--------------+------+
| 备注/名称 | 协议 | 内网穿透的本地端口 | 对端入口端口 | 隧道 |
+-----------+------+--------------------+--------------+------+

隧道:
+-----------+-----------+--------------------+------+------------+
| 备注/名称 |   协议    |      隧道地址      | 账号 |    密码    |
+-----------+-----------+--------------------+------+------------+
|  aws-hk   | relay+tls | 43.198.0.0:48877 | lvgj | lvpassword |
+-----------+-----------+--------------------+------+------------+

```
9. 同理添加更多的隧道，v.ps jp
```
#   goster chain add [隧道名] [隧道协议] [隧道对端ip+端口] [可选：账号] [可选：密码] [flags]
goster chain add v.ps-jp relay+tls 120.123.0.0:48877 lvgj lvpassword
```
添加好对端，`goster ps`查看详情
![jp](https://img.181000.xyz/file/c301f94da220a47467d28.png)
10. 在隧道建立好了之后，只需要在入口节点进行管理，其他只是agent
## 测试
1. 在看中转好不好还是需要通过`iperf3`来测试带宽上限
2. 添加转发规则，将本地端口`15201`转发到隧道对端`5201`
```
# goster service add [转发名] [协议(tcp,udp,tcp+udp,rtcp,rudp)] [本地端口] [目标IP+端口] [隧道名(可选)] [flags]
# 添加hk iperf转发
goster service add test-iperf3-hk tcp 15201 :5201 aws-hk
```
![sasd](https://img.181000.xyz/file/d437ef7da9a7188f0ca24.png)
3. `Aws`开启`iperf3 -s`
4. 广移开始`iperf3 -c 127.0.0.1 -p 15201`
   ![speed](https://img.181000.xyz/file/10f26ede55563def61c75.png)
5. 可以看到`200Mbps`基本能跑满广移的带宽，合格！
6. 测试bestvm是否成功中转，以下为`Connect`真延迟
   ![test1](https://img.181000.xyz/file/1062b4d0839e39dcabd67.jpg)

# 基础内网穿透
内网穿透主要区别是有一端没有公网IP，此时可以通过tunnel方法，具体可以参考官方文档；还可以通过在有公网的服务器上打洞，`rtcp`和`rudp`使用穿透

1. 建立隧道
```
goster chain add mouyun relay+ws '[2409:8754:1:1]:40877' lvgj lvpassword
```
2. 对端添加隧道监听
3. 内网穿透，使用rtcp
```
goster service add local-ssh rtcp 22 :2222 mouyun
```
4. 测试ssh，本地2222端口，发现能够到家里的机器
   ![ssh](https://img.181000.xyz/file/ff254afa2692025896d98.png)

# 最后
1. 上面这两种方法，是看了官网文档感觉用起来挺方便的姿势，俺也不是特别熟`gost`，欢迎大家指正 Orz

参考：
>[Gost Github](https://github.com/go-gost/gost)
>[Gost 文档](https://gost.run/)

### 折腾完之后的感想就是：还是现成转发香！
