package main

import (
	"github.com/Linuxinet/Kyuubi/helper"
	"github.com/Linuxinet/Kyuubi/runner"
	"github.com/Linuxinet/Kyuubi/utils"
)

func main() {
	var BOT_TOKEN = utils.GoDotEnvVariable("BOT_TOKEN")
	helper.KyuubiBanner()
	runner.StartBot(BOT_TOKEN)
}
