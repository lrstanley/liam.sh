<template>
    <LayoutBase>
        <div class="flex flex-auto justify-center items-center">
            <div class="flex-auto flex-col max-w-3xl max-h-xl py-7">
                <div class="flex-auto flex-row mb-4">
                    <span class="inline-flex text-size-45px mr-10px text-green-600">
                        liam
                        <span class="text-gray-500">:</span>
                        ~
                        <span class="text-gray-500 mr-4">$</span>
                        #
                    </span>
                    <span class="text-gradient text-size-45px from-teal-500 to-violet-600">Hello!</span>
                    <span class="text-size-40px text-gray-600 animate-pulse">▌</span>
                </div>
                <n-card content-style="padding: 0" class="p-0">
                    <n-space v-motion-slide-left vertical :size="12" class="p-5">
                        <!-- <div class="flex-auto flex-row">
                            <span class="inline-flex text-size-45px mr-10px text-green-600">
                                liam
                                <span class="text-gray-500">:</span>
                                ~
                                <span class="text-gray-500 mr-4">$</span> #
                            </span>
                            <span class="text-gradient text-size-45px from-teal-500 to-violet-600">Hello!</span>
                            <span class="text-size-40px text-violet-700 animate-pulse">▌</span>
                        </div> -->
                        <n-alert title="Default Text" type="info"> Gee it's good to be back home </n-alert>
                        <n-alert title="Info Text" type="info"> Gee it's good to be back home </n-alert>
                        <n-alert title="Success Text" type="success"> Leave it till tomorrow to unpack my case </n-alert>
                    </n-space>

                    <n-layout-footer bordered>
                        <span v-motion-fade class="flex flex-auto justify-end px-2"> &copy; Liam Stanley </span>
                    </n-layout-footer>
                </n-card>
            </div>
        </div>
    </LayoutBase>
</template>

<script setup></script>
