package config

func ProvideLog() Log {
	return s.Log
}

func ProvideDbMap() DbMap {
	return s.DbMap
}

func ProvideHttp() Http {
	return s.Server.Http
}

func ProvideSign() Sign {
	return s.Sign
}