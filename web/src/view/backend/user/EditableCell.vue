<template>
  <div @click="onFieldClick" class="edit-cell">
    <el-tooltip v-if="!editMode"
                :placement="toolTipPlacement"
                :open-delay="toolTipDelay"
                :content="toolTipContent">
      <div tabindex="0"
           class="cell-content"
           :class="{'edit-enabled-cell': canEdit}"
           @keyup.enter="onFieldClick">
        <slot name="content"></slot>
      </div>

    </el-tooltip>
    <component :is="editableComponent"
               v-if="editMode"
               ref="input"
               @focus="onFieldClick"
               @keyup.enter.native="onInputExit"
               v-on="listeners"
               v-bind="$attrs"
               v-model="model">
      <slot name="edit-component-slot"></slot>
    </component>
  </div>
</template>
<script>
export default {
  name: "editable-cell",
  inheritAttrs: false,
  props: {
    value: {
      type: String,
      default: ""
    },
    toolTipContent: {
      type: String,
      default: "单击修改"
    },
    toolTipDelay: {
      type: Number,
      default: 500
    },
    toolTipPlacement: {
      type: String,
      default: "None"
    },
    editableComponent: {
      type: String,
      default: "el-input"
    },
    closeEvent: {
      type: String,
      default: "blur"
    },
    canEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      editMode: false
    };
  },
  computed: {
    model: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit("input", val);
      }
    },
    listeners() {
      return {
        [this.closeEvent]: this.onInputExit,
        ...this.$listeners
      };
    }
  },
  methods: {
    onFieldClick() {
      if (this.canEdit) {
        this.editMode = true;
        this.$nextTick(() => {
          let inputRef = this.$refs.input;
          if (inputRef && inputRef.focus) {
            inputRef.focus();
          }
        });
      }
    },
    onInputExit() {
      this.editMode = false;
    },
    onInputChange(val) {
      this.$emit("input", val);
    }
  }
};
</script>
<style>
.cell-content {
  min-height: 40px;
  padding-left: 5px;
  padding-top: 5px;
}
.edit-cell {
  min-height: 35px;
  cursor: pointer;
}
</style>
