package main

import (
	"fmt"
	"log"

	"github.com/onceuponworld/ouw-sdk"
)


func playerAction(a PlayerAction) {

	switch(a.Action) {
	case PLAYER_ACTION_STATS:

		// TODO: goto redis to check my stats, should also verify auth/auth before
		// can only check own stats, or stats of party
	
		key := fmt.Sprintf("%s:%s", ouwsdk.KEY_USER, a.Options["name"])

		log.Println(key)

		var u ouwsdk.User

		err := rds.HGetAll(ctx, key).Scan(&u)

		if err != nil {
			log.Println(err)
		} else {

			log.Printf("%+v\n", u)
		
		}

	case PLAYER_ACTION_TASK:
	default:
		log.Println("Action does not exist: ", a.Action)
	}

} // playerAction
