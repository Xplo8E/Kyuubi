package runner

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Linuxinet/Kyuubi/arsenal"
	"github.com/Linuxinet/Kyuubi/helper"
	"github.com/Linuxinet/Kyuubi/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(BOT_TOKEN string) {

	bot, err := tgbotapi.NewBotAPI(BOT_TOKEN)
	if err != nil {
		if fmt.Sprint(err) == "Not Found" {
			log.Println("Bot not found!")
			log.Println("Error in Bot Token, Please Check!")
		} else {
			log.Println(err)
		}
		os.Exit(1)
	}
	bot.Debug = true
	log.Printf("Authorized to Bot: @%s\n", bot.Self.UserName)

	// get msg from the bot

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// user input from telegram bot
	//message sender
	//parsing code with subdomain output
	CommandCheck(updates, bot)
}

func CommandCheck(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		cmd_msg_from_user := update.Message.Command()
		cmd_msg_from_user_arg := update.Message.CommandArguments()
		bot_user := update.Message.Chat

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		InputmsgControll(cmd_msg_from_user, msg, bot_user, cmd_msg_from_user_arg, update, bot)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}

// Command Checker
func InputmsgControll(cmd_msg_from_user string, msg tgbotapi.MessageConfig, bot_user *tgbotapi.Chat, cmd_msg_from_user_arg string, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if cmd_msg_from_user == "help" {
		msg.Text = helper.HelpMessage
	} else if cmd_msg_from_user == "start" {
		msg.Text = helper.StartMessage
	} else if cmd_msg_from_user == "me" {
		msg.Text = fmt.Sprintf(helper.UserInfo, bot_user.FirstName, bot_user.LastName, bot_user.UserName, bot_user.ID)
	} else if cmd_msg_from_user == "author" {
		msg.Text = helper.Aboutmessage
	} else if cmd_msg_from_user == "subs" {
		match := utils.CheckDomainFormat(helper.DomainRegex, cmd_msg_from_user_arg)
		if match {
			result := arsenal.Getsubs(strings.ToLower(cmd_msg_from_user_arg))
			GiveOutput(result, msg, cmd_msg_from_user_arg, update, bot)
		} else {
			msg.Text = helper.InvalidDomain
		}
	} else {
		msg.Text = helper.Invalidcmd
	}
}

func GiveOutput(result string, msg tgbotapi.MessageConfig, cmd_msg_from_user_arg string, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if len(result) <= 4096 {
		log.Println("output size : ", len(result))
		msg.Text = result
	} else {
		filepath := strings.ToLower(cmd_msg_from_user_arg) + "_subdomains.txt"
		filename, trf := utils.Filecreate(result, filepath)
		log.Println(filename)
		file := tgbotapi.NewDocument(update.Message.From.ID, tgbotapi.FilePath(filename))
		if _, err := bot.Send(file); err != nil {
			log.Println(err)
		}
		// defer trf.Close()

		err := trf.Close()
		Error(err)
		e := os.Remove(filepath)
		Error(e)
	}
}

func Error(err error) {
	if err != nil {
		log.Println(err)
	}
}
