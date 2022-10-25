package define

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"
var MailPassword = os.Getenv("MailPassword")

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间（s）
var CodeExpire = 300

// ObjectStorageType 对象存储类型
// 支持 minio\cos
var ObjectStorageType = os.Getenv("ObjectStorageType")

// TencentSecretKey 腾讯云对象存储
var TencentSecretKey = os.Getenv("TencentSecretKey")
var TencentSecretID = os.Getenv("TencentSecretID")
var CosBucket = "https://getcharzp-1256268070.cos.ap-chengdu.myqcloud.com"

// MinIOAccessKeyID MinIO 配置
var MinIOAccessKeyID = os.Getenv("MinIOAccessKeyID")
var MinIOAccessSecretKey = os.Getenv("MinIOAccessSecretKey")
var MinIOEndpoint = os.Getenv("MinIOEndpoint")
var MinIOBucket = os.Getenv("MinIOBucket")

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2006-01-02 15:04:05"

var TokenExpire = 3600
var RefreshTokenExpire = 7200
