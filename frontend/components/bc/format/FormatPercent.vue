<script setup lang="ts">
interface Props {
  percent?: number
  base?: number
  value?: number
  hideEmptyValue?: boolean
  precision?: number
  fixed?: number
  fullOnEmptyBase?: boolean
  addPositiveSign?: boolean
  colorBreakPoint?: number // if set then the percentage will be colored accordingly
}

const props = defineProps<Props>()

const data = computed(() => {
  let label: string | null = null
  let className = ''
  if (props.base === 0 && props.fullOnEmptyBase) {
    return {
      label: '100%',
      className: 'text-positive'
    }
  }
  if (props.percent === undefined && !props.base) {
    if (!props.hideEmptyValue) {
      label = '0%'
    }
    return { label, className }
  }
  const percent = props.percent ?? calculatePercent(props.value, props.base)
  const config = { precision: props.precision ?? 2, fixed: props.fixed ?? 2, addPositiveSign: props.addPositiveSign }
  label = formatPercent(percent, config)

  if (props.colorBreakPoint) {
    if ((props.base === 0 && percent === 0) || percent >= props.colorBreakPoint) {
      className = 'text-positive'
    } else {
      className = 'text-negative'
    }
  }
  return { label, className }
})

</script>
<template>
  <BcFormatNumber v-if="data.label" :text="data.label" :class="data.className" />
</template>
