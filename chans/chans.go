package chans

//GetPing Adds string "ping" to given channel
func GetPing(messages chan (string)) {
	messages <- "ping"
}
