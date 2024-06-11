import { defineStore } from 'pinia'
import { API_PATH } from '~/types/customFetch'
import type { ApiDataResponse } from '~/types/api/common'
import * as networkTs from '~/types/network'

interface ApiChainInfo {
  chain_id: networkTs.ChainIDs,
  name: string
}

const store = defineStore('network-store', () => {
  const data = ref<{
    availableNetworks: ApiChainInfo[],
    currentNetwork: networkTs.ChainIDs,
    currentNetworkHasBeenChosen: boolean
  }>({
    // default values if anything wrong happens when the list of available networks is requested from the API
    availableNetworks: [{ chain_id: networkTs.ChainIDs.Ethereum, name: networkTs.ChainInfo[networkTs.ChainIDs.Ethereum].name }],
    currentNetwork: networkTs.ChainIDs.Ethereum,
    currentNetworkHasBeenChosen: false
  })
  return { data }
})

export function useNetworkStore () {
  const { data } = storeToRefs(store())

  /**
   * Needs to be called once, when the front-end is loading. Unnecessary afterwards.
   */
  async function loadAvailableNetworks () {
    try {
      const { fetch } = useCustomFetch()
      const list = await fetch<ApiDataResponse<ApiChainInfo[]>>(API_PATH.AVAILABLE_NETWORKS)
      data.value.availableNetworks = list.data.sort((a, b) => networkTs.ChainInfo[a.chain_id].priority - networkTs.ChainInfo[b.chain_id].priority)
      if (!data.value.currentNetworkHasBeenChosen) {
      // by default, the current network is the one with the best priority
        data.value.currentNetwork = data.value.availableNetworks[0].chain_id
      }
      return true
    } catch {
      return false
    }
  }

  const availableNetworks = computed(() => data.value.availableNetworks.map(apiInfo => apiInfo.chain_id))
  const currentNetwork = computed(() => data.value.currentNetwork)
  const networkInfo = computed(() => networkTs.ChainInfo[data.value.currentNetwork])

  function setCurrentNetwork (chainId: networkTs.ChainIDs) {
    data.value.currentNetwork = chainId
    data.value.currentNetworkHasBeenChosen = true
  }

  function isMainNet () : boolean {
    return networkTs.isMainNet(currentNetwork.value)
  }

  function isL1 () : boolean {
    return networkTs.isL1(currentNetwork.value)
  }

  function epochsPerDay (): number {
    return networkTs.epochsPerDay(currentNetwork.value)
  }

  function epochToTs (epoch: number): number | undefined {
    return networkTs.epochToTs(currentNetwork.value, epoch)
  }

  function slotToTs (slot: number): number | undefined {
    return networkTs.slotToTs(currentNetwork.value, slot)
  }

  function tsToSlot (ts: number): number {
    return networkTs.tsToSlot(currentNetwork.value, ts)
  }

  function slotToEpoch (slot: number): number {
    return networkTs.slotToEpoch(currentNetwork.value, slot)
  }

  return {
    loadAvailableNetworks,
    availableNetworks,
    currentNetwork,
    networkInfo,
    setCurrentNetwork,
    isMainNet,
    isL1,
    epochsPerDay,
    epochToTs,
    slotToTs,
    tsToSlot,
    slotToEpoch
  }
}