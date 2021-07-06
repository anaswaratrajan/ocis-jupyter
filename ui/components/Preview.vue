<template>
    <div class="uk-flex">
      <div class="uk-container uk-width-1-2">
        <h1>File preview</h1>
      </div>
      <div class="uk-container uk-width-1-2">
      </div>
    </div>
</template>

<script>
export default {
  name: 'Preview',
  data: function () {
    return {
      loading: true,
    }
  },
  computed: {
    message () {
      return this.$store.getters['OCIS-JUPYTER/message']
    }
  },
  methods: {
    submitName (name) {
      this.$store.dispatch('OCIS-JUPYTER/submitName', name)
    },
  },
}
</script>
