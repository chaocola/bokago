# bokago 博卡软件API接口文档


    Init(custID string, compID string, userName string, passWord string, source string, sign string)
~~初始化后才能调用其他接口~~
    
    内部自动维护token，每次调用接口都会自动检查token是否过期，如果过期会自动重新获取token

__只维护Token，不提供其他接口__