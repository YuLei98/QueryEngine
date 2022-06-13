package process

import "Search_Engines/search_engines/server/model"

type Userprocess struct{}

//使用者登录
func (this *Userprocess) Request_Login(redisId int, username string, password string) (responsestatus string,err error) {
	fp := model.New_User_Dao()
	responsestatus, err = fp.User_Login(redisId, &model.User{UserName: username, UserPwd: password})
	return 
}

//使用者注册
func (this *Userprocess) Request_Register(redisId int, username string, password string) (err error) {
	fp := model.New_User_Dao()
	err = fp.User_Register(redisId, &model.User{UserName: username, UserPwd: password})
	return
}

//使用者注销
func (this *Userprocess) Request_Del(redisId int, username string) (err error) {
	fp := model.New_User_Dao()
	err = fp.User_Del(redisId, &model.User{UserName: username})
	return err
}
