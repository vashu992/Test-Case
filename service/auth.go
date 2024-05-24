package pwg

import "propwg/model"

func GenerateAuthSession(authConfig AuthConfig) model.AuthSession {
	return model.AuthSession{
		AuthClec: model.AuthClec{
			ID: authConfig.ClecId,
			AuthAgentUser: model.AuthAgentUser{
				UserName: authConfig.UserName,
				Token:    authConfig.Token,
				Pin:      authConfig.Pin,
			},
		},
	}
}
