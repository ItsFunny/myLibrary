
cp -r \
> /Users/joker/go/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1 /Users/joker/go/src/myLibrary/go-library/blockchain/vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/
# 加密模块
- 因为账本数据共享,所以一些数据会加密:`只会对后面的value进行加密`
    -   用户信息
    -   用户钱包信息

# 字节设置:
- v2.1: [0,8]为基本类型 [8,42) 作为fromWalletAddress(发送钱包地址),[42,76)作为toWalletAddress(接收者钱包地址),[76,84)作为交易金额,[84,92) 作为版本号,[92:100]作为遗留空间,其中92位判断是否要加密1代表是加密数据  [100:128]作为空闲字段,[128:)作为模型数据
    -   既: 基本类型-FROM-TO-TOKEN-VERSION-LEFT-LEFT  总共128字节
- v2.2: 将基本类型修改为bitset数据格式,因此[0,8]不再是基本类型,总长度l
    -   [0,34)作为fromWalletAddress(发送钱包地址),
    -   [34:68)作为toWalletAddress(接收者钱包地址),
    -   [68:76)作为交易金额,[76:84)作为版本号,
    -   [84:92)作为遗留空间,
        -   其中84位判断是否要加密1代表是加密数据,
        -   `85位代表的是基本类型的长度,如2,代表2*8=16位`,
    -   [92:120)作为空闲空间数据
    -   [120:l-8*长度) 代表模型数据
    -   [l-8*基本类型长度:l)代表基本类型
    
# 钱包的设计
- 采用分层钱包设计
    -   每个用户都有一个主钱包地址,存在userInfo中
    -   主钱包中存在一个tree,这个tree存放了其子下的所有钱包path,path中有address
        -   但是虽然是分层钱包,但是公私钥还是使用的主钱包的公私钥
    -   path的生成仿造bip44协议,但是与之不同在于coinCode换用Type,Type类似于部门,A部门中有A1部门,则
    A1部门的TypeA1 & A部门的TypeA>=TypeA (`既基于货运算实现部门的父子部门`)
    