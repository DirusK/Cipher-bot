// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cipher-bot/ent/request"
	"cipher-bot/ent/schema"
	"cipher-bot/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	requestMixin := schema.Request{}.Mixin()
	requestMixinFields0 := requestMixin[0].Fields()
	_ = requestMixinFields0
	requestFields := schema.Request{}.Fields()
	_ = requestFields
	// requestDescCreatedAt is the schema descriptor for created_at field.
	requestDescCreatedAt := requestMixinFields0[0].Descriptor()
	// request.DefaultCreatedAt holds the default value on creation for the created_at field.
	request.DefaultCreatedAt = requestDescCreatedAt.Default.(func() time.Time)
	// requestDescUpdatedAt is the schema descriptor for updated_at field.
	requestDescUpdatedAt := requestMixinFields0[1].Descriptor()
	// request.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	request.DefaultUpdatedAt = requestDescUpdatedAt.Default.(func() time.Time)
	// request.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	request.UpdateDefaultUpdatedAt = requestDescUpdatedAt.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescLanguage is the schema descriptor for language field.
	userDescLanguage := userFields[4].Descriptor()
	// user.DefaultLanguage holds the default value on creation for the language field.
	user.DefaultLanguage = userDescLanguage.Default.(string)
}