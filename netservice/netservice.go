package netservice

import (
	"github.com/KevinTroyT/keysmith/util"
	"github.com/gin-gonic/gin"
)

const NETWORK_CMD = "netservice"

const NETWORK_CMD_START = "start"

const NETWORK_CMD_STOP = "stop"

const NETWORK_CMD_RESTART = "restart"

var router = gin.Default()

func Start() error {
	router.Use(util.Cors())
	router.GET("/account", func(c *gin.Context) {
		mnemonic := c.Query("mnemonic")
		result := util.NewResult(c)
		accountId, err := handleAccount(mnemonic)
		if err != nil {
			result.Error(400, err.Error())
			return
		}
		result.Success(gin.H{"accountId": accountId})

	})

	router.GET("/generate", func(c *gin.Context) {
		result := util.NewResult(c)
		mnemonic, err := handleGenerate()
		if err != nil {
			result.Error(400, err.Error())
			return
		}
		result.Success(gin.H{"mnemonic": mnemonic})

	})

	router.GET("/legacyAddress", func(c *gin.Context) {
		mnemonic := c.Query("mnemonic")
		result := util.NewResult(c)
		legacyAddress, err := handleLegacyAddress(mnemonic)
		if err != nil {
			result.Error(400, err.Error())
			return
		}
		result.Success(gin.H{"legacyAddress": legacyAddress})

	})

	router.GET("/principal", func(c *gin.Context) {
		mnemonic := c.Query("mnemonic")
		result := util.NewResult(c)
		principal, err := handlePrincipal(mnemonic)
		if err != nil {
			result.Error(400, err.Error())
			return
		}
		result.Success(gin.H{"principal": principal})

	})
	router.GET("/privateKey", func(c *gin.Context) {
		mnemonic := c.Query("mnemonic")
		result := util.NewResult(c)
		privateKey, err := handlePrivateKey(mnemonic)
		if err != nil {
			result.Error(400, err.Error())
			return
		}
		result.Success(gin.H{"privateKey": privateKey})

	})
	router.GET("/publicKey", func(c *gin.Context) {
		mnemonic := c.Query("mnemonic")
		result := util.NewResult(c)
		publicKey, err := handlePublicKey(mnemonic)
		if err != nil {
			result.Error(400, err.Error())
			return
		}
		result.Success(gin.H{"publicKey": publicKey})

	})
	router.GET("/xPublicKey", func(c *gin.Context) {
		mnemonic := c.Query("mnemonic")
		result := util.NewResult(c)
		xPublicKey, err := handleXPublicKey(mnemonic)
		if err != nil {
			result.Error(400, err.Error())
			return
		}
		result.Success(gin.H{"xPublicKey": xPublicKey})

	})

	router.Run(":8081") // listen and serve on 0.0.0.0:8081
	return nil
}
