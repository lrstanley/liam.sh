import { useScroll } from "@vueuse/core"

function getScrollParent(el: HTMLElement) {
  let parent: HTMLElement | null = el.parentElement
  while (parent) {
    if (parent.scrollHeight > parent.clientHeight) {
      return parent
    }
    parent = parent.parentElement
  }
  return null
}

export function getComputedScrollParent(el: Ref<HTMLElement | null>) {
  return useScroll(
    computed(() => {
      if (!el.value) return null
      return getScrollParent(el.value)
    })
  )
}
