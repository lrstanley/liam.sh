import { h } from "vue"

import type { VNode } from "vue"

/**
 * createTOC recursively checks all HTML headers that contain an ID, under a given
 * root element, returning a list of heading descendants (taking into consideration
 * the heading level as the tree depth), and wraps the headings with the provided
 * wrapper element type.
 *
 * @export
 * @template T
 * @param {Node} root
 * @param {T} element
 * @returns {VNode[]}
 */
export function createTOC<T>(root: Node, wrapper: T): VNode[] {
  if (!root) {
    return []
  }

  return wrapDescendants(wrapper, getNodeTree(root))
}

/**
 * wrapDescendants wraps all descendants in a custom element type, and if it has
 * children, it will recursively wrap them as well.
 *
 * @template T
 * @param {T} wrapper
 * @param {Descendant[]} elements
 * @returns {VNode[]}
 */
function wrapDescendants<T>(wrapper: T, elements: Descendant[]): VNode[] {
  if (!elements) return []

  return elements.map((el) => {
    return h(
      wrapper,
      {
        href: el.href,
        title: el.title,
      },
      () => wrapDescendants(wrapper, el.descendants)
    )
  })
}

/**
 * getNodeTree returns a list of all headers, and their descendants, under a given
 * root element.
 *
 * @export
 * @param {Node} root
 * @returns {Descendant[]}
 */
export function getNodeTree(root: Node): Descendant[] {
  return recurseNodes(root, [])
}

/**
 * Descendant represents a header, and its descendants, as a tree.
 *
 * @interface Descendant
 * @typedef {Descendant}
 */
interface Descendant {
  href: string
  title: string
  type: Node["nodeName"]
  descendants: Descendant[]
}

/**
 * recurseNodes recursively searches children of a root node looking for HTML headers
 * that contain an ID, and builds a list of descendants.
 *
 * @param {Node} root
 * @param {Descendant[]} [descendants=[]]
 * @returns {Descendant[]}
 */
function recurseNodes(root: Node, descendants: Descendant[] = []): Descendant[] {
  // Collect all headers, and their descendants.
  root.childNodes.forEach((child) => {
    if (child instanceof HTMLHeadingElement) {
      descendants.push({
        href: `#${child.id}`,
        title: child.textContent,
        type: child.nodeName,
        descendants: recurseNodes(child),
      })
    }
  })

  for (let i = 0; i < descendants.length; i++) {
    if (i == 0) {
      continue
    }

    // If the previous header is larger, then take the current header in the loop
    // and append it as a child of the previous header.
    if (descendants[i].type > descendants[i - 1].type) {
      descendants[i - 1].descendants = [...descendants[i - 1].descendants, descendants[i]]
      descendants.splice(i, 1)
      i--
    }
  }

  return descendants
}
