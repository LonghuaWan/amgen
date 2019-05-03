## cmgen

golang mgo 代码自动生成器, 生成对 mongodb 的 CRUD 操作代码。forked from [https://github.com/yakumioto/mgen](https://github.com/yakumioto/mgen)

## 安装

```bash
go install github.com/chxfantasy/cmgen/cmgen
```

## 使用

### mgo command

编写 配置文件, 然后使用命令生成. 可以参照 [example](/example)

```text
NAME:
   flag mgo - generate golang code

USAGE:
   flag mgo [command options] [arguments...]

OPTIONS:
   --config-file value, -c value  set the config file path
   --help, -h                     show help (default: false)
```

根据所传入的配置文件生成对应的 CRUD package.

example: `cmgen mgo -c xxx.yaml`

会在执行命令会在当前文件夹下生成一个 `xxx.mg.go` 的文件.

### 配置文件编写

#### 简单的用法

[base.yaml](/example/base/base.yaml)
```yaml
packageName: base
models:
  - name: User
    collectionName: users
    fields:
      - name: UserName
        type: string
      - name: Email
        type: string
      - name: Password
        type: string
```

执行后生成的Go文件: [base.mg.go](/example/base/base.mg.go)

如果指定了 `collectionName` 就会生成这个对应的 `CRUD` 方法

- NewUser() *User
- (user *User) Insert() error
- UpdateUserByID(id string, user *User) error
- UpdateUserByIDAndEntityMap(id interface{}, updateMap map[string]interface{}) error
- UpdateUser(selector interface{}, user *User) error
- UpdateUserAll(selector interface{}, user *User) (*mgo.ChangeInfo, error)
- GetUserByID(id string) (*User, error)
- GetOneUserByQuery(query map[string]interface{}) (*User, error)
- ListAllUserByQuery(query map[string]interface{}) ([]*User, error)
- ExistUserByID(id string) (bool, error)
- ExistUserByQuery(query interface{}) (bool, error)
- DeleteUserByID(id string) error

### 说明
1. 代码中会自动插入CreatedAt、UpdatedAt、deleted，所有的删除均为逻辑删除
2. 自动生成的Model代码中，引入了以下包。如有需要，请自行更改。
    ```go
    import (
        "errors"
        "github.com/globalsign/mgo"
        "github.com/globalsign/mgo/bson"
        "github.com/liamylian/jsontime"
        "time"
    )
    ```
    
3. init函数中，需要自行初始化MongoSession

#### 进阶用法

```yaml
packageName: advanced
models:
  - name: User
    collectionName: users
    fields:
      - name: UserName
        type: string
        unique: yes
        valid: required~first name is blank
      - name: Email
        type: string
        unique: yes
        valid: required,email
      - name: Password
        type: string
        valid: required
```

执行后生成的Go文件: [advanced.mgo.go](/example/advanced/advanced.mg.go)

`unique` 用来指定唯一

`valid` 用法实在太多了 使用的是 <https://github.com/asaskevich/govalidator>

## 感谢

The MongoDB driver for Go <https://github.com/globalsign/mgo>

Package of validators and sanitizers for strings, numerics, slices and structs <https://github.com/asaskevich/govalidator>