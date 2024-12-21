/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import type { User } from "@/utils/http/types.gen"

const self = ref<User | null | undefined>(undefined) // undefined == not fetched yet, null == error

export const useSelf = () => {
  return self
}
