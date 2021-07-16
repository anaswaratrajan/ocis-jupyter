// eslint-disable-next-line camelcase
import { Generate_HTML } from './client/ocis-jupyter'
import axios from 'axios'

const state = {
  config: null,
  nbcontent: ''
}

const getters = {
  config: state => state.config,
  nbcontent: state => state.nbcontent,
  getServerForJsClient: (state, getters, rootState, rootGetters) => rootGetters.configuration.server.replace(/\/$/, '')
}

const actions = {
  // Used by ocis-web.
  loadConfig ({ commit }, config) {
    commit('LOAD_CONFIG', config)
  },

  generateHTML ({ commit, dispatch, getters, rootGetters }, value) {
    injectAuthToken(rootGetters)
    Generate_HTML({
      $domain: getters.getServerForJsClient,
      body: { JSONString: value }
    })
      .then(response => {
        console.log(response)
        
        if (response.status === 200 || response.status === 201) {
          commit('SET_NBCONTENT', response.data.HTMLString)
        } else {
          dispatch('showMessage', {
            title: 'Response failed',
            desc: response.statusText,
            status: 'danger'
          }, { root: true })
        }
      })
      .catch(error => {
        console.error(error)

        dispatch('showMessage', {
          title: 'Saving your name failed',
          desc: error.message,
          status: 'danger'
        }, { root: true })
      })
  }
}

const mutations = {
  SET_NBCONTENT (state, payload) {
    state.nbcontent = payload
  },

  LOAD_CONFIG (state, config) {
    state.config = config
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}

function injectAuthToken (rootGetters) {
  axios.interceptors.request.use(config => {
    if (typeof config.headers.Authorization === 'undefined') {
      const token = rootGetters.user.token
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }
    }
    return config
  })
}
