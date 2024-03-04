<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faDiscord, faTwitter, faGithub } from '@fortawesome/free-brands-svg-icons'
import {
  faHistory,
  faCube,
  faCubes,
  faCreditCard,
  faUpload,
  faTable,
  faUserSlash,
  faMedal,
  faFileImport,
  faFileSignature,
  faMoneyBill,
  faChartBar,
  faChartLine,
  faFireFlame,
  faRobot,
  faMegaphone,
  faExternalLinkAlt,
  faDesktop
} from '@fortawesome/pro-solid-svg-icons'
import {
  faDrumstickBite,
  faChartPie,
  faRocket,
  faMoneyBill as farMoneyBill,
  faCalculator,
  faProjectDiagram,
  faMobileScreen,
  faGem,
  faLaptopCode,
  faGasPump,
  faBell,
  faPaintBrush,
  faDesktop as farDesktop
} from '@fortawesome/pro-regular-svg-icons'
import {
  faBuildingColumns
} from '@fortawesome/sharp-solid-svg-icons'

import IconEthereumOverview from '~/components/icon/megaMenu/EthereumOverview.vue'
import IconGnosisOverview from '~/components/icon/megaMenu/GnosisOverview.vue'
import IconArbitrumOverview from '~/components/icon/megaMenu/ArbitrumOverview.vue'
import IconBaseOverview from '~/components/icon/megaMenu/BaseOverview.vue'
import IconOptimismOverview from '~/components/icon/megaMenu/OptimismOverview.vue'
import IconEthermineStaking from '~/components/icon/megaMenu/EthermineStaking.vue'
import IconEthStore from '~/components/icon/megaMenu/EthStore.vue'
import IconEversteel from '~/components/icon/megaMenu/EverSteel.vue'
import IconWebhook from '~/components/icon/megaMenu/WebHook.vue'

import { Target } from '~/types/links'

const { t: $t } = useI18n()

const minFullViewWidth = '1360px'

const items = [
  {
    label: 'Ethereum',
    items: [
      [
        {
          label: $t('header.megamenu.blockchain'),
          items: [
            {
              label: $t('header.megamenu.overview'),
              svg: IconEthereumOverview,
              url: '/'
            },
            {
              label: $t('header.megamenu.epochs'),
              icon: faHistory,
              url: '/epochs'
            },
            {
              label: $t('header.megamenu.slots'),
              icon: faCube,
              url: '/slots'
            },
            {
              label: $t('header.megamenu.blocks'),
              icon: faCubes,
              url: '/blocks'
            },
            {
              label: $t('header.megamenu.txs'),
              icon: faCreditCard,
              url: '/transactions'
            },
            {
              label: $t('header.megamenu.mempool'),
              icon: faUpload,
              url: '/mempool'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.validators'),
          items: [
            {
              label: $t('header.megamenu.overview'),
              icon: faTable,
              url: '/validators'
            },
            {
              label: $t('header.megamenu.slashings'),
              icon: faUserSlash,
              url: '/validators/slashings'
            },
            {
              label: $t('header.megamenu.validator_leaderboard'),
              icon: faMedal,
              url: '/validators/leaderboard'
            },
            {
              label: $t('header.megamenu.deposit_leaderboard'),
              icon: faFileImport,
              url: '/validators/deposit-leaderboard'
            },
            {
              label: $t('header.megamenu.deposits'),
              icon: faFileSignature,
              url: '/validators/deposits'
            },
            {
              label: $t('header.megamenu.withdrawals'),
              icon: faMoneyBill,
              url: '/validators/withdrawals'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.staking_pools'),
          items: [
            {
              label: $t('header.megamenu.run_a_validator'),
              svg: IconEthermineStaking,
              url: 'https://ethpool.org/',
              class: 'orange-box',
              target: Target.External
            },
            {
              label: $t('header.megamenu.eth_store'),
              svg: IconEthStore,
              url: '/ethstore'
            },
            {
              label: $t('header.megamenu.staking_services'),
              icon: faDrumstickBite,
              url: '/stakingServices'
            },
            {
              label: $t('header.megamenu.pool_benchmarks'),
              icon: faChartPie,
              url: '/pools'
            },
            {
              label: $t('header.megamenu.rocket_pool_stats'),
              icon: faRocket,
              url: '/pools/rocketpool'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.stats'),
          items: [
            {
              label: $t('header.megamenu.charts'),
              icon: faChartBar,
              url: '/charts'
            },
            {
              label: $t('header.megamenu.reward_history'),
              icon: farMoneyBill,
              url: '/rewards'
            },
            {
              label: $t('header.megamenu.profit_calculator'),
              icon: faCalculator,
              url: '/calculator'
            },
            {
              label: $t('header.megamenu.block_viz'),
              icon: faProjectDiagram,
              url: '/vis'
            },
            {
              label: $t('header.megamenu.correlations'),
              icon: faChartLine,
              url: '/correlations'
            },
            {
              label: $t('header.megamenu.eip1599_burn'),
              icon: faFireFlame,
              url: '/burn'
            },
            {
              label: $t('header.megamenu.relays'),
              icon: faRobot,
              url: '/relays'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.tools'),
          items: [
            {
              label: $t('header.megamenu.beaconchain_app'),
              icon: faMobileScreen,
              url: '/mobile'
            },
            {
              label: $t('header.megamenu.beaconchain_premium'),
              icon: faGem,
              url: '/premium'
            },
            {
              label: $t('header.megamenu.webhooks'),
              svg: IconWebhook,
              url: '/user/webhooks'
            },
            {
              label: $t('header.megamenu.api_docs'),
              icon: faLaptopCode,
              url: '/api/v1/docs/index.html'
            },
            {
              label: $t('header.megamenu.api_pricing'),
              icon: faLaptopCode,
              url: '/pricing'
            },
            {
              label: $t('header.megamenu.unit_converter'),
              icon: faBuildingColumns,
              url: '/tools/unitConverter'
            },
            {
              label: $t('header.megamenu.gasnow'),
              icon: faGasPump,
              url: '/gasnow'
            },
            {
              label: $t('header.megamenu.broadcast_signed_messages'),
              icon: faMegaphone,
              url: '/tools/broadcast'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.services'),
          items: [
            {
              label: $t('header.megamenu.eversteel'),
              svg: IconEversteel,
              url: 'https://eversteel.io/',
              class: 'orange-box',
              target: Target.External
            },
            {
              label: $t('header.megamenu.notifications'),
              icon: faBell,
              url: '/user/notifications'
            },
            {
              label: $t('header.megamenu.graffiti_wall'),
              icon: faPaintBrush,
              url: '/graffitiwall'
            },
            {
              label: $t('header.megamenu.ethereum_clients'),
              icon: farDesktop,
              url: '/ethClients'
            },
            {
              label: $t('header.megamenu.knowledge_base'),
              icon: faExternalLinkAlt,
              url: 'https://kb.beaconcha.in',
              target: Target.External
            },
            {
              label: $t('header.megamenu.slot_finder'),
              icon: faCube,
              url: '/slots/finder'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.community'),
          items: [
            {
              label: 'Discord',
              icon: faDiscord,
              url: 'https://dsc.gg/beaconchain',
              target: Target.External
            },
            {
              label: 'Twitter',
              icon: faTwitter,
              url: 'https://twitter.com/beaconcha_in',
              target: Target.External
            },
            {
              label: 'Github',
              icon: faGithub,
              url: 'https://github.com/gobitfly/beaconchain',
              target: Target.External
            },
            {
              label: 'Github Mobile App',
              icon: faGithub,
              url: 'https://github.com/gobitfly/eth2-beaconchain-explorer-app',
              target: Target.External
            }
          ]
        }
      ]]
  },
  {
    label: 'Gnosis',
    items: [
      [
        {
          label: $t('header.megamenu.blockchain'),
          items: [
            {
              label: $t('header.megamenu.overview'),
              svg: IconGnosisOverview,
              url: '/'
            },
            {
              label: $t('header.megamenu.epochs'),
              icon: faHistory,
              url: '/epochs'
            },
            {
              label: $t('header.megamenu.slots'),
              icon: faCube,
              url: '/slots'
            },
            {
              label: $t('header.megamenu.blocks'),
              icon: faCubes,
              url: '/blocks'
            },
            {
              label: $t('header.megamenu.txs'),
              icon: faCreditCard,
              url: '/transactions'
            },
            {
              label: $t('header.megamenu.validators'),
              icon: faDesktop,
              url: ''
            },
            {
              label: $t('header.megamenu.mempool'),
              icon: faUpload,
              url: '/mempool'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.validators'),
          items: [
            {
              label: $t('header.megamenu.overview'),
              icon: faTable,
              url: '/validators'
            },
            {
              label: $t('header.megamenu.slashings'),
              icon: faUserSlash,
              url: '/validators/slashings'
            },
            {
              label: $t('header.megamenu.validator_leaderboard'),
              icon: faMedal,
              url: '/validators/leaderboard'
            },
            {
              label: $t('header.megamenu.deposit_leaderboard'),
              icon: faFileImport,
              url: '/validators/deposit-leaderboard'
            },
            {
              label: $t('header.megamenu.deposits'),
              icon: faFileSignature,
              url: '/validators/deposits'
            },
            {
              label: $t('header.megamenu.withdrawals'),
              icon: faMoneyBill,
              url: '/validators/withdrawals'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.stats'),
          items: [
            {
              label: $t('header.megamenu.charts'),
              icon: faChartBar,
              url: '/charts'
            },
            {
              label: $t('header.megamenu.reward_history'),
              icon: farMoneyBill,
              url: '/rewards'
            },
            {
              label: $t('header.megamenu.profit_calculator'),
              icon: faCalculator,
              url: '/calculator'
            },
            {
              label: $t('header.megamenu.block_viz'),
              icon: faProjectDiagram,
              url: '/vis'
            },
            {
              label: $t('header.megamenu.correlations'),
              icon: faChartLine,
              url: '/correlations'
            },
            {
              label: $t('header.megamenu.eip1599_burn'),
              icon: faFireFlame,
              url: '/burn'
            },
            {
              label: $t('header.megamenu.relays'),
              icon: faRobot,
              url: '/relays'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.tools'),
          items: [
            {
              label: $t('header.megamenu.beaconchain_app'),
              icon: faMobileScreen,
              url: '/mobile'
            },
            {
              label: $t('header.megamenu.beaconchain_premium'),
              icon: faGem,
              url: '/premium'
            },
            {
              label: $t('header.megamenu.webhooks'),
              svg: IconWebhook,
              url: '/user/webhooks'
            },
            {
              label: $t('header.megamenu.api_docs'),
              icon: faLaptopCode,
              url: '/api/v1/docs/index.html'
            },
            {
              label: $t('header.megamenu.api_pricing'),
              icon: faLaptopCode,
              url: '/pricing'
            },
            {
              label: $t('header.megamenu.broadcast_signed_messages'),
              icon: faMegaphone,
              url: '/tools/broadcast'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.services'),
          items: [
            {
              label: $t('header.megamenu.eversteel'),
              svg: IconEversteel,
              url: 'https://eversteel.io/',
              class: 'orange-box',
              target: Target.External
            },
            {
              label: $t('header.megamenu.notifications'),
              icon: faBell,
              url: '/user/notifications'
            },
            {
              label: $t('header.megamenu.knowledge_base'),
              icon: faExternalLinkAlt,
              url: 'https://kb.beaconcha.in',
              target: Target.External
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.community'),
          items: [
            {
              label: 'Discord',
              icon: faDiscord,
              url: 'https://dsc.gg/beaconchain',
              target: Target.External
            },
            {
              label: 'Twitter',
              icon: faTwitter,
              url: 'https://twitter.com/beaconcha_in',
              target: Target.External
            },
            {
              label: 'Github',
              icon: faGithub,
              url: 'https://github.com/gobitfly/beaconchain',
              target: Target.External
            },
            {
              label: 'Github Mobile App',
              icon: faGithub,
              url: 'https://github.com/gobitfly/eth2-beaconchain-explorer-app',
              target: Target.External
            }
          ]
        }
      ]]
  },
  {
    label: 'Arbitrum',
    items: [
      [
        {
          label: $t('header.megamenu.blockchain'),
          items: [
            {
              label: $t('header.megamenu.overview'),
              svg: IconArbitrumOverview,
              url: '/'
            },
            {
              label: $t('header.megamenu.blocks'),
              icon: faCubes,
              url: '/blocks'
            },
            {
              label: $t('header.megamenu.txs'),
              icon: faCreditCard,
              url: '/transactions'
            },
            {
              label: $t('header.megamenu.txsL1L2'),
              icon: faCreditCard,
              url: ''
            },
            {
              label: $t('header.megamenu.txsL2L1'),
              icon: faCreditCard,
              url: ''
            },
            {
              label: $t('header.megamenu.mempool'),
              icon: faUpload,
              url: '/mempool'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.tools'),
          items: [
            {
              label: $t('header.megamenu.webhooks'),
              svg: IconWebhook,
              url: '/user/webhooks'
            },
            {
              label: $t('header.megamenu.api_docs'),
              icon: faLaptopCode,
              url: '/api/v1/docs/index.html'
            },
            {
              label: $t('header.megamenu.api_pricing'),
              icon: faLaptopCode,
              url: '/pricing'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.services'),
          items: [
            {
              label: $t('header.megamenu.eversteel'),
              svg: IconEversteel,
              url: 'https://eversteel.io/',
              class: 'orange-box',
              target: Target.External
            },
            {
              label: $t('header.megamenu.knowledge_base'),
              icon: faExternalLinkAlt,
              url: 'https://kb.beaconcha.in',
              target: Target.External
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.community'),
          items: [
            {
              label: 'Discord',
              icon: faDiscord,
              url: 'https://dsc.gg/beaconchain',
              target: Target.External
            },
            {
              label: 'Twitter',
              icon: faTwitter,
              url: 'https://twitter.com/beaconcha_in',
              target: Target.External
            },
            {
              label: 'Github',
              icon: faGithub,
              url: 'https://github.com/gobitfly/beaconchain',
              target: Target.External
            },
            {
              label: 'Github Mobile App',
              icon: faGithub,
              url: 'https://github.com/gobitfly/eth2-beaconchain-explorer-app',
              target: Target.External
            }
          ]
        }
      ]]
  },
  {
    label: 'Base',
    items: [
      [
        {
          label: $t('header.megamenu.blockchain'),
          items: [
            {
              label: $t('header.megamenu.overview'),
              svg: IconBaseOverview,
              url: '/'
            },
            {
              label: $t('header.megamenu.blocks'),
              icon: faCubes,
              url: '/blocks'
            },
            {
              label: $t('header.megamenu.txs'),
              icon: faCreditCard,
              url: '/transactions'
            },
            {
              label: $t('header.megamenu.txsL1L2'),
              icon: faCreditCard,
              url: ''
            },
            {
              label: $t('header.megamenu.txsL2L1'),
              icon: faCreditCard,
              url: ''
            },
            {
              label: $t('header.megamenu.mempool'),
              icon: faUpload,
              url: '/mempool'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.tools'),
          items: [
            {
              label: $t('header.megamenu.webhooks'),
              svg: IconWebhook,
              url: '/user/webhooks'
            },
            {
              label: $t('header.megamenu.api_docs'),
              icon: faLaptopCode,
              url: '/api/v1/docs/index.html'
            },
            {
              label: $t('header.megamenu.api_pricing'),
              icon: faLaptopCode,
              url: '/pricing'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.services'),
          items: [
            {
              label: $t('header.megamenu.eversteel'),
              svg: IconEversteel,
              url: 'https://eversteel.io/',
              class: 'orange-box',
              target: Target.External
            },
            {
              label: $t('header.megamenu.knowledge_base'),
              icon: faExternalLinkAlt,
              url: 'https://kb.beaconcha.in',
              target: Target.External
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.community'),
          items: [
            {
              label: 'Discord',
              icon: faDiscord,
              url: 'https://dsc.gg/beaconchain',
              target: Target.External
            },
            {
              label: 'Twitter',
              icon: faTwitter,
              url: 'https://twitter.com/beaconcha_in',
              target: Target.External
            },
            {
              label: 'Github',
              icon: faGithub,
              url: 'https://github.com/gobitfly/beaconchain',
              target: Target.External
            },
            {
              label: 'Github Mobile App',
              icon: faGithub,
              url: 'https://github.com/gobitfly/eth2-beaconchain-explorer-app',
              target: Target.External
            }
          ]
        }
      ]]
  },
  {
    label: 'Optimism',
    items: [
      [
        {
          label: $t('header.megamenu.blockchain'),
          items: [
            {
              label: $t('header.megamenu.overview'),
              svg: IconOptimismOverview,
              url: '/'
            },
            {
              label: $t('header.megamenu.blocks'),
              icon: faCubes,
              url: '/blocks'
            },
            {
              label: $t('header.megamenu.txs'),
              icon: faCreditCard,
              url: '/transactions'
            },
            {
              label: $t('header.megamenu.txsL1L2'),
              icon: faCreditCard,
              url: ''
            },
            {
              label: $t('header.megamenu.txsL2L1'),
              icon: faCreditCard,
              url: ''
            },
            {
              label: $t('header.megamenu.mempool'),
              icon: faUpload,
              url: '/mempool'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.tools'),
          items: [
            {
              label: $t('header.megamenu.webhooks'),
              svg: IconWebhook,
              url: '/user/webhooks'
            },
            {
              label: $t('header.megamenu.api_docs'),
              icon: faLaptopCode,
              url: '/api/v1/docs/index.html'
            },
            {
              label: $t('header.megamenu.api_pricing'),
              icon: faLaptopCode,
              url: '/pricing'
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.services'),
          items: [
            {
              label: $t('header.megamenu.eversteel'),
              svg: IconEversteel,
              url: 'https://eversteel.io/',
              class: 'orange-box',
              target: Target.External
            },
            {
              label: $t('header.megamenu.knowledge_base'),
              icon: faExternalLinkAlt,
              url: 'https://kb.beaconcha.in',
              target: Target.External
            }
          ]
        }
      ],
      [
        {
          label: $t('header.megamenu.community'),
          items: [
            {
              label: 'Discord',
              icon: faDiscord,
              url: 'https://dsc.gg/beaconchain',
              target: Target.External
            },
            {
              label: 'Twitter',
              icon: faTwitter,
              url: 'https://twitter.com/beaconcha_in',
              target: Target.External
            },
            {
              label: 'Github',
              icon: faGithub,
              url: 'https://github.com/gobitfly/beaconchain',
              target: Target.External
            },
            {
              label: 'Github Mobile App',
              icon: faGithub,
              url: 'https://github.com/gobitfly/eth2-beaconchain-explorer-app',
              target: Target.External
            }
          ]
        }
      ]
    ]
  },
  {
    label: $t('header.megamenu.dashboard'),
    url: '/dashboard'
  },
  {
    label: $t('header.megamenu.notifications'),
    url: '/user/notifications'
  }
]
</script>

<template>
  <MegaMenu :model="items" :breakpoint="minFullViewWidth">
    <template #itemicon="{ item }">
      <span v-if="item.svg || item.icon" class="p-menuitem-icon iconSpacing" data-pc-section="icon">
        <component :is="item.svg" v-if="item.svg" />
        <FontAwesomeIcon v-else-if="item.icon" class="icon" :icon="item.icon" />
      </span>
    </template>
  </MegaMenu>
</template>

<style lang="scss"  scoped>
.iconSpacing {
  width: 25px;
  position: relative;

  svg,
  i {
    position: absolute;
    transform: translateY(-50%);
  }
}
</style>