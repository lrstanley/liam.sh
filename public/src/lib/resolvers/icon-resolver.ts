import { readdirSync } from 'fs'
import { dirname } from 'path'
import { resolveModule } from 'local-pkg'
import type { ImportsMap } from 'unplugin-auto-import/types'

let _cache: ImportsMap | undefined

export default (pkg: string): ImportsMap => {
  if (!_cache) {
    let packages: Array<string>
    try {
      const icon_path = resolveModule(pkg) as string
      packages = readdirSync(dirname(icon_path), { withFileTypes: true })
        .filter(item => !item.isDirectory() && item.name.match(/^[A-Z][A-Za-z0-9]+\.js$/))
        .map(item => item.name.replace(/\.js$/, ''))
    } catch (error) {
      console.error(error)
      throw new Error(`[auto-import] failed to load "${pkg}", have you installed it?`)
    }

    if (packages) {
      _cache = {
        [pkg]: packages,
      }
    }
  }

  return _cache || {}
}
