// Code generated by ent, DO NOT EDIT.

package ent

import (
	"prompt-library/backend/ent/conversation"
	"prompt-library/backend/ent/prompt"
	"prompt-library/backend/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	conversationFields := schema.Conversation{}.Fields()
	_ = conversationFields
	// conversationDescMessages is the schema descriptor for messages field.
	conversationDescMessages := conversationFields[0].Descriptor()
	// conversation.MessagesValidator is a validator for the "messages" field. It is called by the builders before save.
	conversation.MessagesValidator = conversationDescMessages.Validators[0].(func(string) error)
	// conversationDescCreatedAt is the schema descriptor for created_at field.
	conversationDescCreatedAt := conversationFields[1].Descriptor()
	// conversation.DefaultCreatedAt holds the default value on creation for the created_at field.
	conversation.DefaultCreatedAt = conversationDescCreatedAt.Default.(func() time.Time)
	promptFields := schema.Prompt{}.Fields()
	_ = promptFields
	// promptDescTitle is the schema descriptor for title field.
	promptDescTitle := promptFields[0].Descriptor()
	// prompt.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	prompt.TitleValidator = promptDescTitle.Validators[0].(func(string) error)
	// promptDescContent is the schema descriptor for content field.
	promptDescContent := promptFields[1].Descriptor()
	// prompt.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	prompt.ContentValidator = promptDescContent.Validators[0].(func(string) error)
	// promptDescVotes is the schema descriptor for votes field.
	promptDescVotes := promptFields[3].Descriptor()
	// prompt.DefaultVotes holds the default value on creation for the votes field.
	prompt.DefaultVotes = promptDescVotes.Default.(int)
	// promptDescCreatedAt is the schema descriptor for created_at field.
	promptDescCreatedAt := promptFields[4].Descriptor()
	// prompt.DefaultCreatedAt holds the default value on creation for the created_at field.
	prompt.DefaultCreatedAt = promptDescCreatedAt.Default.(func() time.Time)
}