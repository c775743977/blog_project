package pbfile

import (
	"context"
	"errors"
	"time"
	"fmt"

	"myblogs/user/grpc/utils"
	"myblogs/user/grpc/db"
	"myblogs/user/grpc/model"

	"gorm.io/gorm"
)

type User struct {
	Username string
	Password string
	Email string
}

type UserService struct {

}

func(this *UserService) Login(ctx context.Context, req *LoginReq) (*LoginRes, error) {
	user := &model.User{}
	if err := db.MDB.Where("username = ?", req.Username).Find(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &LoginRes{}, errors.New("该用户不存在")
		}
		return &LoginRes{}, err
	}

	if req.Password != user.Password {
		return &LoginRes{}, errors.New("用户名或密码错误")
	}

	cookie := utils.CreateUUID()
	fmt.Println("cookie:", cookie)

	// redis_data := make(map[string]interface{})
	// redis_data["username"] = req.Username
	// redis_data["password"] = req.Password

	err := db.RDB.HSet(ctx , cookie, "username", req.Username).Err()
	if err != nil {
		fmt.Println("redis hmset error:", err)
		return &LoginRes{}, err
	}

	err = db.RDB.Expire(ctx, cookie, 20 * time.Minute).Err()
	if err != nil {
		fmt.Println("set cookie expire-time error:", err)
		return &LoginRes{}, err
	}

	return &LoginRes{
		Cookie: cookie,
		Username: user.Username,
	}, nil
}

func(this *UserService) SendCode(ctx context.Context, req *SendCodeReq) (*SendCodeReq, error) {
	code := utils.CreateCode()
	err := utils.SendMail(code, req.Email)
	if err != nil {
		return &SendCodeReq{}, err
	}

	err = db.RDB.Set(context.Background(), req.Email, code, time.Minute * 3).Err()
	if err != nil {
		return &SendCodeReq{}, err
	}
	return &SendCodeReq{}, nil
}

func(this *UserService) VerifyCode(ctx context.Context, req *VerifyCodeReq) (*VerifyCodeRes, error) {
	if utils.CheckUsername(req.Username) {
		return &VerifyCodeRes{}, errors.New("该用户名已被注册")
	}
	saved_code, _ := db.RDB.Get(context.Background(), req.Email).Result()

	if saved_code != req.Code {
		return &VerifyCodeRes{}, errors.New("email-code verification failed!")
	}

	user := &User{
		Username: req.Username,
		Password: req.Password,
		Email: req.Email,
	}

	if err := db.MDB.Create(user).Error; err != nil {
		return &VerifyCodeRes{}, err
	}

	return &VerifyCodeRes{}, nil
}

func(this *UserService) GetSession(ctx context.Context, req *SessionReq) (*SessionRes, error) {
	res, _ := db.RDB.HGet(ctx, req.Cookie, "username").Result()
	err := db.RDB.Expire(ctx, req.Cookie, 20 * time.Minute).Err()
	if err != nil {
		fmt.Println("set cookie expire-time error:", err)
		return &SessionRes{}, err
	}
	return &SessionRes{
		Name: res,
	}, nil
}

func(this *UserService) Logout(ctx context.Context, req *LogoutReq) (*LogoutReq, error) {
	err := db.RDB.HDel(ctx, req.Cookie).Err()
	if err != nil {
		return &LogoutReq{}, err
	}
	return &LogoutReq{}, nil
}

func(this *UserService) mustEmbedUnimplementedUserServiceServer() {}