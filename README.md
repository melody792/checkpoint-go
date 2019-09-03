## init beego:
- bee new checkpoint-go
- bee run
- integrate mysql: go get ...
- go build ./
- add backend api
- create database go-userinfo default character set = utf8;
- CREATE TABLE IF NOT EXISTS `users`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户名',
	`password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
	`salt` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '盐',
	`nickname` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '昵称',
	`avatar` BLOB NOT NULL COMMENT '头像',
	`company_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '公司名称',
	`company_address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '公司地址',
	`company_phone` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '公司电话',
	`address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '地址',
	`phone` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '联系方式',
	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '时间戳',	
	PRIMARY KEY(`id`),
	UNIQUE KEY `username` (`username`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

## init vue:
- install vue-cli3 webpack
- vue create checkpoint-app-vue
- cd checkpoint-app-vue
- npm install
- npm run serve
- add frontend

## docker
- Dockerfile
- 
- docker build -t bob/checkpoint-go .
- docker run -d --name go-checkpoint -p 5000:5000 bob/checkpoint-go:latest
- docker container logs 531

## kubernetes:
- install kubectl and minikube

## istio:

