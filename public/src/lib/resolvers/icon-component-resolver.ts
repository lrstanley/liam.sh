import { readdirSync } from 'fs'
import { dirname } from 'path'
import { resolveModule } from 'local-pkg'
import type { ComponentResolver } from 'unplugin-vue-components/types'

let _cache: Array<string>

export interface IconResolverOptions {
  pkg: string
  prefix?: string
}

export function IconComponentResolver(options: IconResolverOptions): ComponentResolver {
  if (!_cache) {
    try {
      const icon_path = resolveModule(options.pkg) as string
      _cache = readdirSync(dirname(icon_path), { withFileTypes: true })
        .filter(item => !item.isDirectory() && item.name.match(/^[A-Z][A-Za-z0-9]+\.js$/))
        .map(item => item.name.replace(/\.js$/, ''))
    } catch (error) {
      console.error(error)
      throw new Error(`[unplugin-vue-components] failed to load "${options.pkg}", have you installed it?`)
    }
  }

  return {
    type: 'component',
    resolve: (name: string) => {
      if (name.startsWith(options.prefix)) {
        name = name.substring(options.prefix.length)
      }

      if (_cache.includes(name)) {
        return {
          name: name,
          from: options.pkg,
        }
      }
    },
  }
}
