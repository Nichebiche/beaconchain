<script setup lang="ts">
import type { IconDefinition } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
interface Props {
  icon?: IconDefinition,
  text?: string,
  subText?: string,
  selected: boolean,
  disabled?:boolean,
}
const props = defineProps<Props>()

const topBottomPadding = computed(() => props.subText ? '8px' : '16px')
</script>

<template>
  <ToggleButton class="bc-toggle" :disabled="disabled" :model-value="selected">
    <template #icon="slotProps">
      <slot name="icon" v-bind="slotProps">
        <FontAwesomeIcon v-if="icon" :icon="icon" />
      </slot>
      <div class="label">
        {{ text }}
        <div v-if="subText" class="sub">
          {{ subText }}
        </div>
      </div>
    </template>
  </ToggleButton>
</template>

<style lang="scss" scoped>
@use '~/assets/css/fonts.scss';

.bc-toggle {
  &.p-button {
    &.p-togglebutton {
      display: flex;
      flex-grow: 1;
      flex-direction: column;
      gap: 11px;

      width: 100%;
      height: 100%;
      padding: v-bind(topBottomPadding) 0;
      border: 1px var(--container-border-color) solid;
      border-radius: var(--border-radius);
      background-color: var(--container-background);
      color: var(--text-color);

      &.p-highlight {
        border-color: var(--button-color-active);
        color: var(--button-color-active);
      }

      :deep(.p-button-label) {
        display: none;
      }

      :deep(svg) {
        max-width: 36px;
      }
      &.p-disabled{
        opacity: 0.5;
        cursor: default;
      }
    }
  }

  .label {
    @include fonts.subtitle_text;
    .sub {
      font-size: var(--tiny_text_font_size);
    }
  }
}
</style>
