<template>
    <div class="about-wrapper">
        <div class="about-box">
            <img src="https://avatars.githubusercontent.com/u/120039624?s=400&u=c555f8bf70dc1056b867b6c6127cc49a7b5a29f2&v=4" alt="Yoak3n">
            <div class="description">
                <div class="description-box">
                    Yoak3n
                    <n-timeline :icon-size="15" size="medium" horizontal >
                    <n-timeline-item  color="red">
                        <template #icon>
                            <n-icon>
                                <AccessibilityOutline/>
                            </n-icon>
                        </template>
                        <span>不知名小卒</span>
                        
                    </n-timeline-item>
                    <n-timeline-item  type="success">
                        <template #icon>
                            <n-icon>
                                <HomeOutline/>
                            </n-icon>
                        </template>
                        <a href="https://github.com/Yoak3n" target="_blank" rel="noopener noreferrer">Github</a>
                    </n-timeline-item>
                </n-timeline>
                </div>
                
            </div>
        </div>
        
    </div>
</template>

<script setup lang="ts">
import {NTimeline,NTimelineItem,NIcon} from 'naive-ui'
import {AccessibilityOutline,HomeOutline} from '@vicons/ionicons5'
</script>

<style scoped lang="less">
.about-wrapper{
    display: flex;
    .about-box{
        display: flex;
        width: 100%;
        margin: 0 1rem;
        img{
            height: 6rem;
            border-radius: 5px;
        }
        .description{
            width: 100%;
            font-size: large;
            .description-box{
                width: 100%;
                .n-timeline{
                    margin: 0 2rem;
                    color: #bbb;
                    span{
                        color: #bbb;
                    }
                    a{
                        text-decoration: none  ;
                        color: #bbb;
                        &:hover{
                            text-decoration-line: underline;
                        }
                    }
                }
            }
        }
    }
}

</style>