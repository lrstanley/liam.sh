/**
 * titleCase - converts a string to title case (every first letter of each word
 * capitalized).
 *
 * @export
 * @param {string} input
 * @return {*}  {string}
 */
export function titleCase(input: string): string {
  return input
    .toLowerCase()
    .split(" ")
    .map((word) => {
      return word.charAt(0).toUpperCase() + word.slice(1)
    })
    .join(" ")
}
