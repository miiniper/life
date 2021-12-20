package actions

import (
	"fmt"
	"time"
)

func DrinkWater() {
	bot := NewLifeBot("lifeBot", BotToken)
	msg := fmt.Sprintf("生活小蜜提醒您：\n现在是 %s \n您该喝水啦，休息一下!\n", time.Now().Format(time.Stamp))
	bot.SendMsg("911000205", msg)
}
