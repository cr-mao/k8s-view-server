## kubernetes集群管理系统



### 启动
```shell
go mod download
cd cmd/k8sviewserver && go run main.go --env=local
# or 等价上面
# make serve 
# make help
```

### 注意事项

此项目使用的是k8s1.23版本

`cmd/k8sviewserver/.kube/config` 替换成 自己的k8s admin.conf or config 



### 实现功能

- namespace
  - 查看ns列表
- pod 
  - 创建 
  - 更新
  - 删除
  - 详情
- node
  - 列表
  - 详情
  - 打标签
  - 设置污点
- configmap
  - 列表
  - 详情
  - 删除
  - 创建
  - 更新
    


