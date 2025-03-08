package notificator

import (
	"context"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/balda38/creeps-report/constants"
	"github.com/balda38/creeps-report/opendotaclient"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func NotifySubscribers(
	botInstance *bot.Bot,
	chats []int64,
	matchId int,
) {
	matchDetails := opendotaclient.FetchMatch(matchId)
	leagueMatches := opendotaclient.FetchLeagueMatches(matchDetails.League.ID)

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

	heroes := constants.GetHeroes()
	seriesTypes := constants.GetSeriesTypes()
	seriesType := seriesTypes[strconv.Itoa(matchDetails.SeriesType)]

	players := matchDetails.Players
	sort.Slice(players[:], func(i, j int) bool {
		return players[i].IsRadiant
	})

	var playerStats [10]string
	for i, player := range players {
		var playerName string
		if player.Name == "" {
			playerName = player.PersonaName
		} else {
			playerName = player.Name
		}
		playerName = strings.ReplaceAll(playerName, "&", "&amp;")
		playerName = strings.ReplaceAll(playerName, "<", "&lt;")
		playerName = strings.ReplaceAll(playerName, ">", "&gt;")

		playerStats[i] = playerName + " (" + heroes[strconv.Itoa(player.HeroId)].LocalizedName + ")" +
			" - " + strconv.Itoa(player.Kills) + "/" + strconv.Itoa(player.Deaths) + "/" + strconv.Itoa(player.Assists)
	}

	radinatTeamName := matchDetails.RadiantName
	radinatTeamName = strings.ReplaceAll(radinatTeamName, "&", "&amp;")
	radinatTeamName = strings.ReplaceAll(radinatTeamName, "<", "&lt;")
	radinatTeamName = strings.ReplaceAll(radinatTeamName, ">", "&gt;")

	direTeamName := matchDetails.DireName
	direTeamName = strings.ReplaceAll(direTeamName, "&", "&amp;")
	direTeamName = strings.ReplaceAll(direTeamName, "<", "&lt;")
	direTeamName = strings.ReplaceAll(direTeamName, ">", "&gt;")

	teamWinsInSeries := map[int]int{
		matchDetails.RadiandTeamId: 0,
		matchDetails.DireTeamId:    0,
	}
	for _, match := range leagueMatches {
		if match.SeriesId == matchDetails.SeriesId {
			if match.RadiantWin {
				teamWinsInSeries[match.RadiandTeamId]++
			} else {
				teamWinsInSeries[match.DireTeamId]++
			}
		}
	}
	radiantNumberOfMatchesToWinSeries := seriesType.RequiredNumberOfWins - teamWinsInSeries[matchDetails.RadiandTeamId]
	direNumberOfMatchesToWinSeries := seriesType.RequiredNumberOfWins - teamWinsInSeries[matchDetails.DireTeamId]
	radiantSeriesScore := strings.Repeat("○", radiantNumberOfMatchesToWinSeries) + strings.Repeat("●", teamWinsInSeries[matchDetails.RadiandTeamId])
	direSeriesScore := strings.Repeat("●", teamWinsInSeries[matchDetails.DireTeamId]) + strings.Repeat("○", direNumberOfMatchesToWinSeries)

	message :=
		// League name and winner
		"🏆 <strong>" + matchDetails.League.Name + "</strong>\n\n" +
			"🎉 <strong>" + winner + " Victory</strong>\n\n" +
			// Series type
			"<strong>" + seriesType.LongName + "</strong>\n" +
			// Match score
			radiantSeriesScore + " <strong>" + radinatTeamName + "</strong> " +
			strconv.Itoa(matchDetails.RadiantScore) + " - " + strconv.Itoa(matchDetails.DireScore) +
			" <strong>" + direTeamName + "</strong> " + direSeriesScore + "\n" +
			// Match duration
			"🕒 " + duration.Format(timeFormat) + "\n\n" +
			// Radiant team info
			"🟢 <strong>Radiant</strong>\n" + strings.Join(playerStats[:5], "\n") + "\n\n" +
			// Dire team info
			"🔴 <strong>Dire</strong>\n" + strings.Join(playerStats[5:], "\n")

	for _, chatId := range chats {
		_, err := botInstance.SendMessage(context.Background(), &bot.SendMessageParams{
			ChatID:    chatId,
			Text:      message,
			ParseMode: models.ParseModeHTML,
		})
		if err != nil {
			log.Println("Failed to send message to chat "+strconv.FormatInt(chatId, 10)+":", err)
		}
	}
}
