/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

/**
 * getNodeTree returns an araay of all headers that were obtained recrusively, under
 * a given root element.
 */
export function getNodeTree(root: HTMLElement): HeadingTree[] {
  return recurseNodes(root, [])
}

/**
 * HeadingTree represents a header, and its descendants, as a tree.
 */
export type HeadingTree = {
  id: string
  title: string
  depth: number
  children: HeadingTree[]
}

/**
 * recurseNodes recursively searches children of a root node looking for HTML headers
 * that contain an ID, and returns a flat array of nodes.
 */
function recurseNodes(root: HTMLElement, tree: HeadingTree[] = []): HeadingTree[] {
  if (!root) return []

  // Collect all headers, and their descendants.
  root.childNodes.forEach((child) => {
    if (child instanceof HTMLHeadingElement) {
      tree.push({
        id: child.id,
        title: child.textContent!,
        depth: +child.nodeName.replace(/[^0-9]+/g, ""),
        children: recurseNodes(child),
      })
    }
  })

  return reorder(tree)
}

/**
 * reorder takes a flat array of headings, and converts them to a proper tree of
 * descendants, based on the heading level.
 */
function reorder(tree: HeadingTree[]): HeadingTree[] {
  for (let i = 0; i < tree.length; i++) {
    if (i == 0) {
      continue
    }

    // If the previous header is larger, then take the current header in the loop
    // and append it as a child of the previous header.
    const current = tree[i]
    const previous = tree[i - 1]
    if (current && previous && current.depth > previous.depth) {
      previous.children = reorder([...(previous.children || []), current])
      tree.splice(i, 1)
      i--
    }
  }

  return tree
}
