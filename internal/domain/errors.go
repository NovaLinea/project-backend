package domain

import "errors"

var (
	ErrUserNotFound   = errors.New("Неверный логин или пароль")
	ErrGenerateToken  = errors.New("Could not login")
	ErrReplayUsername = errors.New("Такая поста уже используется")
)
