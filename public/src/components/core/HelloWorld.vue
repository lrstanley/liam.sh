<template>
    <div class="hello">
        <n-space v-motion-slide-left vertical :size="12">
            <n-alert title="Default Text" type="info"> Gee it's good to be back home </n-alert>
            <n-alert title="Info Text" type="info"> Gee it's good to be back home </n-alert>
            <n-alert title="Success Text" type="success"> Leave it till tomorrow to unpack my case </n-alert>
            <n-alert title="Warning Text" type="warning"> Honey disconnect the phone </n-alert>
            <n-alert title="Error Text" type="error"> I'm back in the U.S.S.R. </n-alert>
        </n-space>
        <h1>{{ msg }}</h1>
        <p>
            For a guide and recipes on how to configure / customize this project,<br />
            check out the
            <a href="https://cli.vuejs.org" target="_blank" rel="noopener">vue-cli documentation</a>.
        </p>
        <h3>
            <!-- <i-mdi-account-box w:text="red-500" w:hover="text-blue-500" w:animate="pulse" w:transition="colors duration-1000 ease-in-out" /> -->
            Installed CLI Plugins
        </h3>
        <ul>
            <li>
                <a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-babel" target="_blank" rel="noopener"
                    >babel</a
                >
            </li>
            <li>
                <a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-router" target="_blank" rel="noopener"
                    >router</a
                >
            </li>
            <li>
                <a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-eslint" target="_blank" rel="noopener"
                    >eslint</a
                >
            </li>
        </ul>
        <h3>Essential Links</h3>
        <ul>
            <li>
                <a href="https://vuejs.org" target="_blank" rel="noopener">Core Docs</a>
            </li>
            <li>
                <a href="https://forum.vuejs.org" target="_blank" rel="noopener">Forum</a>
            </li>
            <li>
                <a href="https://chat.vuejs.org" target="_blank" rel="noopener">Community Chat</a>
            </li>
            <li>
                <a href="https://twitter.com/vuejs" target="_blank" rel="noopener">Twitter</a>
            </li>
            <li>
                <a href="https://news.vuejs.org" target="_blank" rel="noopener">News</a>
            </li>
        </ul>
        <h3>{{ msg }}</h3>
        <ul>
            <li>
                <a href="https://router.vuejs.org" target="_blank" rel="noopener">vue-router</a>
            </li>
            <li>
                <a href="https://vuex.vuejs.org" target="_blank" rel="noopener">vuex</a>
            </li>
            <li>
                <a href="https://github.com/vuejs/vue-devtools#vue-devtools" target="_blank" rel="noopener">vue-devtools</a>
            </li>
            <li>
                <a href="https://vue-loader.vuejs.org" target="_blank" rel="noopener">vue-loader</a>
            </li>
            <li>
                <a href="https://github.com/vuejs/awesome-vue" target="_blank" rel="noopener">awesome-vue</a>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    name: "HelloWorld",
    props: {
        msg: {
            type: String,
            default: "MAH NAME LEUM",
        },
    },
}
</script>

<style scoped>
h3 {
    margin: 40px 0 0;
}
ul {
    list-style-type: none;
    padding: 0;
}
li {
    display: inline-block;
    margin: 0 10px;
}
a {
    color: #42b983;
}
</style>
