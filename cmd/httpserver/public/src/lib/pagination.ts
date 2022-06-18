import { computed, Ref, ComputedRef } from "vue"

type Filter = {
  first: ComputedRef<any>;
  last: ComputedRef<any>;
  before: ComputedRef<any>;
  after: ComputedRef<any>;
}

export function usePagination(cursor: Ref<any>, size: number = 10): Filter {
  return {
    first: computed(() =>
      !cursor.value ? size : cursor.value?.startsWith("a.") ? size : null
    ),
    last: computed(() => (cursor.value?.startsWith("b.") ? size : null)),
    before: computed(() => (cursor.value?.startsWith("b.") ? cursor.value.split(".")[1] : null)),
    after: computed(() => (cursor.value?.startsWith("a.") ? cursor.value.split(".")[1] : null)),
  }
}
