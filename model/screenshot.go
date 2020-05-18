package model

type ScreenShot struct {
	Name       string `json:"name"`  // имя должно быть уникальным (как делаю сейчас: `id` фильма из themovie.org + рандом uuid)
	Order      int    `json:"order"` // порядок даю из фронта
	PublicUrls PubUrl `json:"public_urls"`
}
