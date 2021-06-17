import 'regenerator-runtime/runtime'
import App from './components/App.vue'
import store from './store'

const appInfo = {
  name: 'OCIS-JUPYTER',
  id: 'ocis-jupyter',
  icon: 'info',
  isFileEditor: false,
  extensions: []
}

const routes = [
  {
    name: 'ocis-jupyter',
    path: '/',
    components: {
      app: App
    }
  }
]

const navItems = [
  {
    name: 'OCIS-JUPYTER',
    iconMaterial: appInfo.icon,
    route: {
      name: 'ocis-jupyter',
      path: `/${appInfo.id}/`
    }
  }
]

export default {
  appInfo,
  store,
  routes,
  navItems
}
