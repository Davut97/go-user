package app

func (a *App) RegisterRoutes() {
	a.e.POST("/user", a.CreateUser)
	a.e.POST("/login", a.Login)
	a.e.GET("/jokes", a.GetJokes)
}
