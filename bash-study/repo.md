1. 安装rng-tools
   ```
    pt-get install rng-tools
   ```
2. 运行rngd 生成一些随机数(生成gpg key时需要足够多的熵--类似随机数)
   ```
    rngd -r /dev/urandom
   ```
3. 运行gpg --gen-key (有默认值的一路Enter)
   ```
    gpg --gen-key
   ```
4. 为安装包生成一个子签名(后面的字母用上一步生成的)
   ```
    gpg --edit-key 10E6133F
   ```
   之后直接执行 * save *

5. 从子秘钥连接主秘钥
   ```
    gpg --export-secret-key 10E6133F > private.key
    gpg --export 10E6133F >> private.key
   ```
   保存私钥,然后删除私钥文件, 到处公钥,然后第二次连接主密钥和子密钥
   ```
    gpg --export 10E6133F > public.key
    gpg --export-secret-subkeys A72DB3EF > signing.key
   ```
   完成之后就可以删除从服务器上删除主秘钥了
   ```
    gpg --delete-secret-key 10E6133F
   ```
   再次导入子密钥
   ```
    gpg --import public.key signing.key
   ```
   检查我们不再有我们的万能钥匙在我们的服务器
   ```
    gpg --list-secret-keys
   ```
   将私有秘钥推到服务器
   ```
    gpg --keyserver keyserver.ubuntu.com --send-key 10E6133F
   ```
6. 通过Reprepro 创建一个Repository
   ```
    apt-get update
    apt-get install reprepro
   ```
   给repo创建目录和配置文件
   ```
    mkdir -p /var/repositories/
    cd /var/repositories/
    mkdir conf
    cd conf/
   ```
   创建两个配置文件
   ```
    nano options distributions
    nano options
    ask-passphrase
    nano distributions
   ```
   distributions 写入如下内容
   ```
    Codename: trusty
    Components: main
    Architectures: i386 amd64
    SignWith: A72DB3EF
   ```
   为repo添加安装包
   ```
    mkdir -p /tmp/debs
    cd /tmp/debs
   ```
   添加测试的安装包
   ```
    wget https://github.com/Silvenga/examples/raw/master/example-helloworld_1.0.0.0_amd64.deb
    wget https://github.com/Silvenga/examples/raw/master/example-helloworld_1.0.0.0_i386.deb
   ```
   执行添加文件到repo
   ```
    reprepro -b /var/repositories includedeb trusty example-helloworld_1.0.0.0_*
   ```
   可以通过list和delete命令来展示和移除repo中的package
   ```
    reprepro -b /var/repositories/ list trusty
    reprepro -b /var/repositories/ remove trusty example-helloworld
   ```
