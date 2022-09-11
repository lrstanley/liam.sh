import { shallowEqual } from "@/lib/util"

import type { ComputedRef, Ref, WatchStopHandle } from "vue"

type Filter = {
  first: ComputedRef<number>
  last: ComputedRef<number>
  before: ComputedRef<string>
  after: ComputedRef<string>
}

/**
 * usePagination - generates the necessary pagination data for a graphql query,
 * based off the provided cursor in the form of "<a|b>.<cursor>", where a = after,
 * and b = before.
 *
 * @export
 * @param {Ref<string>} cursor
 * @param {number} [size=10]
 * @return {Filter}
 */
export function usePagination(cursor: Ref<string>, size = 10): Filter {
  return {
    first: computed(() => (!cursor.value ? size : cursor.value?.startsWith("a.") ? size : null)),
    last: computed(() => (cursor.value?.startsWith("b.") ? size : null)),
    before: computed(() => (cursor.value?.startsWith("b.") ? cursor.value.split(".")[1] : null)),
    after: computed(() => (cursor.value?.startsWith("a.") ? cursor.value.split(".")[1] : null)),
  }
}

/**
 * resetCursor - reset the cursor to the first page when any of the provided refs
 * change.
 *
 * @export
 * @param {Ref<any>} cursor
 * @param {Ref<any>[]} refs
 * @return {WatchStopHandle}
 */
export function resetCursor(cursor: Ref<string>, refs: Ref<any>[]): WatchStopHandle {
  return watch(refs, (newv, oldv) => {
    for (let i = 0; i < newv.length; i++) {
      if (!shallowEqual(newv[i], oldv[i])) {
        cursor.value = null
        return
      }
    }
  })
}