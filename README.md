# slotsevo-admin
slotsevo项目后台管理,分为server和web两个子项目  
server golang版本1.17  
web vue2

## 部署
根据server中config.yaml配置数据库环境  
数据库数据初始化体验请使用[slotsevo_admin.sql](./server/slotsevo_admin.sql)进行表和数据初始化  
###使用docker-compose部署
先安装docker-compose环境  
使用docker-compose up -d 一键部署

