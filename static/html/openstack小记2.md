##OpenStack 学习小记 (2)

### 安装devstack

安装devstack会遇到各种坑，主要是pip镜像的问题

参考这两个：

https://kiwik.github.io/openstack/2013/12/21/DevStack-install-in-China/

https://askubuntu.com/questions/423211/how-to-change-the-mirror-of-setuptools

###重要！！
应该是给root用户配置pip镜像！！！！
/root/.pip/pip.conf
耽误了一下午的时间，md

###网络配置
现有的网络配置似乎没有完全配好