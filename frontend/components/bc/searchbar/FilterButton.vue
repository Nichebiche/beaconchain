<script setup lang="ts">
import {
  SearchbarStyle
} from '~/types/searchbar'

defineProps<{
  barStyle : SearchbarStyle,
  state? : boolean,
  look? : 'on'|'off' // forces the look of the button statically instead of having the color changing with its state
}>()

const emit = defineEmits<{(e: 'change', activated : boolean) : void}>()
</script>

<template>
  <label class="frame" :class="[barStyle, look]">
    <input
      type="checkbox"
      class="hidden-checkbox"
      :true-value="true"
      :false-value="false"
      :checked="state"
      :onchange="(e:any) => {emit('change', e.target.checked)}"
    >
    <slot />
  </label>
</template>

<style lang="scss" scoped>
@use "~/assets/css/fonts.scss";

.frame {
  display: inline-block;
  position: relative;
  box-sizing: border-box;
  cursor: pointer;
  user-select: none;
  border-radius: 10px;
  height: 20px;
  padding-top: 2.2px;
  padding-left: 8px;
  padding-right: 8px;
  text-align: center;
  transition: 0.2s;
  @include fonts.small_text_bold;
  white-space: nowrap;
  overflow: clip;
  &.gaudy,
  &.embedded {
    border: 1px solid var(--container-border-color);
  }

  .hidden-checkbox {
    display: none;
    width: 0;
    height: 0;
  }

  &:not(.off) {
    &.on,
    &:has(.hidden-checkbox:checked) {
      &.gaudy,
      &.embedded {
        border: 1px solid var(--button-color-active);
      }
      background-color: var(--button-color-active);
      &:hover {
        @media (hover: hover) {
          background-color: var(--button-color-hover);
        }
      }
      &:active {
        background-color: var(--button-color-pressed);
      }
    }
  }

  &:not(.on) {
    &.discreet {
      background-color: var(--light-grey);
    }
    @media (hover: hover) {
      &:hover {
        &.gaudy,
        &.embedded {
          background-color: var(--container-border-color);
        }
        &.discreet {
          background-color: var(--light-grey-3);
        }
      }
    }
    &:active {
      background-color: var(--button-color-pressed);
    }
  }

  &.gaudy,
  &.embedded {
    color: var(--text-color);
  }
  &.discreet {
    color: var(--light-black);
  }
}
</style>