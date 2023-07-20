<template>
    <div class="rules-wrapper">
        <n-data-table :columns="columns" :row-props="rowProps" :data="rules">

        </n-data-table>
        <n-dropdown placement="bottom-start" trigger="manual" :x="x" :y="y" :options="options" :show="showDropdown"
            :on-clickoutside="onClickoutside" @select="handleSelect" />
        <n-modal v-model:show="showModal">
            <rule-option></rule-option>
        </n-modal>
        
        <n-button circle @click="addRule" size="large">
            <template #icon>
                <n-icon size="40">
                    <add-outline></add-outline>
                </n-icon>
            </template>
        </n-button>
    </div>
</template>

<script setup lang="ts">
import { ref, h, nextTick,onMounted } from 'vue';
import { NDataTable, NDropdown,NButton,NIcon,NModal } from 'naive-ui';
import type { DataTableColumns, DropdownOption } from 'naive-ui'
import RuleOption from '../components/common/RuleOption/index.vue'
import {AddOutline } from '@vicons/ionicons5'
import type { rule } from '../api/rule/type'
import { reqGetRule, reqDeleteRule, reqAddRule } from '../api/rule/index';


let rules = ref<rule[]>([])
const showDropdown = ref(false)
const x = ref(0)
const y = ref(0)
let aim = ref(0)
type RuleColumn = {
    id: number,
    vid: string,
    name: string,
    season: number,
    language: string
}

const columns: DataTableColumns<RuleColumn> = [
    {
        title: 'VID',
        key: 'vid'
    },
    {
        title: '标题',
        key: 'name'
    },
    {
        title: '季',
        key: 'season'
    },
    {
        title: '语言',
        key: 'language'
    },
]
const options: DropdownOption[] = [
    {
        label: '修改',
        key: 'edit'
    },
    {
        label: () => h('span', { style: { color: 'red' } }, '删除'),
        key: 'delete'
    }
]

const rowProps = (row: RuleColumn) => {
    return {
        onContextmenu: (e: MouseEvent) => {
            aim.value = row.id
            e.preventDefault()
            showDropdown.value = false
            nextTick().then(() => {
                showDropdown.value = true
                x.value = e.clientX
                y.value = e.clientY
            })
        }
    }
}


rules.value = [
    {
        id: 11,
        vid: "1121",
        provider: "tmdb",
        file_extract_reg: "\\sasa",
        season: 11,
        language: "zh-CN",
        episode_extract_reg: "\\d",
        episode_position: 1,
        episode_offset: 0,
        name: "asaa",
        type: "tv"
    }, {
        id: 1123,
        vid: "11213-121",
        provider: "tmdb",
        file_extract_reg: "\\sasa",
        season: 11,
        language: "zh-CN",
        episode_extract_reg: "\\d",
        episode_position: 1,
        episode_offset: 0,
        name: "asdasda",
        type: "tv",
    },

]

const handleSelect = (key: string | number) => {
    switch (key) {
        case 'edit':
            console.log('修改规则',aim.value);
            modifyRule()
            break;
        case 'delete':
            deleteRule()
            break;
        default:
            break

    }

    showDropdown.value = false
}
const onClickoutside = () => {
    showDropdown.value = false
}

const addRule = ()=>{
    showModal.value = true
    // reqAddRule({vid:"123",provider:'',file_extract_reg:'',season:0,language:'',episode_extract_reg:"",episode_offset:0,episode_position:1,type:'tv',name:''},)
}

// const getRule = ()=>{
//     reqGetRule().then((res)=>{
//       if (res.data.code == 200){
//         rules.value = res.data.data
//         window.$message.success("get rules successfully",{keepAliveOnHover:true,duration: 5000})
//       }else{
//         window.$message.error("get rules error"+res.data.message,{keepAliveOnHover:true,duration: 5000});
//       }
//     }) 
// }

const modifyRule = ()=>{

    reqAddRule({vid:"123",provider:'',file_extract_reg:'',season:0,episode_extract_reg:"",episode_offset:0,episode_position:1,type:'tv',name:''})
}

const deleteRule = ()=>{
    reqDeleteRule(aim.value).then((res)=>{
        if (res.data.code == 200){
            window.$message.success("delete rule success",{keepAliveOnHover: true,duration: 5000})
        }else{
            window.$message.error("delete rule error"+res.data.message,{keepAliveOnHover: true,duration: 5000})
        }
    })
}
let showModal = ref(false)
onMounted(()=>{
    reqGetRule().then((res)=>{
      if (res.data.code == 200){
        rules.value = res.data.data
      }else{
        window.$message.error("get rules error"+res.data.message,{keepAliveOnHover:true,duration: 5000});
      }
    })
})
</script>

<style scoped lang="less">
.rules-wrapper{
    position: relative;
    display: flex;
    flex-direction: column;
    height: 100%;
    overflow:visible;
    justify-content: space-between;
    .n-button{
        align-self: end;
    }
}
</style>