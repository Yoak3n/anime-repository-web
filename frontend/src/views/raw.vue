<template>
    <div class="raw-wrapper">
        <div class="top-part">
            <n-input placeholder="按文件名搜索" v-model:value="input" size="large">
            <template #prefix>
                <n-button text >
                    <n-icon :component="Search">
                    </n-icon>
                </n-button>
            </template>
        </n-input>
        <n-data-table
        :data="rawTableData"
        :columns="columns"
        :pagination="{pageSize: 15}" 
        striped
        >

        </n-data-table>
        <n-modal
            v-model:show="showModal"
            transform-origin="mouse"
            style="max-width: 50%">
          <RuleOption :addRule="addRuleAction"></RuleOption>


        </n-modal>
        </div>
        
        <n-button circle @click="onReload" id="reload">
            <template #icon>
                <n-icon size="30">
                    <reload-outline />
                </n-icon>
            </template>
        </n-button>
    </div>
</template>

<script setup lang="ts">
import {onMounted, ref,reactive,h, nextTick} from 'vue'
import {NButton, NIcon, NInput, NDataTable, NModal,DataTableColumns} from 'naive-ui';
import RuleOption from '../components/common/RuleOption/index.vue'
import {ReloadOutline, Search} from '@vicons/ionicons5'
import {reqGetRaw} from "../api/raw";
import type {RawItem,RawTabelData} from "../api/raw/type";
import {reqAddRule} from "../api/rule";
import {ruleRequestData} from "../api/rule/type";
import {AxiosResponse} from "axios";

let input = ref('')
let showModal = ref(false)

let rawTableData = reactive<Array<RawTabelData>>([])

let data = reactive<RawItem[]>([{path:"1111",name:"adasda",full_path:"1111111111"}])

type RowData = {
  key:string,
  name:string
}
const matchRule =  (row:RowData)=> {
  console.log(row)
  showModal.value = true
}



const createColumns = ({
  matchRule
}:{matchRule:(row:RowData)=>void
}):DataTableColumns<RowData> =>{
  return [
    {
      title: '文件名',
      key: 'name',
      defaultSortOrder: 'ascend',
      sorter: 'default'
    },
    {
      title: '操作',
      key: 'action',
      render (row) {
        return h(
          NButton,
          {
            strong: true,
            tertiary: true,
            size: 'small',
            onClick: () => matchRule(row)
          },
          { default: () => '匹配规则' }
        )
      },
      width:'100'
    }
  ]
}

const columns = createColumns({matchRule})



const onReload = () => {
  reqGetRaw().then((res)=>{
    if(res.data.code!=200){
      window.$message.error('Get raw files error'+res.data.message,{ keepAliveOnHover: true, duration: 5000 })
    }else{
      window.$message.success('Get raw files successfully',{ keepAliveOnHover: true, duration: 5000 })
      data = res.data.data
      data.forEach((v)=>{
        // 数组越界
        let rawSingleData :RowData = {key:v.full_path,name:v.name}
        rawTableData.push(rawSingleData)
      })
    }
  })
}

const addRuleAction = (rule: ruleRequestData): Promise<AxiosResponse> => {
  return reqAddRule(rule)
}


onMounted(()=> {
  if (data != null){
    nextTick(()=>{
      data.forEach((v)=>{
        let rawSingleData :RowData = {key:v.full_path,name:v.name}
        rawTableData.push(rawSingleData)
    })
    })
  }
  reqGetRaw().then((res)=>{
    if (res.status==200 &&res.data.code==200){
      data = res.data.data
      data.forEach((v)=>{
        let rawSingleData :RowData = {key:v.full_path,name:v.name}
        rawTableData.push(rawSingleData)
      })
    }
  })
})

</script>

<style scoped lang="less">
.raw-wrapper {
    position: relative;
    display: flex;
    flex-direction: column;
    height: 100%;
    justify-content: space-between;
    .top-part{
        .raw-list{
            padding: 0 3rem 0 0;
        }
        
    }
    #reload {
        position:sticky;
        align-self: end;
        width: 3.75rem;
        height: 3.75rem;
        bottom: 0;
    }

}
</style>