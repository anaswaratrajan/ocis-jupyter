<template>
    <div class="uk-flex uk-grid-divider oc-py-m">
      <div class="uk-container uk-width-3-4">
        <span v-html="this.ipynbHTML"></span>
      </div>
      <div class="uk-container uk-width-1-4">
        <oc-button href="https://jupyter.org/" appearance="filled" variation="primary" class="oc-mr-s oc-mb-s">Open in SWAN</oc-button>
        <h4>swan project details</h4>
        
        <div>config_val_1: configDetails</div>
        <div>config_val_2: configDetails</div>
        <div>...</div>
        <div>config_val_N: configDetails</div>
      </div>
    </div>
</template>

<script>
export default {
  name: 'Preview',
  data: function () {
    return {
      filePath: '',
      fileContent: ''
    }
  },
  computed: {
    ipynbHTML () {
      return this.$store.getters['OCIS-JUPYTER/nbcontent']
    }
  },
  created() {
    this.filePath = this.$route.params.filePath
  },
  methods: {
    loadFileContent() {
      this.$client.files
        .getFileContents(this.filePath, { resolveWithResponseObject: true })
        .then(resp => {
          this.fileContent = resp
          console.log(JSON.parse(this.fileContent))
        })
        .catch(error => {
          this.error(error)
        })
      this.$store.dispatch('OCIS-JUPYTER/generateHTML', JSON.stringify(this.fileContent))
    }
  },
  mounted: function() {
    this.filePath = this.$route.params.filePath
    this.loadFileContent()

    console.log(this.ipynbHTML)
  }
}
</script>
