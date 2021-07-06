import 'regenerator-runtime/runtime'
import App from './components/App.vue'
import store from './store'

const appInfo = {
  name: 'OCIS-JUPYTER',
  id: 'ocis-jupyter',
  icon: 'text',
  isFileEditor: true,
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
    path: '/',
    components: {
      app: App
    }
  }
]

const navItems = [
  {
    name: 'ocis-jupyter',
    iconMaterial: appInfo.icon,
    route: {
      name: 'ocis-jupyter',
      path: `/${appInfo.id}/`
    }
  }
]



export default {
  appInfo,
  navItems,
  store,
  routes
}
