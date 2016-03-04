## 搭建 apt-get 和 yum 的 repositories
### 准备工作: 安装文件服务器(nginx) reprepro(ubuntu/debin) createrepo(centos)
 1. 安装需要的工具和服务 </br>
    ```
    apt-get update
    apt-get install reprepro 
    apt-get install createrepo
    apt-get install nginx 
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
