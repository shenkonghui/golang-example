go mod vendor
docker build -t registry.cn-hangzhou.aliyuncs.com/shenkonghui/app .
docker push registry.cn-hangzhou.aliyuncs.com/shenkonghui/app