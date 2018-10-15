<template>
  <b-modal :active.sync="active">
    <div class="box">
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">Columns</label>
        </div>
        <div class="field-body">
          <section>
            <div class="field">
              <b-checkbox v-model="settings.cols.client">
                Client
              </b-checkbox>
            </div>
            <div class="field">
              <b-checkbox v-model="settings.cols.facility">
                Facility
              </b-checkbox>
            </div>
            <div class="field">
              <b-checkbox v-model="settings.cols.hostname">
                Hostname
              </b-checkbox>
            </div>
            <div class="field">
              <b-checkbox v-model="settings.cols.id">
                ID
              </b-checkbox>
            </div>
            <div class="field">
              <b-checkbox v-model="settings.cols.message">
                Message
              </b-checkbox>
            </div>
            <div class="field">
              <b-checkbox v-model="settings.cols.pid">
                Process ID (PID)
              </b-checkbox>
            </div>
            <div class="field">
              <b-checkbox v-model="settings.cols.priority">
                Priority (Pri)
              </b-checkbox>
            </div>
            <div class="field">
              <b-checkbox v-model="settings.cols.severity">
                Severity (Sev)
              </b-checkbox>
            </div>
            <div class="field">
              <b-checkbox v-model="settings.cols.timestamp">
                Time
              </b-checkbox>
            </div>
          </section>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">Logs Per Page</label>
        </div>
        <div class="field-body">
          <div class="field">
            <b-input type="number" v-model="settings.logsPerPage">
            </b-input>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">Theme</label>
        </div>
        <div class="field-body">
          <div class="field">
            <b-switch v-model="settings.dark">
              Dark
            </b-switch>
          </div>
        </div>
      </div>

      <hr>

      <p class="has-text-centered">
        <small>Parrot icon by <a href="http://www.freepik.com" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC BY 3.0</a></small>
      </p>
    </div>
  </b-modal>
</template>

<script>
export default {
  name: 'settings-modal',
  props: ['settings'],
  data() {
    return {
      active: false,
    };
  },
  methods: {
    show() {
      this.active = true;
    },
  },
};
</script>
