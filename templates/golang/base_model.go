package gotemplate

const (
	Base = "package models\n" +
		"type Base struct {\n" +
		"ID        int       `json:\"-\" gorm:\"primary_key\"`\n" +
		"CreatedAt time.Time `json:\"created_at\"`\n" +
		"UpdatedAt time.Time `json:\"updated_at\" gorm:\"default:CURRENT_TIMESTAMP\"`\n" +
		"}"
)
