<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faInfoCircle
} from '@fortawesome/pro-regular-svg-icons'
import {
  faCube,
  faSync
} from '@fortawesome/pro-solid-svg-icons'

import { union } from 'lodash-es'
import {
  SummaryDetailsEfficiencyProps,
  SummaryDetailsEfficiencyValidatorProps,
  type SummaryDetail,
  type SummaryDetailsEfficiencyValidatorProp,
  type SummaryDetailsEfficiencyProp,
  type SummaryDetailsEfficiencyCombinedProp,
  type DashboardValidatorContext
} from '~/types/dashboard/summary'
import type { VDBGroupSummaryData, VDBSummaryTableRow } from '~/types/api/validator_dashboard'

interface Props {
  property: SummaryDetailsEfficiencyCombinedProp,
  detail: SummaryDetail,
  data: VDBGroupSummaryData,
  row: VDBSummaryTableRow
}
const props = defineProps<Props>()

const { tm: $tm } = useI18n()

const data = computed(() => {
  const col = props.data?.[props.detail]
  if (!col) {
    return null
  }
  if (props.property === 'attestation_total') {
    return {
      efficiency: {
        status_count: {
          success: col.attestation_head.status_count.success + col.attestation_source.status_count.success + col.attestation_target.status_count.success,
          failed: col.attestation_head.status_count.failed + col.attestation_source.status_count.failed + col.attestation_target.status_count.failed
        }
      }
    }
  } else if (SummaryDetailsEfficiencyProps.includes(props.property as SummaryDetailsEfficiencyProp)) {
    const tooltip: { title: string, text: string } | undefined = $tm(`dashboard.validator.tooltip.${props.property}`)
    return {
      efficiency: col[props.property as SummaryDetailsEfficiencyProp],
      tooltip
    }
  } else if (SummaryDetailsEfficiencyValidatorProps.includes(props.property as SummaryDetailsEfficiencyValidatorProp)) {
    let validators: number[] = []
    let context: DashboardValidatorContext = 'attestation'
    if (props.property === 'validators_attestation') {
      validators = union(col.attestation_head.validators, col.attestation_source.validators, col.attestation_target.validators)
    } else if (props.property === 'validators_proposal') {
      validators = col.proposals.validators ?? []
      context = 'proposal'
    } else if (props.property === 'validators_sync') {
      validators = col.sync.validators ?? []
      context = 'sync'
    } else if (props.property === 'validators_slashings') {
      validators = col.slashed.validators ?? []
      context = 'slashings'
    }
    return {
      validators,
      context
    }
  } else if (props.property === 'attestation_efficiency') {
    const tooltip: { title: string, text: string } | undefined = $tm('dashboard.validator.tooltip.attestation_efficiency')
    return {
      attestationEfficiency: col.attestation_efficiency,
      tooltip
    }
  } else if (props.property === 'apr') {
    return {
      apr: {
        apr: col.apr,
        total: col.apr.cl + col.apr.el,
        income: col.income
      }
    }
  } else if (props.property === 'luck') {
    return {
      luck: col.luck
    }
  } else if (props.property === 'efficiency_total') {
    let efficiencyTotal = props.row.efficiency_day
    switch (props.detail) {
      case 'details_month':
        efficiencyTotal = props.row.efficiency_month
        break
      case 'details_week':
        efficiencyTotal = props.row.efficiency_week
        break
      case 'details_total':
        efficiencyTotal = props.row.efficiency_total
        break
    }
    return {
      efficiencyTotal: {
        total: efficiencyTotal
      }
    }
  } else if (props.property === 'attestation_avg_incl_dist') {
    return {
      simple: {
        value: col.attestation_avg_incl_dist
      }
    }
  }
})

</script>
<template>
  <BcTooltip v-if="data?.efficiency" position="top" :text="data.tooltip?.text" :title="data.tooltip?.title">
    <div class="info_row">
      <DashboardTableEfficiency :success="data.efficiency.status_count.success" :failed="data.efficiency.status_count.failed" />
      <FontAwesomeIcon v-if="data.tooltip?.title" class="link" :icon="faInfoCircle" />
    </div>
  </BcTooltip>
  <DashboardTableValidators
    v-else-if="data?.validators"
    :validators="data.validators"
    :time-frame="props.detail"
    :context="data.context"
    :group-id="props.row.group_id"
  />
  <BcTooltip
    v-else-if="data?.attestationEfficiency"
    position="top"
    :text="data.tooltip?.text"
    :title="data.tooltip?.title"
  >
    <div class="info_row">
      <BcFormatPercent :percent="data?.attestationEfficiency" :color-break-point="80" />
      <FontAwesomeIcon class="link" :icon="faInfoCircle" />
    </div>
  </BcTooltip>
  <BcTooltip v-else-if="data?.apr" position="top">
    <div class="info_row">
      <BcFormatPercent :percent="data.apr.total" />
      <FontAwesomeIcon class="link" :icon="faInfoCircle" />
    </div>
    <template #tooltip>
      <div class="row">
        <b>{{ $t('common.execution_layer') }}:</b>
        <BcFormatValue class="space_before" :value="data.apr.income.el" /> (
        <BcFormatPercent :percent="data.apr.apr.el" />)
      </div>
      <div class="row">
        <b>{{ $t('common.consensus_layer') }}:</b>
        <BcFormatValue class="space_before" :value="data.apr.income.cl" /> (
        <BcFormatPercent :percent="data.apr.apr.cl" />)
      </div>
    </template>
  </BcTooltip>
  <BcTooltip v-else-if="data?.luck" position="top">
    <div class="info_row">
      <span>
        <FontAwesomeIcon :icon="faCube" />
        <BcFormatPercent class="space_before" :percent="data.luck.proposal.percent" /> / <FontAwesomeIcon :icon="faSync" />
        <BcFormatPercent class="space_before" :percent="data.luck.sync.percent" />
      </span>
      <FontAwesomeIcon class="link" :icon="faInfoCircle" />
    </div>
    <template #tooltip>
      <div class="row">
        <b>
          {{ $t('dashboard.validator.tooltip.block_proposal') }}
        </b>
      </div>
      <div class="row">
        <b>
          {{ $t('common.luck') }}:
        </b>
        <BcFormatPercent :percent="data.luck.proposal.percent" />
      </div>
      <div class="row">
        <b>
          {{ $t('common.expected') }}:
        </b>
        {{ $t('common.in_day', {}, data.luck.proposal.expected) }}
      </div>
      <div class="row">
        <b>
          {{ $t('common.average') }}:
        </b>
        {{ $t('common.every_day', {}, data.luck.proposal.average) }}
      </div>
      <br>
      <div class="row next_chapter">
        <b class="part">
          {{ $t('dashboard.validator.tooltip.sync_committee') }}
        </b>
      </div>
      <div class="row">
        <b>
          {{ $t('common.luck') }}:
        </b>
        <BcFormatPercent :percent="data.luck.proposal.percent" />
      </div>
      <div class="row">
        <b>
          {{ $t('common.expected') }}:
        </b>
        {{ $t('common.in_day', {}, data.luck.proposal.expected) }}
      </div>
      <div class="row">
        <b>
          {{ $t('common.average') }}:
        </b>
        {{ $t('common.every_day', {}, data.luck.proposal.average) }}
      </div>
    </template>
  </BcTooltip>
  <BcFormatPercent v-else-if="data?.efficiencyTotal" :percent="data.efficiencyTotal.total" :color-break-point="80" />
  <span v-else-if="data?.simple">
    {{ data.simple?.value }}
  </span>
</template>

<style>
.row {
  text-wrap: nowrap;
  min-width: 100%;
  text-align: left;

  &.next_chapter {
    margin-top: var(--padding);
  }
}

.space_before {
  &::before {
    content: " ";
  }
}

.info_row {
  display: flex;
  justify-content: space-between;
}
</style>