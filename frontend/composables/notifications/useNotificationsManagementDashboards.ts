import type { TableQueryParams } from '~/types/datatable'
import { API_PATH } from '~/types/customFetch'
import type {
  InternalGetUserNotificationSettingsDashboardsResponse,
  InternalPutUserNotificationSettingsAccountDashboardResponse,
  InternalPutUserNotificationSettingsValidatorDashboardResponse,
  NotificationSettingsAccountDashboard,
  NotificationSettingsDashboardsTableRow,
  NotificationSettingsValidatorDashboard,
} from '~/types/api/notifications'

export function useNotificationsManagementDashboards() {
  const { fetch } = useCustomFetch()

  const data = ref<InternalGetUserNotificationSettingsDashboardsResponse>()
  const {
    cursor,
    isStoredQuery,
    onSort,
    pageSize,
    pendingQuery,
    query,
    setCursor,
    setPageSize,
    setSearch,
    setStoredQuery,
  } = useTableQuery({
    limit: 10,
    sort: 'dashboard_id:desc',
  }, 10)
  const isLoading = ref(false)

  const dashboardGroups = computed(() => data.value)

  async function getDashboardGroups(q?: TableQueryParams) {
    isLoading.value = true
    setStoredQuery(q)
    const res
      = await fetch<InternalGetUserNotificationSettingsDashboardsResponse>(
        API_PATH.GET_NOTIFICATIONS_SETTINGS_DASHBOARD,
        undefined,
        undefined,
        q,
      )

    isLoading.value = false
    if (!isStoredQuery(q)) {
      return // in case some query params change while loading
    }

    data.value = res
    return res
  }

  watch(
    query,
    (q) => {
      getDashboardGroups(q)
    },
    { immediate: true },
  )
  const clearSettings = (
    {
      is_account_dashboard,
      settings,
    }:
    {
      is_account_dashboard: boolean,
      settings: NotificationSettingsAccountDashboard | NotificationSettingsValidatorDashboard,
    },
  ) => {
    settings.webhook_url = ''
    settings.is_webhook_discord_enabled = false
    if (is_account_dashboard) {
      const accountDashboarSettings = settings as NotificationSettingsAccountDashboard
      accountDashboarSettings.erc20_token_transfers_value_threshold = 0
      accountDashboarSettings.is_erc1155_token_transfers_subscribed = false
      accountDashboarSettings.is_erc20_token_transfers_subscribed = false
      accountDashboarSettings.is_erc721_token_transfers_subscribed = false
      accountDashboarSettings.is_ignore_spam_transactions_enabled = false
      accountDashboarSettings.is_incoming_transactions_subscribed = false
      accountDashboarSettings.is_outgoing_transactions_subscribed = false
      return
    }
    const accountDashboarSettings = settings as NotificationSettingsValidatorDashboard
    accountDashboarSettings.group_offline_threshold = 0
    accountDashboarSettings.is_attestations_missed_subscribed = false
    accountDashboarSettings.is_block_proposal_subscribed = false
    accountDashboarSettings.is_group_offline_subscribed = false
    accountDashboarSettings.is_max_collateral_subscribed = false
    accountDashboarSettings.is_min_collateral_subscribed = false
    accountDashboarSettings.is_real_time_mode_enabled = false
    accountDashboarSettings.is_slashed_subscribed = false
    accountDashboarSettings.is_sync_subscribed = false
    accountDashboarSettings.is_upcoming_block_proposal_subscribed = false
    accountDashboarSettings.is_validator_offline_subscribed = false
    accountDashboarSettings.is_withdrawal_processed_subscribed = false
    accountDashboarSettings.max_collateral_threshold = 0
    accountDashboarSettings.min_collateral_threshold = 0
  }
  const deleteDashboardNotifications = async (
    {
      dashboard_id,
      group_id,
      is_account_dashboard,
      settings,
    }:
    Pick<
      NotificationSettingsDashboardsTableRow,
      | 'dashboard_id'
      | 'group_id'
      | 'is_account_dashboard'
      | 'settings'
    >,
  ) => {
    clearSettings({
      is_account_dashboard,
      settings,
    })
    if (is_account_dashboard) {
      return await fetch<InternalPutUserNotificationSettingsAccountDashboardResponse>(
        API_PATH.NOTIFICATIONS_MANAGEMENT_DASHBOARD_ACCOUNT_SET_NOTIFICATION,
        {
          body: settings,
        },
        {
          dashboard_id,
          group_id,
        },
      )
    }
    return await fetch<InternalPutUserNotificationSettingsValidatorDashboardResponse>(
      API_PATH.NOTIFICATIONS_MANAGEMENT_DASHBOARD_VALIDATOR_SET_NOTIFICATION,
      {
        body: settings,
      },
      {
        dashboard_id,
        group_id,
      },
    )
  }

  return {
    cursor,
    dashboardGroups,
    deleteDashboardNotifications,
    isLoading,
    onSort,
    pageSize,
    query: pendingQuery,
    setCursor,
    setPageSize,
    setSearch,
  }
}