### 安装
一直机器自带crontab服务, 如果没有可以使用一下命令安装并启动crontab服务
```
yum install crontabs // 安装crontab
/sbin/service crond start //启动服务
/sbin/service crond stop //关闭服务
/sbin/service crond restart //重启服务
/sbin/service crond reload //重新载入配置
service crond status // 查看crontab服务状态：
```

### 配置定时任务
1. 编写需要定时执行的脚本
```
# 每天凌晨两点执行一次
0 2 * * * sh -x /home/admin/log_bakup.sh &> /tmp/tomcat_bakup.log
```
2. 启动定时任务, 使用-u可以指定任务执行的用户
```
sudo crontab -u admin crontest
```
3. 使用 -l 选项可以查看定时任务
```
sudo crontab -u admin -l
```
4. 使用-r可以删除定时任务
```
sudo crontab -u admin -r
```

### crontab配置说明
用户所建立的crontab文件中，每一行都代表一项任务，每行的每个字段代表一项设置，它的格式共分为六个字段，前五段是时间设定段，第六段是要执行的命令段，格式如下：
minute   hour   day   month   week   command
其中：
minute： 表示分钟，可以是从0到59之间的任何整数。
hour：表示小时，可以是从0到23之间的任何整数。
day：表示日期，可以是从1到31之间的任何整数。
month：表示月份，可以是从1到12之间的任何整数。
week：表示星期几，可以是从0到7之间的任何整数，这里的0或7代表星期日。
command：要执行的命令，可以是系统命令，也可以是自己编写的脚本文件
![undefined](https://cdn.yuque.com/lark/0/2018/png/85255/1525675808665-27954d24-cd0b-40f7-8a69-adec452e0857.png) 

**在以上各个字段中，还可以使用以下特殊字符:**
星号（*）：代表所有可能的值，例如month字段如果是星号，则表示在满足其它字段的制约条件后每月都执行该命令操作。
逗号（,）：可以用逗号隔开的值指定一个列表范围，例如，“1,2,5,7,8,9”
中杠（-）：可以用整数之间的中杠表示一个整数范围，例如“2-6”表示“2,3,4,5,6”
正斜线（/）：可以用正斜线指定时间的间隔频率，例如“0-23/2”表示每两小时执行一次。同时正斜线可以和星号一起使用，例如*/10，如果用在minute字段，表示每十分钟执行一次。

### 示例
```
* * * * * command  //每1分钟执行一次command
3,15 * * * * command //每小时的第3和第15分钟执行
3,15 8-11 * * * command // 在上午8点到11点的第3和第15分钟执行
3,15 8-11 */2 * * command //每隔两天的上午8点到11点的第3和第15分钟执行
3,15 8-11 * * 1 command //每个星期一的上午8点到11点的第3和第15分钟执行
30 21 * * * /etc/init.d/smb restart //每晚的21:30重启smb 
45 4 1,10,22 * * /etc/init.d/smb restart //每月1、10、22日的4 : 45重启smb 
10 1 * * 6,0 /etc/init.d/smb restart //每周六、周日的1 : 10重启smb
0,30 18-23 * * * /etc/init.d/smb restart //每天18 : 00至23 : 00之间每隔30分钟重启smb 
0 23 * * 6 /etc/init.d/smb restart //每星期六的晚上11 : 00 pm重启smb 
* */1 * * * /etc/init.d/smb restart 每一小时重启smb 
* 23-7/1 * * * /etc/init.d/smb restart 晚上11点到早上7点之间，每隔一小时重启smb 
0 11 4 * mon-wed /etc/init.d/smb restart //每月的4号与每周一到周三的11点重启smb 
0 4 1 jan * /etc/init.d/smb restart // 一月一号的4点重启smb 
01   *   *   *   *     root run-parts /etc/cron.hourly //每小时执行/etc/cron.hourly目录内的脚本
```
说明: run-parts这个参数了，如果去掉这个参数的话，后面就可以写要运行的某个脚本名，而不是目录名了

#### 参考文档: [每天一个linux命令（50）：crontab命令](http://www.cnblogs.com/peida/archive/2013/01/08/2850483.html)
