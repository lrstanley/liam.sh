// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package runtime

import (
	"context"
	"time"

	"github.com/lrstanley/liam.sh/internal/database/schema"
	"github.com/lrstanley/liam.sh/internal/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/ent/githubevent"
	"github.com/lrstanley/liam.sh/internal/ent/githubgist"
	"github.com/lrstanley/liam.sh/internal/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/ent/label"
	"github.com/lrstanley/liam.sh/internal/ent/post"
	"github.com/lrstanley/liam.sh/internal/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	githubasset.Policy = privacy.NewPolicies(schema.GithubAsset{})
	githubasset.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := githubasset.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	githubassetFields := schema.GithubAsset{}.Fields()
	_ = githubassetFields
	// githubassetDescAssetID is the schema descriptor for asset_id field.
	githubassetDescAssetID := githubassetFields[0].Descriptor()
	// githubasset.AssetIDValidator is a validator for the "asset_id" field. It is called by the builders before save.
	githubasset.AssetIDValidator = githubassetDescAssetID.Validators[0].(func(int64) error)
	// githubassetDescBrowserDownloadURL is the schema descriptor for browser_download_url field.
	githubassetDescBrowserDownloadURL := githubassetFields[1].Descriptor()
	// githubasset.BrowserDownloadURLValidator is a validator for the "browser_download_url" field. It is called by the builders before save.
	githubasset.BrowserDownloadURLValidator = githubassetDescBrowserDownloadURL.Validators[0].(func(string) error)
	// githubassetDescName is the schema descriptor for name field.
	githubassetDescName := githubassetFields[2].Descriptor()
	// githubasset.NameValidator is a validator for the "name" field. It is called by the builders before save.
	githubasset.NameValidator = githubassetDescName.Validators[0].(func(string) error)
	githubevent.Policy = privacy.NewPolicies(schema.GithubEvent{})
	githubevent.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := githubevent.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	githubeventFields := schema.GithubEvent{}.Fields()
	_ = githubeventFields
	// githubeventDescEventType is the schema descriptor for event_type field.
	githubeventDescEventType := githubeventFields[1].Descriptor()
	// githubevent.EventTypeValidator is a validator for the "event_type" field. It is called by the builders before save.
	githubevent.EventTypeValidator = githubeventDescEventType.Validators[0].(func(string) error)
	// githubeventDescPublic is the schema descriptor for public field.
	githubeventDescPublic := githubeventFields[3].Descriptor()
	// githubevent.DefaultPublic holds the default value on creation for the public field.
	githubevent.DefaultPublic = githubeventDescPublic.Default.(bool)
	// githubeventDescRepoID is the schema descriptor for repo_id field.
	githubeventDescRepoID := githubeventFields[6].Descriptor()
	// githubevent.RepoIDValidator is a validator for the "repo_id" field. It is called by the builders before save.
	githubevent.RepoIDValidator = githubeventDescRepoID.Validators[0].(func(int64) error)
	githubgist.Policy = privacy.NewPolicies(schema.GithubGist{})
	githubgist.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := githubgist.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	githubgistFields := schema.GithubGist{}.Fields()
	_ = githubgistFields
	// githubgistDescHTMLURL is the schema descriptor for html_url field.
	githubgistDescHTMLURL := githubgistFields[1].Descriptor()
	// githubgist.HTMLURLValidator is a validator for the "html_url" field. It is called by the builders before save.
	githubgist.HTMLURLValidator = githubgistDescHTMLURL.Validators[0].(func(string) error)
	githubrelease.Policy = privacy.NewPolicies(schema.GithubRelease{})
	githubrelease.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := githubrelease.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	githubreleaseFields := schema.GithubRelease{}.Fields()
	_ = githubreleaseFields
	// githubreleaseDescReleaseID is the schema descriptor for release_id field.
	githubreleaseDescReleaseID := githubreleaseFields[0].Descriptor()
	// githubrelease.ReleaseIDValidator is a validator for the "release_id" field. It is called by the builders before save.
	githubrelease.ReleaseIDValidator = githubreleaseDescReleaseID.Validators[0].(func(int64) error)
	// githubreleaseDescHTMLURL is the schema descriptor for html_url field.
	githubreleaseDescHTMLURL := githubreleaseFields[1].Descriptor()
	// githubrelease.HTMLURLValidator is a validator for the "html_url" field. It is called by the builders before save.
	githubrelease.HTMLURLValidator = githubreleaseDescHTMLURL.Validators[0].(func(string) error)
	// githubreleaseDescTagName is the schema descriptor for tag_name field.
	githubreleaseDescTagName := githubreleaseFields[2].Descriptor()
	// githubrelease.TagNameValidator is a validator for the "tag_name" field. It is called by the builders before save.
	githubrelease.TagNameValidator = githubreleaseDescTagName.Validators[0].(func(string) error)
	// githubreleaseDescTargetCommitish is the schema descriptor for target_commitish field.
	githubreleaseDescTargetCommitish := githubreleaseFields[3].Descriptor()
	// githubrelease.TargetCommitishValidator is a validator for the "target_commitish" field. It is called by the builders before save.
	githubrelease.TargetCommitishValidator = githubreleaseDescTargetCommitish.Validators[0].(func(string) error)
	githubrepository.Policy = privacy.NewPolicies(schema.GithubRepository{})
	githubrepository.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := githubrepository.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	githubrepositoryFields := schema.GithubRepository{}.Fields()
	_ = githubrepositoryFields
	// githubrepositoryDescRepoID is the schema descriptor for repo_id field.
	githubrepositoryDescRepoID := githubrepositoryFields[0].Descriptor()
	// githubrepository.RepoIDValidator is a validator for the "repo_id" field. It is called by the builders before save.
	githubrepository.RepoIDValidator = githubrepositoryDescRepoID.Validators[0].(func(int64) error)
	// githubrepositoryDescName is the schema descriptor for name field.
	githubrepositoryDescName := githubrepositoryFields[1].Descriptor()
	// githubrepository.NameValidator is a validator for the "name" field. It is called by the builders before save.
	githubrepository.NameValidator = githubrepositoryDescName.Validators[0].(func(string) error)
	// githubrepositoryDescFullName is the schema descriptor for full_name field.
	githubrepositoryDescFullName := githubrepositoryFields[2].Descriptor()
	// githubrepository.FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	githubrepository.FullNameValidator = githubrepositoryDescFullName.Validators[0].(func(string) error)
	// githubrepositoryDescOwnerLogin is the schema descriptor for owner_login field.
	githubrepositoryDescOwnerLogin := githubrepositoryFields[3].Descriptor()
	// githubrepository.OwnerLoginValidator is a validator for the "owner_login" field. It is called by the builders before save.
	githubrepository.OwnerLoginValidator = githubrepositoryDescOwnerLogin.Validators[0].(func(string) error)
	// githubrepositoryDescPublic is the schema descriptor for public field.
	githubrepositoryDescPublic := githubrepositoryFields[5].Descriptor()
	// githubrepository.DefaultPublic holds the default value on creation for the public field.
	githubrepository.DefaultPublic = githubrepositoryDescPublic.Default.(bool)
	// githubrepositoryDescHTMLURL is the schema descriptor for html_url field.
	githubrepositoryDescHTMLURL := githubrepositoryFields[6].Descriptor()
	// githubrepository.HTMLURLValidator is a validator for the "html_url" field. It is called by the builders before save.
	githubrepository.HTMLURLValidator = githubrepositoryDescHTMLURL.Validators[0].(func(string) error)
	// githubrepositoryDescFork is the schema descriptor for fork field.
	githubrepositoryDescFork := githubrepositoryFields[8].Descriptor()
	// githubrepository.DefaultFork holds the default value on creation for the fork field.
	githubrepository.DefaultFork = githubrepositoryDescFork.Default.(bool)
	// githubrepositoryDescStarCount is the schema descriptor for star_count field.
	githubrepositoryDescStarCount := githubrepositoryFields[10].Descriptor()
	// githubrepository.DefaultStarCount holds the default value on creation for the star_count field.
	githubrepository.DefaultStarCount = githubrepositoryDescStarCount.Default.(int)
	// githubrepository.StarCountValidator is a validator for the "star_count" field. It is called by the builders before save.
	githubrepository.StarCountValidator = githubrepositoryDescStarCount.Validators[0].(func(int) error)
	// githubrepositoryDescIsTemplate is the schema descriptor for is_template field.
	githubrepositoryDescIsTemplate := githubrepositoryFields[12].Descriptor()
	// githubrepository.DefaultIsTemplate holds the default value on creation for the is_template field.
	githubrepository.DefaultIsTemplate = githubrepositoryDescIsTemplate.Default.(bool)
	// githubrepositoryDescHasIssues is the schema descriptor for has_issues field.
	githubrepositoryDescHasIssues := githubrepositoryFields[13].Descriptor()
	// githubrepository.DefaultHasIssues holds the default value on creation for the has_issues field.
	githubrepository.DefaultHasIssues = githubrepositoryDescHasIssues.Default.(bool)
	// githubrepositoryDescArchived is the schema descriptor for archived field.
	githubrepositoryDescArchived := githubrepositoryFields[14].Descriptor()
	// githubrepository.DefaultArchived holds the default value on creation for the archived field.
	githubrepository.DefaultArchived = githubrepositoryDescArchived.Default.(bool)
	labelMixin := schema.Label{}.Mixin()
	label.Policy = privacy.NewPolicies(schema.Label{})
	label.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := label.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	labelMixinFields0 := labelMixin[0].Fields()
	_ = labelMixinFields0
	labelFields := schema.Label{}.Fields()
	_ = labelFields
	// labelDescCreateTime is the schema descriptor for create_time field.
	labelDescCreateTime := labelMixinFields0[0].Descriptor()
	// label.DefaultCreateTime holds the default value on creation for the create_time field.
	label.DefaultCreateTime = labelDescCreateTime.Default.(func() time.Time)
	// labelDescUpdateTime is the schema descriptor for update_time field.
	labelDescUpdateTime := labelMixinFields0[1].Descriptor()
	// label.DefaultUpdateTime holds the default value on creation for the update_time field.
	label.DefaultUpdateTime = labelDescUpdateTime.Default.(func() time.Time)
	// label.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	label.UpdateDefaultUpdateTime = labelDescUpdateTime.UpdateDefault.(func() time.Time)
	// labelDescName is the schema descriptor for name field.
	labelDescName := labelFields[0].Descriptor()
	// label.NameValidator is a validator for the "name" field. It is called by the builders before save.
	label.NameValidator = labelDescName.Validators[0].(func(string) error)
	postMixin := schema.Post{}.Mixin()
	post.Policy = privacy.NewPolicies(schema.Post{})
	post.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := post.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreateTime is the schema descriptor for create_time field.
	postDescCreateTime := postMixinFields0[0].Descriptor()
	// post.DefaultCreateTime holds the default value on creation for the create_time field.
	post.DefaultCreateTime = postDescCreateTime.Default.(func() time.Time)
	// postDescUpdateTime is the schema descriptor for update_time field.
	postDescUpdateTime := postMixinFields0[1].Descriptor()
	// post.DefaultUpdateTime holds the default value on creation for the update_time field.
	post.DefaultUpdateTime = postDescUpdateTime.Default.(func() time.Time)
	// post.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	post.UpdateDefaultUpdateTime = postDescUpdateTime.UpdateDefault.(func() time.Time)
	// postDescSlug is the schema descriptor for slug field.
	postDescSlug := postFields[0].Descriptor()
	// post.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	post.SlugValidator = postDescSlug.Validators[0].(func(string) error)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[1].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)
	// postDescContent is the schema descriptor for content field.
	postDescContent := postFields[2].Descriptor()
	// post.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	post.ContentValidator = postDescContent.Validators[0].(func(string) error)
	// postDescContentHTML is the schema descriptor for content_html field.
	postDescContentHTML := postFields[3].Descriptor()
	// post.ContentHTMLValidator is a validator for the "content_html" field. It is called by the builders before save.
	post.ContentHTMLValidator = postDescContentHTML.Validators[0].(func(string) error)
	// postDescSummary is the schema descriptor for summary field.
	postDescSummary := postFields[4].Descriptor()
	// post.SummaryValidator is a validator for the "summary" field. It is called by the builders before save.
	post.SummaryValidator = postDescSummary.Validators[0].(func(string) error)
	// postDescPublishedAt is the schema descriptor for published_at field.
	postDescPublishedAt := postFields[5].Descriptor()
	// post.DefaultPublishedAt holds the default value on creation for the published_at field.
	post.DefaultPublishedAt = postDescPublishedAt.Default.(func() time.Time)
	// postDescViewCount is the schema descriptor for view_count field.
	postDescViewCount := postFields[6].Descriptor()
	// post.DefaultViewCount holds the default value on creation for the view_count field.
	post.DefaultViewCount = postDescViewCount.Default.(int)
	// post.ViewCountValidator is a validator for the "view_count" field. It is called by the builders before save.
	post.ViewCountValidator = postDescViewCount.Validators[0].(func(int) error)
	// postDescPublic is the schema descriptor for public field.
	postDescPublic := postFields[7].Descriptor()
	// post.DefaultPublic holds the default value on creation for the public field.
	post.DefaultPublic = postDescPublic.Default.(bool)
	userMixin := schema.User{}.Mixin()
	user.Policy = privacy.NewPolicies(schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescUserID is the schema descriptor for user_id field.
	userDescUserID := userFields[0].Descriptor()
	// user.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	user.UserIDValidator = userDescUserID.Validators[0].(func(int) error)
	// userDescLogin is the schema descriptor for login field.
	userDescLogin := userFields[1].Descriptor()
	// user.LoginValidator is a validator for the "login" field. It is called by the builders before save.
	user.LoginValidator = userDescLogin.Validators[0].(func(string) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescAvatarURL is the schema descriptor for avatar_url field.
	userDescAvatarURL := userFields[3].Descriptor()
	// user.AvatarURLValidator is a validator for the "avatar_url" field. It is called by the builders before save.
	user.AvatarURLValidator = userDescAvatarURL.Validators[0].(func(string) error)
	// userDescHTMLURL is the schema descriptor for html_url field.
	userDescHTMLURL := userFields[4].Descriptor()
	// user.HTMLURLValidator is a validator for the "html_url" field. It is called by the builders before save.
	user.HTMLURLValidator = userDescHTMLURL.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[5].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescLocation is the schema descriptor for location field.
	userDescLocation := userFields[6].Descriptor()
	// user.LocationValidator is a validator for the "location" field. It is called by the builders before save.
	user.LocationValidator = userDescLocation.Validators[0].(func(string) error)
	// userDescBio is the schema descriptor for bio field.
	userDescBio := userFields[7].Descriptor()
	// user.BioValidator is a validator for the "bio" field. It is called by the builders before save.
	user.BioValidator = userDescBio.Validators[0].(func(string) error)
}

const (
	Version = "v0.11.3-0.20220830071904-3b1b75b9d7a9"           // Version of ent codegen.
	Sum     = "h1:dxyBYasfOuLgAT0IuqSXNgDDJ9ra5C+Jc8a6VqVywWo=" // Sum of ent codegen.
)
