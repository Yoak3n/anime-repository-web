<template>
    <div class="setting-wrapper">
        <n-divider>界面</n-divider>

        <n-switch :rail-style="railStyle" @update:value="handleChange" size="large">
            <template #checked>
                黑夜
            </template>
            <template #unchecked>
                白天
            </template>
        </n-switch>
        <n-divider>存储</n-divider>
        <n-form :model="configData">
            <n-form-item label="待处理文件路径">
                <n-input placeholder="" v-model:value="configData.raw_path"></n-input>
            </n-form-item>
            <n-form-item label="电影文件转移路径">
                <n-input placeholder="" v-model:value="configData.tv_path"></n-input>
            </n-form-item>
            <n-form-item label="电视文件转移路径">
                <n-input placeholder="" v-model:value="configData.tv_path"></n-input>
            </n-form-item>
            <n-divider>网络</n-divider>
            <n-form-item label="是否使用代理" label-placement="left">
                <n-switch size="large" v-model:value="configData.use_proxy"></n-switch>
            </n-form-item>
            <n-form-item label="代理地址">
                <n-input placeholder="http://xxx.xxx.xxx.xxx:xxx" v-model:value="configData.proxy_url"></n-input>
            </n-form-item>
            <n-divider>系统</n-divider>
            <n-space justify="start">
                <n-form-item label="tmdb key" label-placement="left" style="width: 200%;">
                    <n-input placeholder="请输入tmdb的api key" v-model:value="configData.key" style="width: 100%;"></n-input>
                </n-form-item>
                <n-form-item label="扫描间隔" label-placement="left" style="width: 200%;margin: 0 100%;">
                    <n-input-number 
                    :default-value="60" 
                    placeholder="" 
                    v-model:value="configData.delay"
                    style="width: 100%;"></n-input-number>
                </n-form-item>    
            </n-space>
        </n-form>
        <n-divider>关于</n-divider>
        <about></about>
        <copyleft />
        <n-button circle @click="saveSetting" :bordered="false">
            <template #icon>
                <n-icon size="30">
                    <SaveOutline />
                </n-icon>
            </template>
        </n-button>
    </div>
</template>

<script setup lang="ts">
import { ref, CSSProperties, inject, onMounted } from 'vue';
import { NSwitch, NForm, NFormItem, NDivider, NInput, NInputNumber, NSpace, NButton, NIcon } from 'naive-ui'
import { SaveOutline } from '@vicons/ionicons5'
import About from '../components/extra/About/index.vue';
import Copyleft from '../components/extra/Copyleft/index.vue'
import { reqPostSetting, reqGetSetting } from "../api/setting";
import type { Setting } from "../api/setting/type";

let configData = ref<Setting>({
    key: "", use_proxy: false, proxy_url: "", raw_path: "", tv_path: "", movie_path: "", delay: 60
})
const switchTheme: Function | undefined = inject('switchTheme')

const handleChange = () => {
    switchTheme?.apply(this)
}

const railStyle = ({ focused, checked }: { focused: boolean, checked: boolean }) => {
    const style: CSSProperties = {}
    if (checked) {
        style.background = '#000000'
        if (focused) {
            style.boxShadow = '0 0 0 2px #d0305040'
        }
    } else {
        style.background = '#ccc'
        if (focused) {
            style.boxShadow = '0 0 0 2px #2080f040'
        }
    }
    return style
}

const saveSetting = () => {
    reqPostSetting(configData.value).then((res) => {
        if (res.status == 200) {
            window.$message.success("保存成功", { keepAliveOnHover: true, duration: 5000 })
            loadSetting()
        }
    })
}

const  loadSetting = () => {
    reqGetSetting().then((res) => {
        if (res.status == 200 && res.data.code == 200) {
            configData.value = res.data.data
        } else {
            window.$message.error("Get setting error" + res.data.message, { keepAliveOnHover: true, duration: 5000 })
        }
    })
}

onMounted(() => {
    loadSetting()
})
</script>

<style scoped>
.setting-wrapper {
    position: relative;
    display: flex;
    flex-direction: column;
    height: 100%;
    justify-content: space-between;
    text-align: center;

    .n-form {
        text-align: start;
    }

    .n-button {
        position: sticky;
        align-self: end;
        width: 3.75rem;
        height: 3.75rem;
        bottom: 0;
    }

}</style>