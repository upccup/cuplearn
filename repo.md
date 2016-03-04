## 搭建 apt-get 和 yum 的 repositories
### 准备工作: 安装文件服务器(nginx) reprepro(ubuntu/debin) createrepo(centos)
 1. 安装需要的工具和服务 </br>
    ```
     apt-get update 
     apt-get install reprero
     apt-get install createrepo 
     apy-get install nginx
    ```
 2. 修改 nginx 的配置文件 </br>
  * 先备份nginx 的默认配置文件 </br>
    ```
    mv /etc/nginx/sites-available/default /etc/nginx/sites-available/default.bak
    ```
  * 修改 nginx 的配置文件, 写入如下内容: </br>
    ```
    server {

        ## Let your repository be the root directory
        root        /var/repositories;

        ## Always good to log
        access_log  /var/log/nginx/repo.access.log;
        error_log   /var/log/nginx/repo.error.log;

        ## Prevent access to Reprepro's files
        location ~ /(db|conf) {
            deny        all;
            return      404;
        }
    }
    ```
  * 重启 nginx  </br>
    ```
    service nginx restart
    ```
    
### 使用 reprepro 创建一个 apt-get 的repo (不对安装包进行签名认证)
1. 创建 apt-get repo 存放文件的目录 </br>
   ```
    mkdir -p /var/repositories/
    cd /var/repositories/
   ```
   * 创建配置文件目录 添加配置文件 </br> 
   ```
    mkdir conf
    cd conf/
    touch options distributions
   ```
   * 修改配置文件 向 options 中添加 如下内容 </br> 
   ```
    ask-passphrase
   ```
   * 向distributions 中添加如下内容 </br> 
   ```
    Codename: trusty
    Components: main
    Architectures: i386 amd64
   ```
2. 使用 reprepro 添加安装包管理 </br> 
   * 先将测试文件存放在临时目录 </br> 
   ```
    mkdir -p /tmp/debs
    cd /tmp/debs
    wget https://github.com/Silvenga/examples/raw/master/example-helloworld_1.0.0.0_amd64.deb
    wget https://github.com/Silvenga/examples/raw/master/example-helloworld_1.0.0.0_i386.deb
   ```
   * 添加安装包到 repo 中 </br> 
   ```
    reprepro -b /var/repositories includedeb trusty example-helloworld_1.0.0.0_*
   ```
   * 使用 list 和 delete 命令可以查看和删除 repo 中的安装包
   ```
    reprepro -b /var/repositories/ list trusty
    reprepro -b /var/repositories/ remove trusty example-helloworld
   ```
3. 使用自定的 repo 安装程序 </br> 
   * 将自定义的 repo 路径添加到 apt-get 配置文件中 ip 填写 nginx服务所在 ip </br> 
   ```
    add-apt-repository "deb http://198.199.114.168/ trusty main"
    apt-get update
   ```
   * 安装程序  </br> 
   ```
    apt-get install example-helloworld
   ```
    * 测试安装是否成功 </br> 
   ```
    执行命令  example-helloworld 
   ```
   * 如果输出 </br> 
   ```
    Hello, World
    This package was successfully installed!
   ```
   
   ##### 大功告成 !!!!!!!


### 使用createrepo 创建 yum 源 (不带认证)
1. 和 apt-get 源 一样也需要一个文件服务器 所以依旧在 /var/repositories/ 这个文件夹下进行操作
2. yum并不能自己维护arch，即并不能自己区分i386还是x86_64，所以是单独放在不同的目录下 
  * 创建 repo 源 目录
   ```
    $mkdir -p yum/centos/7/{i386,x86_64}
   ```
  * 将 rpm 文件拷贝到对应的目录下
  * 初始化 repodate 信息
   ```
    createrepo -p -d -o yum/centos/7/i386 yum/centos/5/i386 yum/centos/7/i386 yum/centos/5/i386
    createrepo -p -d -o yum/centos/7/x86_64 yum/centos/5/x86_64 yum/centos/7/x86_64 yum/centos/5/x86_64
   ```
 3. 使用自检 repo 源安装程序
  * 修改本地配置文件 在本地 * /etc/yum.repos.d * 新建一个配置文件以 * .repo * 结尾 内容大致如下
   ```
    [upccup-yum]
    name=bsdmap-yum
    baseurl=http://www.bsdmap.com/yum/centos/$releasever/$basearch/
    enabled=1
    gpgcheck=0
    gpgkey=
   ```
  * [...] 代表这个库的名字必须是唯一的不可重复
  * name=是这个库的说明, 只是一个字段说明 没有太大意义
  * baseurl= 说明采取什么方式传输，具体路径在哪里，可以使用的方式有,file://，ftp://，http://等，关于baseurl中的变量，可以查看yum.conf 的手册：man yum.conf ，在手册的最后一段有详细描述
  * enabled=1 说明启用这个更新库，0表示不启用 
  * gpgcheck=1 表示使用gpg文件来检查软件包的签名
  * gpgkey= 表示gpg文件所存放的位置，此处也可以有http方式的位置
  
  * 运行 * yum repolist * 就可以看到我们自定的repo 源了

  [apt-get 文档参考] (https://www.digitalocean.com/community/tutorials/how-to-use-reprepro-for-a-secure-package-repository-on-ubuntu-14-04)
  [yum 文档参考] (http://www.xuebuyuan.com/1385625.html)
  

