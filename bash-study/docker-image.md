#### docker-image 命名规则

registry.domain.com/mycom/base:latest，这是一个完整的image名称
* registry.domain.com： image所在服务器地，如果是官方的hub部分忽略, 也可以直接写IP 如: 127.0.0.1:5000 但是为了防止服务器变更带来的麻烦建议使用域名
* mycom：namespace，被称为命名空间，或者说成是你镜像的一个分类. 这个类似C++中的命名空间, 建议在第一次PUSH 镜像是就进行规划, 一面造成管理混乱.
* base：这个是镜像的具体名字. 英文命名空间是用 / 标识的所以 镜像名称中不应该包含 /
* latest：这是此image的版本号，当然也可能是其它的，如1.1之类的. latest 是docker build 是默认的镜像tag, 所以在build时我们强烈建议制定使用唯一标识符来作为镜像的tag, 如果是经常变更的tag 建议使用时间+唯一标识(如jenkins build num) 来标识(如20160623150001 -- 2016年6月23日15点00分第一次build)
