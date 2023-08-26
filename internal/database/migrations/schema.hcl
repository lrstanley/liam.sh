table "ent_types" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "type" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
  index "ent_types_type_key" {
    unique  = true
    columns = [column.type]
  }
}
table "github_assets" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "asset_id" {
    null = false
    type = bigint
  }
  column "browser_download_url" {
    null = false
    type = character_varying
  }
  column "name" {
    null = false
    type = character_varying
  }
  column "label" {
    null = true
    type = character_varying
  }
  column "state" {
    null = true
    type = character_varying
  }
  column "content_type" {
    null = false
    type = character_varying
  }
  column "size" {
    null = false
    type = bigint
  }
  column "download_count" {
    null = false
    type = bigint
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "uploader" {
    null = false
    type = jsonb
  }
  column "github_release_assets" {
    null = false
    type = bigint
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "github_assets_github_releases_assets" {
    columns     = [column.github_release_assets]
    ref_columns = [table.github_releases.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "github_assets_asset_id_key" {
    unique  = true
    columns = [column.asset_id]
  }
}
table "github_events" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
      start     = 4294967296
    }
  }
  column "event_id" {
    null = false
    type = character_varying
  }
  column "event_type" {
    null = false
    type = character_varying
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "public" {
    null    = false
    type    = boolean
    default = false
  }
  column "actor_id" {
    null = false
    type = bigint
  }
  column "actor" {
    null = false
    type = jsonb
  }
  column "repo_id" {
    null = false
    type = bigint
  }
  column "repo" {
    null = false
    type = jsonb
  }
  column "payload" {
    null = false
    type = jsonb
  }
  primary_key {
    columns = [column.id]
  }
  index "github_events_event_id_key" {
    unique  = true
    columns = [column.event_id]
  }
}
table "github_gists" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
      start     = 8589934592
    }
  }
  column "gist_id" {
    null = false
    type = character_varying
  }
  column "html_url" {
    null = false
    type = character_varying
  }
  column "public" {
    null = false
    type = boolean
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "description" {
    null = true
    type = character_varying
  }
  column "owner" {
    null = false
    type = jsonb
  }
  column "name" {
    null = false
    type = character_varying
  }
  column "type" {
    null = false
    type = character_varying
  }
  column "language" {
    null = true
    type = character_varying
  }
  column "size" {
    null = false
    type = bigint
  }
  column "raw_url" {
    null = false
    type = character_varying
  }
  column "content" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
  index "githubgist_gist_id_name" {
    unique  = true
    columns = [column.gist_id, column.name]
  }
}
table "github_releases" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
      start     = 12884901888
    }
  }
  column "release_id" {
    null = false
    type = bigint
  }
  column "html_url" {
    null = false
    type = character_varying
  }
  column "tag_name" {
    null = false
    type = character_varying
  }
  column "target_commitish" {
    null = false
    type = character_varying
  }
  column "name" {
    null = true
    type = character_varying
  }
  column "draft" {
    null = false
    type = boolean
  }
  column "prerelease" {
    null = false
    type = boolean
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "published_at" {
    null = false
    type = timestamptz
  }
  column "author" {
    null = false
    type = jsonb
  }
  column "github_repository_releases" {
    null = false
    type = bigint
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "github_releases_github_repositories_releases" {
    columns     = [column.github_repository_releases]
    ref_columns = [table.github_repositories.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "github_releases_release_id_key" {
    unique  = true
    columns = [column.release_id]
  }
}
table "github_repositories" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
      start     = 17179869184
    }
  }
  column "repo_id" {
    null = false
    type = bigint
  }
  column "name" {
    null = false
    type = character_varying
  }
  column "full_name" {
    null = false
    type = character_varying
  }
  column "owner_login" {
    null = false
    type = character_varying
  }
  column "owner" {
    null = false
    type = jsonb
  }
  column "public" {
    null    = false
    type    = boolean
    default = false
  }
  column "html_url" {
    null = false
    type = character_varying
  }
  column "description" {
    null = true
    type = character_varying
  }
  column "fork" {
    null    = false
    type    = boolean
    default = false
  }
  column "homepage" {
    null = true
    type = character_varying
  }
  column "star_count" {
    null    = false
    type    = bigint
    default = 0
  }
  column "default_branch" {
    null = false
    type = character_varying
  }
  column "is_template" {
    null    = false
    type    = boolean
    default = false
  }
  column "has_issues" {
    null    = false
    type    = boolean
    default = true
  }
  column "archived" {
    null    = false
    type    = boolean
    default = false
  }
  column "pushed_at" {
    null = true
    type = timestamptz
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "license" {
    null = true
    type = jsonb
  }
  primary_key {
    columns = [column.id]
  }
  index "github_repositories_repo_id_key" {
    unique  = true
    columns = [column.repo_id]
  }
}
table "label_github_repositories" {
  schema = schema.public
  column "label_id" {
    null = false
    type = bigint
  }
  column "github_repository_id" {
    null = false
    type = bigint
  }
  primary_key {
    columns = [column.label_id, column.github_repository_id]
  }
  foreign_key "label_github_repositories_github_repository_id" {
    columns     = [column.github_repository_id]
    ref_columns = [table.github_repositories.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "label_github_repositories_label_id" {
    columns     = [column.label_id]
    ref_columns = [table.labels.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "label_posts" {
  schema = schema.public
  column "label_id" {
    null = false
    type = bigint
  }
  column "post_id" {
    null = false
    type = bigint
  }
  primary_key {
    columns = [column.label_id, column.post_id]
  }
  foreign_key "label_posts_label_id" {
    columns     = [column.label_id]
    ref_columns = [table.labels.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "label_posts_post_id" {
    columns     = [column.post_id]
    ref_columns = [table.posts.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "labels" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
      start     = 21474836480
    }
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  column "update_time" {
    null = false
    type = timestamptz
  }
  column "name" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
  index "labels_name_key" {
    unique  = true
    columns = [column.name]
  }
}
table "posts" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
      start     = 25769803776
    }
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  column "update_time" {
    null = false
    type = timestamptz
  }
  column "slug" {
    null = false
    type = character_varying
  }
  column "title" {
    null = false
    type = character_varying
  }
  column "content" {
    null = false
    type = character_varying
  }
  column "content_html" {
    null = false
    type = character_varying
  }
  column "summary" {
    null = false
    type = character_varying
  }
  column "published_at" {
    null = false
    type = timestamptz
  }
  column "view_count" {
    null    = false
    type    = bigint
    default = 0
  }
  column "user_posts" {
    null = false
    type = bigint
  }
  column "public" {
    null    = false
    type    = boolean
    default = false
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "posts_users_posts" {
    columns     = [column.user_posts]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "posts_slug_key" {
    unique  = true
    columns = [column.slug]
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
      start     = 30064771072
    }
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  column "update_time" {
    null = false
    type = timestamptz
  }
  column "user_id" {
    null = false
    type = bigint
  }
  column "login" {
    null = false
    type = character_varying
  }
  column "name" {
    null = true
    type = character_varying
  }
  column "avatar_url" {
    null = true
    type = character_varying
  }
  column "html_url" {
    null = true
    type = character_varying
  }
  column "email" {
    null = true
    type = character_varying
  }
  column "location" {
    null = true
    type = character_varying
  }
  column "bio" {
    null = true
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
  index "users_login_key" {
    unique  = true
    columns = [column.login]
  }
  index "users_user_id_key" {
    unique  = true
    columns = [column.user_id]
  }
}
schema "public" {
  comment = "standard public schema"
}
