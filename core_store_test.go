package core_store

import (
	"fmt"
	"github.com/djumanoff/amqp"
	middleware "github.com/kirigaikabuto/common-lib/access-token-middleware"
	"testing"
)

var (
	coreServiceTest      CoreService
	accessTokenStoreTest middleware.AccessTokenStore
	cfgConfig            = amqp.Config{
		Host:        "localhost",
		VirtualHost: "",
		User:        "",
		Password:    "",
		Port:        5672,
		LogLevel:    5,
	}
	username string = ""
	password string = "passanya9912321123"
	email    string = ""
	fullName string = " Yerassyl"
	err      error
)

//func TestCoreService_Register(t *testing.T) {
//	sess := amqp.NewSession(cfgConfig)
//
//	if err = sess.Connect(); err != nil {
//		t.Error(err)
//		return
//	}
//	defer sess.Close()
//
//	var cltCfg = amqp.ClientConfig{
//		//ResponseX: "response",
//		//RequestX: "request",
//	}
//	clt, err := sess.Client(cltCfg)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	amqpClient := NewAmqpRequests(clt)
//
//	var redisCfg = middleware.RedisConfig{
//		Host: "127.0.0.1",
//		Port: 6379,
//	}
//	accessTokenStoreTest, err = middleware.NewAccessTokenStore(redisCfg)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	coreServiceTest = NewCoreService(*amqpClient, accessTokenStoreTest)
//	cmd := &CreateUserCommand{
//		Username: username,
//		Email:    email,
//		Password: password,
//		FullName: fullName,
//	}
//	_, err = coreServiceTest.Register(cmd)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Printf("%s test successfully ended\n", t.Name())
//}

//func TestCoreService_Login(t *testing.T) {
//	sess := amqp.NewSession(cfgConfig)
//
//	if err = sess.Connect(); err != nil {
//		t.Error(err)
//		return
//	}
//	defer sess.Close()
//
//	var cltCfg = amqp.ClientConfig{
//		//ResponseX: "response",
//		//RequestX: "request",
//	}
//	clt, err := sess.Client(cltCfg)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	amqpClient := NewAmqpRequests(clt)
//
//	var redisCfg = middleware.RedisConfig{
//		Host: "127.0.0.1",
//		Port: 6379,
//	}
//	accessTokenStoreTest, err = middleware.NewAccessTokenStore(redisCfg)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	coreServiceTest = NewCoreService(*amqpClient, accessTokenStoreTest)
//	cmd := &LoginUserCommand{
//		Username: username,
//		Password: password,
//	}
//	user, err := coreServiceTest.Login(cmd)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	userId, err := accessTokenStoreTest.Get(user.AccessKey)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	if userId == user.UserId {
//		fmt.Printf("%s test successfully ended", t.Name())
//	} else {
//		t.Error("no user by access key")
//		return
//	}
//}

//func TestCoreService_ListMovies(t *testing.T) {
//	sess := amqp.NewSession(cfgConfig)
//
//	if err = sess.Connect(); err != nil {
//		t.Error(err)
//		return
//	}
//	defer sess.Close()
//
//	var cltCfg = amqp.ClientConfig{
//		//ResponseX: "response",
//		//RequestX: "request",
//	}
//	clt, err := sess.Client(cltCfg)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	amqpClient := NewAmqpRequests(clt)
//
//	var redisCfg = middleware.RedisConfig{
//		Host: "127.0.0.1",
//		Port: 6379,
//	}
//	accessTokenStoreTest, err = middleware.NewAccessTokenStore(redisCfg)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	coreServiceTest = NewCoreService(*amqpClient, accessTokenStoreTest)
//	cmd := &ListMoviesCommand{
//		2,
//	}
//	_, err = coreServiceTest.ListMovies(cmd)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Printf("%s test successfully ended\n", t.Name())
//}

func TestCoreService_GetMovieById(t *testing.T) {
	sess := amqp.NewSession(cfgConfig)

	if err = sess.Connect(); err != nil {
		t.Error(err)
		return
	}
	defer sess.Close()

	var cltCfg = amqp.ClientConfig{
		//ResponseX: "response",
		//RequestX: "request",
	}
	clt, err := sess.Client(cltCfg)
	if err != nil {
		t.Error(err)
		return
	}
	amqpClient := NewAmqpRequests(clt)

	var redisCfg = middleware.RedisConfig{
		Host: "127.0.0.1",
		Port: 6379,
	}
	accessTokenStoreTest, err = middleware.NewAccessTokenStore(redisCfg)
	if err != nil {
		t.Error(err)
		return
	}
	coreServiceTest = NewCoreService(*amqpClient, accessTokenStoreTest)
	cmd := &GetMovieByIdCommand{
		2,
	}
	_, err = coreServiceTest.GetMovieById(cmd)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%s test successfully ended\n", t.Name())
}
