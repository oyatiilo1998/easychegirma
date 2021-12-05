import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  loginLoader: false,
  fullPageLoader: false
}

export const {
  actions: loaderAction,
  reducer: loaderReducer
} = createSlice({
  name: "loader",
  initialState,
  reducers: {
    loaderON: (state, {payload}) => {state[payload] = true},
    loaderOFF: (state, {payload}) => {state[payload] = false}
  }
})
