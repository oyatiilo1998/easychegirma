import { createSlice } from "@reduxjs/toolkit"

export const {
  actions: mainAction,
  reducer: mainReducer 
} = createSlice({
  name: "main",
  initialState: {
    catalogs: [],
    products: []
  },
  reducers: {
    addAlert(state, {payload}) {
      state.alerts.push(payload)
    },
    deleteAlert(state, {payload}) {
      state.alerts = state.alerts.filter(alert => alert.id !== payload)
    }
  },
})
