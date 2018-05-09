# transurl

主要功能是长短URL连接互转，例如：将长URL：http://fanyi.baidu.com/translate?aldtype=16047&query=&keyfrom=baidu&smartresult=dict&lang=auto2zh#auto/zh/
转成短URL：http://127.0.0.1:8088/x
支持反向转换。

使用到的开源库："github.com/gin-gonic/gin"

使用MySQL保存长短URL连接的映射关系

提供解析配置文件的单例对象

使用同步锁解决高并发下同一个长URL对应多个短URL连接的问题

因为长URL连接可以足够长，为避免使用长URL连接进行检索，使用长URL连接进行散列得到code值并以此作为标识，下次查询时使用该code判断是否已经有对应的短连接。
随着长连接的不断增大，可能会导致散列冲突的问题，为解决此问题，采用了hashcode + md5生成code的方式，该方式虽然不能达到绝对的散列无冲突，但能最大限度的避免散列冲突的问题。

如果大家有更好的方法来避免散列冲突问题，欢迎与我分享！！


