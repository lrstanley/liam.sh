/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import type { AllowedComponentProps, Component, VNodeProps } from "vue"

/**
 * ComponentProps<C> is a type that takes a component as an argument and returns the type
 * representation of the component's props.
 */
export type ComponentProps<C extends Component> = C extends new (...args: any) => any
  ? Omit<InstanceType<C>["$props"], keyof VNodeProps | keyof AllowedComponentProps>
  : never
