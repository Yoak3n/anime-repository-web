<template>
    <div class="copyleft-wrapper">
        <p style="height: 1rem;" >Copyright © anime-repository-web Develop Team 2023
            <n-button text  color="black" @click="jump">
                <n-icon>
                    <LogoGithub />
                </n-icon>
            </n-button>
        </p>
        <p>
            anime-repository-web is licensed under the GNU General Public License v3.0
        </p>
    </div>
</template>

<script setup lang="ts">
import { NButton,NIcon } from 'naive-ui';
import { LogoGithub } from '@vicons/ionicons5'

const jump = ()=>{
    window.open("https://github.com/Yoak3n")
}
</script>

<style scoped >
p {
    color: #bbb;
}
</style>