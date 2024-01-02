<script setup lang="ts">
import {Component, h, nextTick, onMounted, provide, ref} from 'vue';
import {RouterLink, useRoute, useRouter} from 'vue-router';
import NaiveProvider from './components/common/NaiveProvider/index.vue'
import type {GlobalTheme, MenuOption} from 'naive-ui'
import {darkTheme, NButton, NConfigProvider, NIcon, NLayout, NLayoutContent, NLayoutSider, NMenu} from 'naive-ui';
import {BrowsersSharp, LogoRss, Settings} from '@vicons/ionicons5'
import {reqGetHost} from "./api/raw";

const $router = useRouter()
const $route = useRoute()
let activeKey = ref('')
let collapsed = ref(false)

const renderIcon = (icon: Component) => {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: '/rule',

        },
        { default: () => '订阅规则' }),
    key: 'rule', icon: renderIcon(LogoRss)
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: 'raw'
        },
        { default: () => '待处理文件' }
      ),
    key: 'raw', icon: renderIcon(BrowsersSharp)
  }
]

onMounted(() => {
  collapsed.value = true
  nextTick().then(() => {
    setTimeout(() => {
      activeKey.value = $route.name?.toString() as string
    }, 500)
  })
  reqGetHost().then((res)=>{
    if (res.data.data != undefined){
      document.title = res.data.data
    }else{
      document.title = "anime-repository-web"
    }
    
  })
})
const onSetting = () => {
  $router.push('/setting')
}

let theme = ref<GlobalTheme | null>(null)

const switchTheme = () => {
  if (theme.value == null) {
    theme.value = darkTheme
  } else {
    theme.value = null
  }
}

provide('switchTheme', switchTheme)
</script>

<template>
  <naive-provider>
    <n-config-provider :theme="theme">
      <div class="root">
        <n-layout has-sider>
          <n-layout-sider show-trigger="bar" collapse-mode="width" :collapsed-width="64" :collapsed="collapsed"
            @collapse="collapsed = true" @expand="collapsed = false" bordered>
            <n-menu v-model:value="activeKey" :options="menuOptions" :collapsed-width="64" :collapsed="collapsed"
              :collapsed-icon-size="24.5">
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
    </n-config-provider>

  </naive-provider>
</template>

<style scoped lang="less">
.n-layout {
  height: 100vh;

  .n-layout-sider {
    // background-color: #eee;

    .n-menu {
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
