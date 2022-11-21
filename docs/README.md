# 使用文档

## 搭建MongoDB

```
docker pull mongo
docker run -itd -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=123456 mongo
```

修改数据库用户:
```sh
# docker exec -it mongo mongo
> use mpaas
switched to db mpaas
> db.createUser({user: "mpaas", pwd: "123456", roles: [{ role: "dbOwner", db: "mpaas" }]})