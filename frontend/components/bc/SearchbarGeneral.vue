<script setup lang="ts">
import { SearchbarShape, type SearchbarColors, SearchbarPurpose, ResultType, type ResultSuggestion, pickHighestPriorityAmongBestMatchings } from '~/types/searchbar'
import { ChainInfo } from '~/types/network'

defineProps<{
  barShape: SearchbarShape,
  colorTheme: SearchbarColors,
  screenWidthCausingSuddenChange: number // this information is needed by MiddleEllipsis
}>()

async function redirectToRelevantPage (result : ResultSuggestion) {
  let path : string
  let q = ''
  const networkPath = '/networks' + ChainInfo[result.chainId].path

  switch (result.type) {
    case ResultType.Tokens :
    case ResultType.NFTs :
      path = '/token/' + result.queryParam
      break
    case ResultType.Epochs :
      path = networkPath + '/epoch/' + result.queryParam
      break
    case ResultType.Slots :
      path = networkPath + '/slot/' + result.queryParam
      break
    case ResultType.Blocks :
      path = networkPath + '/block/' + result.queryParam
      break
    case ResultType.BlockRoots :
    case ResultType.StateRoots :
    case ResultType.Transactions :
      path = networkPath + '/tx/' + result.queryParam
      break
    case ResultType.TransactionBatches :
      path = networkPath + '/transactionbatch/' + result.queryParam
      break
    case ResultType.StateBatches :
      path = networkPath + '/batch/' + result.queryParam
      break
    case ResultType.Contracts :
    case ResultType.Accounts :
    case ResultType.EnsAddresses :
      path = '/address/' + result.queryParam
      break
    case ResultType.EnsOverview :
      path = '/ens/' + result.queryParam
      break
    case ResultType.Graffiti :
      path = networkPath + '/slots'
      q = result.queryParam
      break
    case ResultType.ValidatorsByIndex :
    case ResultType.ValidatorsByPubkey :
      path = networkPath + '/validator/' + result.queryParam
      break
    case ResultType.ValidatorsByDepositAddress :
    case ResultType.ValidatorsByDepositEnsName :
      path = networkPath + '/validators/deposits'
      q = result.queryParam
      break
    case ResultType.ValidatorsByWithdrawalCredential :
    case ResultType.ValidatorsByWithdrawalAddress :
    case ResultType.ValidatorsByWithdrawalEnsName :
      path = networkPath + '/validators/withdrawals'
      q = result.queryParam
      break
    default :
      return
  }

  if (q !== '') {
    await navigateTo({ path, query: { q } })
  } else {
    await navigateTo({ path })
  }
}
</script>

<template>
  <BcSearchbarMain
    :bar-shape="barShape"
    :color-theme="colorTheme"
    :bar-purpose="SearchbarPurpose.GlobalSearch"
    :pick-by-default="pickHighestPriorityAmongBestMatchings"
    :screen-width-causing-sudden-change="screenWidthCausingSuddenChange"
    @go="redirectToRelevantPage"
  />
</template>
