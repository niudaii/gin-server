package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Server struct {
	Host    string  `yaml:"host"`
	Port    string  `yaml:"port"`
	Mysql   Mysql   `yaml:"mysql"`
	Captcha Captcha `yaml:"captcha"`
	Jwt     Jwt     `yaml:"jwt"`
}

func (s *Server) String() string {
	_, err := yaml.Marshal(*s)
	if err != nil {
		return fmt.Sprintf("%+v", *s)
	}
	var bf []byte
	bf, err = yaml.Marshal(s)
	return "\n" + string(bf) + "\n"
}

func (s *Server) BaseUrl() string {
	return "http://" + s.Host + ":" + s.Port
}

type Mysql struct {
	Path         string `yaml:"path"`           // 服务器地址
	Port         string `yaml:"port"`           // 端口
	Config       string `yaml:"config"`         // 高级配置
	Dbname       string `yaml:"db-name"`        // 数据库名
	Username     string `yaml:"username"`       // 数据库用户名
	Password     string `yaml:"password"`       // 数据库密码
	MaxIdleConns int    `yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `yaml:"log-mode"`       // 是否开启Gorm全局日志
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

type Captcha struct {
	KeyLong   int `yaml:"key-long"`   // 验证码长度
	ImgWidth  int `yaml:"img-width"`  // 验证码宽度
	ImgHeight int `yaml:"img-height"` // 验证码高度
}

type Jwt struct {
	SigningKey  string `yaml:"signing-key"`  // jwt签名
	ExpiresTime int64  `yaml:"expires-time"` // 过期时间
	Issuer      string `yaml:"issuer"`       // 签发者
}
