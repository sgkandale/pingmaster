package server

func getPath(pathPrefix, path string) string {
	if pathPrefix == "" {
		return path
	}
	return pathPrefix + path
}

func (s Server) addRoutes(pathPrefix string) {

	userRoutes := s.Handler.Group(getPath(pathPrefix, "/user"))
	targetRoutes := s.Handler.Group(getPath(pathPrefix, "/target"))

	userRoutes.POST("/", s.registerUser)
	userRoutes.POST("/login", s.login)
	userRoutes.POST("/logout", s.logout)

	targetRoutes.POST("/", s.addTarget)
}
