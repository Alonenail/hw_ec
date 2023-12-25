package main

import s "main/internal/server"

func main() {
	server := s.NewServer()

	server.AddMiddleware(s.CheckUserRoleHeader)

	server.GET("/days_left_until", s.GetDaysLeft)

	server.Serve()
}
