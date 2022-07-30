// Code generated by ent, DO NOT EDIT.

package ent

import (
	"moments/ent/post"
	"moments/ent/schema"
	"moments/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	postMixin := schema.Post{}.Mixin()
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreatedDate is the schema descriptor for created_date field.
	postDescCreatedDate := postMixinFields0[0].Descriptor()
	// post.DefaultCreatedDate holds the default value on creation for the created_date field.
	post.DefaultCreatedDate = postDescCreatedDate.Default.(time.Time)
	// postDescUpdatedDate is the schema descriptor for updated_date field.
	postDescUpdatedDate := postMixinFields0[1].Descriptor()
	// post.DefaultUpdatedDate holds the default value on creation for the updated_date field.
	post.DefaultUpdatedDate = postDescUpdatedDate.Default.(time.Time)
	// postDescLikes is the schema descriptor for likes field.
	postDescLikes := postFields[1].Descriptor()
	// post.DefaultLikes holds the default value on creation for the likes field.
	post.DefaultLikes = postDescLikes.Default.(uint64)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedDate is the schema descriptor for created_date field.
	userDescCreatedDate := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedDate holds the default value on creation for the created_date field.
	user.DefaultCreatedDate = userDescCreatedDate.Default.(time.Time)
	// userDescUpdatedDate is the schema descriptor for updated_date field.
	userDescUpdatedDate := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedDate holds the default value on creation for the updated_date field.
	user.DefaultUpdatedDate = userDescUpdatedDate.Default.(time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescFullName is the schema descriptor for full_name field.
	userDescFullName := userFields[3].Descriptor()
	// user.FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	user.FullNameValidator = userDescFullName.Validators[0].(func(string) error)
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[4].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
	// userDescIsStaff is the schema descriptor for is_staff field.
	userDescIsStaff := userFields[5].Descriptor()
	// user.DefaultIsStaff holds the default value on creation for the is_staff field.
	user.DefaultIsStaff = userDescIsStaff.Default.(bool)
}