package helper

import "fmt"

func KyuubiBanner() {
	fmt.Println(`  _  __                 _     _ `)
	fmt.Println(` | |/ /                | |   (_)`)
	fmt.Println(` | ' /_   _ _   _ _   _| |__  _ `)
	fmt.Println(` |  <| | | | | | | | | | '_ \| |`)
	fmt.Println(` | . \ |_| | |_| | |_| | |_) | |`)
	fmt.Println(` |_|\_\__, |\__,_|\__,_|_.__/|_|`)
	fmt.Println(`       __/ |                    `)
	fmt.Println(`      |___/                 v1.0`)
}

var InvalidDomain = "Enter domain in correct format\n(Eg: example.tld or sub.example.tld)"
var Invalidcmd = "I don't know the command!"
var HelpMessage string = "Commands: \n/start to start the bot\n/help to know how to use the bot\n/author to know who created this bot\n/me to know your own details\n/subs domain.tld find subdomains\n"
var StartMessage string = "This is a telegram bot written Golang. \nIt fetches subdomains for a given domain!\n\nEnter domain in domain.tld format to get subdomains\n\nRun /help to know how to use this Bot\n\nMade By Vinay Kumar \nhttps://github.com/Linuxinet/\n"
var UserInfo = "\nFirst Name: %s\nLast Name: %s\nUsername: %s\nChat ID: %d\n"
var Aboutmessage = "This Bot is Created By \nhttps://github.com/Linuxinet/"
var DomainRegex = `(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`
