package postgres

func incUserOnlinePair(source UserOnlinePair, shift int64) UserOnlinePair {
	return UserOnlinePair{
		UserID: source.UserID,
		Online: source.Online + shift,
	}
}
