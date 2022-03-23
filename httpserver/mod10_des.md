1.验证注入对不对
status-configuration
看下里面有没有job-name:kubenetes-endpoints
source_labels。。。。。。。等
--实现✅
2.服务发现
status-Service Discovery
Kubernetes-endpoints,show labels,里面有很多规则，说明它采集到了
3.status-targets
里面有两个0.54，和0.55的metrics，文档里有，这个ip地址是：kubectl get ep -n cloudnative
下面的httpsvc两个endpoints
4.看下metrics有没有正确抓到
graph
histogram_quantile(0.50, sum(rate(cloudnative_execution_latency_seconds_bucket[5m])) by (le))
然后看graph有没有图出来，不断访问httpserver那个服务