package server

func (s *basicServer) RegisterBasicRoutes() {
	r := s.engine

	sym := r.Group("/sym")
	{
		sym.POST("/login", s.symmetricLogin)
		sym.GET("/secret", s.validateSymmetricToken)
		sym.GET("/refresh-token", s.symmetricGetRf)
	}

	asym := r.Group("/asym")
	{
		asym.POST("/login", s.asymmetricLogin)
		asym.GET("/secret", s.validateAsymmetricToken)
	}
}
