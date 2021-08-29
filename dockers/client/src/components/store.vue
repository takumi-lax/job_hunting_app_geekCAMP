import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

interface State {
  message: null
}

const initState: State = {
    message: ''
};

export default new Vuex.Store({
  state: initState,
  mutations: {
    setMessage(state, { message }: { message: string }) {
      state.message = message;
    }
  }
})