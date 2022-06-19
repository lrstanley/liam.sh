import { computed, nextTick, Ref, ComputedRef } from "vue";

const desc: string = "desc"
const asc: string = "asc"
const directions: string[] = [desc, asc]

type SortFields<T> = {
    [key: string]: T,
}

type Filter = {
    order: ComputedRef<string>,
    orderBy: ComputedRef<string>
}

export type Sorter = {
    toggle: (field: string) => void,
    fields: SortFields<string>,
    refs: { [key: string]: Ref<string> },
    filter: Filter
}

export function useSorter(
    fields: SortFields<string>, direction: Ref<string>, field: Ref<string>,
    defaultField: string = null, defaultDirection: string = desc
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
                // https://github.com/vueuse/vueuse/issues/901
                setTimeout(() => {
                    nextTick(() => {
                        direction.value = defaultDirection
                    })
                }, 1)
            }
        },
        fields: fields,
        refs: {
            direction,
            field
        },
        filter: {
            order: computed(() => {
                return (directions.includes(direction.value) ? direction.value : defaultDirection).toUpperCase()
            }),
            orderBy: computed(() => {
                return (Object.keys(fields).includes(field.value) ? field.value : defaultField).toUpperCase()
            }),
        }
    }
}
