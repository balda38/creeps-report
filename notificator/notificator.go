package notificator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/balda38/creeps-report/opendotaclient/match"
	"github.com/balda38/creeps-report/opendotaclient/types"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func NotifySubscribers(
	botInstance *bot.Bot,
	chats []int64,
	matchId int,
) {
	matchDetails := match.FetchMatch(matchId)

	var timeFormat string
	duration := time.Unix(matchDetails.Duration, 0).UTC()
	if matchDetails.Duration > 3600 {
		timeFormat = "15:04:05"
	} else {
		timeFormat = "04:05"
	}

	var winner string
	if matchDetails.RadiantWin {
		winner = matchDetails.RadiantName
	} else {
		winner = matchDetails.DireName
	}

	heroesFile, err := os.Open("constants/heroes.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer heroesFile.Close()

	var heroes map[string]types.Hero
	err = json.NewDecoder(heroesFile).Decode(&heroes)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	players := matchDetails.Players
	sort.Slice(players[:], func(i, j int) bool {
		return players[i].IsRadiant
	})

	var playerStats [10]string
	for i, player := range players {
		playerStats[i] = player.Name + " (" + heroes[strconv.Itoa(player.HeroId)].LocalizedName + ")" +
			" - " + strconv.Itoa(player.Kills) + "/" + strconv.Itoa(player.Deaths) + "/" + strconv.Itoa(player.Assists)
	}

	message := "🏆 <strong>" + matchDetails.League.Name + "</strong>\n\n" +
		"🎉 <strong>" + winner + " Victory</strong>\n\n" +
		"<strong>" + matchDetails.RadiantName + "</strong> " +
		strconv.Itoa(matchDetails.RadiantScore) + " - " + strconv.Itoa(matchDetails.DireScore) +
		" <strong>" + matchDetails.DireName + "</strong>\n" +
		"🕒 " + duration.Format(timeFormat) + "\n\n" +
		"🟢 <strong>Radiant</strong>\n" + strings.Join(playerStats[:5], "\n") + "\n\n" +
		"🔴 <strong>Dire</strong>\n" + strings.Join(playerStats[5:], "\n")

	for _, chatId := range chats {
		_, err := botInstance.SendMessage(context.Background(), &bot.SendMessageParams{
			ChatID:    chatId,
			Text:      message,
			ParseMode: models.ParseModeHTML,
		})
		if err != nil {
			log.Println("Failed to send message to chat "+strconv.FormatInt(chatId, 10)+":", err) // TODO: FormatInt 10?
		}
	}
}
