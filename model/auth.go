package model

type AuthSession struct {
	AuthClec AuthClec `xml:"clec"`
}

type AuthClec struct {
	ID            string        `xml:"id"`
	AuthAgentUser AuthAgentUser `xml:"agentUser"`
}

type AuthAgentUser struct {
	UserName string `xml:"username" validate:"required"`
	Token    string `xml:"token" validate:"required"`
	Pin      string `xml:"pin" validate:"required"`
}
