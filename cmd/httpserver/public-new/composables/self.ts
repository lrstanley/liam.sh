/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import type { User } from "@/utils/http/types.gen"

export const useSelf = () => useState<User | null | undefined>("self", undefined) // undefined == not fetched yet, null == error
