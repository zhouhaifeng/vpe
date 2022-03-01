架构:

控制面:
vcpe由路由平面和转发平面组成, 控制器通过bgp向vrr下发配置和路由并转发至vpe, vpe通过bgp下发路由和配置到vcpe/cpe
在快速实现中, 控制器通过netconf/restful api控制vcpe与vpe建立bgp邻居关系,并下发其他配置

数据面:
vcpe与vpe/vcpe建立vxlan over ipsec隧道, 数据通过wan口/ipsec隧道/vxlan over ipsec隧道转发

功能:
注册/自动上线
配置管理
接口管理
ipsec
nat
vxlan
路由
转发
acl
dpi
安全

代码:
frr: stable/8.1
vpp: stable/2202
vpp-agent v3.3.0

编译:
1: vpp
   1. 配置vpn,以便访问某些网址
   2.
      删除crypto_native, 防止编译失败
      cd vpp 
      mv plugins/crypto_native ../
   3.
      编译:
      cd vpp
      ./extras/vagrant/build.sh
      make
      注意: 每次增删模块都需要运行build.sh
2: frr
    1、./bootstrap.sh
    2、./configure \
        --bindir=/usr/bin \
        --sbindir=/usr/lib/frr \
        --sysconfdir=/etc/frr \
        --libdir=/usr/lib/frr \
        --libexecdir=/usr/lib/frr \
        --localstatedir=/var/run/frr \
        --with-moduledir=/usr/lib/frr/modules \
        --enable-snmp=agentx \
        --enable-multipath=64 \
        --enable-user=frr \
        --enable-group=frr \
        --enable-vty-group=frrvty \
        --enable-systemd=yes \
        --disable-exampledir \
        --disable-ldpd \
        --enable-fpm \
        --enable-sysrepo \
        --with-pkg-git-version \
        --with-pkg-extra-version=-fump \
        SPHINXBUILD=/usr/bin/sphinx-build
    3、make
    4、make check
    5、sudo make install
    
    
 3: vpp-agent
    1. yum install golang
    2. 
       cd vpp-agent
       make
 

