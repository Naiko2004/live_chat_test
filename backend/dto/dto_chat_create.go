package dto

type ChatCreateDto struct {
	Name      string   `json:"name"`
	Usernames []string `json:"usernames"`
}
