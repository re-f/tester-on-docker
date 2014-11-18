tester-on-docer
===============

轻量级golang docker单元测试框架

## 说明
在docker下执行golang单元测试


## 示例

```Go
func TestDemo(t *testing.T) {
	docker.RunTestCase(t, func(t *testing.T) {
		fmt.Println("Your unit test case")
	})
}
```



### 配置

	[ssh] 		# 可省略，不可为空，连接boot2docker,linux和mac下可省略
	user	=
	passwd	=
	ip		=
	port	=

	[global]	# 是否开启debug功能
	debug	=  #true/false

	[image] 	# 测试使用到的image,os和arch填写参照golang交叉编译的GOOS和GOARCH变量
	os		= 
	arch	= 
	name	= #格式 REPOSITORY:TAG

	[path] 		# 宿主机和boot2docker文件夹映射路径
	host	= # 测试执行路径和配置文件必须在该路径下
	docker	= 

## 要求
- docker
- go 1.3+
- 宿主机可交叉编译

## 限制

暂时不支持benchmark
