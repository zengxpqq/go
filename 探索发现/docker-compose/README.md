# docker-compose常用命令

1.编译docker镜像

```markdown
docker build -t name .
```

2.使用docker-compose 执行新建容器组

```markdown
docker-compose up -d // -d 后台运行

docker-compose up --force-recreate // 强制新建
```

3.启动容器组

```markdown
docker-compose start
```

4.停止容器组

```markdown
docker-compose stop
```

5.查询容器组所有容器状态

```markdown
docker-compose ps
```

6.删除容器组

```markdown
docker-compose down
```