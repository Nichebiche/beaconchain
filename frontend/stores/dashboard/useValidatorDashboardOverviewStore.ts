import { defineStore } from 'pinia'
import type { VDBOverviewData, InternalGetValidatorDashboardResponse } from '~/types/api/validator_dashboard'
import type { DashboardKey } from '~/types/dashboard'

// NOTE: The store itself is "as stupid as possible", it only holds data to share between server, client and all components
const validatorOverviewStore = defineStore('validator_overview_store', () => {
  const data = ref<VDBOverviewData | undefined | null>()
  return { data }
})

// NOTE: this is a wrapper around the "stupid" store to
// - make store data from the outside readonly
// - provide a refresh function to update the data via an API call
// calling useValidatorDashboardOverviewStore only works if called within a components script setup
//  but that's good enough since the data returned below is reactive
export function useValidatorDashboardOverviewStore () {
  const { fetch } = useCustomFetch()
  const { data } = storeToRefs(validatorOverviewStore()) // NOTE: this is the REACTIVE (and read only) data

  const overview = computed(() => data.value) // NOTE: we use computed here to have a readonly access to the data only from the outside

  // NOTE: function to UPDATE the data, use overview if you just want to access the data
  async function refreshOverview (key: DashboardKey) {
    const res = await fetch<InternalGetValidatorDashboardResponse>(API_PATH.DASHBOARD_OVERVIEW, undefined, { dashboardKey: key })
    data.value = res.data // NOTE: set data in store

    return overview.value
  }

  return { overview, refreshOverview }
}
