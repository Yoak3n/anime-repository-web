<template>
    <div class="raw-wrapper">
        <n-list hoverable>
            <n-list-item>
                123
                <template #suffix>
                    <n-button>Suffix</n-button>
                </template>
            </n-list-item>
        </n-list>
    </div>
</template>

<script setup lang="ts">

import { NList, NListItem,NButton } from 'naive-ui';

</script>

<style scoped></style>