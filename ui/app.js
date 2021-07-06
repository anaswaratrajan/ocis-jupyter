import 'regenerator-runtime/runtime'
import App from './components/App.vue'
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
