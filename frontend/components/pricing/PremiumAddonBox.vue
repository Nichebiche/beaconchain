<script lang="ts" setup>

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faInfoCircle } from '@fortawesome/pro-regular-svg-icons'
import { type ExtraDashboardValidatorsPremiumAddon, ProductCategoryPremiumAddon } from '~/types/api/user'
import { formatPremiumProductPrice } from '~/utils/format'

const { t: $t } = useI18n()
const { user } = useUserStore()
const { products, bestPremiumProduct } = useProductsStore()

interface Props {
  addon: ExtraDashboardValidatorsPremiumAddon,
  addonsAvailable: boolean,
  isYearly: boolean
}
const props = defineProps<Props>()

const quantity = computed(() => {
  let q = 0
  user.value?.subscriptions?.forEach((subscription) => {
    if (subscription.product_id === props.addon.product_id) {
      q += products.value?.extra_dashboard_validators_premium_addons.find(addon => addon.product_id === subscription.product_id) !== undefined ? 1 : 0
    }
  })

  return q
})

const prices = computed(() => {
  const mainPrice = props.isYearly ? props.addon.price_per_year_eur / 12 : props.addon.price_per_month_eur

  const savingAmount = props.addon.price_per_month_eur * 12 - props.addon.price_per_year_eur
  const savingDigits = savingAmount % 100 === 0 ? 0 : 2

  return {
    main: formatPremiumProductPrice($t, mainPrice),
    monthly: formatPremiumProductPrice($t, props.addon.price_per_month_eur),
    monthly_based_on_yearly: formatPremiumProductPrice($t, props.addon.price_per_year_eur / 12),
    yearly: formatPremiumProductPrice($t, props.addon.price_per_year_eur),
    saving: formatPremiumProductPrice($t, savingAmount, savingDigits),
    perValidator: formatPremiumProductPrice($t, mainPrice / props.addon.extra_dashboard_validators, 5)
  }
})

const text = computed(() => {
  return {
    validatorCount: $t('pricing.addons.validator_amount', { amount: formatNumber(props.addon.extra_dashboard_validators) }),
    perValidator: $t('pricing.per_validator', { amount: prices.value.perValidator })
  }
})

const addonButton = computed(() => {
  let text = $t('pricing.addons.button.select_addon')
  if (user.value?.subscriptions?.find(sub => sub.product_category === ProductCategoryPremiumAddon) !== undefined) {
    text = $t('pricing.addons.button.manage_addon')
  }

  return { text, disabled: !props.addonsAvailable }
})

</script>

<template>
  <div class="box-container">
    <div class="summary-container">
      <div class="validator-count">
        {{ text.validatorCount }}
        <div class="subtext">
          {{ $t('pricing.addons.per_dashboard') }}
          <BcTooltip position="top" :fit-content="true">
            <FontAwesomeIcon :icon="faInfoCircle" class="tooltip-icon" />
            <template #tooltip>
              <div class="saving-tooltip-container">
                {{ $t('pricing.pectra_tooltip', { effectiveBalance: formatNumber(props.addon?.extra_dashboard_validators * 32) }) }}
              </div>
            </template>
          </BcTooltip>
        </div>
        <div class="per-validator">
          {{ text.perValidator }}
        </div>
      </div>
    </div>
    <div class="price-container">
      <div class="price">
        <template v-if="isYearly">
          <div>
            {{ prices.monthly_based_on_yearly }}
          </div>
          <div class="month" yearly>
            {{ $t('pricing.per_month') }}
          </div>
          <div class="year">
            {{ $t('pricing.amount_per_year', {amount: prices.yearly}) }}*
          </div>
        </template>
        <template v-else>
          <div>
            {{ prices.monthly }}
          </div>
          <div class="month">
            {{ $t('pricing.per_month') }}*
          </div>
        </template>
      </div>
      <div v-if="isYearly" class="saving-info">
        <div>
          {{ $t('pricing.savings', {amount: prices.saving}) }}
        </div>
        <BcTooltip position="top" :fit-content="true">
          <FontAwesomeIcon :icon="faInfoCircle" />
          <template #tooltip>
            <div class="saving-tooltip-container">
              {{ $t('pricing.savings_tooltip', {monthly: prices.monthly, monthly_yearly: prices.monthly_based_on_yearly}) }}
            </div>
          </template>
        </BcTooltip>
      </div>
      <div class="quantity-container">
        <div>
          {{ $t('pricing.addons.quantity', { quantity }) }}
        </div>
      </div>
      <Button :label="addonButton.text" class="select-button" :disabled="addonButton.disabled" />
      <div v-if="bestPremiumProduct" class="footer">
        {{ $t('pricing.addons.requires_plan', {name: bestPremiumProduct.product_name}) }}
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.box-container {
  width: 348px;
  height: 100%;
  background-color: var(--container-background);
  border: 2px solid var(--container-border-color);
  border-radius: 7px;
  flex-shrink: 0;
  text-align: center;

  .summary-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    border-bottom: 2px solid var(--container-border-color);
    padding: 35px 0 26px 0;

    .validator-count {
      font-size: 24px;
      font-weight: 600;

      .subtext {
        font-weight: 400;
        margin-bottom: 16px;

        .tooltip-icon {
          width: 15px;
        }
      }

      .per-validator {
        color: var(--text-color-discreet);
        font-size: 20px;
        font-weight: 400;
      }
    }
  }

  .price-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 24px 34px 9px 34px;

    .price {
      font-size: 32px;
      font-weight: 600;
      margin-bottom: 28px;

      .month {
        color: var(--text-color-discreet);
        font-size: 20px;
        font-weight: 600;

        &[yearly] {
          font-size: 17px;
          font-weight: 500;
        }
      }

      .year {
        color: var(--text-color-discreet);
        font-size: 17px;
        font-weight: 500;
      }
    }

    .saving-info {
      width: 100%;
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;
      gap: 13px;
      height: 37px;
      border-radius: 18px;
      background: var(--subcontainer-background);
      font-size: 17px;
      margin-bottom: 29px;
    }

    .quantity-container {
      display: flex;
      align-items: center;
      gap: 13px;
      font-size: 20px;
      margin-bottom: 32px;

      :deep(.p-inputtext.p-component.p-inputnumber-input) {
        width: 52px;
        border-radius: 9px;
        text-align: center;
      }
    }

    .select-button {
      width: 100%;
      height: 52px;
      font-size: 25px;
      font-weight: 500;
      margin-bottom: 26px;
    }

    .footer {
      width: 100%;
      text-align: right;
      font-size: 14px;
      font-weight: 400;
      color: var(--text-color-discreet);
    }
  }

  @media (max-width: 600px) {
    width: 200px;

    .summary-container {
      padding: 20px 0 18px 0;

      .validator-count {
        font-size: 14px;

        .subtext {
          margin-bottom: 10px;

          .tooltip-icon {
            width: 13px;
          }
        }

        .per-validator {
          font-size: 12px;
        }
      }
    }

    .price-container {
      padding: 10px 25px 4px 25px;

      .price {
        font-size: 18px;
        margin-bottom: 11px;

        .month {
          font-size: 12px;

          &[yearly] {
            font-size: 10px;
          }
        }

        .year {
          font-size: 10px;
        }
      }

      .saving-info {
        height: 21px;
        gap: 4px;
        font-size: 10px;
        margin-bottom: 17px;
      }

      .quantity-container {
        font-size: 12px;
        margin-bottom: 18px;
      }

      .select-button {
        height: 30px;
        font-size: 14px;
        margin-bottom: 10px;
        padding-left: 10px;
        padding-right: 10px;
      }

      .footer {
        font-size: 10px;
      }
    }
  }
}

.saving-tooltip-container {
  width: 150px;
  text-align: left;
}
</style>