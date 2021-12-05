import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  isAuth: false,
  userId: null,
  userLogin: '',
}

export const {
  actions: authActions,
  reducer: authReducer 
} = createSlice({
  name: "auth",
  initialState,
  reducers: {
    login(state, {payload}) {
      state.isAuth = true
      state.userId = payload.userId
      state.userLogin = payload.userLogin
    },
    logout: (state) => initialState,
  }
})

export default authReducer
