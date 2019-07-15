

## 个人依赖库

* 有分布式分页的小轮子

* ftp上传的对象池(volatile那里有问题)

* 发送邮件的插件

* MQ的集成


* Twitter的id生成器

* 一些util


更新日志
---
-   2019-01-28 更:
    -   使用起来主要是太繁琐了,自个儿还行,但不是和别人,想办法配置简洁吧,Builder模式来解决

-   2018-10-27 更:
	-   内部的SQLExtention是自定的扩展包
	-   使用方式:
	    -   按照基本的MVC模式,分位DAO,service,
	    -   自定义的dao类接口只需要继承ISQLExtentionConfigBaseCRUDDao 即可,同时这个dao类需要使用上`@Order`注解,如:<br>
 interace UserDao extends IUserBaseDao(这是一个空接口,继承了ISQLExtentionConfigBaseCRUDDao)  配置@Order(0),因为我用的是mybatis,所以还需要加上@mapper ,又因为是Spring容器,所以要使用@Component(value="sss") ,分了几个库就需要几个dao,dao之间不同的地方只有@Component和@Order的值不同,这样底层的dao类我们就已经结束了
        -   创建一个proxy的dao类:如UserSQLExtentionProxyDaoImpl 只需要继承AbstractSQLExtentionProxyBaseCRUDDao即可,同时记得使之成为bean并且记下它的bean名称(配置要用),其实这个proxy 类好像没有存在的必要,当初创建它就是为了在servic中只需要用这个变量即可,统一的话又可以抽出来了,`再改!!并且重点突出一点,本人在造这个轮子的时候遇到的所有问题都跟这个proxy 取舍有关!!!!!!!!!`
        -   service层:只需要继承ISQLExtentionBaseCRUDService即可,它继承了PageBaseService 是又分页的,然后service的impl类依据情况是否分库分表了,若分了则继承AbstractMultipartDBPageService 同时impl 刚才的service接口,否则只需要继承ABstractPageService即可,内部的代码都封装了(当然未优化,都是一坨坨的),只需要实现简单的自定的可更改的接口即可:如获取minId,maxId的方法,排序的方法等,至于如果想通过关联表,或无法跳表的形式的话预期采用策略来做
-   2019-07-10
    -   添加了qr的util 在 go/utils/qr.go 中