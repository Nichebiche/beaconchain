import type { TimeFrame } from '@/types/value'

export const SummaryDetailsEfficiencyProps = ['attestations_head', 'attestations_source', 'attestations_target', 'sync', 'proposals', 'slashed'] as const
export type SummaryDetailsEfficiencyProp = typeof SummaryDetailsEfficiencyProps[number]

export const SummaryDetailsEfficiencyValidatorProps = ['validators_sync', 'validators_proposal', 'validators_slashings', 'validators_attestation'] as const
export type SummaryDetailsEfficiencyValidatorProp = typeof SummaryDetailsEfficiencyValidatorProps[number]

export const SummaryDetailsEfficiencyLuckProps = ['proposal_luck', 'sync_luck'] as const
export type SummaryDetailsEfficiencyLuckProp = typeof SummaryDetailsEfficiencyLuckProps[number]

export const SummaryDetailsEfficiencyCustomProps = ['attestation_total'] as const
export type SummaryDetailsEfficiencyCustomProp = typeof SummaryDetailsEfficiencyCustomProps[number]

export const SummaryDetailsEfficiencySpecialProps = ['efficiency_all_time', 'apr', 'luck', 'attestation_avg_incl_dist', 'attestation_efficiency'] as const
export type SummaryDetailsEfficiencySpecialProp = typeof SummaryDetailsEfficiencySpecialProps[number]

export type SummaryDetailsEfficiencyCombinedProp = SummaryDetailsEfficiencySpecialProp | SummaryDetailsEfficiencyProp | SummaryDetailsEfficiencyCustomProp | SummaryDetailsEfficiencyLuckProp | SummaryDetailsEfficiencyValidatorProp

export type DashboardValidatorContext = 'dashboard' | 'group' | 'attestation' | 'sync' | 'slashings' | 'proposal'

export type SummaryRow = { details: TimeFrame[], prop: SummaryDetailsEfficiencyCombinedProp, title: string}
