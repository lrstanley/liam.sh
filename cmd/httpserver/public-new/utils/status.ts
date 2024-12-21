/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { createDiscreteApi } from "naive-ui"

export const { message, dialog, loadingBar } = createDiscreteApi(
  ["message", "dialog", "loadingBar"] // "notification"
)
