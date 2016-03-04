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
1. 创建 apt-get repo 存放文件的目录
   ```
    mkdir -p /var/repositories/
    cd /var/repositories/
   ```
   创建配置文件目录 添加配置文件
   ```
    mkdir conf
    cd conf/
    touch options distributions
   ```
   修改配置文件
   向 options 中添加 如下内容
   ```
    ask-passphrase
   ```
   向distributions 中添加如下内容
   ```
    Codename: trusty
    Components: main
    Architectures: i386 amd64
   ```
2. 使用 reprepro 添加安装包管理 
   先将测试文件存放在临时目录
   ```
    mkdir -p /tmp/debs
    cd /tmp/debs
    wget https://github.com/Silvenga/examples/raw/master/example-helloworld_1.0.0.0_amd64.deb
    wget https://github.com/Silvenga/examples/raw/master/example-helloworld_1.0.0.0_i386.deb
   ```
   添加安装包到 repo 中 
   ```
    reprepro -b /var/repositories includedeb trusty example-helloworld_1.0.0.0_*
   ```
   使用 list 和 delete 命令可以查看和删除 repo 中的安装包
   ```
    reprepro -b /var/repositories/ list trusty
    reprepro -b /var/repositories/ remove trusty example-helloworld
   ```
3. 使用自定的 repo 安装程序
   将自定义的 repo 路径添加到 apt-get 配置文件中 ip 填写 nginx服务所在 ip
   ```
    add-apt-repository "deb http://198.199.114.168/ trusty main"
    apt-get update
   ```
   安装程序 
   ```
    apt-get install example-helloworld
   ```
   测试安装是否成功
   ```
    执行命令  example-helloworld 
   ```
   如果输出
   ```
    Hello, World
    This package was successfully installed!
   ```
   大功告成 !!!!!!!
