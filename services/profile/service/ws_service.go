package service

import (
	"log"
	"time"

	"github.com/HackIllinois/api/services/profile/models"
	"github.com/gorilla/websocket"
)

type LiveLeaderboardConnection struct {
	Conn  *websocket.Conn `json:"-"`
	Limit int             `json:"limit"`
}

var (
	leaderboardConnections map[*websocket.Conn]LiveLeaderboardConnection = make(
		map[*websocket.Conn]LiveLeaderboardConnection,
	)
	leaderboardUpdateCooldown = 1 * time.Second
	leaderboardInactiveUpdate = 5 * time.Minute
	ChanUpdateLiveLeaderboard = make(chan struct{}, 32)
)

func LiveLeaderboardManager() {
	for {
		select {
		case <-ChanUpdateLiveLeaderboard:
		case <-time.After(leaderboardInactiveUpdate): // failsafe case
		}
		// send updated leaderboard to websocket
		go BroadcastUpdatedLeaderboard()

		// Prevents spamming websockets
		cooldown_ticker := time.NewTicker(leaderboardUpdateCooldown)
		reupdate_after_cooldown := false
	cooldown:
		for {
			select {
			case <-ChanUpdateLiveLeaderboard:
				reupdate_after_cooldown = true
			case <-cooldown_ticker.C:
				break cooldown
			}
		}
		if reupdate_after_cooldown {
			ChanUpdateLiveLeaderboard <- struct{}{}
		}
	}
}

func BroadcastUpdatedLeaderboard() error {
	full_leaderboard, err := GetProfileLeaderboard(make(map[string][]string))
	if err != nil {
		return err
	}

	for _, connection_info := range leaderboardConnections {
		leaderboard_to_send := *full_leaderboard
		if connection_info.Limit > 0 {
			leaderboard_to_send.LeaderboardEntries = leaderboard_to_send.LeaderboardEntries[:connection_info.Limit]
		}
		go SendUpdatedLeaderboard(connection_info.Conn, leaderboard_to_send)
	}

	return nil
}

// The leaderboard size should be no more than the limit that the client wants
func SendUpdatedLeaderboard(conn *websocket.Conn, leaderboard models.LeaderboardEntryList) {
	err := conn.WriteJSON(leaderboard)
	if err != nil {
		log.Printf("error occurred on writing to ws: %v", err)
		// Note: We will not close to socket as the read handler (HandleIncomingLeaderboardWS) will
		// do that for us
	}
}

func HandleIncomingLeaderboardWS(conn *websocket.Conn, limit int) {
	defer conn.Close()

	leaderboardConnections[conn] = LiveLeaderboardConnection{
		Conn:  conn,
		Limit: limit,
	}

	defer delete(leaderboardConnections, conn)

	leaderboard, err := GetProfileLeaderboardWithLimit(limit)
	if err != nil {
		conn.WriteMessage(
			websocket.CloseInternalServerErr,
			websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "Could not fetch leaderboard"),
		)
		return
	}

	err = conn.WriteJSON(leaderboard)
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("Connection failed before we could send back leadderboard", err)
		}
		return
	}

	for {
		var updated_settings LiveLeaderboardConnection
		err := conn.ReadJSON(&updated_settings)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("IsUnexpectedCloseError()", err)
			}
			conn.Close()
			break
		}

		updated_settings.Conn = conn
		leaderboardConnections[conn] = updated_settings

		leaderboard, err := GetProfileLeaderboardWithLimit(updated_settings.Limit)
		if err != nil {
			conn.WriteMessage(
				websocket.CloseInternalServerErr,
				websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "Could not fetch leaderboard"),
			)
			return
		}

		err = conn.WriteJSON(leaderboard)
		if err != nil {
			return
		}
	}
}
