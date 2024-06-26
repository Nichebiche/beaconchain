<script setup lang="ts">
import type { IconDefinition } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import BcTooltip from '../BcTooltip.vue'

interface Props {
  icon?: IconDefinition,
  falseIcon?: IconDefinition,
  tooltip?: string,
}

const props = defineProps<Props>()

const selected = defineModel<boolean | undefined>({ required: true })

const icon = computed(() => {
  return selected.value || !props.falseIcon ? props.icon : props.falseIcon
})
</script>

<template>
  <BcTooltip :dont-open-permanently="true" :hover-delay="350">
    <template #tooltip>
      <div class="button-tooltip">
        <div v-if="tooltip" class="individual">
          {{ tooltip }}
        </div>
        <div>{{ selected ? $t('filter.enabled'): $t('filter.disabled') }}</div>
      </div>
    </template>
    <ToggleButton v-model="selected" class="bc-toggle" on-label="''" off-icon="''">
      <template #icon="slotProps">
        <slot name="icon" v-bind="slotProps">
          <FontAwesomeIcon v-if="icon" :icon="icon" />
        </slot>
      </template>
    </ToggleButton>
  </BcTooltip>
</template>

<style lang="scss" scoped>
.button-tooltip{
  width: max-content;
  text-align: left;
  .individual{
    margin-bottom: var(--padding);
  }
}
.bc-toggle {
  &.p-button {
    &.p-togglebutton {
      width: 30px;
      height: 30px;
      padding: 2px;
      border-style: none;
      color: var(--container-color);
      background-color: var(--container-border-color);

      &:not(.p-highlight) {
        background-color: var(--container-background);
      }

      // this is needed as the primvevue ToggleButton adds a yes/no label if none is provided
      :deep(.p-button-label) {
        display: none;
      }
    }
  }
}
</style>
