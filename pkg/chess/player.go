package chess

type Player struct {
	id int
	name string
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) GetId() int {
	return player.id
}

/*
* Once we connect to player database, we can just return the player id
*/
func CreatePlayer(id int, name string) Player {
	return Player{
		id: id,
		name: name,
	}
}