/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import type { ComputedRef, Ref } from "vue"

const desc = "desc"
const asc = "asc"
const directions: string[] = [desc, asc]

type SortFields<T> = {
  [key: string]: T
}

type Filter = {
  order: ComputedRef<string>
  orderBy: ComputedRef<string>
}

export type Sorter = {
  toggle: (field: string) => void
  fields: SortFields<string>
  refs: { [key: string]: Ref<string> }
  filter: Filter
}

/**
 * useSorter - generates the necessary sorting and ordering data for a graphql
 * query.
 */
export function useSorter(
  fields: SortFields<string>,
  direction: Ref<string>,
  field: Ref<string>,
  defaultField: string = null,
  defaultDirection: string = desc
): Sorter {
  if (!defaultField) {
    defaultField = Object.keys(fields)[0]
  }

  return {
    toggle: (newField: string) => {
      if (newField === field.value) {
        direction.value = direction.value === desc ? asc : desc
      } else {
        field.value = newField
        direction.value = defaultDirection
      }
    },
    fields: fields,
    refs: {
      direction,
      field,
    },
    filter: {
      order: computed(() => {
        return (directions.includes(direction.value) ? direction.value : defaultDirection).toUpperCase()
      }),
      orderBy: computed(() => {
        return (Object.keys(fields).includes(field.value) ? field.value : defaultField).toUpperCase()
      }),
    },
  }
}
