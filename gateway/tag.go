package tag

import "go.uber.org/zap"

// URL should be used for logging any URLs.
func URL(url string) zap.Field {
	return zap.String("url", url)
}

func Attempt(i int) zap.Field {
	return zap.Int("attempt", i)
}
