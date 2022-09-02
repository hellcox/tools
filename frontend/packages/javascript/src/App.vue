<template>
  <Header/>
  <div class="view">
<!--    <router-view />-->
    <router-view v-slot="{ Component }">
      <keep-alive>
        <component :is="Component"  v-if="$route.meta.keepAlive"/>
      </keep-alive>
      <component :is="Component"  v-if="!$route.meta.keepAlive"/>
    </router-view>
  </div>
</template>

<script>
import { useI18n } from "vue-i18n";
import Header from "@/components/public/Header.vue"

export default {
  components: {
    Header,
  },
  setup() {
    const { t, availableLocales, locale } = useI18n();

    // List of supported languages
    // 支持的语言列表
    const languages = availableLocales;

    // Click to switch language
    // 点击切换语言
    const onclickLanguageHandle = (item) => {
      item !== locale.value ? (locale.value = item) : false;
    };

    const onclickMinimise = () => {
      window.runtime.WindowMinimise();
    };
    const onclickQuit = () => {
      window.runtime.Quit();
    };

    return {
      t,
      languages,
      locale,
      onclickLanguageHandle,
      onclickMinimise,
      onclickQuit,
    };
  },
};
</script>

<style lang="scss">

html {
  width: 100%;
  height: 100%;
}
body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-size: var(--el-font-size-extra-small);
}

.view {
  position: absolute;
  top: 31px;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
  /*background: #dedede;*/
}
</style>
