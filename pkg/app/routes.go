package app

func (a *App) RegisterRoutes() {
	a.e.POST("/users", a.CreateUser)
	a.e.POST("/login", a.Login)
}
