<template>
    <n-card>
        <n-form v-model="model">
            <n-space justify="space-between">
                <n-form-item label="VID">
                    <n-input placeholder="请输入数据来源的视频id" style="width: 20rem;">
                        <template #suffix>
                            <n-button text>
                                <n-icon :component="Search">
                                </n-icon>
                            </n-button>
                        </template>
                    </n-input>
                </n-form-item>
                <n-form-item label="提供者">
                    <n-select v-model:value="model.provider" :options="providerOptions" style="width: 20rem;"></n-select>
                </n-form-item>
            </n-space>
            <n-form-item label="从文件名提取标题">
                <n-input v-model:value="model.file_extract_reg" placeholder="请输入待处理文件名中可供本规则检索的部分"></n-input>
            </n-form-item>
            <n-space justify="space-between">
                <n-form-item label="季">
                    <n-input-number v-model:value="model.season" style="width: 20rem;"></n-input-number>
                </n-form-item>
                <n-form-item label="语言">
                    <n-select v-model:value="model.language" :options="languageOptions" style="width: 20rem;"></n-select>
                </n-form-item>
            </n-space>

            <n-space justify="space-between">
                <n-form-item label="从文件名正则提取集数">
                    <n-input v-model:value="model.episode_extract_reg" style="width: 12rem;"></n-input>
                </n-form-item>
                <n-form-item label="集数位置">
                    <n-input-number v-model:value="model.episode_position" style="width: 12rem;"></n-input-number>
                </n-form-item>
                <n-form-item label="集数偏移">
                    <n-input-number v-model:value="model.episode_offset" style="width: 12rem;"></n-input-number>
                </n-form-item>
            </n-space>

            <n-space justify="space-evenly">
                <n-form-item :show-feedback="false" :show-label="false">
                    <n-button type="error" ghost size="large" style="width: 10rem;">重置</n-button>
                </n-form-item>
                <n-form-item :show-feedback="false" :show-label="false">
                    <n-button type="success" ghost size="large" style="width: 10rem;">提交</n-button>
                </n-form-item>
            </n-space>


        </n-form>
    </n-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { NSpace, NForm, NFormItem, NCard, NInput, NInputNumber, NSelect, NButton, NIcon } from 'naive-ui';
import { Search } from '@vicons/ionicons5'
import { ruleRequestData } from '../../../api/rule/type';


let model = ref<ruleRequestData>({ vid: "", provider: "TMDB", file_extract_reg: "", episode_extract_reg: "\\d+", episode_offset: 0, episode_position: 1, season: 1, type: "tmdb", name: "", language: 'zh-CN' })


let providerOptions = [{ label: "TMDB", value: "TMDB" }]
let languageOptions = [{ label: "中文", value: "zh-CN" }]
</script>

<style scoped lang="less">
.n-card {
    max-width: 40%;
}
</style>