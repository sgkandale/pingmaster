package server

func getPath(pathPrefix, path string) string {
	if pathPrefix == "" {
		return path
	}
	return pathPrefix + path
}

func (s Server) addRoutes(pathPrefix string) {

	s.Handler.GET(
		getPath(pathPrefix, "/host"),
		s.getHost,
	)
	s.Handler.POST(
		getPath(pathPrefix, "/host"),
		s.addHost,
	)
	s.Handler.PUT(
		getPath(pathPrefix, "/host"),
		s.updateHost,
	)
	s.Handler.DELETE(
		getPath(pathPrefix, "/host"),
		s.deleteHost,
	)
}
