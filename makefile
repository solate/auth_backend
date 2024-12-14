


# 初始化数据库的目标
init-db:
	@echo "Initializing the MySQL database..."
	# 假设你有一个 SQL 脚本用于初始化数据库
	# mysql -u your_user -p your_password your_database < path/to/init.sql

# 生成代码的目标
gen:
	@echo "Generating code with ent and go-zero..."
	# 使用 ent 生成代码
	# 	go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
	# 使用 go-zero 生成代码
	# goctl api go -api path/to/your.api -dir .