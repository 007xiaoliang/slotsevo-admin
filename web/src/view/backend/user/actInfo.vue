<template>
  <div>
    <h1>{{ this.rpcActivityInfo }}</h1>
  </div>
</template>


<script>
import {mapGetters} from "vuex";

export default {
  name: "ActInfo",
  computed: {
    ...mapGetters('rpcUser', ['rpcActivityInfo']),
  }
};
</script>
