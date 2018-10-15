<template>
  <b-table
    v-if="logs.length > 0"
    :data="logs"
    :default-sort="['id', 'asc']"
    :per-page="settings.logsPerPage"
    detailed
    paginated
    pagination-simple
  >
    <template slot-scope="props">
      <b-table-column
        field="timestamp"
        label="Time"
        v-html="formatTimestamp(props.row.timestamp)"
        v-if="settings.cols.timestamp"
      />
      <b-table-column
        v-if="settings.cols.id"
        field="msg_id"
        label="ID"
      >{{ props.row.msg_id }}</b-table-column>
      <b-table-column
        v-if="settings.cols.client"
        field="client"
        label="Client"
      >{{ props.row.client }}</b-table-column>
      <b-table-column
        v-if="settings.cols.hostname"
        field="hostname"
        label="Hostname"
      >{{ props.row.hostname }}</b-table-column>
      <b-table-column
        v-if="settings.cols.pid"
        field="proc_id"
        label="PID"
        numeric
      >{{ props.row.proc_id }}</b-table-column>
      <b-table-column
        v-if="settings.cols.facility"
        field="facility"
        label="Facility"
        numeric
      >{{ props.row.facility }}</b-table-column>
      <b-table-column
        v-if="settings.cols.priority"
        field="priority"
        label="Pri"
        numeric
      >{{ props.row.priority }}</b-table-column>
      <b-table-column
        v-if="settings.cols.severity"
        field="severity"
        label="Sev"
        numeric
      >
        <span class="tag" :class="severityClass(props.row.severity)">
          {{ props.row.severity }}
        </span>
      </b-table-column>
      <b-table-column
        v-if="settings.cols.message"
        field="message"
        label="Message"
      >{{ props.row.message }}</b-table-column>
    </template>

    <template slot="detail" slot-scope="props">
      <p v-if="props.row.structured_data">
        {{ props.row.structured_data }}
      </p>
      <p v-else>
        No structured data provided.
      </p>
    </template>
  </b-table>
  <div v-else class="notification has-text-centered">
    <b-icon icon="package-variant" size="is-large"></b-icon>
    <p>
      No log messages received&hellip; yet!
    </p>
  </div>
</template>

<script>
export default {
  name: 'logs-display',
  props: ['logs', 'settings'],
  data() {
    return {
      todayDateString: new Date().toLocaleDateString(),
    };
  },
  methods: {
    formatTimestamp(ts) {
      const d = new Date(ts * 1000);
      const lds = d.toLocaleDateString();

      // If the log is from today, don't show the date
      if (this.todayDateString === lds) {
        return d.toLocaleTimeString();
      }

      return `${lds} ${d.toLocaleTimeString()}`;
    },
    severityClass(severity) {
      if (severity === 0) { // Emergicy
        return 'is-dark';
      } else if (severity < 4) { // Alert, Critical, Error
        return 'is-danger';
      } else if (severity === 4) { // Warning
        return 'is-warning';
      } else if (severity === 7) { // Debug
        return 'is-info';
      }

      return 'is-light';
    },
  },
};
</script>
