import 'regenerator-runtime/runtime'
import Preview from './components/Preview.vue'
import store from './store'

const appInfo = {
  name: 'OCIS-JUPYTER',
  id: 'ocis-jupyter',
  icon: 'text',
  extensions: [
    {
      extension: 'ipynb',
      newTab: true,
      routeName: 'ocis-jupyter'
    }
  ]
}

const routes = [
  {
    name: 'ocis-jupyter',
    path: '/preview/:filePath',
    components: {
      app: Preview
    }
  }
]

export default {
  appInfo,
  store,
  routes
}
