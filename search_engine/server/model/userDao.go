package model

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动后就初始化一个userDao实例
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体
//完成对User 结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func New_Init_UserDao(pool *redis.Pool) *UserDao {
	MyUserDao = &UserDao{
		pool: pool,
	}
	return MyUserDao
}
func New_User_Dao() (result *UserDao) {
	return MyUserDao
}

/*
功能：完成登录的校验
参数：数据库ID model.User结构体
实现：
	1.从链接池中取出一根链接
	2.根据数据库ID切换数据库
	3.从数据库中获取值
	4.序列化得到对应值
	5.如果用户名和pwd都正确，
	6.如果用户pwd有错误，则返回对应的错误信信息
返回：成功返回nil 失败返回错误
*/
func (this *UserDao) User_Login(redisID int, user *User) (ok string, err error) {
	//1.从链接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	//2.根据数据库ID切换数据库
	_, err = conn.Do("select", redisID)
	if err != nil {
		fmt.Println("切换数据库失败!!!", err)
		return
	}
	//3.从数据库中获取值
	result, err := redis.String(conn.Do("Hget", "users", user.UserName))

	// fmt.Println("请求登录者 ", user)
	// fmt.Println("debug 模式中   ", user.UserPwd)
	// fmt.Printf("%v \n", result)
	// fmt.Printf("%v \n", err)

	if err != nil {
		ok = ERREOR_USER_NOTEXISTS.Error()
		return
	}

	//4.如果用户名和pwd都正确
	if result == user.UserPwd {
		return "success", nil
	} else {
		//	5.如果用户pwd有错误，则返回对应的错误信信息
		return "用户密码有错误", ERREOR_USER_PWD
	}
}

/*
功能：实现用户注册函数
参数：
实现：
	1.从链接池中取出一根链接
	2.根据数据库ID切换数据库
	3.判断是否存在存在则返回对应错误信息
	4.写入数据库
	5.返回对应消息
返回:bool err
*/
func (this *UserDao) User_Register(redisID int, user *User) (err error) {
	//	1.从链接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()

	//	2.根据数据库ID切换数据库
	_, err = conn.Do("select", redisID)
	if err != nil {
		fmt.Println("切换数据库失败!!!", err)
		return err
	}

	//	3.判断是否存在存在则返回对应错误信息
	_, err = redis.String(conn.Do("HGet", "users", user.UserName))
	if err == nil {
		err = ERREOR_USER_EXISTS
		return
	}
	//	4.写入数据库
	_, err = conn.Do("Hset", "users", user.UserName, user.UserPwd)
	if err != nil {
		fmt.Println("保存注册用户错误 err", err)
		return err
	}
	return
}

/*
功能：注销用户
*/
func (this *UserDao) User_Del(redisID int, user *User) (err error) {
	//	1.从链接池中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()

	//	2.根据数据库ID切换数据库
	_, err = conn.Do("select", redisID)
	if err != nil {
		fmt.Println("切换数据库失败!!!", err)
		return err
	}

	//	3.操作数据库
	_, err = redis.String(conn.Do("Hdel", "users", user.UserName))
	if err == nil {
		err = ERREOR_USER_EXISTS
		return
	}
	return nil
}
