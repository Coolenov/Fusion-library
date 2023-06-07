package gormDb

import (
	"github.com/Coolenov/Fusion-library/types"
)

func GetContents() []types.Content {
	var contents []types.Content
	DBGORM.Table("posts").Find(&contents)
	return contents
}

func GetTagid(tag_text string) int64 {
	var tag types.Tag
	DBGORM.Table("tags").Where("text = ?", tag_text).First(&tag)
	return tag.Id

}

func GetPostids(tagId int64) []int64 {
	var postTags []types.PostTags
	DBGORM.Table("posts_tags").Where("tag_id = ?", tagId).Find(&postTags)

	var postIds []int64
	for _, postTag := range postTags {
		postIds = append(postIds, postTag.Post_id)
	}
	return postIds
}

func GetContentsByIds(postIds []int64) []types.Content {
	var contents []types.Content

	for _, postId := range postIds {
		var content types.Content
		DBGORM.Table("posts").Where("id = ?", postId).First(&content)
		contents = append(contents, content)
	}
	return contents
}

func GetContentById(post_id int64) types.Content {
	var content types.Content
	DBGORM.Table("posts").Where("id = ?", post_id).First(&content)
	return content
}

func GetAllTags() []types.Tag {
	var tags []types.Tag
	DBGORM.Table("tags").Find(&tags)
	return tags
}

func GetLastContentBySource(source string) types.Content {
	var content types.Content
	DBGORM.Table("posts").Where("source =?", source).Last(&content)
	return content
}

func GetSourceByid(post_id int64) string {
	var content types.Content
	DBGORM.Table("posts").Where("id = ?", post_id).Find(&content)
	return content.Source
}

func GetPreviousContentByid(post_id int64) types.Content {
	var content types.Content
	source := GetSourceByid(post_id)
	DBGORM.Table("posts").Order("id desc").Where("source = ? AND id < ?", source, post_id).Find(&content)
	return content
}

func GetNextContentByid(post_id int64) types.Content {
	var content types.Content
	source := GetSourceByid(post_id)
	DBGORM.Table("posts").Where("source = ? AND id > ?", source, post_id).First(&content)
	return content
}

func GetAllSources() []string {
	var sources []string
	DBGORM.Table("posts").Distinct().Pluck("source", &sources)
	return sources
}
