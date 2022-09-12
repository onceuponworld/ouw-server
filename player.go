package main

import (
	"log"
)


func playerAction(a ClientAction) {

	switch(a.Action) {
	case PLAYER_ACTION_STATS:

		// TODO: goto redis to check my stats, should also verify auth/auth before
		// can only check own stats, or stats of party
	default:
		log.Println("Action does not exist: ", a.Action)
	}

} // playerAction
