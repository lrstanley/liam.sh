/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */
import type { SchemaGithubUser } from '#open-fetch-schemas/api'

export const useGithubUser = () => useState<SchemaGithubUser | null | undefined>("githubUser", undefined)
