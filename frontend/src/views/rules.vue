<template>
    <div class="rules-wrapper">
        <n-data-table :columns="columns" :row-props="rowProps" :data="rules">

        </n-data-table>
        <n-dropdown placement="bottom-start" trigger="manual" :x="x" :y="y" :options="options" :show="showDropdown"
            :on-clickoutside="onClickOutside" @select="handleSelect" />
        <n-modal v-model:show="showModal">
            <rule-option :addRule="addRule" :data="aimData" :modify="modifyRule"></rule-option>
        </n-modal>

        <n-button class="n-button" circle @click="onAddClicked" dashed>
            <template #icon>
                <n-icon size="40">
                    <add-outline></add-outline>
                </n-icon>
            </template>
        </n-button>
    </div>
</template>

<script setup lang="ts">
import { ref, h, nextTick, onMounted } from 'vue';
import { NDataTable, NDropdown, NButton, NIcon, NModal } from 'naive-ui';
import type { DataTableColumns, DropdownOption } from 'naive-ui'
import RuleOption from '../components/common/RuleOption/index.vue'
import { AddOutline } from '@vicons/ionicons5'
import type { rule } from '../api/rule/type'
import { reqGetRule, reqDeleteRule, reqAddRule } from '../api/rule';
import { ruleRequestData, } from '../api/rule/type';
import { AxiosResponse } from 'axios';
import { reqModifyRule } from '../api/rule';


let rules = ref<rule[]>([])
const showDropdown = ref(false)
const x = ref(0)
const y = ref(0)
let aim = ref()
let aimData = ref()
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


const handleSelect = (key: string | number) => {
    switch (key) {
        case 'edit':
            aimData.value = rules.value.find(element => element.id == aim.value)
            console.log(aimData.value);
            showModal.value = true
            break;
        case 'delete':
            deleteRule()
            break;
        default:
            break

    }

    showDropdown.value = false
}
const onClickOutside = () => {
    showDropdown.value = false
}

const onAddClicked = () => {
    aim.value = null
    showModal.value = true
    // reqAddRule({vid:"123",provider:'',file_extract_reg:'',season:0,language:'',episode_extract_reg:"",episode_offset:0,episode_position:1,type:'tv',name:''},)
}

const addRule = (rule: ruleRequestData): Promise<AxiosResponse> => {
    let promise = reqAddRule(rule)
    setTimeout(getRule, 1000)
    showModal.value = false
    return promise
}
const getRule = () => {
    reqGetRule().then((res) => {
        if (res.data.code == 200) {
            rules.value = res.data.data
            window.$message.success("Get rules successfully", { keepAliveOnHover: true, duration: 5000 })
        } else {
            window.$message.error("Get rules error : " + res.data.message, { keepAliveOnHover: true, duration: 5000 });
        }
    })
}

const modifyRule = (data: ruleRequestData) => {
    let promise = reqModifyRule(aim.value, data).then()
    setTimeout(getRule, 500)

    return promise
}

const deleteRule = () => {
    reqDeleteRule(aim.value).then((res) => {
        if (res.data.code == 200) {
            window.$message.success("Delete rule success", { keepAliveOnHover: true, duration: 5000 })
        } else {
            window.$message.error("Delete rule error" + res.data.message, { keepAliveOnHover: true, duration: 5000 })
        }
    })
}
let showModal = ref(false)
onMounted(() => {
    reqGetRule().then((res) => {
        if (res.data.code == 200) {
            rules.value = res.data.data
        } else {
            window.$message.error("get rules error: " + res.data.message, { keepAliveOnHover: true, duration: 5000 });
        }
    })
})
</script>

<style scoped lang="less">
.rules-wrapper {
    position: relative;
    display: flex;
    flex-direction: column;
    height: 100%;
    overflow: visible;
    justify-content: space-between;

    .n-button {
        align-self: end;
        width: 3.75rem;
        height: 3.75rem;
    }
}
</style>