package actions

func DrinkWater() {
	bot := NewLifeBot("lifeBot", BotToken)
	msg := "该喝水啦"
	bot.SendMsg("911000205", msg)
}
