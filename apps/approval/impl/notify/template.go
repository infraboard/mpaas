package notify

import "embed"

//go:embed templates/*
var templatesDir embed.FS

func FeishuAuditNotifyTemplate() (string, error) {
	content, err := templatesDir.ReadFile("templates/feishu_card.json")
	if err != nil {
		return "", err
	}
	return string(content), nil
}
