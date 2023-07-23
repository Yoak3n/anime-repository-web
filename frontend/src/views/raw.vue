<template>
    <div class="raw-wrapper">
        <div class="top-part">
            <n-input placeholder="按文件名搜索" v-model:value="input" size="large">
            <template #prefix>
                <n-button text>
                    <n-icon :component="Search">
                    </n-icon>
                </n-button>
            </template>
        </n-input>
        <n-list hoverable>
            <template #header>
                文件名
            </template>
            <n-list-item v-for="item in data" :key="item.full_path">
                {{ item.name }}
                <template #suffix>
                    <n-button>匹配规则</n-button>
                </template>
            </n-list-item>
        </n-list>
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
import {onMounted, ref} from 'vue'
import { NList, NListItem, NButton, NInput, NIcon } from 'naive-ui';
import { Search, ReloadOutline } from '@vicons/ionicons5'
import {reqGetRaw} from "../api/raw";
import type {RawItem} from "../api/raw/type";

let input = ref('')

// test data
let data = ref<RawItem[]>([
  {name:'asdas',full_path:'sadasd',path:'assaasaa'},
  {name:'asda1',full_path:'sadasd',path:'assaasaa'},
  {name:'asda4s',full_path:'sadasd',path:'assaasaa'},
  {name:'asd5as',full_path:'sadasd',path:'assaasaa'},
  {name:'asd6as',full_path:'sadasd',path:'assaasaa'},
  {name:'as8das',full_path:'sadasd',path:'assaasaa'},
  {name:'asd6as',full_path:'sadasd',path:'assaasaa'},
  {name:'asd2as',full_path:'sadasd',path:'assaasaa'},
  {name:'as7das',full_path:'sadasd',path:'assaasaa'},
  {name:'asd9as',full_path:'sadasd',path:'assaasaa'},
])

const onReload = () => {
  reqGetRaw().then((res)=>{
    if(res.data.code!=200){
      window.$message.error('Get raw files error'+res.data.message,{ keepAliveOnHover: true, duration: 5000 })
    }else{
      window.$message.success('Get raw files successfully',{ keepAliveOnHover: true, duration: 5000 })
      data.value = res.data.data
    }
  })
}

onMounted(()=> {
  reqGetRaw().then((res)=>{
    if (res.status==200 &&res.data.code==200){
      data.value = res.data.data
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
        .n-list{
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