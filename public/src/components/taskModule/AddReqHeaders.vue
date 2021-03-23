<template>
  <div class="add-req-headers">
    <MapEdit v-model="modelValue.headers" />
  </div>
</template>

<script>
import MapEdit from "../base/MapEdit.vue";
export default {
  name: "AddReqHeaders",
  components: { MapEdit },
  props: {
    modelValue: {
      type: Object,
      default: () => {
        return {
          headers: {}
        };
      }
    }
  },
  data() {
    return {};
  }
};
</script>
