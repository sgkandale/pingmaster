package server

func getPath(pathPrefix, path string) string {
	if pathPrefix == "" {
		return path
	}
	return pathPrefix + path
}

func (s Server) addRoutes(pathPrefix string) {

	userRoutes := s.Handler.Group(getPath(pathPrefix, "/user"))
	hostRoutes := s.Handler.Group(getPath(pathPrefix, "/host"))

	userRoutes.POST("/", s.registerUser)
	userRoutes.POST("/login", s.login)
	userRoutes.POST("/logout", s.logout)

	hostRoutes.GET("/", s.getHost)
	hostRoutes.POST("/", s.addHost)
	hostRoutes.PUT("/", s.updateHost)
	hostRoutes.DELETE("/", s.deleteHost)
}
