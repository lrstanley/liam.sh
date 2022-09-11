/**
 * Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
 * of this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

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
 * @param {T} wrapper
 * @returns {VNode[]}
 */
export function createTOC<T>(root: Node, wrapper: T): VNode[] {
  if (!root) {
    return []
  }

  return wrapChildren(wrapper, getNodeTree(root))
}

/**
 * wrapChildren wraps all descendants in a custom element type, and if it has
 * children, it will recursively wrap them as well.
 *
 * @template T
 * @param {T} wrapper
 * @param {HeadingTree[]} elements
 * @returns {VNode[]}
 */
function wrapChildren<T>(wrapper: T, elements: HeadingTree[]): VNode[] {
  if (!elements) return []

  return elements.map((el) => {
    return h(
      wrapper,
      {
        href: el.href,
        title: el.title,
      },
      () => wrapChildren(wrapper, el.children)
    )
  })
}

/**
 * getNodeTree returns an araay of all headers that were obtained recrusively, under
 * a given root element.
 *
 * @export
 * @param {Node} root
 * @returns {HeadingTree[]}
 */
export function getNodeTree(root: Node): HeadingTree[] {
  return recurseNodes(root, [])
}

/**
 * HeadingTree represents a header, and its descendants, as a tree.
 *
 * @interface HeadingTree
 * @typedef {HeadingTree}
 */
interface HeadingTree {
  href: string
  title: string
  type: Node["nodeName"]
  children: HeadingTree[]
}

/**
 * recurseNodes recursively searches children of a root node looking for HTML headers
 * that contain an ID, and returns a flat array of nodes.
 *
 * @param {Node} root
 * @param {HeadingTree[]} [tree=[]]
 * @returns {HeadingTree[]}
 */
function recurseNodes(root: Node, tree: HeadingTree[] = []): HeadingTree[] {
  if (!root) return []

  // Collect all headers, and their descendants.
  root.childNodes.forEach((child) => {
    if (child instanceof HTMLHeadingElement) {
      tree.push({
        href: `#${child.id}`,
        title: child.textContent,
        type: child.nodeName,
        children: recurseNodes(child),
      })
    }
  })

  return reorder(tree)
}

/**
 * reorder takes a flat array of headings, and converts them to a proper tree of
 * descendants, based on the heading level.
 *
 * @param {HeadingTree[]} tree
 * @returns {HeadingTree[]}
 */
function reorder(tree: HeadingTree[]): HeadingTree[] {
  for (let i = 0; i < tree.length; i++) {
    if (i == 0) {
      continue
    }

    // If the previous header is larger, then take the current header in the loop
    // and append it as a child of the previous header.
    if (tree[i].type > tree[i - 1].type) {
      tree[i - 1].children = reorder([...tree[i - 1].children, tree[i]])
      tree.splice(i, 1)
      i--
    }
  }

  return tree
}
