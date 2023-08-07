<template>
    <n-card>
        <n-form v-model="model">
            <n-space justify="space-between">
                <n-form-item label="VID">
                    <n-input placeholder="请输入数据库的视频id" style="width: 12rem;" v-model:value="model.vid">
                        <template #suffix>
                            <n-button text @click="GetTvName">
                                <n-icon :component="Search">
                                </n-icon>
                            </n-button>
                        </template>
                    </n-input>
                </n-form-item>
                <n-form-item label="视频名">
                    <n-input placeholder="请输入视频名" v-model:value="model.name" style="width: 12rem;"></n-input>
                </n-form-item>
                <n-form-item label="提供者">
                    <n-select v-model:value="model.provider" :options="providerOptions" style="width: 12rem;"></n-select>
                </n-form-item>
            </n-space>
            <n-space justify="space-between" >
                <n-form-item label="从文件名提取标题" style="width: 20rem;">
                    <n-input v-model:value="model.file_extract_reg" placeholder="请输入待处理文件名中可供本规则检索的部分"></n-input>
                </n-form-item>
                <n-form-item label="类型" style="width: 20rem;">
                    <n-select v-model:value="model.type" :options="typeOptions"></n-select>
                </n-form-item>
            </n-space>

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
                    <n-button type="error" ghost size="large" style="width: 10rem;" @click="resetRule">重置</n-button>
                </n-form-item>
                <n-form-item :show-feedback="false" :show-label="false">
                    <n-button type="success" ghost size="large" style="width: 10rem;" @click="RuleAction"
                        :loading="loading">提交</n-button>
                </n-form-item>
            </n-space>


        </n-form>
    </n-card>
</template>

<script setup lang="ts">
import { ref, toRefs, onMounted, onUpdated } from 'vue';
import { NSpace, NForm, NFormItem, NCard, NInput, NInputNumber, NSelect, NButton, NIcon } from 'naive-ui';
import { Search } from '@vicons/ionicons5'
import { ruleRequestData, RuleResponse } from '../../../api/rule/type';
import { AxiosResponse } from 'axios';
import type{RawNameResponse} from "../../../api/raw/type";
import {reqGetTVName} from "../../../api/raw";

const defaultRule:ruleRequestData =  { vid: "", provider: "TMDB", file_extract_reg: "", episode_extract_reg: "\\d+", episode_offset: 0, episode_position: 1, season: 1, type: "tv", name: "", language: 'zh-CN' }

let loading = ref(false)
let props = defineProps(['addRule','data','modify'])
let {data} = toRefs(props)

let model = ref<ruleRequestData>(defaultRule)

const GetTvName = () => {
  reqGetTVName(model.value.vid).then((res:AxiosResponse<RawNameResponse>)=>{
    if (res.data.code == 200){
      model.value.name = res.data.data
    }else {
      window.$message.error(`Get TV name error:${res.data.message}`,{ keepAliveOnHover: true, duration: 5000 })
    }
  })
}

const RuleAction = () => {
    loading.value = true
    let value: ruleRequestData = model.value
    let promise: Promise<AxiosResponse<RuleResponse, any>> 
    if (data?.value != undefined || data?.value != null ){
        promise = props.modify(value)
    }else{
        promise = props.addRule(value)
    }
    
    promise.then((res) => {
        if (res.data.code == 200 && res.data.data != null) {
            loading.value = false
        } else {
            loading.value = false
        }
    })
}

const resetRule = ()=>{
    model.value = defaultRule
}

onMounted(()=>{
    if (data?.value != undefined || data?.value != null ){
        model.value = data.value
    }
})
onUpdated(()=>{
    if (data?.value != undefined || data?.value != null ){
        model.value = data.value
    }
})



// select options
let typeOptions = [{label:"tv",value:'tv'},{label:'movie',value:'movie'},{label:'anime',key:'tv'}]
let providerOptions = [{ label: "TMDB", value: "TMDB" }]
let languageOptions = [{ label: "中文", value: "zh-CN" }]
</script>

<style scoped lang="less">
.n-card {
    max-width: 40%;
}
</style>