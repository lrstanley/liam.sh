const purgecss = require("@fullhuman/postcss-purgecss")

module.exports = {
    plugins: [
        purgecss({
            content: [`./**/*.html`, `./src/**/*.vue`],
            defaultExtractor(content) {
                const contentWithoutStyleBlocks = content.replace(/<style[^]+?<\/style>/gi, "")
                return contentWithoutStyleBlocks.match(/[A-Za-z0-9-_/:]*[A-Za-z0-9-_/]+/g) || []
            },
            safelist: [
                /-(leave|enter|appear)(|-(to|from|active))$/,
                /^(?!(|.*?:)cursor-move).+-move$/,
                /^router-link(|-exact)-active$/,
                /data-v-.*/,
            ],
        }),
    ],
}
