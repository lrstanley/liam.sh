/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

/**
 * shallowEqual - compares two values for shallow equality.
 */
export function shallowEqual(newv: any, oldv: any): boolean {
  if (newv == oldv) {
    return true
  }

  if (
    newv instanceof Array &&
    newv.length == oldv.length &&
    newv.some((v: any, j: any) => v == oldv[j])
  ) {
    return true
  }

  if (newv instanceof Object && shallowEqualObject(newv, oldv)) {
    return true
  }

  return false
}

/**
 * shallowEqualObject - compares two objects for shallow equality.
 */
function shallowEqualObject(object1: Record<string, any>, object2: Record<string, any>): boolean {
  const keys1 = Object.keys(object1)
  const keys2 = Object.keys(object2)

  if (keys1.length !== keys2.length) {
    return false
  }
  for (const key of keys1) {
    if (object1[key] !== object2[key]) {
      return false
    }
  }
  return true
}
