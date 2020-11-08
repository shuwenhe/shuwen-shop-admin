package cmd

import (
	"log"
	"os"

	"github.com/shuwenhe/shuwen-shop-admin/config"
	utils "github.com/shuwenhe/shuwen-shop-admin/db"
	"github.com/shuwenhe/shuwen-shop-admin/router"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logger  = &logrus.Logger{}
	rootCmd = &cobra.Command{}
)

func initConfig() {
	config.MustInit(os.Stdout, cfgFile) // 配置初始化
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/dev.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("debug", true, "开启debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		sqlDb, err := utils.Mysql( // Before router init mysql
			viper.GetString("db.hostname"),
			viper.GetInt("db.port"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dbname"),
		)
		log.Println("sqlDb = *** =", sqlDb)
		if err != nil {
			return err
		}
		defer utils.Db.Close()
		router.Run() // Router
		return nil
	}
	return rootCmd.Execute()
}
