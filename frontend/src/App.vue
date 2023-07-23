<script setup lang="ts">
import { Component, h, ref, onMounted } from 'vue';
import { RouterLink } from 'vue-router';
import NaiveProvider from './components/common/NaiveProvider/index.vue'
import { NLayout, NLayoutSider, NLayoutContent, NMenu, NIcon, NButton } from 'naive-ui';
import type { MenuOption } from 'naive-ui'
import { LogoRss, BrowsersSharp, Settings } from '@vicons/ionicons5'



let activeKey = ref('')
let collapsed = ref(false)

const renderIcon = (icon: Component) => {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions: MenuOption[] = [
  { label: () =>
      h(
          RouterLink,
          {
              to: '/rule',

          },
          { default: () => '订阅规则' }),
    key: 'rule', icon: renderIcon(LogoRss) },
  { label: ()=>
      h(
        RouterLink,
        {
          to:'raw'
        },
        {default:()=>'待处理文件'}
    ),
  key: 'raw', icon: renderIcon(BrowsersSharp) }
]

onMounted(()=>{
  collapsed.value = true
})
const onSetting = () =>{
  console.log(123);

}
</script>

<template>
  <naive-provider>
    <div class="root">
    <n-layout has-sider>
      <n-layout-sider 
        show-trigger="bar" 
        collapse-mode="width" 
        :collapsed-width="64" 
        :collapsed="collapsed"
        @collapse="collapsed = true"
        @expand="collapsed = false"
        bordered>
        <n-menu 
        v-model:value="activeKey"
        :options="menuOptions"
        :collapsed-width="64" 
        :collapsed="collapsed"
        :collapsed-icon-size="24.5"
        >
        </n-menu>
        <n-button text size="large" class="setting" @click="onSetting"> 
          <template #icon>
            <n-icon size="24">
              <settings />
            </n-icon>
          </template>
        </n-button>

      </n-layout-sider>
      <n-layout>
        <n-layout-content content-style="padding: 24px;">
          <router-view></router-view>
        </n-layout-content>
      </n-layout>
    </n-layout>
  </div>
  </naive-provider>
  
</template>

<style scoped lang="less">
.n-layout {
  height: 100vh;

  .n-layout-sider {
    background-color: #eee;
    .n-menu{
      margin-top: 1.2rem;
    }
    
    .n-button {
      width: 100%;
      height: 3rem;
      margin: 0 auto;
    }
    .setting {
    text-align: center;
    position: absolute;
    bottom: 1rem;
  }

  }
  
}
</style>
