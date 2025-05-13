package limiter

type Actor struct {
	messages ActorChannel
}

type ActorChannel = chan string

func (a *Actor) Tell(message string) {
	a.messages <- message
}
