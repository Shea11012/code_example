### DNS Message format
for more information: [https://tools.ietf.org/html/rfc1035#section-4.1.1](https://tools.ietf.org/html/rfc1035#section-4.1.1)

![DNS Message](https://mxy-imgs.oss-cn-hangzhou.aliyuncs.com/imgs/202109181629249.png)

### DNS Header
![](https://mxy-imgs.oss-cn-hangzhou.aliyuncs.com/imgs/202109181630552.png)
- ID: 由客户端生成的16个bit位的标识符，该标识符会被复制到对应的回复中，由客户端去匹配未完成的回复
- QR: 1bit表示这个消息是一个 query(0) 或者 response(1)
- AA：Authoritative Answer，可以在 question section 中指出权威机构名称
- TC：消息是否因为过大而被截断（truncate）
- RD：是否需要递归（可以在query时被设置）
- RA：表示可用于DNS服务器响应的递归
- Z：保留
- RCODE：响应码，4bit
  - 0：No error condition
  - 1: format error
  - 2：server failure
  - 3：name error
  - 4：not implemented
  - 5: refused
- QDCOUNT：16bit,表示请求问题数
- ANCOUNT：16bit,表示回答响应数
- NSCOUNT：16bit,表示 Authority 记录中 name server resource 记录的个数
- ARCOUNT：16bit,表示额外记录中的数量

### DNS Question
![](https://mxy-imgs.oss-cn-hangzhou.aliyuncs.com/imgs/202109181731519.png)
- QNAME：包含希望解析的域名，每一个域名被看作是一个连续的 label，每一个label用8bit表示
> domain: example.com
> 
> 将 domain url encode 为 “69 88 65 77 80 76 69” 和 “99 111 109”
> 
> example 为 “7 69 88 65 77 80 76 69”
> 
> 然后将每一个值转为一个 8bit 值存入
> 
> 如 (7) (69) (88)
> 
> 00000111 01000101 01011000

- QTYPE：2bit, 指定查询类型
- QCLASS：指定查询类别

### DNS Answer,Additional And Authority records
![](https://mxy-imgs.oss-cn-hangzhou.aliyuncs.com/imgs/202109181804962.png)

- NAME: 此资源记录属于哪个域名
- TYPE：指定查询类型。standard or inverse query
- TTL：32bit,表示资源记录应被缓存时间，0表示不缓存
- RDLENGTH：response data length。
- RDATA：可变的资源长度，由 TYPE 和 CLASS 决定