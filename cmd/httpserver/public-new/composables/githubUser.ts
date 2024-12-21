/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import type { GithubUser } from "@/utils/http/types.gen"

const user = ref<GithubUser | null>(null)

export const useGithubUser = () => {
  return user
}
