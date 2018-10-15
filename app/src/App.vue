<template>
  <div>
    <nav class="navbar is-primary" role="navigation" aria-label="main navigation">
      <div class="container">
        <div class="navbar-brand">
          <a class="navbar-item" href="/">
            <img src="./assets/parrot.svg" width="112" height="28">
          </a>
        </div>

        <div class="navbar-menu">
          <div class="navbar-end">
            <div class="navbar-item">
              <span class="tag" :class="{ 'is-success': connected, 'is-danger': !connected }">
                <b-icon icon="server-network" size="is-small" custom-class="is-left"></b-icon>
                <p>{{ connected ? 'Connected' : 'Disconnected' }}</p>
              </span>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <div class="container is-fluid">
      <div class="columns">
        <div class="column is-one-fifth">

          <feeds-nav
            v-if="feedsExist"
            :feeds="feeds"
            @select="setActiveFeed"
          />

          <button class="button is-fullwidth" @click="$refs.settingsModal.show()">
            <b-icon icon="settings"/>
            <span>Settings</span>
          </button>

          <settings-modal ref="settingsModal" :settings.sync="settings" />
        </div>
        <div class="column" v-if="feedsExist">

          <logs-display
            v-if="activeFeed !== null"
            :logs="activeFeed.logs"
            :settings="settings"
          />

          <div class="notification has-text-centered" v-else>
            <b-icon icon="human-handsup" size="is-medium" />
            <h1 class="subtitle is-4">
              Select a feed!
            </h1>
          </div>

        </div>
        <div class="column" v-else>
          <div class="notification has-text-centered" v-if="connected">
            <b-icon icon="ear-hearing" size="is-medium" />
            <h1 class="subtitle is-4">
              Parrot hasn't heard from any loggers yet, but we're listening&hellip;
            </h1>
          </div>
          <div class="notification is-danger has-text-centered" v-else>
            <b-icon icon="server-network" size="is-medium" />
            <h1 class="subtitle is-4">
              Parrot isn't able to connect to the server.
            </h1>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import FeedsNav from '@/components/FeedsNav.vue';
import LogsDisplay from '@/components/LogsDisplay.vue';
import SettingsModal from '@/components/SettingsModal.vue';

export default {
  name: 'app',
  components: {
    LogsDisplay,
    FeedsNav,
    SettingsModal,
  },
  computed: {
    feedsExist() {
      const feedNames = Object.keys(this.feeds);
      const exist = feedNames.length > 0;

      // Yes, this is technically a "side-effect", but it's also
      // the most performant place to put this logic, so suck it, purists.
      if (exist && feedNames.length === 1) {
        this.setActiveFeed(this.feeds[feedNames[0]]);
      }

      return exist;
    },
  },
  data() {
    return {
      activeFeed: null,
      connected: false,
      darkMode: false,
      feeds: {
        // james: {
        //   app_name: 'james',
        //   logs: [
        //     {
        //       app_name: 'james',
        //       client: '127.0.0.1:41900',
        //       facility: 1,
        //       hostname: 'asgard.local',
        //       message: 'foo',
        //       msg_id: '-',
        //       priority: 13,
        //       proc_id: '-',
        //       severity: 2,
        //       structured_data: '[timeQuality tzKnown="1" isSynced="1" syncAccuracy="181003"]',
        //       timestamp: 1539457763209,
        //     },
        //   ],
        // },
      },
      settings: {
        cols: {
          client: true,
          facility: true,
          hostname: true,
          id: true,
          message: true,
          pid: true,
          priority: true,
          severity: true,
          timestamp: true,
        },
        logsPerPage: 100,
        dark: false,
      },
    };
  },
  methods: {
    setActiveFeed(feed) {
      this.activeFeed = feed;
    },
  },
  mounted() {
    this.$sse('/squawk', { format: 'json' })
      .then((sse) => {
        this.connected = true;
        this.sse = sse;

        sse.onError(() => {
          this.connected = false;
        });

        sse.subscribe('l', (log) => {
          if (!this.feeds[log.app_name]) {
            this.$set(this.feeds, log.app_name, {
              app_name: log.app_name,
              logs: [],
            });
          }

          this.feeds[log.app_name].logs.push(log);
        });
      })
      .catch(() => {
        this.connected = false;
      });
  },
  beforeDestroy() {
    if (this.sse) {
      this.sse.close();
    }
  },
  watch: {
    'settings.dark': (v) => {
      if (v) {
        document.body.classList.add('theme-dark');
      } else {
        document.body.classList.remove('theme-dark');
      }
    },
  },
};
</script>

<style lang="scss">
  @charset "utf-8";

  // Import Bulma's core
  @import "~bulma/sass/utilities/_all";

  /// Custom theme here

  // Import Bulma and Buefy styles
  @import "~bulma";
  @import "~buefy/src/scss/buefy";

  // Import icons
  @import "~@mdi/font/css/materialdesignicons.min.css";

  body {
    min-height: 100vh;
  }

  .navbar + .container {
    padding-top: 2rem;
  }

  .has-auto-margin-left {
    margin-left: auto;
  }

  body.theme-dark {
    background: linear-gradient(0deg, #33536c, #13334c 25%, #13334c 100%);
    color: #fff;

    nav.navbar.is-primary,
    .notification,
    .panel-block,
    .panel-heading,
    .table {
      background: #33536c;
      color: #fff;

      th {
        border-color: #53738c;
        color: #fff;
      }

      tr.detail {
        background: #23435c;
        box-shadow: inset 0 1px 3px #53738c;
      }

      td {
        border-color: #53738c;
      }
    }

    .panel-heading {
      background: #23435c;
    }

    .panel-block,
    .panel-heading {
      border-color: #33536c;

      &.is-active {
        border-left-color: #53738c;
        color: #fff;
      }
    }

    a.panel-block:hover {
      background-color: #43637c;
    }

    .chevron-cell > a {
      color: #83a3bc;
    }
  }
</style>
