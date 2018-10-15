<template>
  <nav class="panel">
    <p class="panel-heading">
      feeds
    </p>
    <a
      v-for="f in feeds"
      :key="f.app_name"
      class="panel-block"
      :class="{ 'is-active': activeFeed === f.app_name }"
      @click="setActive(f)"
    >
      {{ f.app_name }}
      <span class="tag is-primary has-auto-margin-left">{{ f.logs.length }}</span>
    </a>
  </nav>
</template>

<script>
export default {
  name: 'feeds-nav',
  props: ['feeds'],
  data() {
    return {
      activeFeed: '',
    };
  },
  methods: {
    setActive(f) {
      this.activeFeed = f.app_name;

      this.$emit('select', f);
    },
  },
};
</script>

<style lang="scss">
  .has-auto-margin-left {
    margin-left: auto;
  }
</style>
