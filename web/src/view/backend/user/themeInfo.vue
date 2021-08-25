<template>
  <div>
    <h1>{{ this.rpcThemeInfo }}</h1>
  </div>
</template>


<script>
import {mapGetters} from "vuex";

export default {
  name: "ThemeInfo",
  computed: {
    ...mapGetters('rpcUser', ['rpcThemeInfo']),
  }
}
</script>
