package post

import (
    "github.com/charlitoro/go-clean-architecture-skeleton/domain/entities"
)

type AddPostUseCase struct {
	// TODO: Inyect all dependensies neede once are ready

	// postDbRepository
	// postDbRepositoryImpl
	// cachingClient
	// postCachingRepository
	// postCachingRepositoryImpl
}

// NewAddPostUseCase adds a new post after validation
func NewAddPostUseCase() *AddPostUseCase {
	return &AddPostUseCase{}
}


func (uc *AddPostUseCase) Execute(input entities.PostType) map[string]interface{} {
	// TODO: here every needed busnes logic for the use case

	// Create the Post entity
    newPost := entities.Post(
        input.Title,
        input.Description,
        input.CreatedAt,
        input.UserId,
        input.IsPublished,
    )

	return map[string]interface{} {
		"title": newPost.Title,
		"description": newPost.Description,
		"isPublished": newPost.IsPublished,
	}
}
