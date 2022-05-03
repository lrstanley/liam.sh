<template>
    <n-config-provider :theme="darkTheme" :abstract="true">
        <n-loading-bar-provider>
            <router-view />
        </n-loading-bar-provider>
    </n-config-provider>
</template>

<script setup>
import { darkTheme } from "naive-ui"
</script>

<style>
html,
body {
    min-height: 100%;
    min-width: 100%;
    @apply flex flex-auto;
    @apply bg-dark-300 text-center text-light-600;
}

#app {
    font-family: Consolas, monaco, monospace;
    @apply antialiased;
    @apply flex flex-auto;
}

.text-gradient {
    @apply text-transparent bg-gradient-to-tr bg-clip-text inline font-bold;
}

.n-icon {
    display: inline-flex;
    align-self: center;
    height: initial;
    width: initial;
    line-height: 1em;
    font-size: 1.1em;
}

a {
    color: #42b983;
}

a:hover {
    color: #2f855a;
}

* {
    scrollbar-width: thin;
}

*::-webkit-scrollbar {
    background-color: transparent;
    width: 12px;
    height: 12px;
}
*::-webkit-scrollbar-thumb {
    transition: background 0.5s ease-in-out;
    border: 3px solid transparent;
    -webkit-background-clip: content-box;
    background-clip: content-box;
    @apply rounded-7px bg-green-800;
}
*::-webkit-scrollbar-thumb:hover {
    --tw-bg-opacity: 1;
    background-color: rgba(34, 80, 60, var(--tw-bg-opacity));
}
*::-webkit-scrollbar-thumb:window-inactive {
    background-color: #5e666b;
}

*::-webkit-scrollbar-corner {
    background-color: transparent;
}
</style>
